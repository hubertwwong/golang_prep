package main

import "fmt"

type Node struct {
  left *Node
  right *Node
  val int
}

func newNode(val int) (*Node) {
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

func insertList(valList []int) (*Node) {
  var root *Node
  root = nil

  for i:=0 ; i<len(valList) ; i++ {
    root = insert(valList[i], root)
  }

  return root
}

func createDoubleSlice(n int) ([][]int) {
	nDouble := make([][]int, 1)
	nDouble[0] = []int{n}
	return nDouble
}

func appendStart(nVal int, v []int) ([]int) {
	newS := []int{nVal}
	if v == nil {
		return newS
	} else {
		return append(newS, v...)
	}
}

func mergeSlices(s1 [][]int, s2 [][]int) ([][]int) {
	if s1 == nil {
		return s2
	} else if s2 == nil {
		return s1
	}

	finalLen := len(s1) + len(s2)
	finalS := make([][]int, finalLen)

	mergedI := 0
	// left side
	for i:=0 ; i<len(s1) ; i++ {
		finalS[mergedI] = s1[i]
		mergedI++
	}

	// count does not reset.

	// right side
	for i := 0 ; i<len(s2) ; i++ {
		finalS[mergedI] = s2[i]
		mergedI++
	}

	return finalS
}

func printPaths(root *Node) ([][]int) {
  if root == nil {
		return nil
	} else if (*root).left == nil && (*root).right == nil {
		// child.
		v := (*root).val
		i := []int{v}
		j := [][]int{i}
		return j
	}

	left := printPaths((*root).left)
	right := printPaths((*root).right)
	
	merged := mergeSlices(left, right)
	val := (*root).val

	for i := 0 ; i<len(merged) ; i++ {
		merged[i] = appendStart(val, merged[i])
	}
	// append on each slice.
	// return the final result.
	return merged
}

// helper.


// wrapping stuff in test cases.
func testRun(vals []int) {
  root := insertList(vals)
  result := printPaths(root)
  fmt.Println(">", result)
}



func main() {
  testRun([]int{2,1,3})
  testRun([]int{20, 10, 30, 5, 15, 25, 35})
  // testRun(4, []int{2,1,3})
  // testRun(0, []int{2,1,3})
	
	//out := appendStart(4, []int{1,2,3})
	//fmt.Println(out)

	//out2 := createDoubleSlice(3)
	// out3 := createDoubleSlice(4)
	// out4 := mergeSlices(nil, out3)
	// fmt.Println(out4)
}