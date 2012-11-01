The Stammer Programming Language Compiler
=========================================

This is the bootstrap compiler for the stammer programming language created 
by Brendan Higgins. Stammer is designed to be medium to low level programming
language intended for generic application development.  Stammer is a
object-oriented, garbage collected, inferred type programming language which
compiles to assembly.  This compiler is written in the Go programming language
and is being created as a means to boot-strap the language. 

I started writing this compiler for educational purposes; thus, this compiler is completely written from scratch.

Directory list:
---------------
- regex: contains the regular expression compiler and virtual machine
- lexer: contains the code to generate a lexical analyzer.
- parser: contains a parser generater
- test: contains some testing code.
- stammer_code: contains some examples of stammer.
