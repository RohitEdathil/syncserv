package clients

import (
	"log"
	"sync"
	"syncserv/util"
	"time"

	"github.com/google/uuid"
)

// Holds all the TypeSyncs
type ClientsIndex struct {
	Count int
	data  map[string]*Broadcaster
	Lock  *sync.Mutex
}

// Global ClientsIndex instance
var ClientIndexInstance = ClientsIndex{0, make(map[string]*Broadcaster), &sync.Mutex{}}

// Returns a unique id
func (s *ClientsIndex) uniqueId() string {
	id := util.RandString(5)

	s.Lock.Lock()
	_, ok := s.data[id]
	s.Lock.Unlock()

	if ok {
		return s.uniqueId()
	}
	return id
}

// Creates a new TypeSync
func (s *ClientsIndex) CreateNew() *Broadcaster {

	// Obtain a unique id
	id := s.uniqueId()

	// Creating new Broadcaster
	s.Lock.Lock()
	s.data[id] = &Broadcaster{
		Id:        id,
		Secret:    uuid.NewString(),
		Text:      "",
		Lock:      &sync.Mutex{},
		Listeners: map[int]Listener{},
		LastSeen:  time.Now(),
	}
	s.Count++
	s.Lock.Unlock()

	return s.data[id]
}

// Checks if a TypeSync exists
func (s *ClientsIndex) CheckId(id string) bool {
	s.Lock.Lock()
	_, ok := s.data[id]
	s.Lock.Unlock()
	return ok
}

// Returns a TypeSync by id
func (s *ClientsIndex) Get(id string) (*Broadcaster, bool) {
	s.Lock.Lock()
	sharer, found := s.data[id]
	s.Lock.Unlock()
	return sharer, found
}

// Deletes a TypeSync by id
func (s *ClientsIndex) Delete(id string) {
	s.Lock.Lock()
	delete(s.data, id)
	s.Count--
	s.Lock.Unlock()
}

// Purges all the TypeSyncs that have not been seen in a while
func (s *ClientsIndex) Purge(while time.Duration) {
	s.Lock.Lock()
	for id, sharer := range s.data {

		sharer.Lock.Lock()
		// Don't purge if it's been seen within the while duration
		if !(time.Since(sharer.LastSeen) > while) {
			sharer.Lock.Unlock()
			continue
		}

		// Don't purge if there are still listeners
		if !(len(sharer.Listeners) == 0) {
			sharer.Lock.Unlock()
			continue
		}

		// Don't purge if it's still connected
		if sharer.LastSeen == (time.Time{}) {
			sharer.Lock.Unlock()
			continue
		}
		sharer.Lock.Unlock()

		// Purge
		log.Printf("Purged %s", id)
		delete(s.data, id)
		s.Count--
	}
	s.Lock.Unlock()
}
