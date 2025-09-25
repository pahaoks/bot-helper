package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handleMessageTranslateToEnglish handles messages in translate to English mode
func (h *Handler) handleMessageTranslateToEnglish(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	res, err := h.modelPrompt(bot, update, "Переведи на английский, коротко: ")
	if err != nil {
		return err
	}

	err = h.ankiWebRepo.AddNote(
		"English", "Basic (and reversed card)",
		res.GetText(),
		update.Message.Text,
	)
	if err != nil {
		return err
	}

	err = h.ankiWebRepo.Sync()
	if err != nil {
		return err
	}

	_, err = bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"added to anki",
	))

	return err
}
