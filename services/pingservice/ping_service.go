package pingservice

import (
	"errors"
)

type PingService interface {
	PingMessage() (string, error)
}

// hold repos and other dependencies
type pingService struct {
}

func NewPingService() PingService {
	return &pingService{}
}

func (ps *pingService) PingMessage() (string, error) {
	if 0 == 1 {
		return "", errors.New("some error message")
	}
	return "ping return message", nil
}
