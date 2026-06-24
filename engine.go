package main

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	// Run-time
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//Operators
	PLUS = "+"
	MINUS = "-"
	ASSIGN = "="

	//Logical 
	CHECK = "=="
	AND = "&&"
	ORR = "||"
	NOT = "!"

	//Delimiters
	COMMA = ","
	COLON = ":"
	SCOLON = ";"
)

func main() {
	
}