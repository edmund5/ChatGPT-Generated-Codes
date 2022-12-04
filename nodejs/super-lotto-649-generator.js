function generateNumbers(evenCount, oddCount) {
  let output = "";

  for (let i = 0; i < 6; i++) {
    let numbers = [];

    // Generate even numbers
    for (let j = 0; j < evenCount; j++) {
      let even = Math.floor(Math.random() * 48) + 2;
      if (even % 2 !== 0) even--;
      numbers.push(even);
    }

    // Generate odd numbers
    for (let j = 0; j < oddCount; j++) {
      let odd = Math.floor(Math.random() * 49) + 1;
      if (odd % 2 === 0) odd++;
      numbers.push(odd);
    }

    output += numbers.join("-") + "\n";
  }

  // Replace \n with <br>
  //output = output.replace(/\n/g, "<br>");

  return output;
}

console.log(generateNumbers(2, 4));
