package main

import "fmt"

func quickSort(a []int) []int {
  return quick2(a, 0, len(a)-1)
}

func quick2(a []int, low, high int) []int {
  aOut := a
  pi := 0

  // There is a check...
  if low < high {
    // partition.
    pi, aOut = partition(a, low, high);

    // call a quick sort off the partitioned items.
    aOut = quick2(aOut, low, pi - 1);  // Before pi
    aOut = quick2(aOut, pi + 1, high); // After pi
  }

  return aOut
}

/*

*/
func partition(a []int, low, high int) (int, []int) {
  pivot := a[high]
  
  // why -1
  i := low - 1

  for j := low ; j <= high-1 ; j++ {
    if a[j] <= pivot {
      i++
      a[i], a[j] = a[j], a[i]
    }
  }

  // final swap
  a[i+1], a[high] = a[high], a[i+1]
  fmt.Println(i+1, ">", a)
  return i+1, a
}

func main() {
  //p := partition([]int{5,4,1,2,3}, 0, 4)
  p := quickSort([]int{10,1,2,20,11})
  fmt.Println(p)
}

/*
  This is done...
  probably study this...

  Partition
  Note that the parition pivot is different from the quick2 pivot.
  You create another pivot.

  This is going 

  http://quiz.geeksforgeeks.org/quick-sort/
*/