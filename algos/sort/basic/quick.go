package main

import "fmt"

func quickSort(vals []int) ([]int) {
	valsLen := len(vals)
	if valsLen == 0 {
		return vals
	} else if valsLen == 1 {
		return vals
	}

	// picking a pivot.
	// should pick random...
	pivot := vals[valsLen-1]
	fmt.Println("p",pivot)

	var smallerVals []int
	var largerVals []int
	for i:=0 ; i<valsLen ; i++ {
		if vals[i] <= pivot {
			smallerVals = append(smallerVals, vals[i])
		} else {
			largerVals = append(largerVals, vals[i])
		}
	}
	fmt.Println("split", smallerVals, largerVals)
	// p=1 1,2

	smallerVals = quickSort(smallerVals)
	largerVals = quickSort(largerVals)

	mergedVals := append(smallerVals, largerVals...)
	fmt.Println("merged", mergedVals)
	return mergedVals
}

func main() {
	//fmt.Println(quickSort([]int{3,2,1}))
	//fmt.Println(quickSort([]int{1}))
	fmt.Println(quickSort([]int{4,2,1,3}))
	//fmt.Println(quickSort([]int{6,1,2,4,3,5}))
}

/*
so algo fail?
and one typo..
but it runs..

the big...
the splits need to actually split..
the big thing...
i'm missing 2?
i missing piviots on a single list...

you are picking the largest element...
or you are picking the smallest element...

*/