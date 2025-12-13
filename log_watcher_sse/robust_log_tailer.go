package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// RobustTailFile - handles log file rotation and other edge cases
func RobustTailFile(filePath string, broker *RobustBroker) {
	var file *os.File
	var reader *bufio.Reader
	var lastSize int64
	var lastModTime time.Time

	// Function to open file safely
	openFile := func() error {
		if file != nil {
			file.Close()
		}

		f, err := os.Open(filePath)
		if err != nil {
			return err
		}

		// Get file info for rotation detection
		info, err := f.Stat()
		if err != nil {
			f.Close()
			return err
		}

		file = f
		reader = bufio.NewReader(file)
		lastSize = info.Size()
		lastModTime = info.ModTime()

		// Start from end of file
		file.Seek(0, io.SeekEnd)

		return nil
	}

	// Initial file open
	if err := openFile(); err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}

	// Main tailing loop
	for {
		// Check for file rotation
		if err := rb.checkFileRotation(filePath, &lastSize, &lastModTime); err != nil {
			log.Printf("File rotation detected, reopening: %v", err)
			if err := openFile(); err != nil {
				log.Printf("Error reopening file: %v", err)
				time.Sleep(5 * time.Second) // Wait before retry
				continue
			}
		}

		// Read new lines
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// No new data, wait a bit
				time.Sleep(100 * time.Millisecond)
				continue
			}

			// Other error, might be file rotation
			log.Printf("Error reading file: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		// Broadcast the line
		broker.Broadcast(line)
	}
}

// checkFileRotation - detect if log file was rotated
func (rb *RobustBroker) checkFileRotation(filePath string, lastSize *int64, lastModTime *time.Time) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("file not found: %v", err)
	}

	// Check if file was truncated (size decreased)
	if info.Size() < *lastSize {
		return fmt.Errorf("file truncated (rotation detected)")
	}

	// Check if modification time changed significantly
	if info.ModTime().After(*lastModTime.Add(1 * time.Second)) {
		return fmt.Errorf("file modified (rotation detected)")
	}

	// Check if file was recreated (different inode)
	if info.Size() == 0 && *lastSize > 0 {
		return fmt.Errorf("file recreated (rotation detected)")
	}

	return nil
}

// TailFileWithRotation - alternative approach using file watching
func TailFileWithRotation(filePath string, broker *RobustBroker) {
	// This is a simplified version - in production you'd use fsnotify
	for {
		file, err := os.Open(filePath)
		if err != nil {
			log.Printf("Error opening file: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Get file size and seek to end
		info, err := file.Stat()
		if err != nil {
			file.Close()
			time.Sleep(1 * time.Second)
			continue
		}

		// Start from end
		file.Seek(0, io.SeekEnd)
		reader := bufio.NewReader(file)

		// Read until EOF or error
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					// Check if file was rotated by comparing size
					newInfo, err := os.Stat(filePath)
					if err != nil || newInfo.Size() < info.Size() {
						// File was rotated, break and reopen
						file.Close()
						break
					}
					time.Sleep(100 * time.Millisecond)
					continue
				}
				// Other error, close and retry
				file.Close()
				break
			}

			broker.Broadcast(line)
		}

		file.Close()
		time.Sleep(1 * time.Second) // Wait before reopening
	}
}
