# Store data in memory
.literal newline "\n"
.literal vals {} STRING STRING
.literal helloKey "hello"
.literal fooKey "foo"
.literal helloVal "world"
.literal fooVal "bar"
.literal helloVal2 ""
.literal fooVal2 ""
.literal helloVal3 ""
.literal fooVal3 ""

DICT_ASSIGN helloKey helloVal
DICT_ASSIGN fooKey fooVal

DICT_ACCESS helloKey helloVal2
DICT_ACCESS fooKey fooVal2

COPY helloVal2 helloVal3
COPY fooVal2 fooVal3

# Output our message
output helloVal3
output newline
output fooVal3
output newline

# Explicitly bail out of the program
stop 0

