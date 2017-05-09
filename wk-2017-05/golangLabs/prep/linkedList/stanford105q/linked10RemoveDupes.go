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

// assuming its sorted for now in ascending order.
func removeDupes(headP *Node) (*Node) {
  // sanity checks
  if headP == nil {
    return nil
  }

  // edge cases 1 node...
  if (*headP).p == nil {
    return headP
  }

  // normal case
  //var prevNum int

  for curP, prevP := headP, headP ; ; curP = (*curP).p {
    // 1st node. You have a special edge case.
    // previous checks assumes you have 2 nodes in the lined list.
    if curP == headP {
      curP = (*curP).p
    }

    // check current number to previous number.
    curNum := (*curP).v
    prevNum := (*prevP).v
    //fmt.Println("curP", curP, "prevP", prevP)
    //fmt.Println("curNum", curNum, "prevNum", prevNum)

    if curNum == prevNum && (*curP).p != nil {
      //fmt.Println(">1")
      //fmt.Println(">cp1", curP, "prevP", prevP)
      // dupe exist and there is next node.

      // save curP and delete it and there is next node.
      // point prevP to curP.next
      (*prevP).p = (*curP).p
      //fmt.Println(">cp2", curP, "prevP", prevP)
      (*curP).p = nil   // setting this to nil so you don't traverse this.
      //fmt.Println(">cp3", curP, "prevP", prevP)
      curP = prevP // we deleted curP. setting curP to prevP.
      //fmt.Println(">cp4", curP, "prevP", prevP)
    } else if curNum == prevNum {
      //fmt.Println(">2")
      // dupe exist and there is no next node.

      // delete curP.
      // point prevP to nil. 
      (*prevP).p = nil
      // you are done. break out of loop.
      break
    } else if (*curP).p != nil {
      //fmt.Println(">3")
      // dupe does not exist. and there is an next node.
      // assign the new p.
      prevP = curP
    } else {
      //fmt.Println(">4")
      // dupe does not exist and you are in next node.
      break
    }
  }

  return headP
}


func main() {
  fmt.Println("> 1 node")
  longList1 := createListOfNumbers(1)
  result1 := removeDupes(longList1)
  printList(result1)
  fmt.Println("")

  fmt.Println("> 111111111")
  longList2 := createListOfNumbers(111111111)
  result2 := removeDupes(longList2)
  printList(result2)
  fmt.Println("")

  fmt.Println("> 1111112. repeat start")
  longList3 := createListOfNumbers(111111112)
  result3 := removeDupes(longList3)
  printList(result3)
  fmt.Println("")

  fmt.Println("> 1234555 repeat end")
  longList4 := createListOfNumbers(1234555)
  result4 := removeDupes(longList4)
  printList(result4)
  fmt.Println("")

  fmt.Println("> 12222223455567 repeat middle 2")
  longList5 := createListOfNumbers(122345567)
  result5 := removeDupes(longList5)
  printList(result5)
  fmt.Println("")

  fmt.Println("> 111234555677 start middle end")
  longList6 := createListOfNumbers(111234555677)
  result6 := removeDupes(longList6)
  printList(result6)
  fmt.Println("")

  fmt.Println("> 123455567 repeat middle")
  longList7 := createListOfNumbers(123455567)
  result7 := removeDupes(longList7)
  printList(result7)
  fmt.Println("")

  // error handling
  fmt.Println("> nil")
  //longList8 := createListOfNumbers(123455567)
  result8 := removeDupes(nil)
  printList(result8)
  fmt.Println("")
}