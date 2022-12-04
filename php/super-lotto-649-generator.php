<?php

function generateNumbers($evenCount, $oddCount) {
  $output = "";

  for ($i = 0; $i < 6; $i++) {
    $numbers = array();

    // Generate even numbers
    for ($j = 0; $j < $evenCount; $j++) {
      $even = mt_rand(2, 48);
      if ($even % 2 != 0) $even--;

      // Check for duplicates
      while (in_array($even, $numbers)) {
        $even = mt_rand(2, 48);
        if ($even % 2 != 0) $even--;
      }

      // Pad the even number with a leading zero until the total string length is equal to 2
      $even = str_pad($even, 2, "0", STR_PAD_LEFT);

      $numbers[] = $even;
    }

    // Generate odd numbers
    for ($j = 0; $j < $oddCount; $j++) {
      $odd = mt_rand(1, 49);
      if ($odd % 2 == 0) $odd++;

      // Check for duplicates
      while (in_array($odd, $numbers)) {
        $odd = mt_rand(1, 49);
        if ($odd % 2 == 0) $odd++;
      }

      // Pad the odd number with a leading zero until the total string length is equal to 2
      $odd = str_pad($odd, 2, "0", STR_PAD_LEFT);

      $numbers[] = $odd;
    }

    $output .= implode('-', $numbers) . "\n";
  }

  // Replace \n with <br>
  $output = str_replace("\n", "<br>", $output);

  return $output;
}

echo generateNumbers(2, 4);
