Set Height 600
Source '/vhs/.scripts/gifs/base.tape'

Output accessible-mode.gif

# edit accessible_mode value in config file
Set TypingSpeed 0s
Hide
Type "sed -i 's/false/true/g' schema/example.config.json"
Enter
Type "clear"
Enter
Show
Set TypingSpeed 0.16s

Type "pkb new --no-edit"
Enter 
Sleep 0.8s
Type "2"
Enter
Sleep 0.4s
Type "2"
Enter
Sleep 0.4s
Type "1"
Enter
Sleep 0.4s
Type "important file"
Enter
Sleep 3s

# revert accessible_mode value in config file
Set TypingSpeed 0s
Hide
Type "sed -i 's/true/false/g' schema/example.config.json"
Enter
Show
