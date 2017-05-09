package main

import "fmt"

func sum2Array(a, b []int) ([]int) {
  lenA := len(a)
  lenB := len(b)

  if lenA == 0 && lenB == 0 {
    return nil
  } else if lenA == 0 {
    return b
  } else if lenB == 0 {
    return a
  }

  // define the output array.
  var c []int
  if lenA > lenB {
    c = make([]int, lenA)
  } else {
    c = make([]int, lenB)
  }

  if lenA < lenB {
    for i:=0 ; i<lenA ; i++ {
      c[i] = returnTopDigit(a[i] + b[i])
    }
    for i:=(lenB-lenA) ; i<lenB ; i++{
      c[i] = b[i]
    }
  } else if lenA > lenB {
    for i:=0 ; i<lenB ; i++ {
      c[i] = returnTopDigit(a[i] + b[i])
    }
    for i:=(lenA-lenB) ; i<lenA ; i++{
      c[i] = a[i]
    }
  } else if lenA == lenB {
    for i:=0 ; i<lenA ; i++ {
      c[i] = returnTopDigit(a[i] + b[i])
    }
  }

  return c
}

// this assume you are taking 2 numerical digits
func returnTopDigit(a int) (int) {
  if a >= 10 {
    return 1
  } else {
    return a
  }
}

func testRun(a, b []int) {
  c := sum2Array(a,b)
  fmt.Println(c)
}

func main() {
  testRun([]int{1,2,3}, []int{2,3,4})
  testRun([]int{1,2,6}, []int{2,3,4})
  testRun(nil, nil)
}

/*

https://www.careercup.com/question?id=5692987154628608
a few things...
forgot to array make...
probably should be suing this for slices...

and lots of basics...
so this is a good warmup...


*/