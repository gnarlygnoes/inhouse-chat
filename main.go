package main

import (
	"fmt"
	"time"
)

type Message struct {
	message  string
	date     int64
	userName string
}

func main() {
	msg := Message{
		message:  "Hallo Welt",
		date:     time.Now().Unix(),
		userName: "genericName01",
	}
	fmt.Println(msg)
}
