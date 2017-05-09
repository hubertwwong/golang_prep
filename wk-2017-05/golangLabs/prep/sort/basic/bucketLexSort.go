package main

import "fmt"

func printBuckets(vals [][]string) {
  for i := 0 ; i < len(vals) ; i++ {
    fmt.Println(vals[i])
  }
}

func bucketLexSort(vals []string) [][]string {
  // figure out the bucket length.
  maxStrLen := 0
  for i := 0 ; i < len(vals) ; i++ {
    if len(vals[i]) > maxStrLen {
      maxStrLen = len(vals[i])
    }
  }

  // create the buckets
  buckets := make([][]string, maxStrLen+1)

  // stash strings into buckets.
  for i := 0 ; i < len(vals) ; i++ {
    fmt.Println()
    fmt.Println("> inserting >", vals[i])
    lenStr := len(vals[i])

    // bucket does not exist.
    if len(buckets[lenStr]) == 0 {
      fmt.Println("> bucket does not exist")
      curBucket := make([]string,1)
      curBucket[0] = vals[i]
      buckets[lenStr] = curBucket
      fmt.Println("> bucket does not exist > end >", buckets[lenStr])
    } else {
      fmt.Println("> bucket exist.")
      // append to the end.
      buckets[lenStr] = append(buckets[lenStr], vals[i])

      // swap to keep it in lex order.
      for j := len(buckets[lenStr])-1 ; j > 0 ; j-- {
        if buckets[lenStr][j] < buckets[lenStr][j-1] {
          buckets[lenStr][j], buckets[lenStr][j-1] = buckets[lenStr][j-1], buckets[lenStr][j]
        } else {
          break
        }
      }
    }
  }

  return buckets
}

func main() {
  var s []string
  var res [][]string
  s = []string{"cat", "dog", "bull", "bird", "seagull", ""}
  res = bucketLexSort(s)
  printBuckets(res)
}

/*

Perform a bucket sort.
Each bucket is in lex order.

*/