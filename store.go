package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
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

func (s *KeyValueStore) ToString() string {
	var b bytes.Buffer

	s.data.Range(func(key, value interface{}) bool {
		b.WriteString(key.(string) + ":" + value.(string) + "\n")
		return true
	})

	return b.String()
}

func (s *KeyValueStore) Persist() error {
	filePath := "keyvalue.db"

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(s.ToString())

	return err
}

func (s *KeyValueStore) Load() error {
	file, err := os.Open("keyvalue.db")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dbFileLine := scanner.Text()

		words := strings.Split(dbFileLine, ":")

		s.Set(words[0], words[1])
	}

	return nil
}
