<?php

// Define a function that calculates the median of an array of numbers
function median($numbers) {
  // Sort the array in ascending order
  sort($numbers);

  // Calculate the median
  $median = 0;
  $array_size = count($numbers);
  if ($array_size % 2 == 1) {
    // If the array has an odd number of elements, the median is the middle value
    $median = $numbers[($array_size - 1) / 2];
  } else {
    // If the array has an even number of elements, the median is the average of the two middle values
    $median = ($numbers[$array_size / 2] + $numbers[$array_size / 2 - 1]) / 2;
  }

  // Return the median
  return $median;
}

// Define an array of numbers
$numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

// Calculate the median of the array
$median = median($numbers);

// Loop through the array of numbers and apply styles based on whether the number is above, below, or within 0-1% of the median
foreach ($numbers as $number) {
  if ($number > $median) {
    // If the number is above the median, apply a style that sets the text color to green
    echo '<span style="color: green">' . $number . '</span> ';
  } elseif ($number < $median) {
    // If the number is below the median, apply a style that sets the text color to red
    echo '<span style="color: red">' . $number . '</span> ';
  } elseif (abs($number - $median) < 0.01 * $median) {
    // If the number is within 0-1% of the median, apply a style that sets the text color to white
    echo '<span style="color: white">' . $number . '</span> ';
  }
}

// Print the median value
echo 'Median: ' . $median;
