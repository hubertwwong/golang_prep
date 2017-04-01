package main

import "fmt"

func selectSort(vals []int) ([]int) {
	valsLen := len(vals)
	if valsLen <= 1 {
		return vals
	}

	for i:=0 ; i<valsLen ; i++ {
		minPos := -1
		minVal := vals[i]

		//fmt.Println(vals, "i", i, "j", i+1, valsLen, (i+1<valsLen))
		//fmt.Println(i,i+i,(i+1)<valsLen)
		for j:=i+1 ; j<valsLen ; j++ {
			fmt.Println(">in")
			//fmt.Println(">", minVal, vals[j], "i", i, "j", j)
			if minVal > vals[j] {
				minVal = vals[j]
				minPos = j
			}
		}

		// swap is required?
		if minPos >= 0 {
			temp := vals[minPos]
			vals[minPos] = vals[i]
			vals[i] = temp
		}
	}

	return vals
}

func main() {
	//fmt.Println(selectSort([]int{1,2,3}))
	//fmt.Println(selectSort([]int{2,1}))
	//fmt.Println(selectSort([]int{3,2,1}))
	//fmt.Println(selectSort([]int{1}))
	fmt.Println(selectSort([]int{4,1,2,3}))
}


/*

no algo..
go error but algo wrong..
taking way too long for this..

*/