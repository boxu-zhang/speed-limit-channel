# speed-limit-channel
This is a small library for speed limit data transporting. Currently, it limits the write speed. 

## Example
```go
// 100 means the speed is 100 msg/s 
ch, _ := SpeedLimitChannelOpen(100)

// write msg to the channel
for i := 0; i < 1000000; i++ {
   ch.Write("any string message")
}

// read from it
go func() {
   for i := 0; i < 1000000; i++ {
      msg := ch.Read()

      // process the message
   }
}
```
