package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type StrongCharacterState struct {
}

func (state StrongCharacterState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StrongCharacter, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		YoungCoachesState{}.PreviewProcess(ctc)
		return &YoungCoachesState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, StrongCharacter, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		YoungCoachesState{}.PreviewProcess(ctc)
		return &YoungCoachesState{}
	case "Назад":
		ExperiencedPlayersState{}.PreviewProcess(ctc)
		return &ExperiencedPlayersState{}
	default:
		state.PreviewProcess(ctc)
		return &StrongCharacterState{}
	}
}

func (state StrongCharacterState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы любите команды, которые имеют сильный характер и дух борьбы?")
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

func (state StrongCharacterState) Name() string {
	return "StrongCharacterState"
}

// ///////////////////////////////////////////////////////
