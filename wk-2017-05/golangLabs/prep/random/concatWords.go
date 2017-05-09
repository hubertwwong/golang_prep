package main

import "fmt"

type Node struct {
  val byte
  word bool
  children map[byte]*Node
}

type Dict struct {
  root *Node
}

func newNode(val byte, word bool) *Node {
  m := make(map[byte]*Node)
  n := Node{val, word, m}
  return &n
}

func (d *Dict) add(s string) {
  if d.root == nil {
    d.root = newNode(0, false)
  }

  if len(s) != 0 {
    // children map.
    children := (*d.root).children

    for i, v := range s {
      childNode := children[byte(v)]

      // create node
      if childNode == nil {
        endStr := false
        if i+1 == len(s) {
          endStr = true
        }
        
        // create a new node.
        n := newNode(byte(v), endStr)
        
        // assign to correct place in children map.
        children[byte(v)] = n
      }
 
      // traversing to child
      childNode = children[byte(v)]
      children = (*childNode).children
    }
  }
}

func (d *Dict) isComboWord(s string, min int) bool {
  if d.root == nil {
    return false
  }

  splitResult := false
  curResult := false

  // 1 character can't be a combo word
  if len(s) >= 1 {
    // children map.
    children := (*d.root).children
    wordCount := 0
    
    for i, v := range s {
      fmt.Println(i, ">", string(v), ">[wc]", wordCount, ">[i]", i, ">", min, ">[i>=min]", i>=min)
      childNode := children[byte(v)]

      // current chracter in word not in dict. break out.
      if childNode == nil {
        curResult = false
      }

      // break out variables for easy access.
      word := (*childNode).word

      // end of the list...
      if word && i>=min {
        // if its a word. increase word count.
        wordCount++

        // split off. basically if you.
        fmt.Println(">>>>> s > before", i+1)
        splitResult = d.isComboWord(s, i+1)
        fmt.Println(">>>>> s > after", i+1, splitResult)

        // restart the serarch
        children = (*d.root).children
        childNode = children[byte(v)]
      } else if i+1 != len(s) {
        // traverse to child.
        childNode = children[byte(v)]
        children = (*childNode).children
      }

      // you are at the end
      if i+1 == len(s) {
        fmt.Println(">result", splitResult, ">[wc]" ,wordCount, ">[w]", word, ">[w and wc]", (word && wordCount > 1))

        // if word {
        //   wordCount++
        // }

        if (word && wordCount > 1) || splitResult {
          curResult = true
        } else {
          curResult = false
        }
      }
    }
  }

  return curResult || splitResult
}

func findAllConcatenatedWordsInADict(words []string) []string {
  var d Dict
  
  // adding words to the dict.
  for _,v := range words {
    d.add(v)
  }

  // returning words that return true
  result := make([]string, 0)
  for _, v := range words {
    fmt.Println("")
    fmt.Println(">>>>>>>", v, "<<<<<<<<<<<<<<<<")
    fmt.Println("")
    if d.isComboWord(v, 0) {
      result = append(result, v)
    }
  }
  return result
}

func main() {
  //var d Dict
  //d.add("foo")
  //d.add("bar")
  //d.add("foobar")
  
  //fmt.Println(d.isComboWord("foo"))
  //fmt.Println(d.isComboWord("fooba"))
  //fmt.Println(d.isComboWord("bar"))
  //fmt.Println(d.isComboWord("foobar"))
  
  //fmt.Println(findAllConcatenatedWordsInADict([]string{"foo", "bar", "foobar"}))
  //fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"}))
  //fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats"}))
  //fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catsdogcats","dog"}))
  fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catsdog","dog"}))
  //["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]

  // v := d.root
  // c := (*v).children
  // fmt.Println(">c",c)
  // fmt.Println(c["foo"[0]])

  // cn := c["f"[0]]
  // fmt.Println(string((*cn).val))
  // c = (*cn).children
  // cn = c["o"[0]]
  // fmt.Println(string((*cn).val))
}

/*
notes so far
use runes instead of byptes..
for string character parsing..

if you have a map of children...
you need to make it.

need stars...
string[x] is a byptes
range of string return runes...


*/