package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// handleMessageUnknown handles unknown messages
func (h *Handler) handleMessageUnknown(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Sorry, I didn't understand that command. Please try again.",
	))
	return err
}
