📩 Pub/Sub System Design – Problem Statement
🎯 Objective

Design a scalable and efficient Publish-Subscribe (Pub/Sub) messaging system that enables asynchronous communication between publishers and subscribers through topics.

🔧 Core Functional Requirements
1. Topics

The system should support creation of multiple topics.

Each topic can have one or more subscribers.

2. Publishers

A publisher can publish a message to a specific topic.

Once a message is published, it should be delivered to all active subscribers of that topic.

3. Subscribers

Users (or services) can subscribe/unsubscribe from any topic.

Subscribers should receive all messages published after they have subscribed.

Messages should be delivered in the order they were published.

4. Message Delivery Guarantees (Choose at least one)

At-most-once: Best effort, might lose messages.

At-least-once: Messages are delivered one or more times (duplicates possible).

(Optional Advanced) Exactly-once: No duplication and no loss.

📈 Non-Functional Requirements

High throughput and low latency message delivery.

Horizontal scalability (ability to scale publishers and subscribers independently).

Fault tolerance (system should handle process/node failures).

Thread-Safety (system must handle concurrent publishers & subscribers safely).

🚀 Optional Advanced Features

Message persistence to disk (so subscribers can receive past messages).

Support consumer groups (competing consumers).

Provide push-based, pull-based, or hybrid delivery.

Support message replay.

Implement backpressure handling for slow subscribers.

✅ What You Need to Design

Entities and Interfaces (Publisher, Subscriber, Topic, MessageBroker, etc.)

Core flows:

- publish(topic, message)
- subscribe(topic, subscriber)
- unsubscribe(topic, subscriber)




Are we going to design things like that of kafka? 
- one publisher
- many consumers
- each will have its own offset as well? 
- each topic will have stream of message. 
- these messages in the topic can be restreamed as well?
- Is the subscriber supposed to poll the messages or publishers are supposed to push them

asumption: 
- not storing the data into disk as of now. 
- 



New scope: 

One publisher
multiple subscriber

instant message tranfer, no queues.

multiple topics can be there for now. 


Now core entities are like this: 
- Topic 
- Producer
- Subscriber
- PubSubManager


methods like 
publish(topicName, message)
subscribe(userID, topicName) ---
unsubscribe(userID, topicName) --
createTopic(topicName) --

Assumptions, all consumers are already active and when the messages are published, they are able to read them instantly. 

