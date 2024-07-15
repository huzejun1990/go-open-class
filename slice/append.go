// @Author huzejun 2024/7/15 4:12:00
package main

import (
	"fmt"
	"time"
)

// 探究capacity扩容规律
func expansion() {
	s := make([]int, 0, 7)
	prevCap := cap(s)
	for i := 0; i < 100; i++ {
		s = append(s, i)
		currCap := cap(s)
		if currCap > prevCap {
			//每次扩容都是搞到原先的2倍
			fmt.Printf("capacity从%d变成%d\n", prevCap, currCap)
			prevCap = currCap
		}
	}
	fmt.Println("===============")
}

const (
	LOOP = 1000000
)

func huge_append1() {
	begin := time.Now()
	arr := []int{}
	for i := 0; i < LOOP; i++ {
		arr = append(arr, i)
	}
	fmt.Printf("huge_append1 use time %dms arr len %d\n", time.Since(begin).Milliseconds(), len(arr))
}

func huge_append2() {
	begin := time.Now()
	arr := make([]int, 0, LOOP)
	for i := 0; i < LOOP; i++ {
		arr = append(arr, i)
	}
	fmt.Printf("huge_append2 use time %dms arr len %d\n", time.Since(begin).Milliseconds(), len(arr))
}
