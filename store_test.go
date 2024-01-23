package main

import "testing"

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
