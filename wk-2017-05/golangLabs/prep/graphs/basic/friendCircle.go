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



// FIND CIRCLE
// ===========================================================================

func findCircleNum(M [][]int) int {
  var uf UnionFind
  uf.init(len(M))

  // debug
  //printArrayI(M, "input")

  // probably have some exit condition.

  // union everyone in the set.
  // I think you only really have to scan down the bottom half of the screen.
  for i:=1 ; i<len(M) ; i++ {
    for j:=0 ; j<i ; j++ {
      //fmt.Println("> i", i, "j", j)
      if M[i][j] == 1 {
        uf.union(i, j)
      }
    }
  }

  // compress all of the parents.
  uf.findAll()

  // fmt.Println("> parent")
  // fmt.Println(uf.parent)
  // fmt.Println("> rank")
  // fmt.Println(uf.rank)

  return diffItems(uf.parent)
}

func fc1() {
  m := make([][]int, 0)
  m = append(m, []int{1,1,0})
  m = append(m, []int{1,1,1})
  m = append(m, []int{0,1,1})

  fmt.Println(findCircleNum(m))
}

func fc2() {
  m := make([][]int, 0)
  m = append(m, []int{1,1,0})
  m = append(m, []int{1,1,0})
  m = append(m, []int{0,0,1})

  fmt.Println(findCircleNum(m))
}

func fc3() {
  m := make([][]int, 0)
  m = append(m, []int{1,1,0,0,0,0,0,0,0,0,0,1})
  m = append(m, []int{1,1,1,0,0,0,0,0,0,0,0,0})
  m = append(m, []int{0,1,1,1,0,0,0,0,0,0,0,0})
  m = append(m, []int{0,0,1,1,1,0,0,0,0,0,0,0})
  m = append(m, []int{0,0,0,1,1,1,0,0,0,0,0,0})
  m = append(m, []int{0,0,0,0,1,1,1,0,0,0,0,0})
  m = append(m, []int{0,0,0,0,0,1,1,1,0,0,0,0})
  m = append(m, []int{0,0,0,0,0,0,1,1,1,0,0,0})
  m = append(m, []int{0,0,0,0,0,0,0,1,1,1,0,0})
  m = append(m, []int{0,0,0,0,0,0,0,0,1,1,1,0})
  m = append(m, []int{0,0,0,0,0,0,0,0,0,1,1,1})
  m = append(m, []int{1,0,0,0,0,0,0,0,0,0,1,1})


  fmt.Println(findCircleNum(m))
}

func fc4() {
  // this should return 3.
  m := make([][]int, 0)
  m = append(m, []int{1,1,0,0,1})
  m = append(m, []int{1,1,0,0,0})
  m = append(m, []int{0,0,1,0,0})
  m = append(m, []int{0,0,0,1,0})
  m = append(m, []int{1,0,0,0,1})

  fmt.Println(findCircleNum(m))
}



func main() {
  //u1()
  //fc1()
  //fc2()
  //fc3()
  fc4()
}

/*

a
a
a

https://leetcode.com/problems/friend-circles/#/description
Got the question to work. but its not efficient enough.....

optimization 2 was to hack a path compression on the union command.
basically on each union call, call find on a and b.
that brought the answer up to 11/113 to 59/113

i think the optimization is bugged...

sort of see the issue.
the rank should go up if its being pointed to...



GOOD

> i 1 j 0
> union b [1 1 2 3 4] [2 1 1 1 1]
> union f [1 1 2 3 4] [2 1 1 1 1]

> i 4 j 0
> union b [1 1 2 3 1] [2 1 1 1 3]
> union f [1 1 2 3 1] [2 1 1 1 3]

BAD

> i 1 j 0
> union b [1 1 2 3 4] [2 1 1 1 1]
> union f [1 1 2 3 4] [1 1 1 1 1]

> i 4 j 0
> union b [4 1 2 3 4] [2 1 1 1 1]
> union f [4 1 2 3 4] [1 1 1 1 1]

This works.
And copied this over to union find..
The big bugs were that you need to find the top parent of each tree when doing the union operation...

a
a
a
a

*/