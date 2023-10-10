/* The main function simulates a ping-pong game between N PingAgents and PongAgent.
 */

package main

import (
	"fmt"
	"time"
)

type Agent interface {
	Start()
}

const PingString = "ping"
const PongString = "pong"

type Request struct {
	sendId string
	req    string
	c      chan string
}

type PingAgent struct {
	Id   string
	cin  chan string
	cout chan Request
}

type PongAgent struct {
	Id  string
	cin chan Request
}

func NewPingAgent(id string, cout chan Request) *PingAgent {
	cin := make(chan string)
	return &PingAgent{Id: id, cin: cin, cout: cout}
}

func NewPongAgent(id string, cin chan Request) *PongAgent {
	return &PongAgent{Id: id, cin: cin}
}

func (ag *PingAgent) Start() {
	go func() {
		for {
			ag.cout <- Request{ag.Id, PingString, ag.cin} // Send request to PongAgent
			// wait for response
			answer := <-ag.cin
			fmt.Printf("%s received %s\n", ag.Id, answer)
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func (ag *PongAgent) Start() {
	// Receive requests from PingAgents, and then handle them
	go func() {
		for {
			req := <-ag.cin // Receive request from PingAgent
			fmt.Printf("%s received %s from %s\n", ag.Id, req.req, req.sendId)
			go ag.handleAgent(req)
		}
	}()
}

func (ag *PongAgent) handleAgent(req Request) {
	req.c <- PongString
}

func main() {
	c := make(chan Request)

	// Start Ponger
	ponger := NewPongAgent("Ponger", c)
	ponger.Start()

	// Start PingAgents
	for i := 0; i < 3; i++ {
		id := fmt.Sprintf("Pinger %d", i)
		pinger := NewPingAgent(id, c)
		pinger.Start()
	}

	// Run for a while
	time.Sleep(5 * time.Second)
}
