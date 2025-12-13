package main

import (
	"fmt"
	"sync"
)

type Subscriber struct {
	subscriberId int
}

func (s *Subscriber) consume(message string) {
	fmt.Printf("message consumed by %v, message %v\n", s.subscriberId, message)
}

// should be make different consume functions for each topics here?

type Topic struct {
	name        string
	subscribers map[int]*Subscriber // subscriber-id->subscriber
}

// I don't think there will be any struct of publisher right?
// it should be just a function part of our pb sub system.

type PubSubManager struct {
	topics      map[string]*Topic   // topic name -> Topic
	subscribers map[int]*Subscriber // subscriber-id -> Subscriber
	subsCount   int
	mu          sync.RWMutex
}

var (
	instance *PubSubManager
	once     sync.Once
)

func CreatePubSubManager() *PubSubManager {
	once.Do(func() {
		instance = &PubSubManager{
			topics:      make(map[string]*Topic),
			subscribers: make(map[int]*Subscriber),
			subsCount:   0,
		}
	})
	return instance
}

func (p *PubSubManager) CreateTopic(topicName string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, ok := p.topics[topicName]
	if ok {
		fmt.Println("topic already exists")
		return
	}
	p.topics[topicName] = &Topic{
		name:        topicName,
		subscribers: make(map[int]*Subscriber),
	}
}

func (p *PubSubManager) CreateSubscriber() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	subscriber := &Subscriber{
		subscriberId: p.subsCount,
	}
	p.subscribers[p.subsCount] = subscriber
	p.subsCount++
	fmt.Println("subscriber created: ", subscriber.subscriberId)
	return subscriber.subscriberId
}

func (p *PubSubManager) Subscribe(subsId int, topicName string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	subscriber, ok := p.subscribers[subsId]
	if !ok {
		fmt.Println("subscriber doesn't exist id", subsId)
		return
	}

	topic, isTopicExist := p.topics[topicName]
	if !isTopicExist {
		fmt.Println("topic name does not exist")
		return
	}

	_, isAlreadySubs := topic.subscribers[subsId]
	if isAlreadySubs {
		fmt.Printf("topic %v is already subscriber by %v\n", topicName, subsId)
		return
	}

	topic.subscribers[subsId] = subscriber
}

func (p *PubSubManager) Unsubscribe(subsId int, topicName string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, ok := p.subscribers[subsId]
	if !ok {
		fmt.Println("subscriber doesn't exist")
		return
	}

	topic, isTopic := p.topics[topicName]
	if !isTopic {
		fmt.Println("topic name does not exist")
		return
	}

	// just delete it directly even if its not subscribed
	delete(topic.subscribers, subsId)
}

func (p *PubSubManager) Publish(topicName string, message string) {
	topic, isTopic := p.topics[topicName]
	if !isTopic {
		fmt.Println("topic name does not exist")
		return
	}

	for _, subs := range topic.subscribers {
		subs.consume(message)
	}
}

func main() {
	pubSubSystem := CreatePubSubManager()
	pubSubSystem.CreateTopic("topic1")
	subscriberID1 := pubSubSystem.CreateSubscriber()
	subscriberID2 := pubSubSystem.CreateSubscriber()
	pubSubSystem.Subscribe(subscriberID1, "topic1")
	pubSubSystem.Subscribe(subscriberID2, "topic1")
	pubSubSystem.Publish("topic1", "test message")
}
