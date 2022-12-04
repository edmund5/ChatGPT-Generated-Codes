package main

import (
  "fmt"
  "math/rand"
  "strings"
)

func generate3DigitNumber() string {
  digits := [3]int{rand.Intn(10), rand.Intn(10), rand.Intn(10)}
  return fmt.Sprintf("%d-%d-%d", digits[0], digits[1], digits[2])
}

func main() {
  var dataset []string
  
  // Generate 6 random three-digit numbers
  for i := 0; i < 6; i++ {
    dataset = append(dataset, generate3DigitNumber())
  }

  output := strings.Join(dataset, "\n")
  
  fmt.Println(output)
}
