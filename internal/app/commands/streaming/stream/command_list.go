package stream

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	Cursor uint64 = 1
	Limit  uint64 = 2
)

func (c *StreamCommand) List(inputMessage *tgbotapi.Message) {

	streams, err := c.streamService.List(Cursor, Limit)
	if err != nil {
		SendAndLog(c.bot, inputMessage.Chat.ID, err.Error())
		return
	}

	if len(streams) == 0 {
		SendAndLog(c.bot, inputMessage.Chat.ID, "There are no streams")
		return
	}

	var outputMsgText string
	for _, s := range streams {
		outputMsgText += s.String() + "\n\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	last := streams[len(streams)-1].Id
	if last < c.streamService.LastID() { //есть ещё данные - отображаем кнопку
		serializedData, _ := json.Marshal(CallbackListData{
			Cursor: last + 1,
			Limit:  Limit,
		})

		callbackPath := path.CallbackPath{
			Domain:       "streaming",
			Subdomain:    "stream",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}
	c.bot.Send(msg)

}
