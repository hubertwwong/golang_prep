package main

import "fmt"

type NodeS struct {
  next *NodeS
  val int
}

type Stack struct {
  root *NodeS
}

// I am not returning like i did before in the if block
func (s *Stack) Push(val int) {
  if s.root == nil {
    n := NodeS{nil, val}
    s.root = &n
  } else {
    n := NodeS{s.root, val}
    s.root = &n
    //fmt.Println(s.root)
  }
}

// this is the issue....
// i think
func (s *Stack) Pop() (int) {
  if s.root == nil {
    return -1
  }
  // changed the variable names so its more readable...
  poppedNode := s.root
  s.root = (*(s.root)).next
  (*poppedNode).next = nil
  poppedVal := (*poppedNode).val
  return poppedVal
}

type MyQueue struct {
  q1 Stack
  q2 Stack
}

func (q *MyQueue) Add (val int) {
  q.q1.Push(val)
}

// this has a bug of 1 node...
// didn't consider it...
// when designing the algo..
func (q *MyQueue) Remove() (int) {
  // fmt.Println(">remove>")
  var curV int
  
  for {
    // fmt.Println("push to 2")
    // logic here is wrong...
    // you should pop and push and then check.
    curV = q.q1.Pop()
    if curV != -1 {
      q.q2.Push(curV)
      // fmt.Println(">1",curV)
    } else {
      break
    }
  }

  // i think this is bugged....
  retVal := q.q2.Pop()
  curV = q.q2.Pop()
  // this was the fix...

  for {
    // fmt.Println("push to 1")
    // i think this is wrong...
    // you should not pop here.
    //curV = q.q2.Pop()
    if curV != -1 {
      q.q1.Push(curV)
      curV = q.q2.Pop()
    } else {
      break
    }
  }
  
  return retVal
}

func main() {
  var q MyQueue
  q.Add(1)
  //q.Add(2)
  //q.Add(3)
  //q.Add(4)
  r := q.Remove()
  fmt.Println(">o",r)
  r = q.Remove()
  fmt.Println(">o",r)
  r = q.Remove()
  fmt.Println(">o",r)
  r = q.Remove()
  fmt.Println(">o",r)
  r = q.Remove()
  fmt.Println(">o",r)
}

/*

does not work..
03:40p
why...
big thing is to trace you stuff down....
i think my stack logic is wrong..

when using structs and not returing anything. you can't use an if block and fall thru.
- you have to be careful to know that is what you want...

and 1 item...

on the queue remove function
i was double popping in the loop...
this is always probably wrong...

2 big lessons...
if you are using structs, you might not be returning values due to some internal data struct...
- you should not assume that you are returing and program defensively...

iterator.
- should care on how many items you iterate through.
- it should only be one item..


*/