package main

import (
	"sync"
)

type KeyValueStore struct {
	data sync.Map
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{data: sync.Map{}}
}

func (s *KeyValueStore) Set(k string, v string) {
	s.data.Store(k, v)
}

func (s *KeyValueStore) Get(k string) string {
	value, _ := s.data.Load(k)

	if value == nil {
		return ""
	}

	return value.(string)
}

func (s *KeyValueStore) Delete(k string) {
	s.data.Delete(k)
}
