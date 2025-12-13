package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

// SimpleRobustTailFile - simple log tailer with basic rotation handling
func SimpleRobustTailFile(filePath string, broker *SimpleRobustBroker) {
	for {
		file, err := os.Open(filePath)
		if err != nil {
			log.Printf("Error opening file: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Start from end of file
		file.Seek(0, os.SEEK_END)
		reader := bufio.NewReader(file)

		// Read lines until error
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					// No new data, wait a bit
					time.Sleep(100 * time.Millisecond)
					continue
				}

				// Error reading file, might be rotation
				file.Close()
				break
			}

			// Broadcast the line
			broker.Broadcast(line)
		}

		file.Close()
		time.Sleep(1 * time.Second) // Wait before reopening
	}
}
