# Store data in memory
.literal newline "\n"
.literal vals {"hello":"world"}
.literal helloKey "hello"
.literal newHelloVal "universe"

# Test dict assign
DICT_ASSIGN newHelloVal helloKey vals

# Test dict access
DICT_ACCESS vals helloKey outputVal

# Test copy
COPY outputVal message

# Output our message
output message
output newline

# Explicitly bail out of the program
stop 0

