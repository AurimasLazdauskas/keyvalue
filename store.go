package main

import (
	"bufio"
	"bytes"
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

func (s *KeyValueStore) Set(k string, v string) {
	s.data[k] = v
}

func (s *KeyValueStore) Get(k string) string {
	return s.data[k]
}

func (s *KeyValueStore) Delete(k string) {
	delete(s.data, k)
}

func (s *KeyValueStore) ToString() string {
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
