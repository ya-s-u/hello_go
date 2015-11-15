package main

import (
    "bufio"
    "fmt"
    "os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() string {
  sc.Scan()
  return sc.Text()
}

func main() {
  var words []string
  sc.Split(bufio.ScanWords)
  n := 6
  for i := 0; i < n; i++ {
      words = append(words, nextInt())
  }

  fmt.Println(words)
}
