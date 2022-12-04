function generate3DigitNumber() {
  const digits = [Math.floor(Math.random() * 10), Math.floor(Math.random() * 10), Math.floor(Math.random() * 10)];
  return digits.join("-");
}

const dataset = [];

// Generate 6 random three-digit numbers
for (let i = 0; i < 6; i++) {
  dataset.push(generate3DigitNumber());
}

let output = "";
for (const number of dataset) {
  output += `${number}\n`;
}

// Replace \n with <br>
output = output.replace(/\n/g, "<br>");

console.log(output);
