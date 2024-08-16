package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  fmt.Println("Hello world!")
  input := bufio.NewScanner(os.Stdin)
  input.Scan()
}
 