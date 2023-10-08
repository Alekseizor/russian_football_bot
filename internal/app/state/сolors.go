package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type ColorState struct {
}

func (state ColorState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, Colors, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		AttackState{}.PreviewProcess(ctc)
		return &AttackState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, Colors, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		AttackState{}.PreviewProcess(ctc)
		return &AttackState{}
	case "Назад":
		StartState{}.PreviewProcess(ctc)
		return &StartState{}
	default:
		state.PreviewProcess(ctc)
		return &ColorState{}
	}
}

func (state ColorState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы любите команды, которые выступают в ярких цветах?")
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

func (state ColorState) Name() string {
	return "ColorState"
}

// ///////////////////////////////////////////////////////
