package main

import (
    "fmt"
    // "time"
)

func main() {
    const (
        GOROUTINE_COUNT = 30
        TASK_COUNT      = 30
    )
    chReq := make(chan string, GOROUTINE_COUNT)
    chRes := make(chan int, TASK_COUNT)
    // 子协程
    go func() {
        urls := make([]string, TASK_COUNT)
        for i := 0; i < TASK_COUNT; i++ {
            fmt.Sprintf("http://www.%d.com", i)
            urls[i] = fmt.Sprintf("http://www.%d.com", i)
            chReq <- urls[i]
            fmt.Println("new URL is %s", chReq)
        }
    }()
    for i := 0; i < GOROUTINE_COUNT; i++ {
        // time.Sleep(1)
        go func() {
            url := <-chReq
            fmt.Println("new URL is %s", url)
            chRes <- 0
        }()
    }
    // 先执行主进程
    for i := 0; i < TASK_COUNT; i++ {
        d := <-chRes
        fmt.Sprintf("res is %d", d)
    }
}
