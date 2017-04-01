package main

import "fmt"

func maxProfits(prices []int) int {
	lenPrices := len(prices)
	if lenPrices <= 1 {
		return 0
	}

	bc := prices[0]
	bb := prices[0]
	sc := -1
	sb := -1

	for i:=0 ; i<lenPrices ; i++ {
		cp := prices[i]

		// current price less than buy current.
		// could be a new buy price
		if cp <= bc {
			// both sell values not set. just set a new buy price.
			if sb == -1 && sc == -1 {
				bc = cp
				bb = bc
			// sb is set and buy best is greater than current price
			// new possible buy price.
			} else if sb != -1 && bb >= cp {
				// did you find a new best.
				if (sb - bb) < (sc - bc) {
					bb = bc
					sb = sc
				}

				// set a new current buy/sell current.
				bc = cp
				sc = -1
			} 
		// current prices > sell current
		// new possible high.
		} else if cp > sc {
			sc = cp
			// if sell best was never set, just set it.
			if sb == -1 {
				sb = sc
			}
		}
	}

	// return max
	if sb == -1 {
		return 0
	} else if sc == -1 {
		return sb - bb
	} else if (sb - bb) > (sc - bc) {
		return sb - bb
	} else {
		return sc - bc
	}
}

func main() {
	//fmt.Println(maxProfits([]int{1,2}))
	//fmt.Println(maxProfits([]int{3,2,1}))
	//fmt.Println(maxProfits([]int{1,3,2,4}))
	//fmt.Println(maxProfits([]int{1,3,2,4,1,100}))
	fmt.Println(maxProfits([]int{7,5,10,3}))
	fmt.Println(maxProfits([]int{7,2,100,1,3,500}))
}

/*

aaa

buy sell
11:41p
the first 3 examples are good.
and no compile errors
12:32a
still have a logic bug...

*/