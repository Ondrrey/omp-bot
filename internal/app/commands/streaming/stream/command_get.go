package stream

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *StreamCommand) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if args == "" {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Please specify stream id")
		c.bot.Send(msg)
		return
	}

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	stream, err := c.streamService.Describe(idx)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, stream.String())
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamCommander.Help: error sending reply message to chat - %v", err)
	}
}
