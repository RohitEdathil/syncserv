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
}

// Global SyncStore instance
var SyncStoreInstance = ClientsIndex{0, make(map[string]*Broadcaster)}

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

	s.data[id] = &Broadcaster{
		Id:        id,
		Secret:    uuid.NewString(),
		Lock:      &sync.Mutex{},
		Listeners: map[int]Listener{},
	}
	s.Count++
	return s.data[id]
}

// Returns a TypeSync by id
func (s *ClientsIndex) Get(id string) (*Broadcaster, bool) {
	sharer, found := s.data[id]
	return sharer, found
}

// Deletes a TypeSync by id
func (s *ClientsIndex) Delete(id string) {
	delete(s.data, id)
	s.Count--
}
