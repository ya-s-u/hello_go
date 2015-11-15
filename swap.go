package main
import "fmt"

func main(){
  swap := func (a, b int)(int, int) {
    tmp := a
    a = b
    b = tmp
    return a, b
  }

  fmt.Print(swap(1, 2))
}

// func swap (a, b int)(int, int) {
//   tmp := a
//   a = b
//   b = tmp
//   return a, b
// }
