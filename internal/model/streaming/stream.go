package streaming

import (
	"fmt"
	"log"
	"time"
)

const defaultTimeFormat = "2006/01/02 15:04:05"

type Stream struct {
	Id          uint64    `json:"id,omitempty"`
	Title       string    `json:"title"`
	Streamer    string    `json:"streamer"`
	StreamTime  time.Time `json:"stream_time"`
	Description string    `json:"description"`
}

func (stream *Stream) PrintableTime() string {
	msk, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Println("MSK timezone loading error:", err.Error())
		return stream.StreamTime.Format(defaultTimeFormat)
	}

	return stream.StreamTime.In(msk).Format(defaultTimeFormat)
}

func (stream *Stream) String() string {
	return fmt.Sprintf("Streaming title: \"%v\" : Description: \"%v\"\nat %v by %v", stream.Title, stream.Description, stream.PrintableTime(), stream.Streamer)
}
