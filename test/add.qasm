# This is a comment and should be ignored
.literal newline "\n"
.literal message_out "The answer is "
.literal message_in_1 "Adding together "
.literal message_in_2 " and "
.literal a 1
.literal b 2

# Output "Adding together 1 and 2"
output message_in_1
output a
output message_in_2
output b
output newline

# Add variables a and b together and store in variable c
add a b c

# Output "The answer is 3"
output message_out
output c
output newline
