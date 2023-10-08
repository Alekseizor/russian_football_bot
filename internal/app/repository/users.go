package repository

import (
	"context"
	"fmt"

	"github.com/Alekseizor/ordering-bot/internal/app/ds"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func GetUser(Db *sqlx.DB, VkID int) (ds.User, error) {
	var user ds.User
	err := Db.QueryRow("SELECT * from users WHERE vk_id =$1", VkID).Scan(&user.VkID, &user.State, &user.Colors, &user.StyleGame, &user.History, &user.Champions, &user.LocalFans, &user.Region, &user.StarPlayer, &user.StrongDef, &user.FastGame, &user.YoungTalent, &user.StrongAttack, &user.ForeignPlayers, &user.TechnicalGame, &user.ExperiencedPlayers, &user.StrongCharacter, &user.YoungCoaches, &user.TeamPlay, &user.NationwideFans, &user.StrongLeadership, &user.YoungTrainers)
	if err != nil {
		log.Printf("ошибка при получении пользователя: %v", err)
		return ds.User{}, err
	}
	return user, nil
}

func UserUpdateParam(ctx *context.Context, Db *sqlx.DB, VkID int, param string, value bool) error {
	request := fmt.Sprintf("UPDATE users SET %s=$1 WHERE vk_id=$2", param)
	_, err := Db.ExecContext(*ctx, request, value, VkID)
	if err != nil {
		log.Printf("Ошибка при обновлении параметра: %v", err)
	}
	return err
}
