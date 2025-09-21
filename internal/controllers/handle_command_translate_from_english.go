package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// handleCommandTranslateFromEnglish handles the /translate_from_english command
func (h *Handler) handleCommandTranslateFromEnglish(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	h.modeMap[update.Message.Chat.ID] = ModeTranslateFromEnglish

	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"translate from english mode",
	))

	return err
}
