package main

import (
    "fmt"
    "testing"
    "time"
)

func Producer(channel chan<- int) {
    for i := 0; i < 10; i++ {
        channel <- i //写入
        fmt.Println("发送:", i)
        time.Sleep(time.Second)
    }
}

func Consumer(channel <-chan int) {
    for i := 0; i < 10; i++ {
        v := <-channel // 读出
        fmt.Println("接收:", v)
        time.Sleep(time.Second)
    }
}

func Test_channel(t *testing.T) {
    channel := make(chan int, 88)
    go Producer(channel)
    go Consumer(channel)
    time.Sleep(12 * time.Second)
}
