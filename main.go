package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	store := NewKeyValueStore()

	store.Set("Rex", "dog")
	store.Set("Snowball", "cat")
	store.Set("Stuart", "mouse")
	store.Set("Stuart", "mouse")

	fmt.Println(store)

	return

	exit := false
	for exit == false {
		reader := bufio.NewReader(os.Stdin)

		raw_line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		line := strings.TrimRight(raw_line, "\n")

		command := strings.Fields(line)[0]

		switch command {
		case "exit":
			fmt.Println("Exiting...")
			exit = true
		case "insert":
			k := strings.Fields(line)[1]
			v := strings.Fields(line)[2]

			store.Set(k, v)
		case "get":
			k := strings.Fields(line)[1]

			value := store.Get(k)

			fmt.Println(value)

		case "delete":
			k := strings.Fields(line)[1]

			store.Delete(k)
		}

	}
}
