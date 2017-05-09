package main

import "fmt"

func main() {
  n := "32a"
  
  if n[1] > byte("0"[0]) && n[1] < byte("9"[0]) {
    fmt.Println("num")
  } else {
    fmt.Println("not num")
  }
}