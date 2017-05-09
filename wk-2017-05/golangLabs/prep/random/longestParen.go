package main

import "fmt"

func longestValidParentheses(s string) int {
  // adding some constants to make it easier to read.
  const Init = 0
  const Illegal = 1
  const Possible = 2
  const Complete = 3
  const Extend = 4
  const ExtendComplete = 5

  state := Init
  maxLen := 0
  possibleMaxLen := 0
  stack := 0
  stackPop := 0

  // base checks
  // can't have an empty or parenthese that is 1 character long.
  if len(s) <= 1 {
    return 0
  }

  for i, v := range s {
    fmt.Println(">ib", i, ">s", state, "m", maxLen, "p", possibleMaxLen, string(v))

    if string(v) == "(" {
      if state == Init {
        state = Possible
        possibleMaxLen++
        stack++
      } else if state == Illegal {
        state = Possible
        possibleMaxLen++
        stack++
      } else if state == Possible {
        possibleMaxLen++
        stack++
        // i think this is the big thing?
        // basically when you push back into a stack, you erase the pop count.
        stackPop=0
      } else if state == Complete {
        state = Extend
        possibleMaxLen++
        stack++
      } else if state == Extend {
        possibleMaxLen++
        stack++
      } else if state == ExtendComplete {
        state = Extend
        possibleMaxLen++
        stack++
      }
    } else {
      if state == Init {
        state = Illegal
      } else if state == Illegal {
        // do nothing. just putting this in so its obvious.
        // continue
      } else if state == Possible {
        stack--
        stackPop++
        if stack == 0 {
          state = Complete
        } else if stack > 0 {
          state = Extend
        }
        if maxLen < stackPop*2 {
          maxLen = stackPop*2
        }
      } else if state == Complete {
        state = Illegal

        // need to check here if you have a solution.
        if maxLen < possibleMaxLen*2 {
          maxLen = possibleMaxLen*2
        }

        possibleMaxLen = 0
        stack = 0
        stackPop = 0
      } else if state == Extend {
        fmt.Println("extend", stackPop+1)
        stack--
        stackPop++
        if stack == 0 {
          state = ExtendComplete
        }
        
        if maxLen < stackPop*2 {
            maxLen = stackPop*2
          }
      } else if state == ExtendComplete {
        fmt.Println("extendComplete")
        // check if we have some new max len.
        if maxLen < possibleMaxLen*2 {
          maxLen = possibleMaxLen*2
        }

        state = Illegal
        possibleMaxLen = 0
        stack = 0
        stackPop = 0
      }
    }

    fmt.Println(">if", i, ">s", state, "m", maxLen, "p", possibleMaxLen, string(v))
  }

  // some cleanup checks...
  // if maxLen < stackPop*2 {
  //   maxLen = stackPop*2
  // }

  // fmt.Println("sp", stackPop)
  // stackPop = stackPop * 2
  // if stackPop > maxLen {
  //   maxLen = stackPop
  // }

  // printing to shutup compiler
  // fmt.Println(state, maxLen, stack)

  return maxLen
}

func main() {
  //fmt.Println(longestValidParentheses("(()(((()"), ">2")
  fmt.Println(longestValidParentheses("(()()"), ">4")
  //fmt.Println(longestValidParentheses("(()((((("), ">2")
  //fmt.Println(longestValidParentheses(")()())()()("), ">4")
  //fmt.Println(longestValidParentheses(")()()()()((("), ">8")
  //fmt.Println(longestValidParentheses("(()"), ">2")
  //fmt.Println(longestValidParentheses("())"), ">2")

  // final )()())()()(
  // 2 (()(((((
  // 2 ))))))()()()()((()
  // 8 )()()()()(((
  // 2 ())
  // 2 (()
  // )()()()()(((
  // 2 (()(((()
  // 4 (()()

}


/*

)()()(

>ib 0 >s 0 m 0 p 0 )
>if 0 >s 1 m 0 p 0 )

new one....
>ib 1 >s 1 m 0 p 0 (
>if 1 >s 2 m 0 p 1 (
>ib 2 >s 2 m 0 p 1 )
>if 2 >s 3 m 2 p 1 )

extend...
>ib 3 >s 3 m 2 p 1 (
>if 3 >s 4 m 2 p 2 (
>ib 4 >s 4 m 2 p 2 )
extend.. this should put 4..
>if 4 >s 5 m 2 p 2 )

>ib 5 >s 5 m 2 p 2 (
>if 5 >s 4 m 2 p 3 (
2




(()()
>ib 0 >s 0 m 0 p 0 (
>if 0 >s 2 m 0 p 1 (

>ib 1 >s 2 m 0 p 1 (
>if 1 >s 2 m 0 p 2 (
>ib 2 >s 2 m 0 p 2 )
>if 2 >s 2 m 2 p 2 )

>ib 3 >s 2 m 2 p 2 (
>if 3 >s 2 m 2 p 3 (

//  s2....
>ib 4 >s 2 m 2 p 3 )
>if 4 >s 2 m 2 p 3 )
2


// ......
const Init = 0
const Illegal = 1
const Possible = 2
const Complete = 3
const Extend = 4
const ExtendComplete = 5


i think i have the right idea...
but need to rethink this a little...



*/