package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func generateNumbers(evenCount, oddCount int) string {
	var output string

	for i := 0; i < 6; i++ {
		numbers := []string{}

		// Generate even numbers
		for j := 0; j < evenCount; j++ {
			var even int
			for {
				even = rand.Intn(48) + 2
				if even%2 == 0 {
					break
				}
				even--
			}

			// Check for duplicates
			for _, number := range numbers {
				if strconv.Itoa(even) == number {
					even = rand.Intn(48) + 2
					if even%2 == 0 {
						break
					}
					even--
				}
			}

			// Pad the even number with a leading zero until the total string length is equal to 2
			evenStr := strconv.Itoa(even)
			if len(evenStr) < 2 {
				evenStr = "0" + evenStr
			}

			numbers = append(numbers, evenStr)
		}

		// Generate odd numbers
		for j := 0; j < oddCount; j++ {
			var odd int
			for {
				odd = rand.Intn(49) + 1
				if odd%2 != 0 {
					break
				}
				odd++
			}

			// Check for duplicates
			for _, number := range numbers {
				if strconv.Itoa(odd) == number {
					odd = rand.Intn(49) + 1
					if odd%2 != 0 {
						break
					}
					odd++
				}
			}

			// Pad the odd number with a leading zero until the total string length is equal to 2
			oddStr := strconv.Itoa(odd)
			if len(oddStr) < 2 {
				oddStr = "0" + oddStr
			}

			numbers = append(numbers, oddStr)
		}

		output += strings.Join(numbers, "-") + "\n"
	}

	// Replace \n with <br>
	//output = output.replace(/\n/g, "<br>");

	return output
}

func main() {
	fmt.Println(generateNumbers(2, 4))
}
