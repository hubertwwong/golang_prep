package main

import "fmt"

func mergeSort(vals []int) ([]int) {
	valsLen := len(vals)
	if valsLen == 0 {
		return nil
	} else if valsLen == 1 {
		return vals
	} else if valsLen == 2 {
		if vals[0] > vals[1] {
			temp := vals[0]
			vals[0] = vals[1]
			vals[1] = temp
		}
		return vals
	}

	// split list and merge.
	//fmt.Println(valsLen/2)
	//fmt.Println("l", vals[:(valsLen/2)])
	//fmt.Println("r", vals[(valsLen/2):])
	//fmt.Println(vals)
	left := mergeSort(vals[:(valsLen/2)])
	right := mergeSort(vals[(valsLen/2):])
	vals = sortedMerge(left, right)
	return vals
}

// assume sorted merge array.
func sortedMerge(s1, s2 []int) ([]int) {
	lenS1 := len(s1)
	lenS2 := len(s2)

	if lenS1 == 0 {
		return s2
	} else if lenS2 == 0 {
		return s1
	}

	resultA := make([]int, lenS1 + lenS2)
	lenResult := lenS1 + lenS2
	s1i := 0
	s2i := 0
	for i:=0 ; i<lenResult ; i++ {
		if s1i == lenS1 {
			resultA[i] = s2[s2i]
			s2i++
		} else if s2i == lenS2 {
			resultA[i] = s1[s1i]
			s1i++
		} else if s1[s1i] > s2[s2i] {
			resultA[i] = s2[s2i]
			s2i++
		} else {
			resultA[i] = s1[s1i]
			s1i++
		}
	}

	return resultA
}

func main() {
	fmt.Println(mergeSort([]int{3,2,1}))
	//fmt.Println(mergeSort([]int{3,2,1,4,5}))
	//fmt.Println(mergeSort([]int{1}))
	//fmt.Println(mergeSort([]int{2,1}))
	fmt.Println(mergeSort([]int{9,1,2,4,5,10}))
}



/*

Implement a merge sort...
01:00p
a few errors...
one... wrong return statement...
- always check it.
two...
try to write a more elegant solution.
- and got some errors. compressing if else statements..
- and the order is wrong.
three.
- code does not work...
i think the third issue is a mis understanding of slices....
- i dont' think its start and end syntax.
its start and number of elements..
i'm unsure of integer division...
but i got the algo to be clean....

short answer...
for slices
first arg is 0 index position
2nd arg is number of items.
for the end of the array
7/2 = 3.5 or 3
first 3 items.
4 / 2 = 2
a[len/2:]
a[:len/2]
this worsk for both odd and even
you don't have to do the +1 syntax like arrays.
note this is just for slices...

*/