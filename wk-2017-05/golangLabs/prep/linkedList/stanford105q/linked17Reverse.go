package main

import "fmt"

type Node struct {
  p *Node
  v int
}

// This was changed to take pointers.
// You need to modify the actual address and not copies.
func addLink(from, to *Node) (*Node) {
  toAddress := to;
  (*from).p = toAddress;
  return from;
}

// assming its positive for now.
func createListOfNumbers(myNum int) (*Node) {
  var startNode, curNode *Node
  
  // initial cases.
  if myNum < 0 {
    return nil
  } else if myNum < 10 {
    lastNode := Node{nil, myNum}
    curNode = &lastNode
    return curNode
  } else {
    lastNode := Node{nil, myNum%10}
    //fmt.Println(">i", myNum%10)
    curNode = &lastNode
  }

  for curNum := myNum/10 ; ; curNum = curNum/10 {
    if curNum >= 10 {
      //fmt.Println("l", curNum)
      // setup new node.
      newNode := Node{nil, curNum%10}
      
      // link node.
      curNode = addLink(&newNode, curNode)
      
      // update start node
      startNode = curNode
    } else {
      //fmt.Println("f", curNum)
      // setup new node.
      newNode := Node{nil, curNum}
      
      // link node.
      curNode = addLink(&newNode, curNode)

      // update start node
      startNode = curNode

      // number is < 10. break out of loop.
      break
    }
  }

  return startNode
}

func printList(initNode *Node) {
  curNode := initNode
  
  // initial checks
  if initNode == nil {
    fmt.Println("nil")
  } else {
    for {
      // print last node.
      if curNode.p == nil {
        fmt.Println("{nil",(*curNode).v, "}")
        break
      }

      // print anything else
      fmt.Println(*curNode)
      curNode = (*curNode).p
    }
  }
}

func reverse(headP *Node) (*Node) {
  // sanity checks / base cases
  if headP == nil {
    return nil
  } else if (*headP).p == nil {
    return headP
  }

  var newHeadP *Node
  newHeadP = nil

  // Moving this out of the for loop for reability.
  prevP := headP 
  curP := (*headP).p
  nextP := (*curP).p
  nextNextP := headP        // need this to stash a extra pointer.

  // 2 or more.
  // curPoints starts on the second node.
  for i:=0 ; i<10 ; i++ {
    // fmt.Println(">>>>>>>")
    // fmt.Println(prevP)
    // fmt.Println(curP)
    // fmt.Println(nextP)
    // fmt.Println("<<<<<<<")
    
    if prevP == headP {
      //fmt.Println(">prevP==headP")
      // first node.
      (*prevP).p = nil
      (*curP).p = prevP

      // a check for 2 nodes in the list.
      if nextP == nil {
        // curP is the final node.
        return curP
        //break
      } else if (*nextP).p != nil {
        // 3+ nodes
        nextNextP = (*nextP).p
      } else {
        // 3 nodes.
        (*nextP).p = curP
        return nextP
      }
    } else if nextP == nil {
      //fmt.Println(">nextP == nil")
      // last node
      (*curP).p = prevP
      return curP
      //break
    } else {
      //fmt.Println(">else")
      // all other nodes.
      nextNextP = (*nextP).p
      (*curP).p = prevP
      (*nextP).p = curP
    }

    // increments.
    // need to be careful here.
    // previous statements are doing pointer mods...
    // order matters.
    prevP = curP
    curP = nextP
    nextP = nextNextP
  }

  return newHeadP
}

func main() {
  num1 := 1
  list1 := createListOfNumbers(num1)
  final1 := reverse(list1)
  fmt.Println("> 1", num1)
  printList(final1)
  fmt.Println("")

  num2 := 12
  list2 := createListOfNumbers(num2)
  final2 := reverse(list2)
  fmt.Println("> 2", num2)
  printList(final2)
  fmt.Println("")

  num3 := 123
  list3 := createListOfNumbers(num3)
  final3 := reverse(list3)
  fmt.Println("> 3", num3)
  printList(final3)
  fmt.Println("")

  num4 := 123456789
  list4 := createListOfNumbers(num4)
  final4 := reverse(list4)
  fmt.Println("> 3+", num4)
  printList(final4)
  fmt.Println("")

  //num5 := 123456789
  //list5 := createListOfNumbers(num5)
  final5 := reverse(nil)
  fmt.Println("> nil", final5)
  printList(final5)
  fmt.Println("")
}