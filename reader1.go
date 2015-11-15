package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  var lines []string

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }

  fmt.Println(lines)
}
