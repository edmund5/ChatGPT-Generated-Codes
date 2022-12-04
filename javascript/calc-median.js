// Define a function that calculates the median of an array of numbers
function median(numbers) {
  // Sort the array in ascending order
  numbers.sort();

  // Calculate the median
  let median = 0;
  const arraySize = numbers.length;
  if (arraySize % 2 === 1) {
    // If the array has an odd number of elements, the median is the middle value
    median = numbers[(arraySize - 1) / 2];
  } else {
    // If the array has an even number of elements, the median is the average of the two middle values
    median = (numbers[arraySize / 2] + numbers[arraySize / 2 - 1]) / 2;
  }

  // Return the median
  return median;
}

// Define an array of numbers
const numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// Calculate the median of the array
const median = median(numbers);

// Loop through the array of numbers and apply styles based on whether the number is above, below, or within 0-1% of the median
for (const number of numbers) {
  if (number > median) {
    // If the number is above the median, apply a style that sets the text color to green
    console.log(`\x1b[32m${number}\x1b[0m`);
  } else if (number < median) {
    // If the number is below the median, apply a style that sets the text color to red
    console.log(`\x1b[31m${number}\x1b[0m`);
  } else if (Math.abs(number - median) < 0.01 * median) {
    // If the number is within 0-1% of the median, apply a style that sets the text color to white
    console.log(`\x1b[37m${number}\x1b[0m`);
  }
}

// Print the median value
console.log(`Median: ${median}`);
