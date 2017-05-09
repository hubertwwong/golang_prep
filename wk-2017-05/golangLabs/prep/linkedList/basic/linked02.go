package main

import "fmt"

type Node struct {
  p *Node
  v int
}

//pointing from the form node to the to node.
func addLink(from, to Node) (*Node) {
  curPointer := &to;
  from.p = curPointer;
  return &from;
}

// func createList(n int) (*Node) {
//   startP *Node
//   curP *Node
  
//   // create first link.
//   lastNode := Node{nil, 1}

//   // create n other nodes and connect them.
//   // for now it will just add one.
//   for i := 0 ; i<n ; i++ {
//     // handle case 0
//     if i == 0 {
//       curP = &lastNode
//     }

//     // setup new node.
//     // probably should add a check.
//     newNode = Node{nil, curNode.v + 1}

//     // link node. add new curNode.
//     curP = addLink(newNode, &curP)

//     // update start node pointer.
//     start = curP
//   }

//   // return the nodes.
//   return start;
// }

func printList(initNode Node) {
  curNode := initNode
  fmt.Println(curNode)

  for {
    // if curNode != nil && curNode.v != nil {
    //   fmt.Println(curNode.v)
    // }
    if curNode.p == nil {
      fmt.Println(curNode.v);
      break;
    }

    fmt.Println(curNode.v);
    curNode = *curNode.p
  }
}

func main() {
  // node 1.
  foo1 := Node{nil, 2}
  foo2 := addLink(Node{nil, 3}, foo1)
  foo3 := addLink(Node{nil, 4}, *foo2)
  //fmt.Println(*foo2.p);

  // create list
  //result := firstCreateList(10)

  // print list
  printList(*foo3)

  //fmt.Println(&result)
}



/*
Move adding to linked list to a function.....
*/