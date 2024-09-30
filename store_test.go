package main

import (
	"testing"
)

func TestKeyValueStoreSetAndGet(t *testing.T) {
	key := "one"
	value := "1"

	store := NewKeyValueStore()

	store.Set(key, value)

	result := store.Get(key)

	if result != value {
		t.Errorf(key+" should output "+value+" but got: ", result)
	}
}

func TestKeyValueStoreSetExistingKey(t *testing.T) {
	key := "one"
	value := "1"

	store := NewKeyValueStore()

	store.Set(key, value)

	store.Set(key, "2")

	result := store.Get(key)

	if result != "2" {
		t.Errorf(key+" should output 2 but got: ", result)
	}
}

func TestKeyValueStoreFetchNonExistingKey(t *testing.T) {
	key := "one"

	store := NewKeyValueStore()

	result := store.Get(key)

	if result != "" {
		t.Errorf(key+" should output empty but got: ", result)
	}
}

func TestKeyValueStoreDelete(t *testing.T) {
	key := "one"
	value := "1"

	store := NewKeyValueStore()

	store.Set(key, value)

	store.Delete(key)

	result := store.Get(key)

	if result != "" {
		t.Errorf(key+" should output empty but got: ", result)
	}
}

func TestNewKeyValueStoreDeleteNonExistingKey(t *testing.T) {
	key := "one"

	store := NewKeyValueStore()

	store.Delete(key)

	result := store.Get(key)

	if result != "" {
		t.Errorf(key+" should output empty but got: ", result)
	}

}

func TestKeyValueStoreSettingKeyToEmptyString(t *testing.T) {
	key := "one"
	value := ""

	store := NewKeyValueStore()

	store.Set(key, value)

	result := store.Get(key)

	if result != value {
		t.Errorf(key+" should output empty but got: ", result)
	}
}

func TestNewKeyValueStoreUsingEmptyStringAsKey(t *testing.T) {
	key := ""
	value := "1"

	store := NewKeyValueStore()

	store.Set(key, value)

	result := store.Get(key)

	if result != value {
		t.Errorf(key+" should output "+value+" but got: ", result)
	}
}

func TestKeyValueStorePersistAndLoad(t *testing.T) {
	key := "one"
	value := "1"

	store := NewKeyValueStore()

	store.Set(key, value)

	store.Persist()

	newStore := NewKeyValueStore()

	newStore.Load()

	result := newStore.Get(key)

	if result != value {
		t.Errorf("expected "+value+" but got: ", result)
	}
}
