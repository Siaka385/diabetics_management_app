package services

import (
	"sync"
	"time"
)

type SSEvent struct {
	Time string	`json:"time"`
	Data string `json:"data"`
}

func Register(client chan SSEvent) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	clients[client] = true
}

var (
	clients      map[chan SSEvent]bool // registered clients channel
	clientsMutex *sync.Mutex                 // protect the access to the clients channel map
)

func Init() map[chan SSEvent]bool {
	clients = make(map[chan SSEvent]bool)
	clientsMutex = &sync.Mutex{}
	return clients
}

func Broadcast(message string) {
	// create an event
	event := SSEvent{
		Time: time.Now().Format(time.RFC1123),
		Data: message,
	}

	// Send the event to all connected clients
	for clientChan := range clients {
		clientChan <- event
	}
}
