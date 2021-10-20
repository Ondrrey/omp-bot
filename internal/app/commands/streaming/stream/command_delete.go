package stream

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *StreamCommand) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if args == "" {
		Send(c.bot, inputMessage.Chat.ID, "Please specify stream id")
		return
	}

	streamID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		SendAndLog(c.bot, inputMessage.Chat.ID, "wrong args")
		return
	}

	if _, err := c.streamService.Remove(streamID); err != nil {
		SendAndLog(c.bot, inputMessage.Chat.ID, err.Error())
	} else {
		SendAndLog(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Successfully deleted stream with ID=%d", streamID))
	}
}
