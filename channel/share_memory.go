// @Author huzejun 2024/7/10 21:58:00
package main

import (
	"fmt"
	"sync"
	"time"
)

var mp sync.Map //共享内存

func rwGlobalMemory() {
	if value, exists := mp.Load("mykey"); exists {
		fmt.Println(value)
	} else {
		mp.Store("mykey", "myvalue")
	}
}

func main1() {
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()

	//rwGlobalMemory()
	//rwGlobalMemory()
	//rwGlobalMemory()
	//rwGlobalMemory()

	time.Sleep(time.Second)
}
