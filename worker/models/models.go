package models

type ID string

type Message struct {
	Message interface{}
}

type Publisher struct {
	Topic   string
	Message Message
}

type Subscriber struct {
	MessageID ID
	Message   Message
}

type Status struct {
	SubscriberID, MessageID ID
	Status                  interface{}
	Error                   error
}
