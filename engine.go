package main

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	// Run-time
	ILLEGAL	TokenType = "ILLEGAL"
	EOF 	TokenType = "EOF"

	//Operators
	PLUS 	TokenType = "+"
	MINUS 	TokenType = "-"
	ASSIGN 	TokenType = "="

	//Logical 
	CHECK 	TokenType = "=="
	AND 	TokenType = "&&"
	ORR 	TokenType = "||"
	NOT 	TokenType = "!"

	//Delimiters
	COMMA 	TokenType = ","
	COLON 	TokenType = ":"
	SCOLON 	TokenType = ";"

	//Identifiers
	IDENT 	TokenType = "IDENT"
	INT 	TokenType = "INT"
)

type Lexer struct {
	Input string
	Pos int
	Next int
	Ch byte
}

func New(input string) *Lexer {
	//Physical Instance
	l := &Lexer{Input : input}
	l.Pos = 0
	l.Next = 1
	return l
}

func main() {
	
}