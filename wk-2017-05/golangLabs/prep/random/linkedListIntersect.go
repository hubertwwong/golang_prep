package main

import "fmt"

type Node struct {
  val int
  next *Node
}

// =====LINKED LIST

type LinkedList struct {
  root *Node
  tail *Node
}

func (l *LinkedList) add(val int) {
  n := Node{val, nil}
  if l.root == nil {
    //fmt.Println(">root")
    l.root = &n
    l.tail = l.root
  } else {
    //fmt.Println("tail")
    (*l.tail).next = &n
    l.tail = (*l.tail).next
  }
  //fmt.Println(l.root, l.tail)
  //fmt.Println("")
}

func (l *LinkedList) addList(vals []int) {
  for i:= 0 ; i< len(vals) ; i++ {
    l.add(vals[i])
  }
}

func (l *LinkedList) find(val int) *Node {
  for c := l.root ; c != nil ; c = (*c).next {
    if (*c).val == val {
      return c
    }
  }
  return nil
}

// =====STACK

type Stack struct {
  root *Node
}

func (s *Stack) push(val int) {
  n := Node{val, nil}
  if s.root == nil {
    s.root = &n
  } else {
    // didn't have else paper code.
    n.next = s.root
    s.root = &n
  }
}

func (s *Stack) pop() int {
  if s.root == nil {
    return -1
  }

  retItem := s.root
  s.root = (*s.root).next
  (*retItem).next = nil
  return (*retItem).val
}

// =====HELPERS

func pushToStack(l LinkedList) Stack {
  var s Stack
  for current := l.root ; current != nil ; current = (*current).next {
    s.push((*current).val)
  }
  return s
}

// =====CORE FUNC

func hasIntersect(l1, l2 LinkedList) int {
  if l1.root == nil || l2.root == nil {
    return -1
  }
  s1 := pushToStack(l1)
  s2 := pushToStack(l2)
  s1Node := s1.pop()
  s2Node := s2.pop()
  prevVal := -1

  for {
    if s1Node == -1 || s2Node == -1 {
      return -1
    }

    if s1Node == s2Node {
      prevVal = s1Node
      s1Node = s1.pop()
      s2Node = s2.pop()
    } else {
      return prevVal
    }
  }
}

func main() {
  var l1 LinkedList
  var l2 LinkedList
  
  //fmt.Println(">1");
  //l1.addList([]int{1,2,3,4,5,6})
  //fmt.Println(">1.1");
  //l2.addList([]int{7,8})
  //fmt.Println(">2", l2.tail, l1.tail);
  //fmt.Println(">>", l1.root, (*l1.root).next)
  //(*l2.tail).next = l1.tail
  //intersect := l1.find(2)
  //(*l2.tail).next = intersect

  fmt.Println(">3");
  //fmt.Println(">result", hasIntersect(l1, l2))
  fmt.Println(">result", hasIntersect(nil, nil))
} 


/*
  not checking the return types to values that i assign to...
  typos..
  n for nil...
  logic errors.

  linkedlistErrors..
  forgot to add * to methods..



*/