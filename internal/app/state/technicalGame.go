package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type TechnicalGameState struct {
}

func (state TechnicalGameState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, TechnicalGame, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		ExperiencedPlayersState{}.PreviewProcess(ctc)
		return &ExperiencedPlayersState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, TechnicalGame, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		ExperiencedPlayersState{}.PreviewProcess(ctc)
		return &ExperiencedPlayersState{}
	case "Назад":
		ForeignPlayersState{}.PreviewProcess(ctc)
		return &ForeignPlayersState{}
	default:
		state.PreviewProcess(ctc)
		return &TechnicalGameState{}
	}
}

func (state TechnicalGameState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы предпочитаете команды, которые играют техничный футбол?")
	b.PeerID(ctc.User.VkID)
	k := &object.MessagesKeyboard{}
	k.AddRow()
	k.AddTextButton("Да", "", "primary")
	k.AddRow()
	k.AddTextButton("Нет", "", "negative")
	k.AddRow()
	k.AddTextButton("Назад", "", "")
	b.Keyboard(k)
	_, err := ctc.Vk.MessagesSend(b.Params)
	if err != nil {
		log.Println("Failed to get record")
		log.Error(err)
	}
}

func (state TechnicalGameState) Name() string {
	return "TechnicalGameState"
}

// ///////////////////////////////////////////////////////
