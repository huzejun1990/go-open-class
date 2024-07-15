// @Author huzejun 2024/7/15 21:57:00
package main

import "fmt"

func traverseSlice() {
	arr := []int{1, 2, 3}
	//Go语言中 for range取得的是集合中元素的拷贝
	for _, ele := range arr {
		ele = ele + 1
	}
	fmt.Println(arr) //[1 2 3]
	for i, ele := range arr {
		arr[i] = ele + 1
	}
	fmt.Println(arr) //[2,3,4]
}
