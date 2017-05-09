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

// running the find command on all nodes.
// basically compacts everything down to 1 level.
func (uf *UnionFind) compact() {
  for i := 0 ; i<len(uf.parent) ; i++ {
    uf.find(i)
  }
}

// number of islands
// ===========================================================================

// given a point on the grid.
// calculate the legal pairs.
// return it as a list.
// using i as the other array j as the inner array.
func neighborsPairs(grid [][]byte, i, j int) [][]int {
  iMax := len(grid)
  jMax := len(grid[0])
  neigh := make([][]int, 0)
  
  if i+1 < iMax {
    neigh = append(neigh, []int{i+1, j})
  }
  if i-1 >= 0 {
    neigh = append(neigh, []int{i-1, j})
  } 
  if j+1 < jMax {
    neigh = append(neigh, []int{i, j+1})
  }
  if j-1 >= 0 {
    neigh = append(neigh, []int{i, j-1})
  }

  //fmt.Println(">", i, j, "|", neigh)
  return neigh
}

// an optimization might be to not calculate all the neighbors.
// hash of computed neighbors.
func numIslands(grid [][]byte) int {
  if len(grid) == 0 {
    return 0
  } else if len(grid[0]) == 0 {
    return 0
  }

  var uf UnionFind
  uf.init(len(grid) * len(grid[0]))
  
  // loop to figure out if items are connected. 
  // join them if they are.
  for i := 0 ; i < len(grid) ; i++ {
    for j := 0 ; j < len(grid[0]) ; j++ {
      neighPairs := neighborsPairs(grid, i, j)
      for k := 0 ; k < len(neighPairs) ; k++ {
        //fmt.Println("\n> i", i, "j", j)
        // if both values are the same, they are apart of the same set.
        iNeigh := neighPairs[k][0]
        jNeigh := neighPairs[k][1]
        if grid[i][j] == grid[iNeigh][jNeigh] {
          // covert the 2d array to 1d array.
          iPosInRow := i * len(grid[0]) + j
          jPosInRow := iNeigh * len(grid[0]) + jNeigh
          //fmt.Println("i", i, j, "=", iPosInRow)
          //fmt.Println("iN", iNeigh, jNeigh, "=", jPosInRow)
          uf.union(iPosInRow, jPosInRow)
          //fmt.Println("> i1 [", i, "] j1 [", j, "] >", grid[i][j], "| i2 [", iNeigh, "] j2 [", jNeigh, "] >", grid[iNeigh][jNeigh], "|", iPosInRow, jPosInRow)
        }
      }
    }
  }

  // compact the results.
  uf.compact()

  // figure out for each union if its a union of land or water.
  // count if it is.
  numIslands := 0
  islands := make(map[int]bool)
  for i := 0 ; i < len(uf.parent) ; i++ {
    iPos := i / len(grid[0])
    jPos := i % len(grid[0])
    parent := uf.parent[i]
    //fmt.Println(">>>", i, ">>>", iPos, jPos, "=", parent, ">", grid[iPos][jPos], ">", len(grid), len(grid[0]))
    if grid[iPos][jPos] == 1 {
      if _, ok := islands[parent]; !ok {
        //fmt.Println(">i", iPos, "j", jPos)
        islands[parent] = true
        numIslands++
      }
      //fmt.Println("> result > parent >", parent, ">", islands)
    }
    //fmt.Println("0"[0], "1"[1])
  }

  //fmt.Println(uf.parent)

  return numIslands
}

// problem setup.
func n1() {
  grid := make([][]byte, 0)
  grid = append(grid, []byte{1,1,0,0,0})
  grid = append(grid, []byte{1,1,0,0,0})
  grid = append(grid, []byte{0,0,1,0,0})
  grid = append(grid, []byte{0,0,0,1,1})

  fmt.Println("> num islands", numIslands(grid))
}

// problem setup.
func n2() {
  grid := make([][]byte, 0)
  grid = append(grid, []byte{1,1,1,1,0})
  grid = append(grid, []byte{1,1,0,1,0})
  grid = append(grid, []byte{1,1,0,0,0})
  grid = append(grid, []byte{0,0,0,0,0})

  fmt.Println("> num islands", numIslands(grid))
}

func n3() {
  grid := make([][]byte, 0)
  
  fmt.Println("> num islands", numIslands(grid))
}

func n4() {
  grid := make([][]byte, 0)
  grid = append(grid, []byte{1,0})
  grid = append(grid, []byte{0,1})
  
  fmt.Println("> num islands", numIslands(grid))
}

func n5() {
  grid := make([][]byte, 0)
  grid = append(grid, []byte{1,1})
  
  fmt.Println("> num islands", numIslands(grid))
}

func n6() {
  grid := make([][]byte, 0)
  grid = append(grid, []byte{1})
  grid = append(grid, []byte{1})
  
  fmt.Println("> num islands", numIslands(grid))
}

func main() {
  n3()
  n4()
  n5()
  n6()
}

/*

a
a

The basic idea here is that you find merge the islands over and over until you have a set...
Can you use x*y as the length...
Cound the number of islands.


check all 4 directions to merge...
convert a byte to int...
you have check for 0 and 1 and merge both..
you make another pass to check the solution.
you probably want a way to calc all legal edges?

a
a

Solution was accepted but slow..
They are doing 4ms
I'm doing 46ms....

*/