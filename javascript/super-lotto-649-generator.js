function generateNumbers(evenCount, oddCount) {
  let output = "";

  for (let i = 0; i < 6; i++) {
    let numbers = [];

    // Generate even numbers
    for (let j = 0; j < evenCount; j++) {
      let even = Math.floor(Math.random() * 48) + 2;
      if (even % 2 !== 0) even--;

      // Check for duplicates
      while (numbers.includes(even)) {
        even = Math.floor(Math.random() * 48) + 2;
        if (even % 2 !== 0) even--;
      }

      // Pad the even number with a leading zero until the total string length is equal to 2
      even = even.toString().padStart(2, "0");

      numbers.push(even);
    }

    // Generate odd numbers
    for (let j = 0; j < oddCount; j++) {
      let odd = Math.floor(Math.random() * 49) + 1;
      if (odd % 2 === 0) odd++;

      // Check for duplicates
      while (numbers.includes(odd)) {
        odd = Math.floor(Math.random() * 49) + 1;
        if (odd % 2 === 0) odd++;
      }

      // Pad the odd number with a leading zero until the total string length is equal to 2
      odd = odd.toString().padStart(2, "0");

      numbers.push(odd);
    }

    output += numbers.join("-") + "\n";
  }

  // Replace \n with <br>
  //output = output.replace(/\n/g, "<br>");

  return output;
}

console.log(generateNumbers(2, 4));
