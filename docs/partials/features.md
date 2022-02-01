## Features

### md  

Outputs in terminal formated content of an *.md file.\
Example: stylog md -f README.md -s dark

**flags:**
   --f string   An *.md file to output.\
   --s string   Style to use for ouputing fomated MD file. Available options are: dark|ascii|light|notty|dracula (default "dark")
### star
Outputs in terminal formated message, preceeded by star sign symbol.\
Example: log -star "Star Message, Hey!"

**flags:**
   --m string   Stared Message to be printed
### warn
Outputs in terminal formated message, with title "WARN" followed by warning message.\
Example: log -warn "Warning Message, Hey!"

**flags:**
   --m string   Warning Message to be printed

### error
Outputs in terminal formated message, with title "ERROR" followed by error message.\
Example: log -error "Error Message, Fatal Error!"

**flags:**
  --m string   Error Message to be printed

### success
Outputs in terminal formated message, preceeded by green check sign symbol.\
Example: log -success "Success Message Congrats!"

**flags:**
  --m string   Success Message to be printed

### info
Outputs in terminal formated message, preceeded by blue dot symbol.\
Example: log info -m "Informational Message FYI"

**flags:**
  --m string   Informational Message to be printed