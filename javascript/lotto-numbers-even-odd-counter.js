function evenOddCount(numbersString) {
  // Split the string on either a hyphen or whitespace character
  const numbers = numbersString.split(/\-|\s/);

  // Convert each element in the array to an integer
  numbers = numbers.map(Number);

  let evenCount = 0;
  let oddCount = 0;

  for (const number of numbers) {
    if (number % 2 === 0) {
      evenCount++;
    } else {
      oddCount++;
    }
  }

  return { evenCount, oddCount };
}

if (process.argv.length >= 3) {
  const numbersString = process.argv[2];
  const counts = evenOddCount(numbersString);

  console.log(`There are ${counts.evenCount} even numbers and ${counts.oddCount} odd numbers.`);
} else {
  console.log("Error: The 'numbers' parameter is not set in the URL query string.");
}
