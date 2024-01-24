package main

import (
	"testing"
)

func TestKeyValueStoreInsertAndGet(t *testing.T) {
	key := "one"
	value := "1"

	store := NewKeyValueStore()

	store.Insert(key, value)

	result := store.Get(key)

	if result != value {
		t.Errorf(key+" should output "+value+" but got: ", result)
	}
}

func TestKeyValueStorePersistAndLoad(t *testing.T) {
	key := "one"
	value := "1"

	store := NewKeyValueStore()

	store.Insert(key, value)

	store.Persist()

	newStore := NewKeyValueStore()

	newStore.Load()

	result := newStore.Get(key)

	if result != value {
		t.Errorf("expected "+value+" but got: ", result)
	}
}
