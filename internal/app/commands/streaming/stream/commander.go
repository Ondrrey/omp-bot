package stream

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/streaming/stream"
)

type StreamCommandInvoker interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath)

	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type StreamCommand struct {
	bot           *tgbotapi.BotAPI
	streamService stream.StreamServicer
}

func NewCommander(bot *tgbotapi.BotAPI) StreamCommandInvoker {
	return &StreamCommand{
		bot:           bot,
		streamService: stream.NewStreamService(),
	}
}

func (c *StreamCommand) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("StreamCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *StreamCommand) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		fmt.Println(msg)
		c.Help(msg)
	case "get":
		fmt.Println(msg)
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		fmt.Println(msg)
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
