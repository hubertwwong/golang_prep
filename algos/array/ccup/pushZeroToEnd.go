package main

import "fmt"

func pushZeros(vals []int) ([]int) {
	// 2:18
	valsLen := len(vals)
	if valsLen <= 1 {
		return vals
	}

	i:=0
	j:=valsLen-1
	for ; i<valsLen && j>0 ; {
		// find first insert location
		for ; j>0 ; j-- {
			if vals[j] != 0 {
				break
			}
		}

		// find the first zero
		for ; i<valsLen ; i++ {
			if vals[i] == 0 && i<j {
				temp := vals[i]
				vals[i] = vals[j]
				vals[j] = temp
				break
			}
		}
	}

	return vals
}

func main() {
	fmt.Println(pushZeros([]int{1,2,3}))
	fmt.Println(pushZeros([]int{1}))
	fmt.Println(pushZeros([]int{}))
	fmt.Println(pushZeros([]int{0,1,2,3}))
	fmt.Println(pushZeros([]int{1,2,3,0}))
	fmt.Println(pushZeros([]int{1,2,0,3,0}))
	fmt.Println(pushZeros([]int{1,2,0,0}))
	fmt.Println(pushZeros([]int{0,0,0,1}))
}



/*
https://www.careercup.com/question?id=5767394003779584

push some zeros to the end of an array...
12/25/2016
02:01p
02:04 start

push start of zero to end...
think before you start...
00000000000
done
0100000000

01000
01
10000

12345
idea being that you swap on the right?
- swap zero for non zero...

012345
512340

10230
13200
on the end you need to find the first non zero...
this has some nunance..

not too bad...
but variable name mis spelled
variables not assigned..
but other than that it was good..

02:35p

*/