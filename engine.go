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

//Lexer readChar() Method - void
func (l *Lexer) readChar() {
	if l.Next >= len(l.Input) {
		l.Ch = 0
		return
	} 
	l.Ch = l.Input[l.Next]
	l.Pos = l.Next
	l.Next += 1
}

//Lexer skipWhitespace() Method - void
func (l *Lexer) skipWhitespace() {
	for l.Ch == ' ' || l.Ch == '\n' || l.Ch == '\t' || l.Ch == '\r' {
		l.readChar()
	}
}

//Lexer peekChar() Method - byte
func (l *Lexer) peekChar() byte {
	if l.Next >= len(l.Input) {
		return 0
	}
	return l.Input[l.Next]
}

//Lexer NextToken() Method - Token
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	var t Token

	switch l.Ch {
	case 0:
		t.Type = EOF
		t.Literal = string(l.Ch)

	case '+':
		t.Type = PLUS
		t.Literal = string(l.Ch)
	
	case '=':
		if l.peekChar() == '=' {
			t.Type = CHECK
			t.Literal = "=="
			l.readChar()
		} else {
			t.Type = ASSIGN
			t.Literal = string(l.Ch)
		}
	
	case '&':
		if l.peekChar() == '&' {
			t.Type = AND
			t.Literal = "&&"
			l.readChar()
		} else {
			t.Type = BAND
			t.Literal = string(l.Ch)
		}
	
	case '|':
		if l.peekChar() == '|' {
			t.Type = ORR
			t.Literal = "||"
			l.readChar()
		} else {
			t.Type = BORR
			t.Literal = string(l.Ch)
		}
	
	case ';':
		t.Type = SCOLON
		t.Literal = string(l.Ch)
		
	case ':':
		t.Type = COLON
		t.Literal = string(l.Ch)
	
	case ',':
		t.Type = COMMA
		t.Literal = string(l.Ch)
	
	case '!':
		if l.peekChar() == '=' {
			t.Type = NOT_EQ
			t.Literal = "!="
			l.readChar()
		} else {
			t.Type = NOT
			t.Literal = string(l.Ch)
		}
	
	case '-':
		t.Type = MINUS
		t.Literal = string(l.Ch)

	default:
		t.Type = ILLEGAL
		t.Literal = string(l.Ch)
	}

	l.readChar()

	return t
}

func main() {
	
}