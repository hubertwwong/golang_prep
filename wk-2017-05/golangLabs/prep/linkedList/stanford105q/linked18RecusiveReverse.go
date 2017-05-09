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

func recReverse(headP *Node) (*Node, *Node) {
  // base cases
  if headP == nil {
    //fmt.Println("> headP == nil")
    // sanity check. if things are nil.
    return nil, nil
  } else if (*headP).p == nil {
    // 1 node case.
    //fmt.Println("> 1/ last node node case.")
    return headP, headP
  } else {
    //fmt.Println("> else")
    var finalP, newHeadP *Node
    finalP, newHeadP = recReverse((*headP).p)
    
    // need to nil the end of the list.
    if (*headP).p == finalP {
      (*headP).p = nil
    }
    
    // at this point you have 2 pointers.
    // headP - the 2nd to last node
    // finalP - the last node.
    (*finalP).p = headP
    
    return headP, newHeadP
  }
}



func main() {
  num1 := 1
  list1 := createListOfNumbers(num1)
  _, finalH1 := recReverse(list1)
  fmt.Println("> ", num1)
  printList(finalH1)
  fmt.Println("")

  num2 := 12
  list2 := createListOfNumbers(num2)
  _, finalH2 := recReverse(list2)
  fmt.Println("> ", num2)
  printList(finalH2)
  fmt.Println("")
  
  num3 := 123
  list3 := createListOfNumbers(num3)
  _, finalH3 := recReverse(list3)
  fmt.Println("> ", num3)
  printList(finalH3)
  fmt.Println("")

  num4 := 1234567
  list4 := createListOfNumbers(num4)
  _, finalH4 := recReverse(list4)
  fmt.Println("> ", num4)
  printList(finalH4)
  fmt.Println("")

  //num5 := 12
  //list5 := createListOfNumbers(num5)
  _, finalH5 := recReverse(nil)
  fmt.Println("> nil")
  printList(finalH5)
  fmt.Println("")
}