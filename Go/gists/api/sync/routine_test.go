package main

import (
    "fmt"
    "log"
    "luvx/common"
    "sync"
    "testing"
    "time"
)

func Test_return(t *testing.T) {
    N := 2

    wg := sync.WaitGroup{}
    // 定义⼀个⽤于存放结果的channel
    resultCh := make(chan int, N)

    for i := 0; i < N; i++ {
        wg.Add(1)
        // 启动协程计算结果，并将结果发送到channel中
        go func(num int) {
            time.Sleep(time.Second)
            resultCh <- num
            wg.Done()
        }(i)
    }

    // 等待所有协程执⾏完毕
    go func() {
        fmt.Println("point1")
        wg.Wait()
        // 关闭channel
        close(resultCh)
        fmt.Println("point2")
    }()
    fmt.Println("point3")

    // 从channel中读取所有结果，并将结果累加起来
    sum := 0
    for res := range resultCh {
        fmt.Println(res)
        sum += res
    }
    fmt.Println("和:", sum)
}

func Test_02(t *testing.T) {
    f1 := func() string {
        log.Print("task start")
        time.Sleep(time.Second * 2)
        log.Print("task done")
        return "ok"
    }
    f2 := func() string {
        log.Print("task start")
        time.Sleep(time.Second * 3)
        log.Print("task done")
        return "ok"
    }

    start := time.Now()
    fmt.Println(start)

    wg := sync.WaitGroup{}

    r1 := make(chan string)
    r2 := make(chan string)

    common.RunInRoutine(&wg, func() {
        r1 <- f1()
    })
    common.RunInRoutine(&wg, func() {
        r2 <- f2()
    })

    //直接使用会出现死锁或者直接不使用WaitGroup
    //wg.Wait()
    go func() {
        wg.Wait()
    }()

    s1 := <-r1
    s2 := <-r2
    fmt.Println(s1, s2)
    fmt.Println(time.Now().Sub(start))
}
