# MonoLang

> A statically-typed programming language and tree-walking interpreter built from scratch in Go.

MonoLang is a personal systems programming project focused on understanding how programming languages work internally. Instead of relying on parser generators or compiler frameworks, every stage of the language is being implemented manually—from lexical analysis to parsing, evaluation, and type checking.

The goal of this project is to gain a deeper understanding of compiler construction, language design, and runtime systems.

## Current Status

🚧 **Active Development**

### Implemented

* Token representation
* Token type definitions
* Lexer structure
* Character stream traversal
* Basic operator tokenization
* Incremental project structure and documentation
* Identifier tokenization
* Integer literals
* Whitespace handling
* Multi-character operators (`==`, `&&`, `||`, etc.)

### In Progress
* Parser
* Abstract Syntax Tree (AST)


### Planned

* Tree-walking evaluator
* Static type checker
* Variables and expressions
* Functions
* Control flow (`if`, `while`, etc.)
* REPL
* Browser playground

## Motivation

Programming languages are often treated as black boxes.

MonoLang is an attempt to understand every layer involved in language implementation by building each component from first principles.

Rather than focusing on performance, this project prioritizes readability, correctness, and learning.

## Project Structure

```text
monolang/
├── cmd/
│   └── main.go
├── lexer/
│   └── lexer.go
├── token/
│   └── token.go
├── go.mod
├── README.md
└── LICENSE
```

> **Note:** The repository structure will evolve as additional compiler stages are implemented.


## Roadmap

* [x] Project initialization
* [x] Token definitions
* [x] Lexer implementation
* [x] Identifier parsing
* [x] Numeric literals
* [ ] Parser
* [ ] AST
* [ ] Expression evaluation
* [ ] Static type checking
* [ ] Functions
* [ ] Error diagnostics
* [ ] Interactive REPL
* [ ] Web playground

## Technologies

* Go
* Git
* GitHub

## References

This project draws inspiration from several excellent resources on compiler and interpreter construction, while being implemented independently as a learning exercise.

* *Writing an Interpreter in Go* — Thorsten Ball
* *Crafting Interpreters* — Robert Nystrom
* *Compilers: Principles, Techniques, and Tools* ("Dragon Book")



## License

MIT License
