package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type RegionState struct {
}

func (state RegionState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, Region, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		StarPlayerState{}.PreviewProcess(ctc)
		return &StarPlayerState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, Region, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		StarPlayerState{}.PreviewProcess(ctc)
		return &StarPlayerState{}
	case "Назад":
		LocalFansState{}.PreviewProcess(ctc)
		return &LocalFansState{}
	default:
		state.PreviewProcess(ctc)
		return &RegionState{}
	}
}

func (state RegionState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы хотите поддерживать команду, которая выступает в Москве?")
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

func (state RegionState) Name() string {
	return "RegionState"
}

// ///////////////////////////////////////////////////////
