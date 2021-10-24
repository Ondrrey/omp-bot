package streaming

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/streaming/stream"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type StreamingCommander struct {
	bot             *tgbotapi.BotAPI
	streamCommander stream.StreamCommandInvoker
}

func NewCommander(bot *tgbotapi.BotAPI) Commander {
	return &StreamingCommander{
		bot:             bot,
		streamCommander: stream.NewCommander(bot),
	}
}

func (c *StreamingCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "stream":
		c.streamCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("StreamingCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *StreamingCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "stream":
		c.streamCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("StreamingCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
