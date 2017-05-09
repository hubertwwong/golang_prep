package main

import "fmt"

// linked list stuff

type NodeL struct {
  next *NodeL
  val *Node
}

func newNodeL(val *Node) (*NodeL) {
  n := NodeL{nil, val}
  return &n
}

// insert to the start of the queue
// this is the end of the list
func insertQ(val *Node, root *NodeL) (*NodeL) {
  n := newNodeL(val)
  inserted := false

  if root == nil {
    //fmt.Println("> insertQ > root is nil", n)
    //fmt.Println("> insertQ > ", (*(*n).val).val)
    return n
  } else {
    // insert
    for curNode := root ; curNode != nil ; curNode = (*curNode).next {
      // if the next pointer is nil, you are at the end of the list
      // append the last item.
      if (*curNode).next == nil && inserted == false {
        (*curNode).next = n
        inserted = true
      }
    }
    
    // return the root node.
    return root
  }
}

// get the last item off the queue.
func getQ(root *NodeL) (*Node, *NodeL) {
  if root == nil {
    return nil, nil
  } else {
    var newRoot *NodeL

    // extract the current value
    cVal := (*root).val
    
    // new root
    newRoot = (*root).next
    
    // set next pointer to nil.
    // do you have to do this?
    (*root).next = nil
    
    // return the stuff
    return cVal, newRoot
  }
}


// tree stuff

type Node struct {
  left *Node
  right *Node
  val int
}

func newNode(val int) (*Node){
  n := Node{nil, nil, val}
  return &n
}

func insert(val int, root *Node) (*Node) {
  if root == nil {
    return newNode(val)
  }

  if val <= (*root).val {
    (*root).left = insert(val, (*root).left)
  } else {
    (*root).right = insert(val, (*root).right)
  }

  return root
}

func insertList(vals []int) (*Node) {
  if len(vals) == 0 {
    return nil
  }

  var root *Node
  for i:=0 ; i<len(vals) ; i++ {
    root = insert(vals[i], root)
  }

  return root
}

func printBFS(root *Node) {
  var q *NodeL
  //fmt.Println("> BFS >")
    
  // intial setup. add a queue.
  if root != nil {

    // initial setup
    q := insertQ(root, q)
    //fmt.Println("> BFS > ", q)
    //fmt.Println("> BFS > root val", (*root).val)
    //fmt.Println("> BFS > what q is root val", (*(*q).val).val)    

    for {
      //fmt.Println("> BFS > for loop start")
      
      // create a new q.
      var newQ *NodeL

      // check if the q is nil.
      if q != nil {
        
        // dequeue.
        // print result
        // stuff children in new queue.
        for {
          // get child node of the queue which is a tree node.
          treeN := (*q).val
          
          // insert new tree nodes in queue
          if (*treeN).left != nil {
            //fmt.Println("> BFS > left");
            newQ = insertQ((*treeN).left, newQ)
          }
          if (*treeN).right != nil {
            //fmt.Println("> BFS > right");
            newQ = insertQ((*treeN).right, newQ)
          }

          // extract existing.
          // technically you don't need the value.
          // you have it eariler.
          var cVal *Node
          //fmt.Println("> before fail")
          cVal, q = getQ(q)
          //fmt.Println("> after fail")

          // print value
          if cVal != nil {
            //fmt.Println("> BFS > printing")
            fmt.Println("> BFS > final >", (*cVal).val)
          }

          // exit condition.
          if q == nil {
            break
          }
        }
      } else {
        // if queue is empty break.
        break
      }

      // exit conditions for the entire thing.
      if newQ == nil {
        break
      } else {
        q = newQ
      }

      fmt.Println("")
    }
  }
}

func testRun(vals []int) {
  fmt.Println("> test run > ==================================================")
  root := insertList(vals)
  // fmt.Println(root)
  // fmt.Println((*root).left)
  // fmt.Println((*root).right)
  printBFS(root)
  //fmt.Println("> test run end")
  fmt.Println("")
}



func main() {
  testRun([]int{2,1,3})
  testRun([]int{10, 5, 15, 1, 7, 11, 16})
  testRun([]int{2, 3, 4, 5, 6, 1})
  testRun([]int{2})

  //n1 := insertQ(1, nil)
  //n1 = insertQ(2, n1)
  
  // dequeue
  //v, newN := getQ(n1)
  //fmt.Println(v, newN)
}