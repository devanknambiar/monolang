package main

//Parser & its initialization
type Parser struct {
	l			*Lexer
	curToken	Token
	peekToken	Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

//Parser Constructor
func NewParser(l *Lexer) *Parser {
	p := &Parser{l : l}
	p.nextToken() //for p.curToken
	p.nextToken() //for p.nextToken
	return p
} 