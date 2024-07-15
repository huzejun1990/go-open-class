// @Author huzejun 2024/7/13 21:49:00
package main

import (
	"container/list"
	"fmt"
)

var (
	n1    = 0        //记录递归的时间复杂度
	n2    = 0        //记录非递归的时间复杂度
	stack *list.List //List实现了栈的功能
)

func init() {
	stack = list.New()
}

// 斐波那契数列可如下 被递归的方法定义：F(1) = 0,F(2) = 1,F(n) = F(n-1)+F(n-2)
//
// 斐波那契数列前10个数为：0，1，1，2，3，5，8，13，21，34
func FibonacciWithRecursion(n int) int {
	if n == 1 || n == 2 {
		return n - 1 //凡是递归，一定要有终止条件，否则会进入无限循环
	}
	n1++
	return FibonacciWithRecursion(n-1) + FibonacciWithRecursion(n-2) //递归调用函数自身
}

func FibonacciWithStack(n int) int {
	if n == 1 || n == 2 {
		return n - 1 //凡是递归，一定要有终止条件，否会进入无限循环
	}
	stack.PushBack(0)
	stack.PushBack(1)
	for i := 2; i < n; i++ {
		//弹出栈顶的2个元素，分别赋给a和b
		a := stack.Back()
		stack.Remove(a) //从链表上删除一个元素的时间复杂度为0(1)
		b := stack.Back()
		stack.Remove(b)

		//依次压入a和a+b
		stack.PushBack(a.Value.(int))
		stack.PushBack(a.Value.(int) + b.Value.(int))
		n2 += 5
	}
	a := stack.Back()
	result := stack.Remove(a)
	n2 += 5
	return result.(int)
}

func main() {
	//n := 20
	n := 5
	fmt.Println(FibonacciWithRecursion(n))
	fmt.Println(FibonacciWithStack(n))
	fmt.Println(n1, n2)
}
