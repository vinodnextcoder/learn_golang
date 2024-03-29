package main

import "fmt"
// Channels are the pipes that connect concurrent goroutines. 
// You can send values into channels from one goroutine 
// and receive those values into another goroutine.



func mainbuf1() {

    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
    
}


func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
    mainbuf1()
}