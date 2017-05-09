package main

import "fmt"

type Node struct {
  p *Node
  v int
}

type LList struct {
  head *Node
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
    fmt.Println(">i", myNum%10)
    curNode = &lastNode
  }

  for curNum := myNum/10 ; ; curNum = curNum/10 {
    if curNum >= 10 {
      fmt.Println("l", curNum)
      // setup new node.
      newNode := Node{nil, curNum%10}
      
      // link node.
      curNode = addLink(&newNode, curNode)
      
      // update start node
      startNode = curNode
    } else {
      fmt.Println("f", curNum)
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

// it should return the new head node.
func pop(head *Node) (*Node){
  // sanity checks
  if head == nil {
    return nil
  }

  // pop 1 node.
  if (*head).p == nil {
    // dealloc this.
  }

  // general case..
  tempN := (*head).p
  fmt.Println("deleting the head node", tempN)

  return (*head).p
}



func main() {
  myList := createListOfNumbers(123)

  fmt.Println("> initial")
  printList(myList)

  fmt.Println("> pop1")
  myList = pop(myList)
  printList(myList)
  
  fmt.Println("> pop2")
  myList = pop(myList)
  printList(myList)
  
  fmt.Println("> pop3")
  myList = pop(myList)
  printList(myList)
}