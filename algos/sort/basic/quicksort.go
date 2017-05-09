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
  if(p<r)
    {
        int q = Partition(a, p,r);
        qSort(a, p, q-1);
        qSort(a, q+1, r);
    }
  }
}

// partition based.
// this assumes the piviot is 1 right of the right index.
func partition(vals []int, leftI, rightI int) ([]int, int) {
  if len(vals) <= 1 {
    return vals, len(vals)
  }
  // } else if len(vals) == 2 {
  //   if vals[0] > vals[1] {
  //     vals[0], vals[1] = vals[1], vals[0]
  //   }
  //   return vals, 1
  // }

  pivotI := rightI+1
  
  for i,j := leftI, rightI ; ; {
    // left. looking for item greater than pivot
    for ; i<j && i<len(vals) ; i++ {
      if vals[i] > vals[pivotI] {
        break
      } 
    }

    // right. looking for item less than pivot.
    for ; i<j && j>=0 ; j-- {
      if vals[j] < vals[pivotI] {
        break
      } 
    }

    // swap items or exit if you are done
    if i<j {
      fmt.Println(i, j, "<" , vals[i], vals[j])
      vals[i], vals[j] = vals[j], vals[i]
      i++
      j--
    } else {
      // swap the partition index.
      fmt.Println("> final swap >", j, pivotI)
      //i++
      fmt.Println("1", vals)
      vals[j], vals[pivotI] = vals[pivotI], vals[j]
      fmt.Println("2", vals)
      pivotI = i
      break
    }
  }

  return vals, pivotI
}

func main() {
  //fmt.Println(partition([]int{6,5,1,3,2,4}, 0, 4))
  fmt.Println(partition([]int{5}, 0, 0))
  fmt.Println(partition([]int{5,6}, 0, 0))
  fmt.Println(partition([]int{6,5}, 0, 0))
  
  // fmt.Println(quicksort([]int{6,5,4,3,2,1}))
}