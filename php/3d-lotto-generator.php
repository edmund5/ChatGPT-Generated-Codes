<?php

function generate_3digit_number() {
  $digits = array(rand(0, 9), rand(0, 9), rand(0, 9));
  return implode("-", $digits);
}

$dataset = array();

// Generate 6 random three-digit numbers
for ($i = 0; $i < 6; $i++) {
  $dataset[] = generate_3digit_number();
}

$output = "";
foreach ($dataset as $number) {
  $output .= "$number\n";
}

// Replace \n with <br>
$output = str_replace("\n", "<br>", $output);

echo $output;
