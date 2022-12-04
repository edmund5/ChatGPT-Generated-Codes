package main

import (
  "fmt"
  "math"
  "sort"
)

// Define a function that calculates the median of an array of numbers
func Median(numbers []int) int {
  // Sort the array in ascending order
  sort.Ints(numbers)

  // Calculate the median
  var median int
  arraySize := len(numbers)
  if arraySize % 2 == 1 {
    // If the array has an odd number of elements, the median is the middle value
    median = numbers[(arraySize - 1) / 2]
  } else {
    // If the array has an even number of elements, the median is the average of the two middle values
    median = (numbers[arraySize / 2] + numbers[arraySize / 2 - 1]) / 2
  }

  // Return the median
  return median
}

func main() {
  // Define an array of numbers
  numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  // Calculate the median of the array
  median := Median(numbers)

  // Loop through the array of numbers and apply styles based on whether the number is above, below, or within 0-1% of the median
  for _, number := range numbers {
    if number > median {
      // If the number is above the median, apply a style that sets the text color to green
      fmt.Printf("\x1b[32m%d\x1b[0m\n", number)
    } else if number < median {
      // If the number is below the median, apply a style that sets the text color to red
      fmt.Printf("\x1b[31m%d\x1b[0m\n", number)
    } else if math.Abs(float64(number - median)) < 0.01 * float64(median) {
      // If the number is within 0-1% of the median, apply a style that sets the text color to white
      fmt.Printf("\x1b[37m%d\x1b[0m\n", number)
    }
  }

  // Print the median value
  fmt.Printf("Median: %d\n", median)
}
