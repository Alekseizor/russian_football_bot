package state

import (
	"sort"
	"strconv"

	"github.com/Alekseizor/ordering-bot/internal/app/ds"
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

var (
	clubs = []ds.Club{
		{"Спартак", true, true, true, true, true, true, false, true, true, true, true, false, true, true, false, false, true, true, false, true, 0},
		{"ЦСКА", false, true, true, true, true, true, true, true, true, true, true, false, true, true, false, true, true, true, false, true, 0},
		{"Зенит", true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, false, true, true, true, 0},
		{"Локомотив", true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, false, true, 0},
		{"Краснодар", false, true, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, false, false, 0},
		{"Ростов", true, false, false, false, false, false, false, true, true, true, true, true, true, true, true, false, true, true, false, false, 0},
		{"Уфа", false, false, false, false, false, false, false, true, false, true, false, true, false, true, true, false, true, false, false, false, 0},
		{"Арсенал", false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, false, true, false, false, false, 0},
		{"Тамбов", false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, false, false, false, false, false, 0},
		{"Ахмат", false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, false, false, false, false, false, 0},
	}
)

// ////////////////////////////////////////////////////////
type ResultState struct {
}

func (state ResultState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	switch messageText {
	case "Начать сначала":
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	case "Назад":
		YoungTrainersState{}.PreviewProcess(ctc)
		return &YoungTrainersState{}
	default:
		state.PreviewProcess(ctc)
		return &ResultState{}
	}
}

func (state ResultState) PreviewProcess(ctc ChatContext) {
	userClubs, err := countingResults(ctc)
	if err != nil {
		log.Printf("ошибка при подведении итогов:%s", err)
	}
	response := "Вот клубы, которые подходят вам:\n"
	for _, userClub := range userClubs {
		response += userClub.Name + " с вероятностью: " + strconv.Itoa(userClub.Probability*5) + "%\n"
	}
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message(response)
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Назад", "", "")
	k.AddRow()
	k.AddTextButton("Начать сначала", "", "")
	b.Keyboard(k)
	_, err = ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state ResultState) Name() string {
	return "ResultState"
}

// ///////////////////////////////////////////////////////

func countingResults(ctc ChatContext) ([]ds.Club, error) {
	user, err := repository.GetUser(ctc.Db, ctc.User.VkID)
	if err != nil {
		return nil, err
	}
	threeBestClubs := make([]ds.Club, 3)
	probability := 0
	clubProbability := 0
	for _, club := range clubs {
		clubProbability = 0
		if user.Colors == club.Colors {
			clubProbability += 1
		}
		if user.StyleGame == club.StyleGame {
			clubProbability += 1
		}
		if user.History == club.History {
			clubProbability += 1
		}
		if user.Champions == club.Champions {
			clubProbability += 1
		}
		if user.LocalFans == club.LocalFans {
			clubProbability += 1
		}
		if user.Region == club.Region {
			clubProbability += 1
		}
		if user.StarPlayer == club.StarPlayer {
			clubProbability += 1
		}
		if user.StrongDef == club.StrongDef {
			clubProbability += 1
		}
		if user.FastGame == club.FastGame {
			clubProbability += 1
		}
		if user.YoungTalent == club.YoungTalent {
			clubProbability += 1
		}
		if user.StrongAttack == club.StrongAttack {
			clubProbability += 1
		}
		if user.ForeignPlayers == club.ForeignPlayers {
			clubProbability += 1
		}
		if user.TechnicalGame == club.TechnicalGame {
			clubProbability += 1
		}
		if user.ExperiencedPlayers == club.ExperiencedPlayers {
			clubProbability += 1
		}
		if user.StrongCharacter == club.StrongCharacter {
			clubProbability += 1
		}
		if user.YoungCoaches == club.YoungCoaches {
			clubProbability += 1
		}
		if user.TeamPlay == club.TeamPlay {
			clubProbability += 1
		}
		if user.NationwideFans == club.NationwideFans {
			clubProbability += 1
		}
		if user.StrongLeadership == club.StrongLeadership {
			clubProbability += 1
		}
		if user.YoungTrainers == club.YoungTrainers {
			clubProbability += 1
		}
		if clubProbability > probability {
			club.Probability = clubProbability
			if len(threeBestClubs) < 3 {
				threeBestClubs = append(threeBestClubs, club)
				continue
			}
			threeBestClubs[2] = club
			sort.Slice(threeBestClubs, func(i int, j int) bool {
				return threeBestClubs[i].Probability > threeBestClubs[j].Probability
			})
			probability = threeBestClubs[2].Probability
		}
	}
	return threeBestClubs, nil
}
