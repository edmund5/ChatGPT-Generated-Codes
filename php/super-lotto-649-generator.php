<?php

function generateNumbers($evenCount, $oddCount) {
  $output = "";

  for ($i = 0; $i < 6; $i++) {
    $numbers = array();

    // Generate even numbers
    for ($j = 0; $j < $evenCount; $j++) {
      $even = mt_rand(2, 48);
      if ($even % 2 != 0) $even--;
      $numbers[] = $even;
    }

    // Generate odd numbers
    for ($j = 0; $j < $oddCount; $j++) {
      $odd = mt_rand(1, 49);
      if ($odd % 2 == 0) $odd++;
      $numbers[] = $odd;
    }

    $output .= implode('-', $numbers) . "\n";
  }

  // Replace \n with <br>
  $output = str_replace("\n", "<br>", $output);

  return $output;
}

echo generateNumbers(2, 4);