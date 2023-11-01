## ascii-art-output by ayessenb and bbilisbe
* `ayessenb` 
* `bbilisbe` 

### Objectives

Ascii-art-justify is a program which specifies the format of the text alignmnet and outputs the provided string in graphic format using the specified banner aligned to one of four formats.
The argument will be a flag, --align=direction_of_alignmnet, in which --align is the flag and direction_of_alignment is one of four given directions of alignments, i.e. left/right/center/justify in which the output will be aligned to.
Moreover, the program can convert the graphical representation of the letters to text.
The argument will be a flag, --output=<fileName.txt>, in which --output is the flag and  <fileName> is the file name that will contain the output of the program.
The argument will be a flag, --reverse=<fileName.txt>, in which --reverse is the flag and <fileName> is the file name. The program must then print this string in normal text.


- This project should handle an input with numbers, letters, spaces, special characters and `\n`.
- Usage: go run . [OPTION]
- Ascii-art-reverse project implemented, the program accept the correctly formatted [OPTION] [STRING].
- Ascii-art-fs project implemented, the program accept the correctly formatted [STRING] [BANNER].
- Ascii-art-output project implemented, the program accept the correctly formatted [OPTION] [STRING] [BANNER].
- The program must still able to run with a single [STRING] argument.


### Banner Format

* shadow
* standard
* thinkertoy

#### Description:


* The program reads the banner file where the height of each letter is 8 lines
* The first line is not considered
* Each character in banner file is separated by one '\n'
* Since Banner contains the ascii values between 32 and 126 then the other runes are not accepted.
* If the Banner file will be changed the program will give the Error message
* If the program will get wrong input format, the program will give the Error message
* If the second argument is not "standard" or "shadow" or "thinkertoy" then the program will give the Error messages
* If the input string argument doesn't contain ascii value then the program gives the Error message.(between 32 and 126 (both included))


#### Improved skills:
* go lang programming skills 
* usage of string libraries of Go
* handling new line char.
* usage of many functions and files


## Usage/Examples
Cloning storage to your host
```CMD/Terminal 
git clone git@git.01.alem.school:bbilisbe/ascii-art-justify.git
```
Go to the downloaded repository:

```CMD/Terminal 
cd ascii-art-justify
```
Run a program:
```CMD/Terminal 
go run . --align=justify "Hello\nMy name is Kozha!" thinkertoy
```


Output: __HELLO__
```bash         
o  o     o o     
|  |     | |     
O--O o-o | | o-o 
|  | |-' | | | | 
o  o o-o o o o-o 
                 
                 
                                                                                
o   o                                                                    o  o         o         o 
|\ /|                                                  o                 | /          |         | 
| O | o  o             o-o   oo  o-O-o o-o               o-o             OO   o-o o-o O--o  oo  o 
|   | |  |             |  | | |  | | | |-'             |  \              | \  | |  /  |  | | |    
o   o o--O             o  o o-o- o o o o-o             | o-o             o  o o-o o-o o  o o-o- O 
         |                                                                                        
      o--o                                                                                                         
```





## Feedback

If you liked our project, we would be grateful if you could add `Star` to the repository.

Alem Student
02.06.2023.
