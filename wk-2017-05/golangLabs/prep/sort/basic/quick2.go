package main

import "fmt"

func quicksort(vals []int) []int {
  valsLen := len(vals)
  if valsLen <= 1 {
    return vals
  }

  return quick2(vals, 0, (valsLen-1))
}

func quick2(vals []int, left, right int) []int {
  return nil
}

// partition based.
// this assumes the piviot is 1 right of the right index.
func partition(vals []int, leftI, rightI int) ([]int, int) {
  pivotI := rightI+1
  
  for i,j := leftI, rightI ; ; {
    // left. looking for item greater than pivot
    fmt.Println("i inc", i)  
    for ; i<j ; i++ {
      if vals[i] > vals[pivotI] {
        break
      } 
    }

    // right. looking for item less than pivot.
    fmt.Println("j dec", j)
    for ; i<j ; j-- {
      if vals[j] < vals[pivotI] {
        break
      } 
    }

    // swap items or exit if you are done
    fmt.Println("i and j final", i, j)
    if i<j {
      fmt.Println(i, j, "<" , vals[i], vals[j])
      vals[i], vals[j] = vals[j], vals[i]
      i++
      j--
    } else {
      // swap the partition index.
      //i++
      fmt.Println("partition swap", i, pivotI)
      vals[i], vals[pivotI] = vals[pivotI], vals[i]
      
      pivotI = i
      break
    }
  }

  return vals, pivotI
}

func main() {
  //fmt.Println(partition([]int{6,5,1,3,2,4}, 0, 4))
  fmt.Println(partition([]int{2,1}, 0, 0))
  // fmt.Println(quicksort([]int{6,5,4,3,2,1}))
}