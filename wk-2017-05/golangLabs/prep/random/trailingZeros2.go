// memo hash
// func trailingZeroes(n int) int {
//     memo := make(map[int]int)
//     numZeros := 0
//     for i := 5 ; i <= n ; i += 5 {
//         // assuming that counting 5 factors is good enough.
//         if i%5 == 0 {
//             curNumZeros := 1
//             for curNum := i/5 ; curNum % 5 == 0 ; curNum /= 5 {
//                 if val, ok := memo[curNum]; ok {
//                   curNumZeros += val
//                   break
//                 } else {
//                   curNumZeros++
//                 }
//             }
//             numZeros += curNumZeros
//             // store it for later user.
//             memo[i] = curNumZeros
//         }
//     }
//     return numZeros
// }

// MEMO on array
// func trailingZeroes(n int) int {
//     memo := make([]int, n+1)
//     numZeros := 0
//     for i := 5 ; i <= n ; i += 5 {
//         // assuming that counting 5 factors is good enough.
//         if i%5 == 0 {
//             curNumZeros := 1
//             for curNum := i/5 ; curNum % 5 == 0 ; curNum /= 5 {
//                 if memo[curNum] != 0 {
//                   curNumZeros += memo[curNum]
//                   break
//                 } else {
//                   curNumZeros++
//                 }
//             }
//             numZeros += curNumZeros
//             // store it for later user.
//             memo[i] = curNumZeros
//         }
//     }
//     return numZeros
// }