package controllers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handleMessageTranslateToGerman handles messages in translate to German mode
func (h *Handler) handleMessageTranslateToGerman(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	res, err := h.modelPrompt(bot, update, "Переведи на немецкий, коротко: ")
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
		"German", "Basic (and reversed card)",
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
