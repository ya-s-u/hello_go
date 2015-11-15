package main
import "fmt"

func main(){
  m := map[string]int{
    "Toyko": 1,
    "Osaka": 2 ,
    "Nagoya": 3,
  }

  for index, value := range m {
    fmt.Println(index, value)
  }

  for index := range m {
    fmt.Println(index)
  }

  for _, value := range m {
    fmt.Println(value)
  }

  for range m {
    fmt.Print("*")
  }
}
