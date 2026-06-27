package main

type TokenType string //'TokenType' alias

//Token Structure
type Token struct {
	Type TokenType
	Literal string
}

//Token Vocab
const (
	// Run-time
	ILLEGAL	TokenType = "ILLEGAL"
	EOF 	TokenType = "EOF"

	//Operators
	PLUS 	TokenType = "+"
	MINUS 	TokenType = "-"
	ASSIGN 	TokenType = "="

	//Bitwise
	BAND	TokenType = "&"
	BORR	TokenType = "|"

	//Logical 
	CHECK 	TokenType = "=="
	NOT_EQ	TokenType = "!="
	AND 	TokenType = "&&"
	ORR 	TokenType = "||"
	NOT 	TokenType = "!"

	//Delimiters
	COMMA 	TokenType = ","
	COLON 	TokenType = ":"
	SCOLON 	TokenType = ";"

	//Identifiers
	IDENT 	TokenType = "IDENT"
	DIGIT 	TokenType = "DIGIT"

	//Keywords
	LET		TokenType = "LET"
	IF 		TokenType = "IF"
	ELSE	TokenType = "ELSE"
	TRUE	TokenType = "TRUE"
	FALSE	TokenType = "FALSE"
)

//keywords map
var keywords = map[string]TokenType {
	"let" 	: LET,
	"if"	: IF,
	"else"	: ELSE,
	"true"	: TRUE,
	"false"	: FALSE,
}

//helper function LookupIdent(ident string) - TokenType
func LookupIdent(ident string) TokenType {
	value, ok := keywords[ident]
	if ok {
		return value
	}
	return IDENT
}