package main

import "fmt"


func printArrayI(val [][]int, label string) {
  fmt.Println(">", label)
  for i := 0 ; i < len(val) ; i++ {
    for j := 0 ; j < len(val) ; j++ {
      fmt.Printf("%4d", val[i][j])
    }
    fmt.Println("");
  }
  fmt.Println("")
}

// count number of unique items using a hash.
func diffItems(val []int) int {
  hashed := make(map[int]bool)
  for _, v := range val {
    hashed[v] = true
  }

  return len(hashed)
}

// UNION FIND
// ===========================================================================

type UnionFind struct {
  parent []int
  rank []int
}

func (uf *UnionFind) init(val int) {
  uf.parent = make([]int, val)
  uf.rank = make([]int , val)

  // init the parents to the index.
  for i := 0 ; i < len(uf.parent) ; i++ {
    uf.parent[i] = i
  }
}

func (uf *UnionFind) findGP(val int) int {
  currentP := val
  topP := -1        // set it to a non legal value. in cause you want to debug.
  
  // first pass. find the top parent.
  // probably should do some error checking
  for {
    if uf.parent[currentP] != currentP {
      currentP = uf.parent[currentP]
    } else {
      topP = currentP
      break
    }
  }

  return topP
}

func (uf *UnionFind) union(a, b int) {
  aGParent := uf.findGP(a)
  aGPRank := uf.rank[aGParent]
  
  bGParent := uf.findGP(b)
  bGPRank := uf.rank[aGParent]

  // parents are different, you need a union op.
  if aGParent != bGParent {
    // join off the smaller rank to the larger rank to keep a balanced tree.
    if aGPRank < bGPRank {
      uf.parent[aGParent] = bGParent
      uf.rank[aGParent] = bGPRank
      
      //fmt.Println(">1", bRank + 1)
      uf.rank[bGParent] = uf.rank[aGParent] + 1
    } else {
      uf.parent[bGParent] = aGParent
      uf.rank[bGParent] = aGPRank + 1
      //fmt.Println(">2", aRank + 1, uf.rank[b], uf.rank)
      uf.rank[aGParent] = uf.rank[bGParent] + 1
    }
  }
}

func (uf *UnionFind) find(a int) int {
  currentP := a
  topP := -1        // set it to a non legal value. in cause you want to debug.
  
  // first pass. find the top parent.
  // probably should do some error checking
  for {
    if uf.parent[currentP] != currentP {
      currentP = uf.parent[currentP]
    } else {
      topP = currentP
      break
    }
  }

  // update cur node to the top parent.
  uf.parent[a] = topP
  uf.rank[a] = 0

  return topP
}

// conveinence function to compact all of the results down.
// basically calls find on every index.
func (uf *UnionFind) findAll() {
  //fmt.Println(">>>>>>>", uf.parent)
  for i := 0 ; i < len(uf.parent) ; i++ {
    uf.find(i)
  }
}

func u1() {
  var uf UnionFind
  uf.init(4)
  uf.union(2,3)
  fmt.Println(uf.find(3))
  
  fmt.Println(">parent")
  fmt.Println(uf.parent)
  fmt.Println(">rank")
  fmt.Println(uf.rank)
}

func u2() {
  var uf UnionFind
  uf.init(8)
  uf.union(1,2)
  uf.union(2,3)
  uf.union(4,5)
  uf.union(6,7)
  uf.union(5,6)
  // joining 2 larger sets is an issue....
  uf.union(3,7)
  
  //fmt.Println(uf.find(3))
  
  fmt.Println(">parent")
  fmt.Println(uf.parent)
  fmt.Println(">rank")
  fmt.Println(uf.rank)
}

func main() {
  u2()
}

/*

a
a

11:54a
I think this is good..
Has 2 basic optimizations.
1. Path collapse on the find command.
2. Union command favors small trees to connect to the larger tree for balance trees.

This should get you close to nlogn?

NOTE..
THIS HAS BUGS.
See friendCircle for a correct implementation...


a
a

*/