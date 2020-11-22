package main

import (
	"./stack"
	"fmt"
)

func main() {
	gms := new(stack.GetMinStack)
	gms.Init()
	gms.Push(8)
	gms.Push(7)
	gms.Push(9)
	gms.Push(5)
	gms.Push(4)
	gms.Push(10)
	fmt.Println(gms.GetMin())
	gms.Pop()
	fmt.Println(gms.GetMin())
	gms.Pop()
	fmt.Println(gms.GetMin())
	gms.Pop()
	fmt.Println(gms.GetMin())
	gms.Pop()
	fmt.Println(gms.GetMin())
	gms.Pop()
	fmt.Println(gms.GetMin())
	gms.Pop()
	fmt.Println(gms.GetMin())
}
