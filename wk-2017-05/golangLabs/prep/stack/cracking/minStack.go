package main

import "fmt"

type Node struct {
  Next *Node
  Val int
  MinV int
}

type Stack struct {
  Root *Node
}

func (s *Stack) Push(val int) {
  fmt.Println(s.Root)
  if s.Root == nil {
    n := Node{nil, val, val}
    s.Root = &n
    fmt.Println(">>>",s.Root)
  }
  newMin := val
  if val > (*(s.Root)).MinV {
    v := s.Root
    newMin = (*v).MinV
  }
  newRoot := Node{s.Root, val, newMin}
  s.Root = &newRoot
}

func (s *Stack) Pop() (*Node) {
  if s.Root == nil {
    return nil
  }

  retNode := s.Root
  s.Root = (*(s.Root)).Next
  (*retNode).Next = nil
  return retNode
}

func (s *Stack) Min() int {
  if s.Root == nil {
    return -1
  }
  return (*(s.Root)).MinV
}

func main() {
  var stack Stack
  //stack := new(Stack)
  stack.Push(5)
	stack.Push(4)
	//stack.Push(5)
	//stack.Push(2)
	//stack.Push(3)
	//stack.Push(1)
	//stack.Push(4)
	//stack.Push(1)

  fmt.Println(stack.Min())
  stack.Pop()
  fmt.Println(stack.Min())
  stack.Pop()
  fmt.Println(stack.Min())
  stack.Pop()
  fmt.Println(stack.Min())
}

/*

s.Root

func new(Type) *Type
This is the thing...
new you new something you are returning *Type
you need to use that argment in the methods....

*/