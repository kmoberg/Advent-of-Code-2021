# Read the file
with open("input.txt") as f:
    seq = [line.rstrip() for line in f]  # Return the result without the new line character

seq = list(map(int, seq))

# seq = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263] # For making debugging simpler

# Initialize a bunch of variables
index = prev_sum = number_increase = number_decrease = 0
increase = decrease = False # For debugging purposes

# Set the size of the sliding window (how many numbers you want to group together)
window_size = 3

# Loop through the entire sequence of numbers
for i in range(len(seq) - window_size + 1):
    # Slice the through the numbers to add just the three numbers in the current window and add as a new array
    nums = seq[i: i + window_size]
    print(nums)

    # Loop through the array to find the sum of the current array
    window_sum = sum(nums)

    if index > 0:
        if window_sum > prev_sum:
            number_increase += 1
            # increase = True  # For debugging purposes
        elif window_sum < prev_sum:
            number_decrease += 1
            # decrease = True  # For debugging purposes

    index += 1

    # This block makes debugging much easier, but can be removed
    # if increase:
    #     # print("Final Sum:", window_sum, "(Increase - previous:", prev_sum, ")")
    #     increase = False
    # elif decrease:
    #     # print("Final Sum:", window_sum, "(Decrease - previous:", prev_sum, ")")
    #     decrease = False

    prev_sum = window_sum

print("\n\nFinal Tally...")
print("Increase:", number_increase)
print("Decrease:", number_decrease)
