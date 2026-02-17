package main

import (
	"fmt"
)

type Stack struct {
	Data []any
}

func (stack Stack) IsEmpty() bool {
	return len(stack.Data) == 0
}

func (stack *Stack) Add(element any) {
	stack.Data = append(stack.Data, element)
	fmt.Println("item added to stack", element)
}

func (stack *Stack) Remove() {
	if ok := stack.IsEmpty(); !ok {
		stack.Data = stack.Data[:len(stack.Data)-1]
		fmt.Println("item added to stack: ", stack.Data[:len(stack.Data)-1])
	}
	fmt.Println("stack is empty!")
}

func (stack Stack) Read() any {
	if ok := stack.IsEmpty(); !ok {
		return stack.Data[len(stack.Data)-1]
	}
	return nil
}

func (stack *Stack) Clear() {
	stack.Data = make([]any, 0)
	fmt.Println("stack is now clear!")
}

func main() {
	instance := Stack{}

	instance.Add("10")
	instance.Add(struct {
		ID   int
		Name string
	}{1, "Mobile"})
	instance.Add(1.22)

	instance.Remove()

	read_last := instance.Read()
	fmt.Println(read_last)

	instance.Clear()
}
