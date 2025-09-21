package controllers

import (
	"bot-helper/internal/domain/repositories"
)

type Listener struct {
	handler      *Handler
	telegramRepo *repositories.TelegramRepository
	chatGptRepo  *repositories.ChatGPTRepository
}

func NewListener(
	handler *Handler,
	telegramRepo *repositories.TelegramRepository,
	chatGptRepo *repositories.ChatGPTRepository,
) *Listener {
	return &Listener{
		handler:      handler,
		telegramRepo: telegramRepo,
		chatGptRepo:  chatGptRepo,
	}
}

func (l *Listener) Run() error {
	return l.telegramRepo.Run(
		func(bot *repositories.BotAPI, update repositories.Update) error {
			if update.Message == nil {
				return nil
			}

			if update.Message.IsCommand() {
				return l.handler.HandleCommand(bot, update)
			}

			return l.handler.HandleMessage(bot, update)
		},
	)
}
