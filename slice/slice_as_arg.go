// @Author huzejun 2024/7/15 21:45:00
package main

import "fmt"

func changeSlice(arr *[]int) {
	//len和cap要变，需要传切片指针
	*arr = append(*arr, 9) //arr=[1 2 3 9] cap = 6
	//len、pointer和cap要变动，需要传切片指针
	*arr = (*arr)[1:2] //arr=[2]
	//slice结构体里的3个Field都不变，只变底层数组，不需要传切片指针
	(*arr)[0] = 9 //arr=[9] len=1 cap

}

func testChangeSlice() {
	arr := []int{1, 2, 3}
	fmt.Printf("len %d cap %d array %v\n", len(arr), cap(arr), arr)
	changeSlice(&arr)
	fmt.Printf("len %d cap %d array %v\n", len(arr), cap(arr), arr)
}
