package controllers

import (
	"sort"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleStart handles the /start command
func (h *Handler) handleCommandStart(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	commands := []string{}

	for key := range h.commandsMap {
		if key == CmdStart {
			continue
		}
		commands = append(commands, string(key))
	}

	sort.SliceStable(commands, func(i, j int) bool {
		return commands[i] < commands[j]
	})

	cmdList := strings.Join(commands, "\n\n")

	_, err := bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Here are the available commands:\n"+cmdList,
	))

	return err
}
