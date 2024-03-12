package main

import (
    "log"
    "sync"
    "testing"
    "time"
)

func TestTask(t *testing.T) {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        task()
        wg.Done()
    }()
    wg.Wait()
    log.Print("main done")
}

func TestTask1(t *testing.T) {
    ch := make(chan struct{}) // 初始化 chan
    go func() {
        task()
        ch <- struct{}{} // 发送到 chan
    }()
    <-ch // 从 chan 获取
    log.Print("main done")
}

func task() {
    // 模拟耗时任务
    time.Sleep(time.Second)
    log.Print("task done")
}

func TestPanic(t *testing.T) {
    // 没有这段,panic会异常退出
    defer func() {
        if r := recover(); r != nil {
            t.Log(r)
        }
    }()

    // panic只对当前goroutine的defer有效!
    //go func() {
        panic("异常了")
    //}()
    log.Print("end")
}
