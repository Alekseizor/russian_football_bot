package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type StrongLeadershipState struct {
}

func (state StrongLeadershipState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StrongLeadership, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		YoungTrainersState{}.PreviewProcess(ctc)
		return &YoungTrainersState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StrongLeadership, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		YoungTrainersState{}.PreviewProcess(ctc)
		return &YoungTrainersState{}
	case "Назад":
		NationwideFansState{}.PreviewProcess(ctc)
		return &NationwideFansState{}
	default:
		state.PreviewProcess(ctc)
		return &StrongLeadershipState{}
	}
}

func (state StrongLeadershipState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы любите команды, которые имеют сильный лидерский состав?")
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

func (state StrongLeadershipState) Name() string {
	return "StrongLeadershipState"
}

// ///////////////////////////////////////////////////////
