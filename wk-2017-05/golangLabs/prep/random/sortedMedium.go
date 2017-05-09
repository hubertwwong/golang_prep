package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    mergedNums := sortedMerge(nums1, nums2)
    //fmt.Println(mergedNums)
    lenMerged := len(mergedNums)
    
    // even or odd.
    if lenMerged % 2 == 0 {
        v1 := float64(mergedNums[lenMerged/2 -1])
        v2 := float64(mergedNums[lenMerged/2])
        fmt.Println(v1, v2, (v1 + v2)/2.0)
        return float64((v1 + v2)/2.0)
    } else {
        return float64(mergedNums[lenMerged/2])
    }
}

func sortedMerge(nums1, nums2 []int) []int {
    lenNums1 := len(nums1)
    lenNums2 := len(nums2)
    lenNums3 := lenNums1 + lenNums2
    
    // some optimizations...
    if lenNums1 == 0 {
        return nums2
    } else if lenNums2 == 0 {
        return nums1
    }
    nums3 := make([]int, lenNums3)
    i1:=0
    i2:=0
    
    for i:=0 ; i<lenNums3 ; i++ {
        // list 1 is done.
        if i1 == lenNums1 {
            nums3[i] = nums2[i2]
            i2++
        // list 2 is done
        } else if i2 == lenNums2 {
            nums3[i] = nums1[i1]
            i1++
        // list 1 < list 2
        } else if nums1[i1] < nums2[i2] {
            nums3[i] = nums1[i1]
            i1++
        // list 2 < list 1
        } else {
            nums3[i] = nums2[i2]
            i2++
        }
    }

    return nums3
}

func main() {
  fmt.Println(findMedianSortedArrays([]int{}, []int{2,3}))
}