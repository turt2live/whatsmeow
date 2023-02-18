package whatsmeow

import (
	"encoding/json"
	"fmt"
	"sync"

	"go.mau.fi/whatsmeow/types/events"
)

type pendingMessage struct {
	ChatID   string `json:"chat_id"`
	SenderID string `json:"sender_id"`
	Message  string `json:"message"`
}

var timeline = make([]*pendingMessage, 0)
var lock = &sync.Mutex{}

func (cli *Client) SetupTimeline() {
	cli.AddEventHandler(func(evt interface{}) {
		fmt.Printf("%T %s", evt, evt)
		fmt.Println()
		msg, ok := evt.(*events.Message)
		if !ok {
			return
		}

		if msg.Message == nil || msg.Message.Conversation == nil {
			return
		}

		chatId := msg.Info.Chat.String()
		text := msg.Message.GetConversation()
		senderId := msg.Info.Sender.String()
		lock.Lock()
		timeline = append(timeline, &pendingMessage{
			ChatID:   chatId,
			SenderID: senderId,
			Message:  text,
		})
		lock.Unlock()
	})
}

func (cli *Client) PollTimeline() (string, error) {
	lock.Lock()
	vals := timeline
	timeline = make([]*pendingMessage, 0)
	lock.Unlock()

	b, err := json.Marshal(vals)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
