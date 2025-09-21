package repositories

import (
	"bot-helper/pkg/logger"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramConfig holds the configuration for the Telegram bot.
type TelegramConfig struct {
	BotToken string
}

type BotAPI = tgbotapi.BotAPI
type Update = tgbotapi.Update

// TelegramRepository handles interactions with the Telegram Bot API.
type TelegramRepository struct {
	config TelegramConfig
	logger logger.Logger
}

// TelegramMessageCallback defines the signature for the callback function
type TelegramMessageCallback func(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error

// NewTelegramRepository creates a new instance of TelegramRepository.
func NewTelegramRepository(
	config TelegramConfig,
	logger logger.Logger,
) *TelegramRepository {
	return &TelegramRepository{
		config: config,
		logger: logger,
	}
}

// Run starts the Telegram bot and listens for incoming messages.
func (r *TelegramRepository) Run(
	callback TelegramMessageCallback,
) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			r.logger.Error("recovered from panic:", recovered)
			err = fmt.Errorf("recovered from panic: %v", recovered)
		}
	}()

	bot, err := tgbotapi.NewBotAPI(r.config.BotToken)
	if err != nil {
		r.logger.Error("failed to create bot api", err)
		return err
	}

	bot.Debug = true

	r.logger.Info("authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ch := bot.GetUpdatesChan(u)

	for update := range ch {
		err := callback(bot, update)
		if err != nil {
			r.logger.Error("error handling update", err)
			continue
		}
	}

	return nil
}
