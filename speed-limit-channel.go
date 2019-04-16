package main

import (
	"sync/atomic"
	"time"
)

type SpeedLimitChannel struct {
	Limitation   int64
	Channel      chan string
	WriteHistory []time.Time
	Sample       int64
	Counter      int64
}

func SpeedLimitChannelOpen(limitation int64) (*SpeedLimitChannel, error) {
	ch := SpeedLimitChannel{}

	ch.Limitation = limitation
	ch.Channel = make(chan string, 10)
	ch.WriteHistory = make([]time.Time, 0)
	ch.Sample = 10
	ch.Counter = 0

	return &ch, nil
}

func (ch *SpeedLimitChannel) speed() int64 {
	if len(ch.Channel) < 10 {
		return 0
	}

	if int64(len(ch.WriteHistory)) < (ch.Limitation / ch.Sample) {
		return 0
	}

	duration := int64(time.Now().Sub(ch.WriteHistory[0]) / time.Microsecond)

	return int64(len(ch.WriteHistory)) * ch.Sample * int64(time.Millisecond) / duration
}

func (ch *SpeedLimitChannel) Read() string {
	return <-ch.Channel
}

// Restrict write speed only
func (ch *SpeedLimitChannel) Write(msg string) {
	for ch.speed() > ch.Limitation {
		time.Sleep(500 * time.Millisecond)
	}

	counter := atomic.AddInt64(&ch.Counter, 1)

	if (counter % ch.Sample) == 0 {
		// enqueue time
		if int64(len(ch.WriteHistory)) == ch.Sample {
			// pop first
			ch.WriteHistory = ch.WriteHistory[1:]
		}

		ch.WriteHistory = append(ch.WriteHistory, time.Now())
	}

	ch.Channel <- msg
}
