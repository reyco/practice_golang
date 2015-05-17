package main

import(
  "fmt"
)


func main(){
  fmt.Println("Bad", safeDiv(3,0))
  fmt.Println("Good", safeDiv(3,2))
}


func safeDiv(num1, num2 int) int{
  defer func(){
    recover()
  }()
  
  solution := num1 / num2
  return solution
  
  
}
