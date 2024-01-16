package main

import "testing"

func TestKeyValueStoreGet(t *testing.T) {
	key := "one"
	value := "1"

	data := make(map[string]string)

	data[key] = value

	store := KeyValueStore{data: data}

	result := store.Get(key)

	if result != value {
		t.Errorf(key+" should output "+value+" but got: ", result)
	}
}
