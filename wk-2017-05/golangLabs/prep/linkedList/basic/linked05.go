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

// create a list of nodes.
// starts on 1 at the end of the list.
// so it will count backwards.
func createList(numNodes int) (*Node) {
  var startNode, curNode *Node

  // create last linked
  lastNode := Node{nil, numNodes}

  // handle edge cases.
  if numNodes < 0 {
    return nil        
  } else if numNodes == 1 {
    return &lastNode
  }

  // create n-1 other node.
  for i := 0 ; i<numNodes ; i++ {
    //fmt.Println("i", i)
    // handle init case.
    if i == 0 {
      curNode = &lastNode
    }

    // setup new node.
    newNode := Node{nil, numNodes-(i+1)}

    // link node.
    curNode = addLink(&newNode, curNode)

    // update start node
    startNode = curNode
  }

  return startNode
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

// assuming they are list without cycles for now.
func subTwoLists(n1, n2 *Node) (int) {
  n1Int := getIntFromList(n1)
  n2Int := getIntFromList(n2)
  
  // error checks. add them.
  return (n1Int - n2Int)
}

func getIntFromList(initNode *Node) (int) {
  curNode := initNode
  var curNum = 0
  
  // initial checks
  if initNode == nil {
    fmt.Println("nil")
  } else {
    for {
      // print last node.
      if curNode.p == nil {
        curNum = curNum + (*curNode).v
        fmt.Println("{nil",(*curNode).v, "}")
        break
      }

      // calc. Add the value. Mul 10
      curNum = curNum + (*curNode).v
      curNum = curNum * 10 

      // advance the pointer.
      fmt.Println(*curNode)
      curNode = (*curNode).p
    }
  }

  return curNum
}

func addCircularLink(startNode *Node, indexToPoint int) (*Node, bool) {
  var curNode, pointNode, lastNode *Node
  
  // Initial checks
  if startNode == nil {
    return nil, false
  }

  // Boolean to return if if you didn't link the node.
  isLinked := false

  // point curNode to start of list
  curNode = startNode

  // figure out the nodes in question.
  i := 0
  for {
    fmt.Println(curNode)
    // You hit a nil
    if curNode.p == nil {
      lastNode = curNode
      //fmt.Println("lastNode", lastNode);
      break
    }

    // if i>indexToPoint {
    //   lastNode = curNode
    //   break
    // }

    // figure out if you are at the node to point to.
    if i == indexToPoint {      
      pointNode = curNode
      //fmt.Println("indexToPoint", pointNode);
    }
    
    // increment stuff.
    i++
    curNode = curNode.p
  }
  
  // point the last node to the node in question.
  // its odd that you can't check for nils for the actual objects.
  if pointNode != nil && pointNode.p != nil && lastNode != nil && lastNode.p == nil {
    addLink(lastNode, pointNode)
    //fmt.Println(result)
    isLinked = true
  }

  return startNode, isLinked
}

func hasCycle(initNode *Node) (bool) {
  slow, fast := initNode, initNode
  
  // Initial checks.
  if initNode == nil {
    return false
  }

  for {
    // if the next are nil, you didn't hit a cycle.
    // 3 checks needed for exist.
    if (*fast).p == nil || (*(*fast).p).p == nil || (*slow).p == nil {
      return false
    }

    // fast node checks. if its good, move pointer.
    if (*fast).p != nil && (*(*fast).p).p != nil {
      //fmt.Println("fast b", (*fast).v)
      fast = (*fast).p
      fast = (*fast).p
      //fmt.Println("fast f", (*fast).v)
    }

    // slow node checks. if its good, move pointer.
    if (*slow).p != nil {
      //fmt.Println("slow b", (*slow).v)
      slow = (*slow).p
      //fmt.Println("slow f", (*slow).v)
    }

    // cycle checks
    //fmt.Println("slow c", (*slow).v)
    //fmt.Println("fast c", (*fast).v)
    if fast == slow {
      //fmt.Println("cycle detected");
      return true
    }

    // output slow pointer.
    fmt.Println(slow)
  }

  return false
}


func main() {
  list350 := createListOfNumbers(352)
  list200 := createListOfNumbers(200)
  
  result := subTwoLists(list350, list200)
  fmt.Println(">result>", result)
  
  //fmt.Println(">>>>>>>>>>")
  //printList(list350)
  //fmt.Println("list350", list350)
}