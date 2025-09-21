package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// handleCommandTranslateToGerman handles the /translate_to_german command
func (h *Handler) handleCommandTranslateToGerman(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	h.modeMap[update.Message.Chat.ID] = ModeTranslateToGerman

	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"translate to german mode",
	))

	return err
}
