// @Author huzejun 2024/7/15 17:24:00
package main

import "fmt"

func sub_slice() {
	/*
		截取一部分，创建子切片，些时子切片与母切片（或母数组）共享内存空间，母切片的capacity子切片可能直接用
	*/
	s := make([]int, 3, 5)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	}
	fmt.Println("s[1] address %p\n", &s[1])
	sub_slice := s[1:3] //从切片创造子切片，len=2,cap=2
	fmt.Printf("len %d cap %d\n", len(sub_slice), cap(sub_slice))
	/*
		母切片的capacity还允许子切片执行append操作
	*/
	sub_slice = append(sub_slice, 6, 7) //可以一次append多个元素
	sub_slice[0] = 8
	fmt.Printf("s=%v, sub_slice=%v, s[1] address %p, sub_slice[0] address %p\n", s, sub_slice, &s[1], &sub_slice[0])
	/*
		母切片的capacity用完了，子切片再执行append就得申请一片新的内存，把老数据先拷贝过来，在新内存上执行append操作
	*/
	sub_slice = append(sub_slice, 8)
	sub_slice[0] = 9
	fmt.Printf("s=%v,sub_slice=%v,s[1] address %p,sub_slice[0] address %p\n", s, sub_slice, &s[1], &sub_slice[0])
}
