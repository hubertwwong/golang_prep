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

  // create the buckets. you want +1 on the length.
  buckets := make([][]string, maxStrLen+1)

  // stash strings into buckets.
  for i := 0 ; i < len(vals) ; i++ {
    lenStr := len(vals[i])

    // bucket does not exist.
    if len(buckets[lenStr]) == 0 {
      curBucket := make([]string,1)
      curBucket[0] = vals[i]
      buckets[lenStr] = curBucket
    } else {
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

// is the sequence of characters in s in d?
// e.g. (This would return true)
//  s = ale
//  d = azzzzlzzzze
func containsSeq(s, d string) bool {
  if len(s) > len(d) {
    return false
  }

  charCount := 0
  di := 0
  for si := 0 ; si < len(s) ; si++ {
    // match s[si] to d[di]
    for ; di < len(d) ; di++ {
      // check if source character matches in dest.
      if s[si] == d[di] {
        charCount++
        di++
        break
      }
    }
  }
  
  if charCount == len(s) {
    return true
  } else {
    return false
  }
}

func findLongestWord(s string, d []string) string {
  buckets := bucketLexSort(d)

  // iterate thru all bucket
  for l := len(buckets)-1 ; l > 0 ; l-- {
    //fmt.Println(">", l)

    // iterate each item in bucket.
    // strlen should be at least the bucket size.
    if len(s) >= l {
      for i := 0 ; i < len(buckets[l]) ; i++ {
        if containsSeq(buckets[l][i], s) {
          return buckets[l][i]
        }
      }
    }
  }
  
  return ""
}



func main() {
  var d []string
  var res string
  var s string

  // d = []string{"cat", "dog", "bull", "bird", "seagull"}
  // s = "cat"
  // d = []string{"ale","apple","monkey","plea"}
  // s = "abpcplea"
  // d = []string{"a","b"}
  // s = "abpcplea"
  d = []string{"a", "b"}
  s = ""

  res = findLongestWord(s, d)
  fmt.Println(">", res, "|", s, d)

  //fmt.Println(">cs", containsSeq("ale", "elale"))
}

/*

s = "abpcplea", d = ["a","b","c"]

s = "abpcplea", d = ["ale","apple","monkey","plea"]


https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/#/description



*/