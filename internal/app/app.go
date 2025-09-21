package app

import (
	"bot-helper/internal/controllers"
	"bot-helper/internal/domain/repositories"
	"bot-helper/pkg/logger"
	"net/http"
)

type App struct {
	listener *controllers.Listener
}

func New() *App {
	cfg := NewConfig()
	rt := http.DefaultTransport
	logger := logger.NewConsoleLogger()
	telegramRepo := repositories.NewTelegramRepository(cfg.Telegram, logger)
	chatGptRepo := repositories.NewChatGPTRepository(cfg.ChatGPT, rt)
	ankiWebRepo := repositories.NewAnkiWebRepository(cfg.AnkiWeb, rt, logger)

	return &App{
		listener: controllers.NewListener(
			controllers.NewHandler(
				chatGptRepo,
				ankiWebRepo,
			),
			telegramRepo,
			chatGptRepo,
		),
	}
}

func (a *App) Run() error {
	return a.listener.Run()
}
