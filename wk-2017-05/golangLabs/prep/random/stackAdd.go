package main

import "fmt"

type Node struct {
  val int
  op string
  next *Node
}

type Stack struct {
  root *Node
}

func (s *Stack) push(val int, op string) {
  fmt.Println("> stack > push", val, op)
  if s.root == nil {
    s.root = &(Node{val, op, nil})
  } else {
    newRoot := &(Node{val, op, s.root})
    s.root = newRoot
  }
}

func (s *Stack) pop() Node {
  if s.root == nil {
    var n Node
    return n
  } else {
    fmt.Println("> stack > pop > root", s.root)
    retVal := *(s.root)
    s.root = (*s.root).next
    fmt.Println("> stack > pop > new root", s.root)
    retVal.next = nil
    return retVal
  }
}

func (s *Stack) top() Node {
  if s.root == nil {
    var n Node
    return n
  } else {
    return *(s.root)
  }
}

func (s *Stack) empty() bool {
  if s.root == nil {
    return true
  } else {
    return false
  }
}

func isInt(b byte) bool {
   if b >= byte("0"[0]) && b <= byte("9"[0]) {
    return true
  } else {
    return false
  }
}

// so if you have multiple return. you need a paren.
func parseInt(s string, start int, neg bool) (int, int) {
  num := 0
  pos := -1

  i := start
  for ; i<len(s) ; i++ {
    if isInt(s[i]) {
      num = num * 10 + (int(s[i])-48)
    } else {
      pos = i-1
      break
    }
  }

  if neg {
    num = num * -1
  }
  // edge case if the character is at the end of the list
  if i == len(s) {
    pos = len(s) - 1
  }

  return num, pos
}

// This should just add until you hit a paren.. then it checks for a neg -1...
// in case the value was meant to be subtracted.
func (s *Stack) calcCurrent(total int) int {
  for curNode := s.pop() ;   ; curNode = s.pop() {
    //fmt.Println("> calcCurrent > on ", curNode)
    if curNode.op == "" {
      total = total + curNode.val
    } else if curNode.op == "(" {
      // basically want to look for a -1.
      // for the "- (" case
      curNode = s.top()
      if s.top().op == "*" {
        s.pop()
        total = total * -1
      }

      break
    }

    // adding this here.
    if s.empty() {
      break
    }
  }

  return total
}

func calculate(s string) int {
  var stk Stack
  total := 0
  currentTotal := 0

  neg := false

  // use this for manipulation of i.
  for i := 0 ; i<len(s) ; i++ {
    fmt.Println("")
    fmt.Println("i b>", i, ">[", string(s[i]), "]", s)
    if s[i] == byte(" "[0]) {
      continue
    } else if s[i] == byte("("[0]) {
      if neg == true {
        stk.push(-1, "*")
        neg = false
      }
      stk.push(0, "(")
    } else if s[i] == byte("+"[0]) {
      // not sure if this is right
      continue
    } else if s[i] == byte("-"[0]) {
      // set a flag. the next value will determine what happens.
      neg = true
    } else if isInt(s[i]) {
      // push a number into the stack
      curNum, pos := parseInt(s, i, neg)
      //fmt.Println("> pushing", curNum)
      stk.push(curNum, "")
      i = pos

      // reset flag
      neg = false
    } else if s[i] == byte(")"[0]) {
      currentTotal = stk.calcCurrent(currentTotal)
      if stk.empty() {
        total = total + currentTotal
        currentTotal = 0
      }
      fmt.Println("> popping > t", total)
    }
    fmt.Println("i a>", i, "l", len(s))
  }

  // pop off remaining stack.
  //fmt.Println(">c ", total)
  currentTotal = stk.calcCurrent(currentTotal)
  if stk.empty() {
    total = total + currentTotal
    currentTotal = 0
  }
  return total
}



func main() {
  //fmt.Println(">m", calculate("3+2"))
  //fmt.Println(">m", calculate("3+2+(4+6)  "))
  //fmt.Println(">m", calculate("3+2-(3+3)"))
  //fmt.Println(">m", calculate("(1+(4+5+2)-3)+(6+8)"))
  //fmt.Println(">m", calculate("(5-(1+(5)))"), "-1")
  //fmt.Println(">m", calculate("(7)-(0)+(4)"), "11")
  //fmt.Println(">m", calculate("(3-(5-(8)-(2+(9-(0-(8-(2))))-(4))-(4)))"), ">23")
  //fmt.Println(">m", calculate("(3-(5-(8)))-(100)"), ">1")
  fmt.Println(">m", calculate("(4)-(4)"), ">1")
  //g, v := parseInt("300", 0, false)
  //fmt.Println(">m", g, v)
}

/*
multiple return need a paren.
11:57a

lots of syntax error.
pointers can be nil. object cannot
var names before types....
calc current not haing the correct return type.
need to play around with the string to byte thing a bit more
parseInt should take a string...
nil for structs..
- basically just check for the value s inside.
s.top didn't have a func
- if you have a function call... make it..
forgot to add code for pos..

and fixing caused bugs...
and it crashes...
logic errors.....

parse int is broekn...
not sure how to convert a string character to int character.

is int is wrong.
- you need an inclusive check.

parse int didn't have an edge case of having the number at the end of the list....

i don't like the stack design..

data structs...
you have to put them in one group....

*/