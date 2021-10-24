package stream

import (
	"time"

	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

var streamList map[uint64]streaming.Stream = map[uint64]streaming.Stream{
	1: {Id: 1, Title: "Stream #1", Streamer: "Vasya", StreamTime: time.Now(), Description: "Awesome streaming"},
	2: {Id: 2, Title: "Stream #2", Streamer: "Petay", StreamTime: time.Now(), Description: "Awesome streaming"},
	3: {Id: 3, Title: "Stream #3", Streamer: "Valera", StreamTime: time.Now(), Description: "Awesome streaming"},
	4: {Id: 4, Title: "Stream #4", Streamer: "Johnny", StreamTime: time.Now(), Description: "Awesome streaming"},
	5: {Id: 5, Title: "Stream #5", Streamer: "Ololosha", StreamTime: time.Now(), Description: "Awesome streaming"},
	6: {Id: 6, Title: "Stream #6", Streamer: "Artyom", StreamTime: time.Now(), Description: "Awesome streaming"},
	7: {Id: 7, Title: "Stream #7", Streamer: "Ololoev", StreamTime: time.Now(), Description: "Awesome streaming"},
}
