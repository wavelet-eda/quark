Quark Compiler Design
=====================

The quark compiler is broken into a few different stages.

1. QuarkLexer takes a stream of input characters and produces a stream of tokens that represent keywords, complete literals, operators, and identifiers.
2. QuarkParser works on a single file level and produces a syntax tree describing all of the tokens that were passed from the lexer. This does not check references to other files or even checks that the variables being referenced are valid, it just makes sure the syntax is well formed.
3. QuarkLinker stage first looks at the import statements in the file to build a tree of what file dependencies exist in the build. This is then used to refine the syntax tree, with pointers to the appropriate referenced structures, as well as do a complete static evaluation of the syntax. This will not catch things like parameter-driven asserts or combinational loops, but it will catch width mismatches and such.
4. QuarkEvaluation acts as pre-processor that resolves/expands all of the parameter usage in the syntax tree, validates that there are no combinational loops in the expanded syntax tree, checks that unique case statements are fully covered, and makes sure that all design units are properly connected. This essentially takes the syntax tree that was generated by the parser and partially evaluates it.
5. The partially evaluated syntax tree that was produced can now be fed into one of two different "runtimes", one for simulation and the other for hardware generation.
    a. QuarkSimulator takes the partially evaluated syntax tree and produces C/C++ that can be used to simulate the modules you wrote. This module description can then be hooked up to a quark simulator runtime that can either run on its own or possibly be managed through an industry standard VPI interface.
    b. QuarkNetlist is responsible for building a directed graph representing the logic. This is an expansion of the syntax tree since all variables become nodes in the graph that can only ever be assigned to once, completely unrolling for loops and turning if-statement into "multiplexers".
        i. QuarkVerilog is the final stage responsible for turning the fully formed netlist into actual verilog, which is a fairly simple step, just having to manage the netlist names do not conflict with Verilog keywords.

Header Descriptions:

- QuarkTokens.h
    Contains descriptions of the different tokens that are generated by parsing a text file, including literals, symbols, identifiers, and keywords.

- QuarkNodes.h
    Contains descriptions of the syntax tree containing the parsed tokens but now with a hierarchical structure and not just a linear stream of tokens

- QuarkNetlist.h
    Contains descriptions of a netlist 