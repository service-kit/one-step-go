package main

import (
	"./stack"
	"fmt"
)

func main() {
	sq := new(stack.StackQueue)
	sq.Init()
	for i := 0; i < 10; i++ {
		sq.Add(i)
	}
	for i := 0; i < 11; i++ {
		fmt.Println(sq.Peek())
		fmt.Println(sq.Poll())
	}
}
