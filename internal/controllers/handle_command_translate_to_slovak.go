package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// handleCommandTranslateToSlovak handles the /translate_to_slovak command
func (h *Handler) handleCommandTranslateToSlovak(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	h.modeMap[update.Message.Chat.ID] = ModeTranslateToSlovak

	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"translate to slovak mode",
	))

	return err
}
