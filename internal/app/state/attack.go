package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type AttackState struct {
}

func (state AttackState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StyleGame, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		HistoryState{}.PreviewProcess(ctc)
		return &HistoryState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StyleGame, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		HistoryState{}.PreviewProcess(ctc)
		return &HistoryState{}
	case "Назад":
		ColorState{}.PreviewProcess(ctc)
		return &ColorState{}
	default:
		state.PreviewProcess(ctc)
		return &AttackState{}
	}
}

func (state AttackState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы предпочитаете команды, которые играют атакующий футбол?")
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

func (state AttackState) Name() string {
	return "AttackState"
}

// ///////////////////////////////////////////////////////
