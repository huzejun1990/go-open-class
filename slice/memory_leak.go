// @Author huzejun 2024/7/15 19:00:00
package main

import "fmt"

const (
	TOTAL = 1 << 20
	begin = 2
	end   = 7
)

func return_sub_slice1() []int {
	parent := make([]int, TOTAL) //parent切片持有一个长度为1M的int数组，共占内存8M
	child := parent[begin:end]   //只要child没被GC回收，长度为1M的int就一直得不到释放
	return child
}

func return_sub_slice2() []int {
	parent := make([]int, TOTAL)
	length := end - begin
	child := make([]int, length)
	for i := begin; i < end; i++ {
		child[i-begin] = parent[i]
	}
	return child

}

func use_sub_slice() {
	s1 := return_sub_slice1()
	fmt.Printf("init len %d cap %d TOTAL=%d\n", len(s1), cap(s1), TOTAL)
	for i := 0; i < 5; i++ {
		s1 := append(s1, 9)
		fmt.Printf("len %d cap %d\n", len(s1), cap(s1))
	}
	fmt.Println("====================")
	s2 := return_sub_slice2()
	fmt.Printf("init len %d cap %d\n", len(s2), cap(s2))
	for i := 0; i < 5; i++ {
		s2 := append(s2, 9)
		fmt.Printf("len %d cap %d\n", len(s2), cap(s2))
	}

}
