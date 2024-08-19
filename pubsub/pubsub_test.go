package pubsub

import (
	"testing"
)

type test struct {
	name     string
	testcase func(*testing.T)
}

//func Test_PubSub(t *testing.T) {
//	newPubSub := pubsub.NewPubSub()
//
//	tests := []test{
//		{
//			name: "Happy Path I: Publisher",
//			testcase: func(t *testing.T) {
//				// t.Parallel()
//				msgID, status := newPubSub.Publish("test", "hello")
//				t.Log("Message ID for first published message: ", msgID)
//				for msg := range status {
//					t.Log("Status received from subscriber: ", msg)
//					break
//				}
//			},
//		},
//		{
//			name: "Happy Path II: Subscriber",
//			testcase: func(t *testing.T) {
//				// t.Parallel()
//				subID, msgChan, statusChan := newPubSub.Subscribe("test")
//				t.Log("Subcriber ID received for first topic: ", subID)
//				for msg := range msgChan {
//					t.Log("Message received by subscriber for first topic: ", msg)
//					statusChan <- "message received"
//					break
//				}
//			},
//		},
//	}
//
//	for num, test := range tests {
//		t.Logf("Test number: %d \n Test name: %s", num, test.name)
//		t.Run(test.name, test.testcase)
//	}
//}
