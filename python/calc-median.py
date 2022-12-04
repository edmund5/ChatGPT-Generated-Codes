import statistics

# Set the list of numbers
numbers = [1, 2, 3, 4, 5]

# Calculate the median
median = statistics.median(numbers)

# Print the numbers in the list
print("The numbers in the list are:")
for number in numbers:
  print(number)

# Check the value of the median and print the output in the appropriate color
if median < 0.01:
  print("\033[1;37mThe median is", median, "\033[0m")
elif median < 0.5:
  print("\033[1;31mThe median is", median, "\033[0m")
else:
  print("\033[1;32mThe median is", median, "\033[0m")
