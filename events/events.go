// Package events handles sending notifications to users and websockets
// when certain events happen. It also handles firing webhooks on workflow
// transitions.
package events

import (
	"github.com/gorilla/websocket"

	"github.com/praelatus/praelatus/models"
)

var Evm = New()

type EventManager struct {
	Event    chan models.Event
	ActiveWS []WSManager
}

type WSManager struct {
	Socket *websocket.Conn
	InBuf  [256]byte
	OutBuf [256]byte
}

func New() *EventManager {
	return &EventManager{
		Event:    make(chan models.Event),
		ActiveWS: make([]WSManager, 0),
	}
}
