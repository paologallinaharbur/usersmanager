package messagingSystem

import (
	"github.com/sirupsen/logrus"
	"net/url"
	"time"
)

//StubMessageQueue stub implementation of a messaging system
type StubMessageQueue struct {
	Endpoint        url.URL
	Interval        time.Duration
	NumberofRetries int64
	Timeout         time.Duration
	// A channel is useful in order to avoid slowing down api calls to send an internal notification
	queue chan interface{}
}

//StartRoutine starts the goroutine that will periodically grab one message and send it to an other service
func (sm *StubMessageQueue) StartRoutine() {
	sm.queue = make(chan interface{}, 100)
	go sm.readChannelSendMessageWithRetry()
}

//AddMessageToQueue add the message to a queue that will be read by a different goroutine, In this way the process is decoupled
func (sm *StubMessageQueue) AddMessageToQueue(m interface{}) {
	sm.queue <- m
}

//readChannelSendMessageWithRetry grab each tick a message
func (sm *StubMessageQueue) readChannelSendMessageWithRetry() {
	t := time.NewTicker(sm.Interval)
	for {
		<-t.C // This is to avoid too many messages sent at the same time overloading the backend
		m := <-sm.queue
		logrus.Infof("Sending the following message '%v' to %s having %s timeout with %d NumberofRetries", m, sm.Endpoint.Path, sm.Timeout.String(), sm.NumberofRetries)
		time.Sleep(5 * time.Second) //simulating time needed to send the message
	}
}
