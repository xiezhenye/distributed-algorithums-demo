// logic-clock project main.go
package main

import (
	"fmt"
)


type Processor struct {
	channel   chan Msg
	timeSys   TimeSys
}

type Msg struct {
	time TimeMsg
	data interface{}
}

func (self *Processor) Send(msg interface{}) {
	self.channel <- msg
}

func NewProcessor(n int, timeSys TimeSys) Processor {
	var ret Processor
	ret.channel = make(chan Msg)
	ret.timeSys = timeSys
	TimeSys.Init(n)
	return ret
}


type TimeMsg interface {
	AddTo(t *Time)
}

type Time []int

const (
	TIME_LT = '<'
	TIME_GT = '>'
	TIME_EQ = '='
	TIME_ID = '|'
)

type TimeSys interface {
	Init(n int)
	TimeMsgTo(n int) TimeMsg
	Time() Time
}

type SimpleTimeSys struct {
	time Time
}

type SimpleTimeMsg Time

func (self *SimpleTimeSys) Init(n int) {
	self.time = new([]int, n)
}

func main() {
	fmt.Println("Hello World!")
}
