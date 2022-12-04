package main

import (
  "fmt"
  "strings"
  "strconv"
  "os"
)

func evenOddCount(numbersString string) (evenCount int, oddCount int) {
  // Split the string on either a hyphen or whitespace character
  numbers := strings.Split(numbersString, "-")
  numbers = strings.FieldsFunc(numbersString, func(c rune) bool {
    return c == ' ' || c == '-'
  })

  // Convert each element in the array to an integer
  for _, number := range numbers {
    number, err := strconv.Atoi(number)
    if err != nil {
      fmt.Println("Error converting number to integer: ", err)
    }

    if number % 2 == 0 {
      evenCount++
    } else {
      oddCount++
    }
  }

  return
}

func main() {
  if len(os.Args) >= 3 {
    numbersString := os.Args[2]
    evenCount, oddCount := evenOddCount(numbersString)

    fmt.Printf("There are %d even numbers and %d odd numbers.\n", evenCount, oddCount)
  } else {
    fmt.Println("Error: The 'numbers' parameter is not set in the URL query string.")
  }
}
