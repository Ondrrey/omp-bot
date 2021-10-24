package stream

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

func (c *StreamCommand) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	if args == "" {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Please specify stream id")
		c.bot.Send(msg)
		return
	}

	arguments := strings.SplitN(args, " ", 2)

	if len(arguments) != 2 {
		SendAndLog(c.bot, inputMessage.Chat.ID, "Incorrect parameters")
		return
	}

	streamID, err := strconv.ParseUint(arguments[0], 10, 64)
	if err != nil {
		SendAndLog(c.bot, inputMessage.Chat.ID, "Wrong args")
		return
	}

	var stream streaming.Stream
	err = json.Unmarshal([]byte(arguments[1]), &stream)
	if err != nil {
		SendAndLog(c.bot, inputMessage.Chat.ID, "Could not parse input data")
		return
	}

	err = c.streamService.Update(streamID, stream)
	if err != nil {
		SendAndLog(c.bot, inputMessage.Chat.ID, err.Error())
		return
	}

	SendAndLog(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Stream %v has been successfully updated", streamID))
}
