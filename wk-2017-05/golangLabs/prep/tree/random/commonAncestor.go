package main

import "fmt"

// STACK STUFF

type NodeS struct {
  next *NodeS
  val *NodeT
}

func newNodeS(val *NodeT) (*NodeS) {
  n := NodeS{nil, val}
  return &n
}

func push(val *NodeT, root *NodeS) (*NodeS) {
  //fmt.Println("> push >", val)
  if root == nil {
    return newNodeS(val)
  }
  
  newRoot := newNodeS(val)
  (*newRoot).next = root
  return newRoot
}

// returns 2 things
// 1st is the top node. (The popped off item)
// 2nd is the new root. (The rest of the list)
func pop(root *NodeS) (*NodeS, *NodeS) {
  if root == nil {
    return nil, nil
  }

  newRoot := (*root).next
  // clean up..
  curNode := root
  (*curNode).next = nil
  
  return curNode, newRoot
}



// TREE STUFF

type NodeT struct {
  left *NodeT
  right *NodeT
  val int
}

func newNodeT(val int) (*NodeT) {
  n := NodeT{nil, nil, val}
  return &n  
}

func insert(val int, root *NodeT) (*NodeT) {
  if root == nil {
    return newNodeT(val)
  }

  if val <= (*root).val {
    (*root).left = insert(val, (*root).left)
  } else {
    (*root).right = insert(val, (*root).right)
  }

  return root
}

func insertList(vals []int) (*NodeT) {
  valsLen := len(vals)
  
  if valsLen <= 0 {
    return nil
  }

  var root *NodeT
  for i:=0 ; i<valsLen ; i++ {
    root = insert(vals[i], root)
  }
  return root
}

// main problem...

// this take a node.
// push the path on to the stack.
// can you assume its in the tree?
// nil for the initial path.
func dfsPath(root, targetN *NodeT, path *NodeS) (*NodeS) {
  if root == nil || targetN == nil {
    return nil
  }

  fmt.Println("> dfs > start path", path)
  fmt.Println("> dfs > root", root)
  fmt.Println("> dfs > target", targetN)

  // base correct case
  if (*root).val == (*targetN).val {
    fmt.Println("> dfs > FOUND", path)
    return path
  }

  // push current Node into path
  path = push(root, path)

  // figure out what path to take
  var cPath *NodeS
  if (*targetN).val <= (*root).val {
    //fmt.Println("> dfs > left")
    cPath = dfsPath((*root).left, targetN, path)
  } else {
    //fmt.Println("> dfs > right")
    cPath = dfsPath((*root).right, targetN, path)
  }

  return cPath
}


// finds a number in the tree and returns a node
func dfsFind(val int, root *NodeT) (*NodeT) {
  if root == nil {
    return nil
  }

  if val == (*root).val {
    return root
  } else if val < (*root).val {
    return dfsFind(val, (*root).left)
  } else {
    return dfsFind(val, (*root).right)
  }
}

func commonAncestor(root, c1, c2 *NodeT) (*NodeT) {
  var s1, s2 *NodeS

  fmt.Println("> DFS 1", c1)
  c1s := dfsPath(root, c1, s1)
  fmt.Println("> DFS 2", c2)
  c2s := dfsPath(root, c2, s2)

  //fmt.Println("> CA > C1s", c1s)
  //fmt.Println("> CA > C2s", c2s)

  var c1n *NodeS
  var c2n *NodeS

  for {
    // start to pop off stuff
    // only pop off if its not nil.
    if c1s != nil {
      c1n, c1s = pop(c1s)
    }
    if c2s != nil {
      c2n, c2s = pop(c2s)
    }

    // fmt.Println("> CA > C1s", c1s)
    // fmt.Println("> CA > C2s", c2s)

    // if both stacks are nil
    // return nil. did find anything...
    if c1s == nil && c2s == nil {
      // edge case to check
      if c1n != nil && c2n != nil {
        c1TreeNode := (*c1n).val
        c2TreeNode := (*c2n).val

        // if the values of the stack match, they have the same ancestor...
        // does not matter which one it returns.
        if (*c1TreeNode).val == (*c2TreeNode).val {
          return c1TreeNode
        }
      } else {
        return nil
      }
    } else if c1n != nil && c2n != nil {
      c1TreeNode := (*c1n).val
      c2TreeNode := (*c2n).val

      // if the values of the stack match, they have the same ancestor...
      // does not matter which one it returns.
      if (*c1TreeNode).val == (*c2TreeNode).val {
        return c1TreeNode
      }
    }
  }

  return nil
}



// helper...
func testRun(vals []int, vL, vR int) {
  t := insertList(vals)

  // dfs find.
  //r1 := dfsFind(3, t)
  //fmt.Println(r1)

  //lNode := (*(*t).left).left
  //rNode := (*(*t).left).right
  //lNode = nil
  lNode := dfsFind(vL, t)
  rNode := dfsFind(vR, t)
  // rNode = nil

  fmt.Println("> LN", lNode)
  fmt.Println("> RN", rNode)
  //fmt.Println("> testRun >", t, lNode, rNode)
  m := commonAncestor(t, lNode, rNode)
  fmt.Println("> final >", m)
}

func main() {
  //testRun([]int{2,1,3})
  testRun([]int{20,30,10,5,15}, 5, 15)
}


/*

given 2 nodes
find the common ancestory...

so lets start
03:37p
03:45p - 2mins...
03:51p stack done.
03:56p tree basics done...
scan again.
- found a few error on stack.
- using older var names... wrong var names..
nothign on tree. so i hope..
03:59p on the main problem...
04:18p dfs path is done...
04:22 stop for allen...
04:31 back....
04:37 testing again...
04:40 back again..
04:45 back again again...
04:46p done...
not with the main stuff
04:52p
- seeing if the thing runs....
04:53 debugging.....
04:59p nil....
should return 1....
the actual syntax erros were not too bad....
assigning things twich...
05:00p
05:09 tree insertion seem ok..
got the conditional of the dfs backwards...
the value check should always be on the left..
but i think it works....
05:18p
one case works.... but final they always point to root.
05:28p test runner was not working.
but it works....
05:41 another bug...
i'm assuming both stack are of the same length
but i can't assume that...

05:54p
that was the edge case i was missing....
if the stacks were not the same...
and the way i wrote it was bad...

*/