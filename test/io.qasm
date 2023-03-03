# Store data in memory
.literal adminUser "John Carter"
.literal newline "\n"
.literal inputPrompt "Enter your name : "
.literal outputPrompt "Hello, "
.literal adminMessage "Admin access granted"
.literal normalMessage "Non-admin access"
.literal doneMessage "Done!"

# Output input prompt
output inputPrompt

# Get user input
input_block name keyCode

# Output our message
output outputPrompt
output name
output newline

equal name adminUser isAdminUser

branch_positive isAdminUser printAdminUser

# Explicitly bail out of the program
output normalMessage
output newline

goto end

.label printAdminUser

output adminMessage
output newline

.label end
output doneMessage
output newline
