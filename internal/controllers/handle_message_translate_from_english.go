package controllers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// handleMessageTranslateFromEnglish handles messages in translate from English mode
func (h *Handler) handleMessageTranslateFromEnglish(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	res, err := h.chatGptRepo.Prompt("Переведи со английского, на русский, коротко: " + update.Message.Text)
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
		"English", "Basic (and reversed card)",
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
