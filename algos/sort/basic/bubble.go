package main

import "fmt"

func bubble(vals []int) ([]int) {
	valsLen := len(vals)
	if valsLen <= 1 {
		return vals
	}

	for i:=0 ; i<valsLen ; i++ {
		for j:=i+1 ; j<valsLen ; j++ {
			if vals[i] > vals[j] {
				temp := vals[i]
				vals[i] = vals[j]
				vals[j] = temp
			}
		}
	}

	return vals
}

func main() {
	fmt.Println(bubble([]int{3,2,1}))
	fmt.Println(bubble([]int{2,1}))
	fmt.Println(bubble([]int{1,2}))
	fmt.Println(bubble([]int{1}))
	fmt.Println(bubble([]int{1,2,3}))
}

/*

2 errors
1 more...
slow down a bit more...
on the checking...
these are simple errors..

*/