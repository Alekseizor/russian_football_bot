package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type FastGameState struct {
}

func (state FastGameState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, FastGame, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		YoungTalentState{}.PreviewProcess(ctc)
		return &YoungTalentState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, FastGame, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		YoungTalentState{}.PreviewProcess(ctc)
		return &YoungTalentState{}
	case "Назад":
		StrongDefState{}.PreviewProcess(ctc)
		return &StrongDefState{}
	default:
		state.PreviewProcess(ctc)
		return &FastGameState{}
	}
}

func (state FastGameState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы предпочитаете команды, которые играют быстрый и динамичный футбол?")
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

func (state FastGameState) Name() string {
	return "FastGameState"
}

// ///////////////////////////////////////////////////////
