package telegram

import (
	http "git-releaser/internal/integrations"

	"github.com/vicanso/go-axios"
)

type (
	telegram struct {
		instance *axios.Instance
		botToken string
	}

	Telegram interface {
		GetInstance() *axios.Instance
		SendMessage(to, text string) (err error)
		SetBotToken(token string)
	}

	sendMessageReqBody struct {
		ChatID    string `json:"chat_id"`
		Text      string `json:"text"`
		ParseMode string `json:"parse_mode"`
	}
)

func New(url string) (t Telegram) {
	instance := http.New(url)
	instance.Config.ResponseInterceptors = append(instance.Config.ResponseInterceptors, ParseErrorResponse)

	t = &telegram{
		instance: instance,
	}
	return
}

func (t *telegram) GetInstance() *axios.Instance {
	return t.instance
}

func (t *telegram) SetBotToken(token string) {
	t.botToken = token
}

func (t *telegram) SendMessage(to, text string) (err error) {
	reqConfig := &axios.Config{
		Method: "POST",
		URL:    "/bot:bot_token/sendMessage",
		Params: map[string]string{
			"bot_token": t.botToken,
		},
		Body: &sendMessageReqBody{
			ChatID:    to,
			ParseMode: "Markdown",
			Text:      text,
		},
	}

	res, err := t.instance.Request(reqConfig)

	if err == nil {
		var telRes interface{}
		_ = res.JSON(&telRes)
		// fmt.Printf("Telegram Response: %#v\n\n", telRes)
	}

	return
}
