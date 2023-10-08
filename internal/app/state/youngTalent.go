package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type YoungTalentState struct {
}

func (state YoungTalentState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, YoungTalent, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		StrongAttackState{}.PreviewProcess(ctc)
		return &StrongAttackState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, YoungTalent, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		StrongAttackState{}.PreviewProcess(ctc)
		return &StrongAttackState{}
	case "Назад":
		FastGameState{}.PreviewProcess(ctc)
		return &FastGameState{}
	default:
		state.PreviewProcess(ctc)
		return &YoungTalentState{}
	}
}

func (state YoungTalentState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы хотите поддерживать команду, которая имеет много молодых и перспективных игроков?")
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

func (state YoungTalentState) Name() string {
	return "YoungTalentState"
}

// ///////////////////////////////////////////////////////
