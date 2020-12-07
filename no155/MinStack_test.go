package no155

import "testing"

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(12)
	obj.Push(-12)
	obj.Push(2)
	obj.Push(34)
	println(obj.Top())
	println(obj.GetMin())
	obj.Pop()
	println(obj.Top())
}
