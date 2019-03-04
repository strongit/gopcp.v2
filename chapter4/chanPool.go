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
    done := make(chan bool)  //创建一个无缓存的channel，由于该channel是阻塞的，在所有的数据未取出前，主程序就不退出
    // 子协程
    go func() {
        urls := make([]string, TASK_COUNT)
        for i := 0; i < TASK_COUNT; i++ {
            fmt.Sprintf("http://www.%d.com", i)
            urls[i] = fmt.Sprintf("http://www.%d.com", i)
            chReq <- urls[i]
            done <- true
            fmt.Println("new URL is %s", chReq)
        }
    }()
    for i := 0; i < GOROUTINE_COUNT; i++ {
        // time.Sleep(1)
        <-done
        go func() {
            url := <-chReq
            fmt.Println("new URL is %s", url)
            chRes <- 0
        }()
    }
    // 执行主进程
    // for i := 0; i < TASK_COUNT; i++ {
    //     d := <-chRes
    //     fmt.Sprintf("res is %d", d)
    // }
}
