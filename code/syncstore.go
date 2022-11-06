package code

import (
	"syncserv/util"

	"github.com/google/uuid"
)

// Holds all the TypeSyncs
type SyncStore struct {
	Count int
	data  map[string]*TypeSync
}

// Global SyncStore instance
var SyncStoreInstance = SyncStore{0, make(map[string]*TypeSync)}

// Returns a unique id
func (s *SyncStore) uniqueId() string {
	id := util.RandString(5)

	_, ok := s.data[id]

	if ok {
		return s.uniqueId()
	}
	return id
}

// Creates a new TypeSync
func (s *SyncStore) CreateNew() *TypeSync {

	id := s.uniqueId()

	s.data[id] = &TypeSync{
		Id:     id,
		Secret: uuid.NewString(),
	}
	s.Count++
	return s.data[id]
}

// Returns a TypeSync by id
func (s *SyncStore) Get(id string) (*TypeSync, bool) {
	sharer, found := s.data[id]
	return sharer, found
}

// Deletes a TypeSync by id
func (s *SyncStore) Delete(id string) {
	delete(s.data, id)
	s.Count--
}
