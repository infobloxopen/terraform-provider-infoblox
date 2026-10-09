package utils

import (
	"sync"
)

// GlobalMutexStore is a global mutex store.
var GlobalMutexStore = newMutexStore()

// mutexStore is a store for mutexes.
type mutexStore struct {
	mu      sync.Mutex
	mutexes map[string]*sync.Mutex
}

// newMutexStore creates a new mutexStore.
// It is used to store mutexes to serialize operations with the same key,
// e.g. for next-available IP address allocation.
func newMutexStore() *mutexStore {
	return &mutexStore{
		mutexes: make(map[string]*sync.Mutex),
	}
}

// get returns a mutex for the given key.
func (s *mutexStore) get(key string) *sync.Mutex {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.mutexes[key]; !ok {
		s.mutexes[key] = &sync.Mutex{}
	}

	return s.mutexes[key]
}

// Lock locks the mutex for the given key.
func (s *mutexStore) Lock(key string) {
	s.get(key).Lock()
}

// Unlock unlocks the mutex for the given key.
func (s *mutexStore) Unlock(key string) {
	s.get(key).Unlock()
}
