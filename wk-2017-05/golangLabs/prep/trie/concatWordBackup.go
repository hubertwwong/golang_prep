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

func (d *Dict) debug(s string) {
  if d.root == nil {
    d.root = newNode(0, false)
  }

  if len(s) != 0 {
    // children map.
    children := (*d.root).children

    for _, v := range s {
      childNode := children[byte(v)]
      fmt.Println("> debug >", string(childNode.val), "|", childNode)

      // traversing to child
      childNode = children[byte(v)]
      children = (*childNode).children
    }
  }
}

func (d *Dict) isComboWord(s string) bool {
  return d.isComboWordDFS(s, 0, nil)
}

func (d *Dict) isComboWordDFS(s string, wc int, passedCurNode *Node) bool {
  // fmt.Println("> icw > start >", s, wc, passedCurNode)
  if d.root == nil {
    return false
  }

  splitResult := false
  curResult := false

  // 1 character can't be a combo word
  if len(s) >= 1 {
    // children map.
    children := (*d.root).children
    if passedCurNode != nil {
      children = (*passedCurNode).children
    }
    wordCount := wc
   
    // Loop thru each character.
    for i := 0 ; i<len(s) ; i++ {
      v := s[i]
      //fmt.Println("> icw >", i, ">", string(v), ">[wc]", wordCount, "> [i]", i)
      childNode := children[byte(v)]

      // current chracter in word not in dict. break out.
      if childNode == nil {
        //fmt.Println("> icw > Char not in dict. Break out of the loop.")
        curResult = false
        break
      }

      word := (*childNode).word

      // This is the end check.
      if i+1 == len(s) {
        // you should mark a word if its a word.
        if word {
          wordCount++
        }
        
        //fmt.Println(">[result]", splitResult, ">[wc]" ,wordCount, ">[w]", word, ">[curResult]", (word && wordCount > 1), ">[val]", string(val))
        if (word && wordCount > 1) || splitResult {
          //fmt.Println("> icw > is a concat word")
          curResult = true
          break
        } else {
          //fmt.Println("> icw > is not a concat word")
          curResult = false
          break
        }
      } else if word {
        //fmt.Println("> icw > split")
        splitResult = d.isComboWordDFS(s[i+1:], wordCount, childNode)
        wordCount++

        // restart the serarch
        children = (*d.root).children
        childNode = children[byte(v)]
      } else if i+1 != len(s) {
        //fmt.Println("> icw > traversal")
        childNode = children[byte(v)]
        children = (*childNode).children
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
    //fmt.Println("\n> findAllConcatenatedWordInADict >>>", v, "<<<<")
    if d.isComboWord(v) {
      result = append(result, v)
    }
  }
  return result
}

// TESTS and MAIN FUNC.
//////////////////////////////////////////////////////////////////////////////

func dict01() {
  var d Dict
  d.add("cat")
  d.add("cats")
  d.debug("cats")
}

func concat01() {
  words := []string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"}
  fmt.Println("\n> result", findAllConcatenatedWordsInADict(words))
}

func concat02() {
  words := []string{"cat","cats","catsdogcats","dog"}
  fmt.Println("\n> result", findAllConcatenatedWordsInADict(words))
}

//
func main() {
  //dict01()
  concat01()
  //concat02()
}

