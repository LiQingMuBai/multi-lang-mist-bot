package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"ushield_bot/internal/bot"
	"ushield_bot/internal/domain"
)

type HelpHandler struct{}

func NewHelpHandler() *HelpHandler {
	return &HelpHandler{}
}

func (h *HelpHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "📞請聯繫客服 @Ushield001\n",
	}
	_ = b.SendMessage(msg, bot.DefaultChannel)
	return nil
}
