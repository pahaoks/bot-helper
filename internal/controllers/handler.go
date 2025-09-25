package controllers

import (
	"bot-helper/internal/domain/entities"
	"bot-helper/internal/domain/repositories"
	"bot-helper/pkg/voiceconverter"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Cmd string

const (
	CmdStart                Cmd = "/start"
	CmdTranslateToSlovak    Cmd = "/translate_to_slovak"
	CmdTranslateToEnglish   Cmd = "/translate_to_english"
	CmdTranslateToGerman    Cmd = "/translate_to_german"
	CmdTranslateFromSlovak  Cmd = "/translate_from_slovak"
	CmdTranslateFromEnglish Cmd = "/translate_from_english"
	CmdTranslateFromGerman  Cmd = "/translate_from_german"
)

type Mode uint16

const (
	ModeTranslateToSlovak Mode = iota
	ModeTranslateToEnglish
	ModeTranslateToGerman
	ModeTranslateFromSlovak
	ModeTranslateFromEnglish
	ModeTranslateFromGerman
)

type Handler struct {
	chatGptRepo     *repositories.ChatGPTRepository
	ankiWebRepo     *repositories.AnkiWebRepository
	commandsMap     map[Cmd]repositories.TelegramMessageCallback
	modeMap         map[int64]Mode
	messageHandlers map[Mode]repositories.TelegramMessageCallback
	voiceConverter  *voiceconverter.VoiceConverter
	telegramRepo    *repositories.TelegramRepository
}

func NewHandler(
	chatGptRepo *repositories.ChatGPTRepository,
	ankiWebRepo *repositories.AnkiWebRepository,
	voiceConverter *voiceconverter.VoiceConverter,
	telegramRepo *repositories.TelegramRepository,
) *Handler {
	h := &Handler{
		chatGptRepo:    chatGptRepo,
		ankiWebRepo:    ankiWebRepo,
		modeMap:        make(map[int64]Mode),
		voiceConverter: voiceConverter,
		telegramRepo:   telegramRepo,
	}

	h.commandsMap = map[Cmd]repositories.TelegramMessageCallback{
		CmdStart:                h.handleCommandStart,
		CmdTranslateToSlovak:    h.handleCommandTranslateToSlovak,
		CmdTranslateToEnglish:   h.handleCommandTranslateToEnglish,
		CmdTranslateToGerman:    h.handleCommandTranslateToGerman,
		CmdTranslateFromSlovak:  h.handleCommandTranslateFromSlovak,
		CmdTranslateFromEnglish: h.handleCommandTranslateFromEnglish,
		CmdTranslateFromGerman:  h.handleCommandTranslateFromGerman,
	}

	h.messageHandlers = map[Mode]repositories.TelegramMessageCallback{
		ModeTranslateToSlovak:    h.handleMessageTranslateToSlovak,
		ModeTranslateToEnglish:   h.handleMessageTranslateToEnglish,
		ModeTranslateToGerman:    h.handleMessageTranslateToGerman,
		ModeTranslateFromSlovak:  h.handleMessageTranslateFromSlovak,
		ModeTranslateFromEnglish: h.handleMessageTranslateFromEnglish,
		ModeTranslateFromGerman:  h.handleMessageTranslateFromGerman,
	}

	return h
}

func (h *Handler) HandleCommand(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	if command, exists := h.commandsMap[Cmd(update.Message.Text)]; exists {
		return command(bot, update)
	} else {
		return h.handleCommandUnknown(bot, update)
	}
}

func (h *Handler) HandleMessage(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
) error {
	if mode, exists := h.messageHandlers[h.modeMap[update.Message.Chat.ID]]; exists {
		return mode(bot, update)
	}
	return h.handleMessageUnknown(bot, update)
}

func (h *Handler) modelPrompt(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	prefix string,
) (*entities.ChatGPTResponse, error) {
	text := update.Message.Text
	var err error

	if update.Message.Voice != nil {
		text, err = h.translateVoice(update)
		if err != nil {
			return nil, h.logError(bot, update, err)
		}
	}

	res, err := h.chatGptRepo.Prompt(prefix + text)
	if err != nil {
		return nil, h.logError(bot, update, err)
	}
	if res.Error != nil && res.Error.Message != "" {
		return nil, h.logError(
			bot, update, fmt.Errorf("chatgpt error: %s", res.Error.Message))
	}

	return &res, err
}

func (h *Handler) translateVoice(
	update tgbotapi.Update,
) (string, error) {
	fileBytes, err := h.telegramRepo.GetFileContent(update.Message.Voice.FileID)
	if err != nil {
		return "", err
	}
	mp3File, err := h.voiceConverter.OggToMp3(fileBytes)
	if err != nil {
		return "", err
	}

	res, err := h.chatGptRepo.TranscribeAudio(mp3File)
	if err != nil {
		return "", err
	}

	return res.Text, nil
}

func (h *Handler) logError(
	bot *tgbotapi.BotAPI,
	update tgbotapi.Update,
	err error,
) error {
	bot.Send(tgbotapi.NewMessage(
		update.Message.Chat.ID,
		fmt.Sprintf("error handling update: %v", err),
	))
	return err
}
