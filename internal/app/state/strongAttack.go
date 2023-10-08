package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type StrongAttackState struct {
}

func (state StrongAttackState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StrongAttack, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		ForeignPlayersState{}.PreviewProcess(ctc)
		return &ForeignPlayersState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StrongAttack, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		ForeignPlayersState{}.PreviewProcess(ctc)
		return &ForeignPlayersState{}
	case "Назад":
		YoungTalentState{}.PreviewProcess(ctc)
		return &YoungTalentState{}
	default:
		state.PreviewProcess(ctc)
		return &StrongAttackState{}
	}
}

func (state StrongAttackState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы любите команды, которые имеют сильную атаку?")
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

func (state StrongAttackState) Name() string {
	return "StrongAttackState"
}

// ///////////////////////////////////////////////////////
