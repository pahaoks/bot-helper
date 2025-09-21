package app

import (
	"bot-helper/internal/domain/repositories"
	"bot-helper/pkg/config"
)

type Config struct {
	Telegram repositories.TelegramConfig
	ChatGPT  repositories.ChatGPTConfig
	AnkiWeb  repositories.AnkiWebConfig
}

func NewConfig() Config {
	cfg := Config{}
	config.Load(&cfg)
	return cfg
}
