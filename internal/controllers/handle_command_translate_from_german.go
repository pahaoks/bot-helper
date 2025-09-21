package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// handleCommandTranslateFromGerman handles the /translate_from_german command
func (h *Handler) handleCommandTranslateFromGerman(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	h.modeMap[update.Message.Chat.ID] = ModeTranslateFromGerman

	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"translate from german mode",
	))

	return err
}
