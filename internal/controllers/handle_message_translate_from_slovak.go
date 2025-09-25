package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handleMessageTranslateFromSlovak handles messages in translate from Slovak mode
func (h *Handler) handleMessageTranslateFromSlovak(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	res, err := h.modelPrompt(bot, update, "Переведи со словацкого, на русский, коротко: ")
	if err != nil {
		return err
	}

	_, err = bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		res.GetText(),
	))
	if err != nil {
		return err
	}

	err = h.ankiWebRepo.AddNote(
		"Slovak", "Basic (and reversed card)",
		update.Message.Text,
		res.GetText(),
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
