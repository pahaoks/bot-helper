package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// handleMessageTranslateToSlovak handles messages in translate to Slovak mode
func (h *Handler) handleMessageTranslateToSlovak(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	res, err := h.chatGptRepo.Prompt("Переведи на словацкий, коротко: " + update.Message.Text)
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
