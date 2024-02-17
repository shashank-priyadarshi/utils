package main

import (
	"strconv"
	"sync"

	gUUID "github.com/google/uuid"
	"github.com/pborman/uuid"
)

type PubSub interface {
	Publish(string, interface{}) (string, <-chan interface{})
	Subscribe(string) (string, <-chan interface{}, chan<- interface{})
	Unsubscribe(string, string)
}

/*
topics: map[topic][]subcriberChan stores list of all subsribers for a topic
status: map[topic]chan interface allows publisher and subscriber publish and receive status updates
messagesBeforeAnySubscriber: map[topic][]interface{} stores messages when no subscriber was available
messageCount, subscriberCount: int store count of all messages & subscribers of all topics for id
subscribers: map[subsciberID]subcriberChan stores list of all subscribers, helps when removing a subscriber channel at unsubscribe
*/
type pubsub struct {
	mu                            sync.Locker
	topics                        map[string][]chan interface{}
	subscribers                   map[string]chan interface{}
	status                        map[string]chan interface{}
	messagesBeforeAnySubscriber   map[string][]interface{}
	messageCount, subscriberCount int
}

func NewPubSub() PubSub {
	return &pubsub{
		mu:                          &sync.Mutex{},
		topics:                      make(map[string][]chan interface{}),
		subscribers:                 make(map[string]chan interface{}),
		status:                      make(map[string]chan interface{}),
		messagesBeforeAnySubscriber: make(map[string][]interface{}),
		messageCount:                0,
		subscriberCount:             0,
	}
}

func (p *pubsub) Publish(topic string, message interface{}) (messageID string, statusChan <-chan interface{}) {
	p.mu.Lock()

	p.exists(topic)
	p.messageCount += 1

	statusChan = p.status[topic]
	messageID = getID(topic, strconv.Itoa(p.messageCount))

	subscribers := p.topics[topic]

	if len(subscribers) == 0 {
		p.messagesBeforeAnySubscriber[topic] = append(p.messagesBeforeAnySubscriber[topic], message)
		return
	}

	p.publishMessagesBeforeAnySubscriber(topic)
	p.mu.Unlock()

	for _, subscriber := range subscribers {
		subscriber <- message
	}

	return
}

func (p *pubsub) Subscribe(topic string) (subscriberID string, messageChan <-chan interface{}, statusChan chan<- interface{}) {
	p.mu.Lock()
	p.exists(topic)
	p.subscriberCount += 1
	p.mu.Unlock()

	subscriberID = getID(topic, strconv.Itoa(p.subscriberCount))
	statusChan = p.status[topic]

	channel := make(chan interface{})
	messageChan = channel

	p.mu.Lock()
	p.topics[topic] = append(p.topics[topic], channel)
	p.subscribers[subscriberID] = channel // Used when unsubscribing=

	p.publishMessagesBeforeAnySubscriber(topic)
	p.mu.Unlock()

	return
}

func (p *pubsub) publishMessagesBeforeAnySubscriber(topic string) {

	subscribers := p.topics[topic]

	messagesBeforeAnySubscriber := p.messagesBeforeAnySubscriber[topic]
	p.messagesBeforeAnySubscriber[topic] = make([]interface{}, 0)

	/*
		BUG: If a new subscriber comes online when there already are subscribers, duplicate messages will be resent to older subscriber
		This bug will manifest if
	*/

	if len(messagesBeforeAnySubscriber) > 0 {
		for _, message := range messagesBeforeAnySubscriber {
			for _, subscriber := range subscribers {
				subscriber <- message
			}
		}
	}
}

func (p *pubsub) Unsubscribe(topic string, subscriberID string) {

	var subscribers []chan interface{}
	var ok bool

	p.mu.Lock()
	if subscribers, ok = p.topics[topic]; !ok {
		return
	}

	channel := p.subscribers[subscriberID]
	p.mu.Unlock()

	for i, subscriber := range subscribers {
		if subscriber == channel {
			subscribers = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}

	p.mu.Lock()
	p.topics[topic] = subscribers
	close(channel)
	delete(p.subscribers, subscriberID)
	p.mu.Unlock()
}

func (p *pubsub) exists(topic string) {

	if _, ok := p.topics[topic]; !ok {
		p.topics[topic] = make([]chan interface{}, 0)
	}

	if _, ok := p.status[topic]; !ok {
		p.status[topic] = make(chan interface{})
	}

	if _, ok := p.messagesBeforeAnySubscriber[topic]; !ok {
		p.messagesBeforeAnySubscriber[topic] = make([]interface{}, 0)
	}
}

func getID(items ...string) (id string) {
	var entropy string
	for _, item := range items {
		entropy += item
	}

	entropy += gUUID.New().String()

	id = uuid.NewSHA1(uuid.NameSpace_DNS, []byte(entropy)).String()
	return
}
