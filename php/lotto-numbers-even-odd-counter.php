<?php

function even_odd_count($numbers_string) {
  // Split the string on either a hyphen or whitespace character
  $numbers = preg_split('/-|\s/', $numbers_string);

  // Convert each element in the array to an integer
  $numbers = array_map('intval', $numbers);

  $even_count = 0;
  $odd_count = 0;

  foreach($numbers as $number) {
    if($number % 2 == 0) {
      $even_count++;
    } else {
      $odd_count++;
    }
  }

  return compact('even_count', 'odd_count');
}

if (isset($_GET['numbers'])) {
  $numbers_string = $_GET['numbers'];
  $counts = even_odd_count($numbers_string);

  echo "There are {$counts['even_count']} even numbers and {$counts['odd_count']} odd numbers.";
} else {
  echo "Error: The 'numbers' parameter is not set in the URL query string.";
}
