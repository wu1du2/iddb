package main

import (
    "strconv"
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan string)
    go pump1(ch1)
    go pump2(ch2)
    go suck(ch1, ch2)
    time.Sleep(time.Duration(time.Second*30))
}

func pump1(ch chan int) {
    for i := 0; ; i++ {
        ch <- i * 2
        time.Sleep(time.Duration(time.Second))
    }
}

func pump2(ch chan string) {
    for i := 0; ; i++ {
        ch <- strconv.Itoa(i+5)
        time.Sleep(time.Duration(time.Second))
    }
}

func suck(ch1 chan int, ch2 chan string) {
    chRate := time.Tick(time.Duration(time.Second*5)) // 定时器
    for {
        select {
        case v := <-ch1:
            fmt.Printf("Received on channel 1: %d\n", v)
        case v := <-ch2:
            fmt.Printf("Received on channel 2: %s\n", v)
        case <-chRate:
            fmt.Printf("Log log...\n")
        }
    }
}
