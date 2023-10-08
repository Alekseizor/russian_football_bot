package state

import (
	"github.com/Alekseizor/ordering-bot/internal/app/repository"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
	log "github.com/sirupsen/logrus"
)

// ////////////////////////////////////////////////////////
type LocalFansState struct {
}

func (state LocalFansState) Process(ctc ChatContext, msg object.MessagesMessage) State {
	messageText := msg.Text
	var err error
	switch messageText {
	case "Да":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, LocalFans, true)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		RegionState{}.PreviewProcess(ctc)
		return &RegionState{}
	case "Нет":
		err = repository.UserUpdateParam(ctc.Ctx, ctc.Db, ctc.User.VkID, LocalFans, false)
		if err != nil {
			log.Printf("не удалось записать параметр: ")
		}
		RegionState{}.PreviewProcess(ctc)
		return &RegionState{}
	case "Назад":
		ChampionsState{}.PreviewProcess(ctc)
		return &ChampionsState{}
	default:
		state.PreviewProcess(ctc)
		return &LocalFansState{}
	}
}

func (state LocalFansState) PreviewProcess(ctc ChatContext) {
	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.Message("Вы хотите болеть за команду, которая имеет много локальных фанатов?")
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

func (state LocalFansState) Name() string {
	return "LocalFansState"
}

// ///////////////////////////////////////////////////////
