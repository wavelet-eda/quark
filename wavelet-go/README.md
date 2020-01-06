# Quark Go

This is the Go implementation of Quark. Right now Antlr4 is used to generate the parser and the lexer. In the future, we will use a hand written parser for better error reporting.

# Building This Program

In an effort to make building easy, I have written init scripts for *nix and Windows. These download the build system and put the executable in the top level dir. I was inspired by boost :grimacing:.

Here are the steps:

1. Clone this repo 
2. Install Go version 13 (https://golang.org/dl/)
3. With the root directory of this repo as your working directory, run either `init.sh` or `init.bat` depending on you're operating system. This will place either "task.exe" or "task" in your current directory.
4. Run `$ ./task build` to build the project into `<PROJ_DIR>/build/exe/quarkc`

To clean your build area, run `$ ./task clean`. The gitignore has been configured to ignore the task executable so there's no need to delete that file. 