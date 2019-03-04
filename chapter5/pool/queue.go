package main

import (
    "fmt"
    "sync"
)

type Pool struct {
    queue chan int
    wg    *sync.WaitGroup
}

func NewPool(cap, total int) *Pool {
    if cap < 1 {
        cap = 1
    }
    p := &Pool{
        queue: make(chan int, cap),
        wg:    new(sync.WaitGroup),
    }
    p.wg.Add(total)
    return p
}

func (p *Pool) AddOne() {
    p.queue <- 1
}

func (p *Pool) DelOne() {
    <-p.queue
    p.wg.Done()
}

func main() {
    urls := []string{"a", "b", "c", "d", "e", "f"}
    pool := NewPool(6, len(urls))

    for _, v := range urls {
        go func(url string) {
            pool.AddOne()
            err := Download(url)
            if nil != err {
                fmt.Println(err)
            }
            pool.DelOne()
        }(v)
    }
    pool.wg.Wait()
}

func Download(s string) error {
    fmt.Println(s)
    return nil
}
