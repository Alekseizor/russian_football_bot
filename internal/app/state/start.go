package state

import (
	"context"

	"github.com/Alekseizor/ordering-bot/internal/app/ds"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	Colors             = "colors"
	StyleGame          = "style_game"
	History            = "history"
	Champions          = "champions"
	LocalFans          = "local_fans"
	Region             = "region"
	StarPlayer         = "star_player"
	StrongDef          = "strong_def"
	FastGame           = "fast_game"
	YoungTalent        = "young_talent"
	StrongAttack       = "strong_attack"
	ForeignPlayers     = "foreign_players"
	TechnicalGame      = "technical_game"
	ExperiencedPlayers = "experienced_players"
	StrongCharacter    = "strong_character"
	YoungCoaches       = "young_coaches"
	TeamPlay           = "team_play"
	NationwideFans     = "nationwide_fans"
	StrongLeadership   = "strong_leadership"
	YoungTrainers      = "young_trainers"
)

type ChatContext struct {
	User *ds.User
	Vk   *api.VK
	Ctx  *context.Context
	Db   *sqlx.DB
}

type State interface {
	Name() string                                      //получаем название состояния в виде строки, чтобы в дальнейшем куда-то записать(БД)
	Process(ChatContext, object.MessagesMessage) State //нужно взять контекст, посмотреть на каком состоянии сейчас пользователь, метод должен вернуть состояние
	PreviewProcess(ctc ChatContext)
}

// ////////////////////////////////////////////////////////
type StartState struct {
}

func (state StartState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	if messageText == "Начать" {
		ColorState{}.PreviewProcess(ctc)
		return &ColorState{}
	}
	state.PreviewProcess(ctc)
	return &StartState{}
}

func (state StartState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Привет! Для того, чтобы начать выбирать клуб, нажми кнопку \"Начать\"")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Начать", "", "secondary")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state StartState) Name() string {
	return "StartState"
}

// ///////////////////////////////////////////////////////
