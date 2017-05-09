package main

import "fmt"

func maxSubArray(nums []int) int {
  if len(nums) == 0 {
    return 0
  } else if len(nums) == 1 {
    return nums[0]
  }

  // my guess is that you don't need either array.
  bestSoFar := make([]int, len(nums))
  startPos := make([]int, len(nums))
  

  // init col 0. basically if you only have one thing, thats the best solution so far.
  bestSoFar[0] = nums[0]
  startPos[0] = 0
  
  // dyniamically find the best contigious sum.
  // we start at 1 since we started at zero.
  curBestSum := 0
  curBestPos := 0
  bestSoFarI := nums[0] // current best solution.
  
  for i := 1 ; i < len(nums) ; i++ {
    // you need to be a bit more discreet about this.
    if nums[i] > curBestSum {
      curBestSum += nums[i]
    } else {
      curBestSum
    }

    fmt.Println("\n======= i", i, "nums[i]" ,nums[i])
    fmt.Println(">b", bestSoFar)
    fmt.Println(">s", startPos)
    fmt.Println("nums[i]", nums[i], "bestSoFar[i-1]", bestSoFar[i-1], "bestSoFarI", bestSoFarI, "curBestSum", curBestSum)
    
    if nums[i] > bestSoFar[i-1] {
      fmt.Println("> better, startPos i", i-1, "is", (startPos[i-1] + 1))

      if nums[i] > bestSoFarI + curBestSum {
        fmt.Println("> blowout")
        // the current value is better than existing solution. blow out old solution.
        bestSoFar[i] = nums[i]
        bestSoFarI = nums[i]
        curBestSum = 0
      } else {
        fmt.Println("> extend")
        bestSoFar[i] = bestSoFarI + curBestSum
        bestSoFarI = bestSoFarI + curBestSum
        startPos[i] = startPos[i-1] + 1
      }
    } else if bestSoFarI + curBestSum > bestSoFarI {
      fmt.Println("> improved")
      // this improves the list.
      bestSoFar[i] = bestSoFarI + curBestSum
      bestSoFarI = bestSoFarI + curBestSum
      startPos[i] = startPos[i-1] + 1
      curBestSum = 0
    } else {
      fmt.Println("> skip")
      bestSoFar[i] = bestSoFar[i-1]
      startPos[i] = startPos[i-1]
    }
  }

  // return the best possible sum.
  return bestSoFar[len(bestSoFar)-1]
}

func main() {
  var nums []int
  
  // 6
  // nums = []int{-2,1,-3,4,-1,2,1,-5,4}
  
  // 9
  //nums = []int{1,3,5}

  // 1
  // nums = []int{-2,1}

  // 21. pick the last 3
  nums = []int{8,-19,5,-4,20}

  // 110
  // nums = []int{1,2,-1,1,2,3,-1,1,2,100}

  fmt.Println("> output >", maxSubArray(nums))
}

/*
You do need a best pos...
a running best pos..

Passes a 150/200 case on leet code.
I think the issue is that i'm being greedy.
And thats leading to wrong solutions.


*/