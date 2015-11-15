package main
import "fmt"

func main(){
  arr := []int{2, 3, 5, 7, 11, 13}

  // arr := make([]int, 6)
  // arr[0] = 2
  // arr[1] = 3
  // arr[2] = 5
  // arr[3] = 7
  // arr[4] = 11
  // arr[5] = 13

  for i := 0; i < len(arr); i++ {
    fmt.Printf("arr[%d] == %d\n", i, arr[i])
  }

  for index, value := range arr {
    fmt.Println(index, value)
  }

  for index := range arr {
    fmt.Println(index)
  }

  for _, value := range arr {
    fmt.Println(value)
  }

  for range arr {
    fmt.Print("*")
  }
}
