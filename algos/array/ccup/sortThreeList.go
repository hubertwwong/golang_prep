package main

import "fmt"

// 11:15a

// do 2 not 3...
func merge2(a1, a2 []int) ([]int) {
	lenA1 := len(a1)
	lenA2 := len(a2)

	// preconditions
	if lenA1 == 0 && lenA2 == 0 {
		return nil
	} else if lenA1 == 0 {
		return a2
	} else if lenA2 == 0 {
		return a1
	}

	finalArray := make([]int, lenA1+lenA2)
	lenFinalArray := lenA1 + lenA2

	i1:=0
	i2:=0
	for i:=0 ; i<lenFinalArray ; i++ {
		if i1 == lenA1 {
			finalArray[i] = a2[i2]
			i2++
		} else if i2 == lenA2 {
			finalArray[i] = a1[i1]
			i1++
		} else if a2[i2] < a1[i1] {
			finalArray[i] = a2[i2]
			i2++
		} else {
			finalArray[i] = a1[i1]
			i1++
		}
	}

	return finalArray
}

func merge3(a1, a2, a3 []int) ([]int) {
	finalArray := merge2(a1, a2)
	finalArray = merge2(finalArray, a3)
	return finalArray
}

func main() {
	a1 := []int{1,2,3,4,5,20}
	a2 := []int{6}
	//a3 := []int{7}
	fmt.Println(merge3(a1,a2,a3))
}

/*

kinda short... trying to do this once without errors...
https://careercup.com/question?id=5689898495377408
merge 3 sorted list into 1...

so lessons...
did a a redo from 3 to 2 to simplify logic.
didin't check..
but got some syntax and logic errors when switch...
need to force myself to do a once over...
before hitting compile...

*/