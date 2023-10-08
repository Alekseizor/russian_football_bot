package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type StrongDefState struct {
}

func (state StrongDefState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StrongDef, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		FastGameState{}.PreviewProcess(ctc)
		return &FastGameState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StrongDef, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		FastGameState{}.PreviewProcess(ctc)
		return &FastGameState{}
	case "Назад":
		StarPlayerState{}.PreviewProcess(ctc)
		return &StarPlayerState{}
	default:
		state.PreviewProcess(ctc)
		return &StrongDefState{}
	}
}

func (state StrongDefState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы хотите болеть за команду, которая имеет сильную защиту?")
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

func (state StrongDefState) Name() string {
	return "StrongDefState"
}

// ///////////////////////////////////////////////////////
