package clients

import (
	"sync"
	"syncserv/util"

	"github.com/google/uuid"
)

// Holds all the TypeSyncs
type ClientsIndex struct {
	Count int
	data  map[string]*Broadcaster
	Lock  *sync.Mutex
}

// Global SyncStore instance
var ClientIndexInstance = ClientsIndex{0, make(map[string]*Broadcaster), &sync.Mutex{}}

// Returns a unique id
func (s *ClientsIndex) uniqueId() string {
	id := util.RandString(5)

	_, ok := s.data[id]

	if ok {
		return s.uniqueId()
	}
	return id
}

// Creates a new TypeSync
func (s *ClientsIndex) CreateNew() *Broadcaster {

	id := s.uniqueId()

	s.Lock.Lock()
	s.data[id] = &Broadcaster{
		Id:        id,
		Secret:    uuid.NewString(),
		Lock:      &sync.Mutex{},
		Listeners: map[int]Listener{},
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
