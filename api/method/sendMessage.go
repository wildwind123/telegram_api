//https://core.telegram.org/bots/api#sendmessage
package method

import (
	"encoding/json"
	"github.com/wildwind123/telegram_api/api/model"
)

type SendMessage string

type SendMessageResponseStruct struct {
	Ok     bool `json:"ok"`
	Result model.Message
}

type sendMessageParams struct {
	ChatId              string `json:"chat_id"`
	Text                string `json:"text"`
	DisableNotification *bool  `json:"disable_notification,omitempty"`
}

func (method *SendMessage) GetParams(ChatId, Text string) (*sendMessageParams, error) {

	if ChatId == "" {
		return nil, ErrorMsg{Msg: "Chat is required"}
	}

	if Text == "" {
		return nil, ErrorMsg{Msg: "Text is required"}
	}

	sm := &sendMessageParams{
		ChatId: ChatId,
		Text:   Text,
	}

	return sm, nil
}

func (method *sendMessageParams) GetResponseStruct() SendMessageResponseStruct {
	return SendMessageResponseStruct{}
}

func (method *sendMessageParams) GetMethodName() string {
	return "sendMessage"
}

func (method *sendMessageParams) Request() (*SendMessageResponseStruct, error) {
	resp, err := Request(method)
	if err != nil {
		return nil, err
	}
	respStr := method.GetResponseStruct()
	err = json.Unmarshal(resp, &respStr)
	if err != nil {
		return nil, err
	}
	return &respStr, nil
}
