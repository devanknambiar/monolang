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

//Helper function isLetter(ch byte) - bool
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

//Helper function isDigit(ch byte) - bool
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

//Lexer readIdentifier() Method - string
func (l *Lexer) readIdentifier() string {
	startPos := l.Pos
	for isLetter(l.Ch) {
		l.readChar()
	}
	return l.Input[startPos : l.Pos]
}

//Lexer readDigit() Method - string
func (l *Lexer) readDigit() string {
	startPos := l.Pos
	for isDigit(l.Ch) {
		l.readChar()
	}
	return l.Input[startPos : l.Pos]
}

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
		if isLetter(l.Ch) {
			t.Literal = l.readIdentifier()
			t.Type = LookupIdent(t.Literal)
			return t
		} else if isDigit(l.Ch) {
			t.Literal = l.readDigit()
			t.Type = DIGIT
			return t
		} else {
			t.Type = ILLEGAL
			t.Literal = string(l.Ch)
		}
	}

	l.readChar()

	return t
}

func main() {
	input := `let result = 105;`
	l := New(input)
	for {
		t := l.NextToken()
		println("Type: ", t.Type, "& Literal: ", t.Literal)
		if t.Type == EOF {
			break
		}
	}
}