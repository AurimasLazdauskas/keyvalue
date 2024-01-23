package main

import (
	"bufio"
	"fmt"
	"os"
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

	file, err := os.Open("keyvalue.db")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	dbFileLine := scanner.Text()

	expectedResult := key + ":" + value

	if dbFileLine != expectedResult {
		t.Errorf("Should output "+expectedResult+" but got: ", dbFileLine)
	}
}

func TestKeyValueStoreLoad(t *testing.T) {
	key := "one"
	value := "1"

	keyValue := key + ":" + value

	filePath := "keyvalue.db"

	file, err := os.Create(filePath)

	if err != nil {
		return
	}

	defer file.Close()

	_, err = file.WriteString(keyValue)

	store := NewKeyValueStore()

	store.Load()

	result := store.Get(key)

	if store.Get(key) != value {
		t.Errorf("expected "+value+" but got: ", result)
	}
}
