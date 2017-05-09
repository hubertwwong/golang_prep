package main

import "fmt"

func sumTwoVals(vals []int, sum int) (map[int]int) {
  valsLen := len(vals)
  if valsLen <= 1 {
    return nil
  }

  valsHash := valsToHash(vals)
  result := make(map[int]int)

  for i := 0 ; i<valsLen ; i++ {
    // one simple check for if the value in the list is equal to the sum
    if vals[i] == sum {
      continue
    }

    diff := sum - vals[i]
    if valsHash[diff] == 1 {
      //fmt.Println(i,vals[i],diff)
      if result[vals[i]] == 0 && result[diff] == 0 {
        result[vals[i]] = diff
      }
    }
  }

  return result
}

func valsToHash(vals []int) (map[int]int) {
  valsLen := len(vals)
  if valsLen == 0 {
    return nil
  }

  // this is how to make maps.
  vh := make(map[int]int)
  for i:=0 ; i<valsLen; i++ {
    vh[vals[i]] = 1
  }
  return vh
}

func main() {
  fmt.Println(sumTwoVals([]int{1,2,3}, 4))
  fmt.Println(sumTwoVals([]int{1,2,3}, 100))
  fmt.Println(sumTwoVals([]int{1,2,3,4}, 5))
  fmt.Println(sumTwoVals([]int{7,2,3,4,10,1}, 5))
   fmt.Println(sumTwoVals([]int{-1,0,1}, 0))
}

/*

10:54p

i'm in the right ballpark......

2 errors so far.
one is handling larger case.
one is duplicates.
you can't assume the items are in the first half of the list. (You can only assume this if the list is sorted... even then.)
i'm getting blanks...
not sure why.
the blanks are assuming the i is the same....
figure out how to append on to a slice... this was problem 1...
- result = result[:j] this can trim the slice...
that or you can append one on one...

you should store the output in a hash too...
then you iterate on the hash...
11:24a

converted the output array to a hashmap.
and that gets this to a on operation.
so this is a constant operation ...

there is a bug about picking itself...


*/