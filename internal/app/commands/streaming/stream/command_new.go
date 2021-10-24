package stream

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *StreamCommand) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if args == "" {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Please specify stream id")
		c.bot.Send(msg)
		return
	}

	var stream streaming.Stream
	err := json.Unmarshal([]byte(args), &stream)
	if err != nil {
		SendAndLog(c.bot, inputMessage.Chat.ID, "Could not parse input data")
		return
	}

	streamID, err := c.streamService.Create(stream)
	if err != nil {
		SendAndLog(c.bot, inputMessage.Chat.ID, err.Error())
		return
	}

	SendAndLog(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Stream %v has been successfully created", streamID))
}
