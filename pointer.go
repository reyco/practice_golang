package main

import(
  "fmt"
)


func main(){
  x := 0
  changeXValNow(&x)
  fmt.Println("x =", x)
  fmt.Println("Memory Address for x =", &x)
}


func changeXValNow(x *int){
  *x = 2
}
