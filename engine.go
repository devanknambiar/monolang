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

//Lexer Structure
type Lexer struct {
	Input string
	Pos int
	Next int
	Ch byte
}

//Lexer Constructor
func New(input string) *Lexer {
	//Physical Instance
	l := &Lexer{Input : input}
	l.Pos = 0
	l.Next = 0 //Modification: For the 0th character
	l.readChar()
	return l
}

//Lexer readChar() Method
func (l *Lexer) readChar() {
	if l.Next >= len(l.Input) {
		l.Ch = 0
		return
	} 
	l.Ch = l.Input[l.Next]
	l.Pos = l.Next
	l.Next += 1
}

func main() {
	
}