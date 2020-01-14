# Wavelet
Tools for working with digital logic written with modern programming languages and paradigms.

## Quark 
`*.qk`

Quark is an independent HDL language using a syntax largely inspired by SystemVerilog but with a dramatically improved type system and purely sequential event ordering paradigm. While no behavior across clock cycles is inferred by the language compiler, modern language features like Rust's ownership model help to make hardware programming more flexible and easier to see design intent. The compiler is written in Go and supports a variety of language targets, including simulation code and netlists in a variety of formats.

## Neutron 
`*.nu`

Neutron is an HLS language borrowing syntax and operators from Quark, but allowing further and more advanced language features to be used during runtime, as well as inferring clock cycles between lines of code by building implicit state machines. A variety of workflows including async/await, arbitrary memory access, recursion, and clock cycle events are supported and further build functionality into the baseline Quark syntax. Neutron's compiler is also written in Go, but simply produces Quark code and relies on the Quark compiler for simulation or netlist generation.
