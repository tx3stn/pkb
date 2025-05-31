Set Height 600
Source '/vhs/.scripts/gifs/base.tape'

Output new-no-edit.gif

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
Enter 1
Sleep 0.8s
Down 1
Sleep 0.4s
Enter 1
Sleep 0.4s
Enter 1
Sleep 0.4s
Type "amazing project"
Enter 1
Sleep 0.4s
Type meeting minutes
Enter 1
Sleep 0.4s
Down 1
Sleep 0.4s
Space
Down 1
Sleep 0.4s
Space
Sleep 0.3s
Enter
Sleep 3s

Set TypingSpeed 0s
Hide
Type "sed -i 's/true/false/g' schema/example.config.json"
Enter
Show
