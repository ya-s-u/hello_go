package main
import "fmt"

func main(){
  matrix := [][]int{
    {0, 0},
    {0, 1},
    {1, 0},
    {1, 1},
  }

  // matrix := make([][]int, 4)
  // matrix[0] = []int{0, 0}
  // matrix[1] = []int{0, 1}
  // matrix[2] = []int{1, 0}
  // matrix[3] = []int{1, 1}

  for index, value := range matrix {
    fmt.Println(index, value)
  }

  for index := range matrix {
    fmt.Println(index)
  }

  for _, value := range matrix {
    fmt.Println(value)
  }
}
