package main
import "fmt"

func main() {
  var input int
  fmt.Scan(&input)

  if isPrime(input) {
    fmt.Println("is a prime number")
  } else {
    fmt.Println("is not a prime number")
  }
}

func isPrime(n int) bool {
  for v := 2; v * v <= n; v++ {
    if n % v == 0 {
      return false
    }
  }
  return true
}
