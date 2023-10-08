package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type NationwideFansState struct {
}

func (state NationwideFansState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, NationwideFans, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		StrongLeadershipState{}.PreviewProcess(ctc)
		return &StrongLeadershipState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, NationwideFans, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		StrongLeadershipState{}.PreviewProcess(ctc)
		return &StrongLeadershipState{}
	case "Назад":
		TeamPlayState{}.PreviewProcess(ctc)
		return &TeamPlayState{}
	default:
		state.PreviewProcess(ctc)
		return &NationwideFansState{}
	}
}

func (state NationwideFansState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы хотите поддерживать команду, которая имеет много болельщиков по всей России?")
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

func (state NationwideFansState) Name() string {
	return "NationwideFansState"
}

// ///////////////////////////////////////////////////////
