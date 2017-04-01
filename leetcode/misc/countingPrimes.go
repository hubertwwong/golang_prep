import "math"

func trailingZeroes(n int) int {
    cTwos := 0
    cFives := 0
    nTwos := 0
    nTwosPow := 0
    nFives := 0
    nFivesPow := 0
    nTens := 0
    nTensPow := 0
    
    for i:=0 ; i<=n ; i++ {
        if i % 10 == 0 {
            nTens++
            // calc n2 pow
            if i / int(math.Pow(float64(10), float64(nTensPow+1))) == 1 {
                // incremnts of 10^n
                nTensPow++
                cTwos += nTwosPow * nTensPow
                cFives += nFivesPow * nTensPow
            } else if i % 50 == 0 {
                //increments of 50
                cTwos += nTwosPow
                cFives += nFivesPow * 2
            } else if i % 20 == 100 {
                // incremnts of 20
            }else {
                cTwos += nTwosPow
                cFives += nFivesPow
            }
        } else if i % 2 == 0 {
            nTwos++
            // calc n2 pow
            if i / int(math.Pow(float64(2), float64(nTwosPow+1))) == 1 {
                nTwosPow++
            }
            //fmt.Println(i, nTwos, nTwosPow)
            cTwos += nTwosPow
        } else if i % 5 == 0 {
            nFives++
            // calc n2 pow
            if i / int(math.Pow(float64(5), float64(nFivesPow+1))) == 1 {
                nFivesPow++
            }
            fmt.Println(i, nFives, nFivesPow)
            cFives += nFivesPow
        }
    }
    
    // return
    if cFives < cTwos {
        return cFives
    }
    return 0
}
/*

10 = 1 * 10
20 = 2 * 10
30 = 3 * 10
40 = 4 * 10
50 = 5 * 10
60
70
80
90
100 10 * 10
150 15 * 10
200 20 * 10
250 25 * 10
300 30 * 5

*/