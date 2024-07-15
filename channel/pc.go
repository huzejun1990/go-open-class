// @Author huzejun 2024/7/11 2:11:00
package main

import (
	"fmt"
	"time"
)

type Task struct {
	A int
	B int
}

var ch = make(chan Task, 100)

func producer() {
	for i := 0; i < 10; i++ {
		ch <- Task{A: i + 3, B: i - 8} //2i-5
	}
}

func consumer() {
	for i := 0; i < 10; i++ {
		task := <-ch
		sum := task.A + task.B
		fmt.Println(task.A, task.B, sum)
	}
}

func main2() {
	go producer() //生产者协程
	go consumer() //消费者协程
	time.Sleep(time.Second)
}
