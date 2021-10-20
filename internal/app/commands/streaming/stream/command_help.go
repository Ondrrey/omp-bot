package stream

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *StreamCommand) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__streaming__stream — print list of commands\n"+
			"/get__streaming__stream — get a entity\n"+
			"/list__streaming__stream — get a list of your entity\n"+
			"/delete__streaming__stream — delete an existing entity\n"+
			"/new__streaming__stream — create a new entity\n"+
			"/edit__streaming__stream — edit a entity",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamCommander.Help: error sending reply message to chat - %v", err)
	}
}
