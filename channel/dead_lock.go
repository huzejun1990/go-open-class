// @Author huzejun 2024/7/11 3:53:00
package main

import "fmt"

/*func traveseChannel2() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) //删除本行会报deadlock
	for {
		ele, ok := <-ch
	}
}*/

func traveseCannel3() {
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		//close(ch) //删除本行会报deadlock
	}()
	for ele := range ch { //遍历并取走管道中的元素
		fmt.Println(ele)
	}
	fmt.Println("bye bye")
}

func main() {
	traveseCannel3()
}
