package streaming

import (
	"fmt"
	"time"
)

type Stream struct {
	Id          uint64    `json:"id,omitempty"`
	Title       string    `json:"title"`
	Streamer    string    `json:"streamer"`
	StreamTime  time.Time `json:"stream_time"`
	Description string    `json:"description"`
}

func (stream *Stream) PrintableTime() string {
	return stream.StreamTime.UTC().Format("2006/01/02 15:04:05")
}

func (stream *Stream) String() string {
	return fmt.Sprintf("Streaming title: \"%v\" : Description: \"%v\"\nat %v by %v", stream.Title, stream.Description, stream.PrintableTime(), stream.Streamer)
}
