package main

import "fmt"

func preSuffix(s string) []int {
	if len(s) == 0 {
		return nil
	}
	
	pmt := make([]int, len(s))

	for i:=1 ; i<len(s) ; i++ {
		//fmt.Println("")
		//fmt.Println(">i", i)
		preE := 0
		sufS := i
		max := 0

		// iterate thru all prefix suffix
		for {
			// prefix suffix for some length
			for j,k,curMax := 0, sufS, 0 ; ; j,k = j+1, k+1 {
				//fmt.Println("jk",j,k)
				if s[j] == s[k] {
					curMax++
				} else {
					break
				}
				
				if j>=preE {
					if curMax > max {
						max = curMax
					}
					break
				}
			}

			// increment
			preE++
			sufS--

			// you don't want to go full on.
			//fmt.Println(">E", preE, sufS)
			if sufS <= 0 {
				break
			}
		} 

		// stash the value into pmt
		pmt[i] = max
	}

	return pmt
}

func subString(longS, shortS string) int {
	if len(longS) < len(shortS) {
		return -1
	} else if len(shortS) == 0 || len(longS) == 0 {
		return -1
	}
	pmt := preSuffix(shortS)

	pos := -1
	for i:=0 ; i<len(longS) ; i++ {
		fmt.Println("i", i)
		// just skip to next char.
		//if longS[i] != shortS[0] {
		//	fmt.Println("i>")
		//	continue
		//}
		
		// figure out if character matches
		numCharsMatched := 0
		for j,k:=i,0 ; k<len(shortS) && j<len(longS) ; j,k=j+1,k+1 {
			if longS[j] == shortS[k] {
				numCharsMatched++
			} else {
				break
			}
		}

		// return
		if numCharsMatched == len(shortS) {
			return i
		}

		// skip
		i = i + pmt[numCharsMatched]
	}

	return pos
}

func main() {
	//fmt.Println(preSuffix("abababca"))
	fmt.Println(subString("cat in the hat", "e h"))
}