package main

import(
  "fmt"
)


func main(){  
  presAge := make(map[string] int)
  
  presAge["Manuel L Quezon"] = 98
  
  fmt.Println(len(presAge))

  presAge["Fidel Ramos"] = 80
  
  fmt.Println(len(presAge))
  
  
  presAge["Gloria Macapagal Arroyo"] = 60
  
  fmt.Println(len(presAge))  


  presAge["Benigno Simeon \"Noynoy\" Aquino III"] = 58
  
  fmt.Println(len(presAge))  
  
  
  fmt.Println("Ramos age is", presAge["Fidel Ramos"])
  
  
}



