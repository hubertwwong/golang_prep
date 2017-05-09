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

func printList(initNode Node) {
  curNode := initNode
  
  for {
    if curNode.p == nil {
      fmt.Println(curNode.v)
      break
    }

    fmt.Println(curNode)
    curNode = *curNode.p
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

  // LIST CHECKS
  // Creating list of n size and checking to see if they have a cycle.
  // These should all be false.

  // List of -1 items.
  listNeg := createList(-3)
  fmt.Println("listNeg is ", listNeg)
  hasCycleListNeg := hasCycle(listNeg)
  fmt.Println("hasCycleListNeg is ", hasCycleListNeg)
  
  // List of 0 items
  listZero := createList(0)
  fmt.Println("listZero is ", listZero)
  hasCycleListZero := hasCycle(listZero)
  fmt.Println("hasCycleListZero is ", hasCycleListZero)

  // List 1 item.
  listOne := createList(1)
  fmt.Println("listOne is ", listOne)
  hasCycleListOne := hasCycle(listOne)
  fmt.Println("hasCycleListOne is ", hasCycleListOne)

  // List 2 item.
  listTwo := createList(2)
  fmt.Println("listTwo is ", listTwo)
  hasCycleListTwo := hasCycle(listTwo)
  fmt.Println("hasCycleListTwo is ", hasCycleListTwo)

  // List 10 item.
  listTen := createList(10)
  fmt.Println("listTen is ", listTen)
  hasCycleListTen := hasCycle(listTen)
  fmt.Println("hasCycleListTen is ", hasCycleListTen)

  // List 11 items
  listEleven := createList(11)
  fmt.Println("listEleven is ", listEleven)
  hasCycleListEleven := hasCycle(listEleven)
  fmt.Println("hasCycleListEleven is ", hasCycleListEleven)

  // CYCLE CHECKS...
  // Adding in some cycles and seeing if the hasCycle works properly.
  
  cycleNeg, _ := addCircularLink(listNeg, 0)
  cycleNegBool := hasCycle(cycleNeg)
  fmt.Println("> cycleNeg is ", cycleNegBool)
  
  cycleZero, _ := addCircularLink(listZero, 0)
  cycleZeroBool := hasCycle(cycleZero)
  fmt.Println("> cycleZero is ", cycleZeroBool)

  cycleOne, _ := addCircularLink(listOne, 0)
  cycleOneBool := hasCycle(cycleOne)
  fmt.Println("> cycleZero is ", cycleOneBool)

  cycleTwo, _ := addCircularLink(listTwo, 1)
  cycleTwoBool := hasCycle(cycleTwo)
  fmt.Println("> cycleTwo is ", cycleTwoBool)

  cycleTen, _ := addCircularLink(listTen, 5)
  cycleTenBool := hasCycle(cycleTen)
  fmt.Println("> cycleTen is ", cycleTenBool)

  cycleEleven, _ := addCircularLink(listEleven, 20)
  cycleElevenBool := hasCycle(cycleEleven)
  fmt.Println("> List of 11 items. Setting cycle of 20th node. Answer is ", cycleElevenBool)
}

/*

Does a basic cycle test...

fails if the list has 1 item


*/