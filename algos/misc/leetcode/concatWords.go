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
				//fmt.Println("> creating node", string(v))
        endStr := false
        if i+1 == len(s) {
					//fmt.Println("> end str", string(v))
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

// min is i
func (d *Dict) isComboWord(s string, min int, wc int, passedCurNode *Node) bool {
  if d.root == nil {
    return false
  }

  splitResult := false
  curResult := false

  // 1 character can't be a combo word
  if len(s) >= 1 {
    // Children map. Basically if you get a passed node, assign it.
		children := (*d.root).children
    if passedCurNode != nil {
			children = (*passedCurNode).children
		}
		wordCount := wc
   
    //for i, v := range s {
		i:=min
		for ; i<len(s) ; i++ {
			v := s[i]
			fmt.Println("")
      fmt.Println("> For", i, ">", string(v), ">[wc]", wordCount, ">[i]", i, ">", min, ">[i>=min]", i>=min)
      childNode := children[byte(v)]

      // current chracter in word not in dict. break out.
      if childNode == nil {
				fmt.Println("> Char not in dict. Child node is.")
        curResult = false
				break
      }

      // break out variables for easy access.
      word := (*childNode).word
			val := (*childNode).val
			fmt.Println("> On char >", string(val), childNode)

      if i+1 == len(s) {
				// you should mark a word if its a word.
        if word {
					fmt.Println(">> word found at end")
          wordCount++
        }
				
				fmt.Println("> [result]", splitResult, ">[wc]" ,wordCount, ">[w]", word, ">[curResult]", (word && wordCount > 1), ">[val]", string(val))

        //if (word && wordCount > 1) || splitResult {
				if (word && wordCount > 1) {
          curResult = true
					break
        } else {
          curResult = false
					break
        }
      } else if word && i>=min {
        // split off. basically if you.
				fmt.Println("> word found > before i", i+1, childNode, string((*childNode).val))
        //splitResult = d.isComboWord(s, i+1, wordCount)
				// the combo words....
				splitResult = d.isComboWord(s, i+1, wc, childNode)
        fmt.Println("> word found > after i", i+1, splitResult)

				// if its a word. increase word count.
				// the word count should be after the split.
        wordCount++

        // restart the serarch
        children = (*d.root).children
        childNode = children[byte(v)]
      } else if i+1 != len(s) {
				fmt.Println("> t", string(v))
        // traverse to child.
        childNode = children[byte(v)]
        children = (*childNode).children
      }
    }
  }

	fmt.Println("> Returning", ">c", curResult, ">s", splitResult)
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
    if d.isComboWord(v, 0, 0, nil) {
      result = append(result, v)
    }
  }
  return result
}

func main() {
  //var d Dict
  // d.add("a")
	// d.add("b")
  // d.add("ab")
  // d.add("abc")
	// fmt.Println(d.isComboWord("ab",0,0, nil))
  // fmt.Println(d.isComboWord("fooba"))
  // fmt.Println(d.isComboWord("bar"))
  // fmt.Println(d.isComboWord("foobar"))
	//d.add("cat")
	//d.add("cats")
  //fmt.Println(d.isComboWord("cats",0,0,nil))
	//fmt.Println(*(d.root).children["c"])
 
  //fmt.Println(findAllConcatenatedWordsInADict([]string{"foo", "bar", "foobar"}))
  //fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"}))
	//fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catsdogcats","dog","dogcatsdog"}))
	fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catscats"}))
	//fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catscat", "catscat"}))
	//fmt.Println(findAllConcatenatedWordsInADict([]string{"a","b","ab","abc"}))

	//fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats"}))
  //fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catsdogcats","dog"}))
  //fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catsdog","dog"}))
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

the edge case i'm running into this this...
abc
you are trying to hack the system by passing some variables back
and split of a decision..
- i'm guessing that is off...
a,b,ab,abc
and the 2 results that i'm returning is this.
a,b split.
on one of the trees you are passing 2 as the word count.
2 as the min value.
and then you are retraversing...
and then you go abc.

i think i get the issue. splitting logic is off....



*/