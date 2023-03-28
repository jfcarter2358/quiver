# Store data in memory
.literal newline "\n"
.literal vals ["hello ","world"]
.literal firstIdx 0
.literal newVal "universe"
.literal newIdx 2

# Test list assign
LIST_ASSIGN newVal newIdx vals

# Test list access
LIST_ACCESS vals firstIdx output1
LIST_ACCESS vals newIdx output2

# Output our message
output output1
output output2
output newline

# Explicitly bail out of the program
stop 0

