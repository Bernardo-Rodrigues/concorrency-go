package publish_subscribe

import (
	"fmt"
	"sync"
	"time"
)

type PubSub struct {
	mu       sync.Mutex
	channels map[string][]chan string
}

func NewPubSub() *PubSub {
	return &PubSub{
		channels: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ch := make(chan string)
	ps.mu.Lock()
	ps.channels[topic] = append(ps.channels[topic], ch) // add subscriber to topic
	ps.mu.Unlock()
	return ch
}

func (ps *PubSub) Publish(topic, msg string) {
	ps.mu.Lock()
	for _, ch := range ps.channels[topic] { // send message to all subscribers
		ch <- msg
	}
	ps.mu.Unlock()
}

func (ps *PubSub) Close(topic string) {
	ps.mu.Lock()
	for _, ch := range ps.channels[topic] { // close all subscriber channels
		close(ch)
	}
	ps.mu.Unlock()
}

func Execute() {
	// create broker
	ps := NewPubSub()

	// subscribe to topic (returns message channel)
	subscriber1 := ps.Subscribe("news")
	subscriber2 := ps.Subscribe("news")

	var wg sync.WaitGroup
	wg.Add(2)

	// read until channel is closed
	go func() {
		defer wg.Done()
		for msg := range subscriber1 {
			fmt.Println("Subscriber 1 received:", msg)
		}
	}()

	// read until channel is closed
	go func() {
		defer wg.Done()
		for msg := range subscriber2 {
			fmt.Println("Subscriber 2 received:", msg)
		}
	}()

	// publish messages to topic
	ps.Publish("news", "Breaking News!")
	ps.Publish("news", "Another News!")

	time.Sleep(time.Second)
	ps.Close("news") // close all subscriber channels
	wg.Wait()
}
