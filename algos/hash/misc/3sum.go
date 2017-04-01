package main

import "fmt"

func threeSum(nums []int) [][]int {
    sortedNums := mergeSort(nums)
		hashedNums := storeAsHash(sortedNums)
		
		//fmt.Println(">sn",sortedNums)
		//fmt.Println(">hn", hashedNums)

		result := make([][]int, 0)

		for i:=0 ; i<len(sortedNums)-2 ; i++ {
			for j:=i+1 ; j<len(sortedNums)-1 ; j++ {
				curResult := sortedNums[i] + sortedNums[j]
				k := 0 - curResult

				//fmt.Println(sortedNums[i], i, ">", sortedNums[j], j, ">", hashedNums[k], k)
				if hashedNums[k] != 0  && hashedNums[k] != i && hashedNums[k] != j {
					//fmt.Println(">inserting")
					r := minArray(sortedNums[i], sortedNums[j], sortedNums[hashedNums[k]])

					// insert if not in the result set.
					if !isInList(r, result) {
						result = append(result, r)
					}
				}
			}
		}		

    return result
}

func isInList(val []int, list [][]int) bool {
	//fmt.Println(list)
	for i:=0 ; i<len(list) ; i++ {
		if isSliceMatch(val, list[i]) {
			return true
		}
	}
	return false
}

func isSliceMatch(a, b []int) bool {
	if len(a) == 0 && len(b) == 0 {
		return false
	} else if len(a) != len(b) {
		return false
	}

	for i:=0 ; i<len(a) ; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// puts in array as a min array.
// this is really ugly.
func minArray(i, j, k int) []int {
	if i <= j && i <= k {
		if j < k {
			return []int{i, j, k}
		} else {
			return []int{i, k, j}
		}
	} else if j <= i && j <= k {
		if i < k {
			return []int{j, i, k}
		} else {
			return []int{j, k, i}
		}
	} else {
		if i < j {
			return []int{k, i, j}
		} else {
			return []int{k, j, i}
		}
	}
}

// stores values as a hash.
func storeAsHash(nums []int) map[int]int {
	numsLen := len(nums)
	
	result := make(map[int]int)
	for i:=0 ; i<numsLen ; i++ {
		result[nums[i]] = i
	}
	return result
} 

// merge sorting for problem
func mergeSort(nums []int) []int {
    numsLen := len(nums)
    if numsLen <= 1 {
        return nums
    }
    left := mergeSort(nums[:numsLen/2])
    right := mergeSort(nums[numsLen/2:])
		return merge2(left, right)
}

// merges 2 list.
func merge2(a, b[]int) ([]int) {
    lenA := len(a)
    lenB := len(b)
    
    lenResult := lenA + lenB
    if lenA == 0 {
        return b    
    } else if lenB == 0 {
        return a
    }
    
    result := make([]int, lenResult)
    ai := 0
    bi := 0
    for i:=0 ; i<lenResult ; i++ {
        if ai == lenA {
            result[i] = b[bi]
            bi++
        } else if bi == lenB {
            result[i] = a[ai]
            ai++
        } else if a[ai] < b[bi] {
            result[i] = a[ai]
            ai++
        } else {
            result[i] = b[bi]
            bi++
        }
    }
    
    return result
}

func main() {
	//fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	//fmt.Println(threeSum([]int{1,2,-2,-1}))
	r := []int{-7,2,1,10,9,-10,-5,4,13,-9,-4,6,11,-12,-6,-9,-6,-9,-11,-4,10,10,-3,-1,-4,-7,-12,-15,11,5,14,11,-7,-8,6,9,-2,9,-10,-12,-15,2,10,4,5,11,10,6,-13,6,-13,12,-7,-9,-12,4,-9,13,-4,10,4,-12,6,4,-5,-10,-2,0,14,4,4,6,13,-9,-5,-5,-13,12,-14,11,3,10,8,11,-13,4,-8,-7,2,4,10,13,7,2,2,9,-1,8,-5,-10,-3,6,3,-5,12,6,-3,6,3,-2,2,14,-7,-13,10,-13,-2,-12,-4,8,-1,13,6,-9,0,-14,-15,6,9}
	fmt.Println(threeSum(r))
}