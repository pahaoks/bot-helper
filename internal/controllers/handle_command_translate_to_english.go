package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// handleCommandTranslateToEnglish handles the /translate_to_english command
func (h *Handler) handleCommandTranslateToEnglish(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	h.modeMap[update.Message.Chat.ID] = ModeTranslateToEnglish

	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"translate to english mode",
	))

	return err
}
