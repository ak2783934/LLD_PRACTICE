package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func TailFile(filePath string, broker *Broker) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.Seek(0, os.SEEK_END)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		broker.Messages <- line
	}
}
