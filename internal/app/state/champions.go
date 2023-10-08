package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type ChampionsState struct {
}

func (state ChampionsState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, Champions, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		LocalFansState{}.PreviewProcess(ctc)
		return &LocalFansState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, Champions, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		LocalFansState{}.PreviewProcess(ctc)
		return &LocalFansState{}
	case "Назад":
		HistoryState{}.PreviewProcess(ctc)
		return &HistoryState{}
	default:
		state.PreviewProcess(ctc)
		return &ChampionsState{}
	}
}

func (state ChampionsState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы хотите поддерживать команду, которая выступает в Лиге Чемпионов?")
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

func (state ChampionsState) Name() string {
	return "ChampionsState"
}

// ///////////////////////////////////////////////////////
