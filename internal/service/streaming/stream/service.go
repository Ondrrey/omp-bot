package stream

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ozonmp/omp-bot/internal/model/streaming"
)

type StreamServicer interface {
	Describe(streamID uint64) (*streaming.Stream, error)
	List(cursor uint64, limit uint64) ([]streaming.Stream, error)
	Create(streaming.Stream) (uint64, error)
	Update(streamID uint64, stream streaming.Stream) error
	Remove(streamID uint64) (bool, error)
	LastID() uint64
}

type StreamService struct {
	lastID uint64
	mu     sync.RWMutex
}

func NewStreamService() StreamServicer {
	last := len(streamList) //так как у нас dummy данные, то такая ересь сойдёт
	return &StreamService{
		lastID: uint64(last),
	}
}

func (s *StreamService) LastID() uint64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	last := s.lastID
	return last
}

func (s *StreamService) newId() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastID++
	newId := s.lastID
	return newId
}

func validateStreamId(streamID uint64) error {
	if streamID < 1 {
		return errors.New("incorrect stream Id")
	}
	return nil
}

func validateStream(stream streaming.Stream) error {
	if len(stream.Title) == 0 {
		return errors.New("title cannot be empty")
	}
	if len(stream.Streamer) == 0 {
		return errors.New("streamer is not specified")
	}
	return nil
}
func (s *StreamService) Describe(streamID uint64) (*streaming.Stream, error) {

	if err := validateStreamId(streamID); err != nil {
		return nil, err
	}

	if item, found := streamList[streamID]; found {
		return &item, nil
	}

	return nil, errors.New("stream not found")
}

func (s *StreamService) List(cursor uint64, limit uint64) ([]streaming.Stream, error) {

	streamsList := make([]streaming.Stream, 0, limit)

	lastId := s.LastID()
	if cursor > lastId {
		return nil, errors.New("index out of range")
	}

	var listed uint64
	for i := 0; listed < limit && cursor <= lastId; i++ {

		if item, found := streamList[cursor]; found {
			streamsList = append(streamsList, item)
			listed++
		}
		cursor++
	}
	return streamsList, nil
}

func (s *StreamService) Create(stream streaming.Stream) (uint64, error) {
	if err := validateStream(stream); err != nil {
		return 0, err
	}

	stream.Id = s.newId()
	streamList[stream.Id] = stream
	return stream.Id, nil
}

func (s *StreamService) Update(streamID uint64, stream streaming.Stream) error {
	if err := validateStreamId(streamID); err != nil {
		return err
	}
	if _, found := streamList[streamID]; !found {
		return fmt.Errorf("stream with id = %d not found", streamID)
	}
	if err := validateStream(stream); err != nil {
		return err
	}
	stream.Id = streamID
	streamList[streamID] = stream
	return nil
}

func (s *StreamService) Remove(streamID uint64) (bool, error) {
	if err := validateStreamId(streamID); err != nil {
		return false, err
	}
	if _, found := streamList[streamID]; !found {
		return false, fmt.Errorf("stream with id = %d not found", streamID)
	}
	delete(streamList, streamID)
	return true, nil
}
