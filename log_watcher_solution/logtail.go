package main

import (
	"bufio"
	"io"
	"os"
	"sync"
	"time"
)

type LogTail struct {
	filePath    string
	lastLines   []string
	mu          sync.RWMutex
	subscribers map[int]chan string
	subIDCount  int
}

func NewLogTail(filePath string) *LogTail {
	return &LogTail{
		filePath:    filePath,
		lastLines:   make([]string, 0),
		subscribers: make(map[int]chan string),
	}
}

// Adds a subscriber and returns a channel to send messages
func (l *LogTail) AddSubscriber() (int, chan string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	id := l.subIDCount
	ch := make(chan string, 100)
	l.subscribers[id] = ch
	l.subIDCount++
	return id, ch
}

// Removes subscriber
func (l *LogTail) RemoveSubscriber(id int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.subscribers, id)
}

// Broadcast line to all subscribers
func (l *LogTail) broadcast(line string) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	for _, ch := range l.subscribers {
		ch <- line
	}
}

// Keep tailing the log file
func (l *LogTail) Start() {
	go func() {
		file, err := os.Open(l.filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := bufio.NewReader(file)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					time.Sleep(500 * time.Millisecond)
					continue
				} else {
					panic(err)
				}
			}

			// Keep last 10 lines
			l.mu.Lock()
			if len(l.lastLines) >= 10 {
				l.lastLines = l.lastLines[1:]
			}
			l.lastLines = append(l.lastLines, line)
			l.mu.Unlock()

			l.broadcast(line)
		}
	}()
}

// Get last 10 lines for new client
func (l *LogTail) GetLastLines() []string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	lines := make([]string, len(l.lastLines))
	copy(lines, l.lastLines)
	return lines
}
