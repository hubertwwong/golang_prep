package main

import "fmt"

func rowPrint(vals [][]int) {
  for k,v := range vals {
    fmt.Println(k, ">", v)
  }
}

// going to store it in reverse
// 0 place goes as the min.
// on the append fn back to a string, it will work.
func numStrToIntArray(num string) []int {
  row := make([]int, len(num))
  
  for k,v := range num {
    row[(len(num)-1)-k] = int(byte(v) - "0"[0])
  }

  return row
}

// convert int back to a string
// assume lowest place value in in the zero position.
// will do some slicing for zeros..
func numArrayToStr(nums []int) string {
  result := make([]byte, 0)

  // slice off initial zeros.
  numsFinal := nums
  for i:=len(nums)-1 ; i> 0 ; i-- {
    if nums[i] == 0 {
      numsFinal = numsFinal[:len(numsFinal)-1]
    } else {
      break
    }
  }

  // convert int back to a string.
  for i := 0 ; i < len(numsFinal) ; i++ {
    cur := byte(nums[(len(numsFinal)-1)-i]) + "0"[0]
    result = append(result, cur)
  }

  return string(result)
}

func multiply(num1 string, num2 string) string {
  rowLen := len(num1) + len(num2)
  num1IntA := numStrToIntArray(num1)
  num2IntA := numStrToIntArray(num2)
  mulA := make([][]int, 0)

  // debug
  //fmt.Println(num1IntA)
  //fmt.Println(num2IntA)

  // multiply. 
  // num1 on top. num2 on bottom.
  for i := 0 ; i < len(num2) ; i++ {
    curRow := make([]int, rowLen)
    // prepend zeros.
    // just use i to prepend.
    
    // optimization later is to cache this value.
    // for each bottom val, iter on each top value and multply.
    curOver := 0
    for j := 0 ; j < len(num1) ; j++ {
      curVal := num1IntA[j] * num2IntA[i] + curOver
      curOver = curVal/10
      curRow[i+j] = curVal%10
      
      // add in the last digit.
      if j + 1 == len(num1) {
        curRow[i+j+1] = curOver
      }
    }

    // append row that you multipled.
    mulA = append(mulA, curRow)
  }

  // adding rows of the initial mul.
  curColVal := 0
  curOver := 0
  resultInt := make([]int, 0)
  for i := 0 ; i < rowLen ; i++ {
    // adding all values of the column.
    for j := 0 ; j < len(mulA) ; j++ {
      curColVal += mulA[j][i]
    }
    // adding the overfloow
    curColVal += curOver

    // adding values up
    curOver = curColVal / 10
    //fmt.Println("curCol", curColVal, "curOver", curOver)
    resultInt = append(resultInt, curColVal%10)

    // reset
    curColVal = 0
  }
  
  for i := curOver ; i > 0 ; i /= 10 {
    resultInt = append(resultInt, i%10)
  }

  return numArrayToStr(resultInt)
}

func main() {
  //fmt.Println(multiply("9999999999999999999", "99999999999999999999999999999999999999999999999999999999999999"))
  fmt.Println(multiply("0", "0"))
}

/*
10:12
not sure how delcare 2d arrays as a var.
do you have to use make?

10:17
numStrToIntArray
for k,v on a string return runes... not bytes.
good to know.

10:24
thinking on how to multiply.

10:51p
10:57
on the adding the multiplciation.
11:18a on the addition part...

11:21 raw algo is done.
finish it first...
11:29 back

11:45am
got a string together
prepending zeros...

11:54 done.
9ms..


*/