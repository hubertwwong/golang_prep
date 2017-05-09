package main

import "fmt"

func threeSumClosest(nums []int, target int) int {
    bestDiff := target - (nums[0] + nums[1] + nums[2])
    bestSum := nums[0] + nums[1] + nums[2]
    
    if len(nums) == 3 {
        return bestSum
    }
    
    for i:= 0 ; i<len(nums)-2 ; i++ {
        for j:=i+1 ; j<len(nums)-1 ; j++ {
            for k:=j+1 ; k<len(nums) ; k++ {
                curDiff := target - (nums[i] + nums[j] + nums[k])
                //curSum := nums[i] + nums[j] + nums[k]
                //fmt.Println(i,j,k, ">", curSum, ">", curDiff, bestDiff)
                if smallerDist(curDiff, bestDiff) {
                    bestDiff = curDiff
                    bestSum = nums[i] + nums[j] + nums[k]
                }
            }
        }
    }
    
    return bestSum
}

// if a is smaller than b distance wise, return true
func smallerDist(a, b int) bool {
  finalA := a
  finalB := b

  if a <= 0 {
    finalA = finalA - (2*a)
  }
  if b <= 0 {
    finalB = finalB - (2*b)
  }

  return finalA < finalB
}

func main() {
  //fmt.Println(threeSumClosest([]int{1,1,1,0}, -100))
  fmt.Println(threeSumClosest([]int{-1,2,1,-4}, 1)) // 2
}
// [-1,2,1,-4]
//1


/*
the 2 big things
typo with i and 1
not diffing....
*/