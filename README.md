# Quark 
[![Actions Status](https://github.com/wavelet-eda/quark/workflows/Go%20Tests/badge.svg)](https://github.com/wavelet-eda/quark/actions)

[Read The Docs](https://quark-docs.readthedocs.io/en/latest/?)

`*.qk`

Quark is an independent HDL language using a syntax largely inspired by SystemVerilog but with a dramatically improved type system and purely sequential event ordering paradigm. While no behavior across clock cycles is inferred by the language compiler, modern language features like Rust's ownership model help to make hardware programming more flexible and easier to see design intent. The compiler is written in Go and supports a variety of language targets, including simulation code and netlists in a variety of formats.

This is the Go implementation of Quark. Right now Antlr4 is used to generate the parser and the lexer. In the future, we will use a hand written parser for better error reporting.

# Building This Program

In an effort to make building easy, I have written init scripts for *nix and Windows. These download the build system and put the executable in the top level dir. I was inspired by boost :grimacing:.

Here are the steps:

1. Clone this repo 
2. Install Go version 13 (https://golang.org/dl/)
3. With the root directory of this repo as your working directory, run either `init.sh` or `init.bat` depending on your operating system. This will place either "task.exe" or "task" in your current directory.
4. Run `$ ./task build` to build the project into `<PROJ_DIR>/build/exe/quarkc` (if on Windows the executable will have the `.exe` file extension)

To clean your build area, run `$ ./task clean`. The gitignore has been configured to ignore the task executable so there's no need to delete that file. 

# How to Run

No setup is required if Golang 1.13.5 is installed.

1. `$ ./init.sh` (or `$ init.bat` on windows). You only have to do this once.
2. `$ ./task run "INFILE=<file>"` where <file> is a valid quark file.

That's it. The compiler will be appropriately rebuilt to reflect source changes
between runs.

# Running Tests

`$ ./task test` compiles and runs all tests. Do this before pushing please.

# Viewing the Docs

Just run `$godoc -http localhost:8080` and open that address in a browser.

# Package Structure

The compiler has 3 packages right now.

- `pkg/quarkc` is where the application itself lives. This is the main package. Most code shouldn't live here.
- `quark` is where front end language data structures are. This includes the AST and relevant metadata
- `parser` is where the parser and lexer live. The go code is generated from the antlr g4 files. The parse_tree_converter
module is for converting raw parse trees to proper `quark.AST`. 
