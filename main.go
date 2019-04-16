package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("test SpeedLimitChannel")

	ch, _ := SpeedLimitChannelOpen(100)

	// read proc
	go func() {
		begin := time.Now()

		for i := 0; i < 1000000; i++ {
			ch.Read()

			if i == 0 {
				begin = time.Now()
			}

			if (i % 100) == 0 {
				duration := time.Now().Sub(begin).Nanoseconds() / int64(time.Millisecond)

				if duration != 0 {
					fmt.Printf("read speed is %d / %d = %d\n", int64(i), int64(duration), int64(i)*1000/int64(duration))
				}
			}
		}
	}()

	// write proc
	for i := 0; i < 1000000; i++ {
		ch.Write("abcdefghijklmn")
	}
}
