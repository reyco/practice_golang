package main

import(
  "fmt"
)


func main(){
  yPtr := new(int)
  changeYValNow(yPtr)
  fmt.Println("y =", *yPtr)
  fmt.Println("Memory Address for y =", yPtr)
}


func changeYValNow(yPtr *int){
  *yPtr = 2
}
