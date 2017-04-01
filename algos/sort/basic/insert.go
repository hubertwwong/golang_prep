package main

import "fmt"

func insertSort(vals []int) ([]int) {
	valsLen := len(vals)
	if valsLen <= 1 {
		return vals
	}

	for i:=1 ; i<valsLen ; i++ {
		for j,k:=i-1,i ; j>=0 ; j,k=j-1,k-1 {
			if vals[j] > vals[k] {
				temp := vals[j]
				vals[j] = vals[k]
				vals[k] = temp
			}
		}
	}

	return vals
}

func main() {
	fmt.Println(insertSort([]int{1,2,3}) )
	fmt.Println(insertSort([]int{2,1,3}) )
	fmt.Println(insertSort([]int{1}))
	fmt.Println(insertSort([]int{2,1}) )
	fmt.Println(insertSort([]int{1,2}) )
}

/*

should be more vals..
basically check as you go..
algo is wrong....
when you go to the start of an array you are going to zero... >=0
instead of <len

*/