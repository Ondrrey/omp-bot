package stream

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *StreamCommand) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("StreamCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	streams, err := c.streamService.List(parsedData.Cursor, parsedData.Limit)
	if err != nil {
		SendAndLog(c.bot, callback.Message.Chat.ID, err.Error())
		return
	}

	if len(streams) == 0 {
		SendAndLog(c.bot, callback.Message.Chat.ID, "There are no streams")
		return
	}

	var outputMsgText string
	for _, s := range streams {
		outputMsgText += s.String() + "\n\n"
	}

	//кнопки
	last := streams[len(streams)-1].Id
	first := streams[0].Id
	buttons := make([]tgbotapi.InlineKeyboardButton, 0, 2)

	if parsedData.Cursor > 0 && parsedData.Limit < first {
		start := first - parsedData.Limit
		button := CreateButton(start, parsedData.Limit, "Prev page")
		buttons = append(buttons, *button)
	}

	if last < c.streamService.LastID() {
		start := last + 1
		button := CreateButton(start, parsedData.Limit, "Next page")
		buttons = append(buttons, *button)
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(buttons...),
		)
	}
	c.bot.Send(msg)

}

func CreateButton(newCursor uint64, limit uint64, name string) *tgbotapi.InlineKeyboardButton {
	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: newCursor,
		Limit:  limit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "streaming",
		Subdomain:    "stream",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}
	button := tgbotapi.NewInlineKeyboardButtonData(name, callbackPath.String())
	return &button
}
