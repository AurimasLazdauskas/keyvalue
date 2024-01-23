package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

type KeyValueStore struct {
	data map[string]string
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{data: make(map[string]string)}
}

func (s *KeyValueStore) Insert(k string, v string) error {
	_, exists := s.data[k]

	if exists {
		return errors.New("Key already exists")
	}

	s.data[k] = v

	return nil
}

func (s *KeyValueStore) Get(k string) string {
	return s.data[k]
}

func (s *KeyValueStore) Delete(k string) {
	delete(s.data, k)
}

func ToString(s *KeyValueStore) string {
	var b bytes.Buffer

	for key, value := range s.data {
		b.WriteString(key + ":" + value + "\n")
	}

	return b.String()
}

func (s *KeyValueStore) Persist() error {
	filePath := "keyvalue.db"

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(ToString(s))

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

		s.Insert(words[0], words[1])
	}

	return nil
}
