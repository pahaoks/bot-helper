package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// handleCommandTranslateFromSlovak handles the /translate_from_slovak command
func (h *Handler) handleCommandTranslateFromSlovak(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	h.modeMap[update.Message.Chat.ID] = ModeTranslateFromSlovak

	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"translate from slovak mode",
	))

	return err
}
