package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

// SimpleTailFile - much simpler log tailer
func SimpleTailFile(filePath string, broker *SimpleBroker) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	// Start from end of file
	file.Seek(0, os.SEEK_END)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// No new data, wait a bit
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Broadcast the new log line to all clients
		broker.Broadcast(line)
	}
}
