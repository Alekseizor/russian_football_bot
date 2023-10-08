package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type HistoryState struct {
}

func (state HistoryState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, History, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		ChampionsState{}.PreviewProcess(ctc)
		return &ChampionsState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, History, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		ChampionsState{}.PreviewProcess(ctc)
		return &ChampionsState{}
	case "Назад":
		AttackState{}.PreviewProcess(ctc)
		return &AttackState{}
	default:
		state.PreviewProcess(ctc)
		return &HistoryState{}
	}
}

func (state HistoryState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы болеете за команды, которые имеют большую историю и традиции?")
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

func (state HistoryState) Name() string {
	return "HistoryState"
}

// ///////////////////////////////////////////////////////
