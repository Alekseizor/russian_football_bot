package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type TeamPlayState struct {
}

func (state TeamPlayState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, TeamPlay, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		NationwideFansState{}.PreviewProcess(ctc)
		return &NationwideFansState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, TeamPlay, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		NationwideFansState{}.PreviewProcess(ctc)
		return &NationwideFansState{}
	case "Назад":
		YoungCoachesState{}.PreviewProcess(ctc)
		return &YoungCoachesState{}
	default:
		state.PreviewProcess(ctc)
		return &TeamPlayState{}
	}
}

func (state TeamPlayState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы предпочитаете команды, которые играют командный футбол?")
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

func (state TeamPlayState) Name() string {
	return "TeamPlayState"
}

// ///////////////////////////////////////////////////////
