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

//parseStatement() method - Statement
func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case LET:
		return p.parseLetStatement()
	
	case RETURN:
		return p.parseReturnStatement()
	
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *LetStatement {
	return nil
}

func (p *Parser) parseReturnStatement() *LetStatement {
	return nil
}