package main

import "fmt"

func quicksort(vals []int) ([]int) {
  valsLen := len(vals)
  if valsLen <= 1 {
    return vals
  }

  return quick2(vals, 0, (valsLen-1))
}

// p is start and q is end.
func quick2(vals []int, p, q int) ([]int) {
  fmt.Println(">a", vals, p, q, vals[q])
  valsLen := len(vals)
  if valsLen < 1 {
    return vals
  }

  i := p
  j := q-1
  pivot := q
  for ; i<j ; {
    fmt.Println(i, j, ">>", vals[i], vals[j], ">>", pivot)
    
    if vals[i] < vals[pivot] && i+1 != j {
      fmt.Println("i", i)
      i++
      continue
    }
    if vals[j] > vals[pivot] && i+1 != j {
      fmt.Println("j", j)
      j--
      continue
    }

    // swap
    if vals[i] > vals[j] {
      temp := vals[i]
      vals[i] = vals[j]
      vals[j] = temp
    }

     if i+1 == j {
       break
     }
  }

  fmt.Println(">b", vals, i, j)

  // final swap
  vals[pivot], vals[i] = vals[i], vals[pivot]

  fmt.Println("> c", vals, i, j)

  // PIVOT..
  // not sure but i think you need a check to see if p<q
  fmt.Println("> p i >", p, i, "> j q >",j, q)
  if p < i {
    quick2(vals, p, i)
    
  }
  if j < q {
    quick2(vals, j, q)
  }

  return vals
}

func main() {
  fmt.Println(quicksort([]int{6,5,4,3,2,1}))
}

/*
02:04
quick sort attempt...
syntax error in params
wrong type.
typo on vals and vals len...

*/