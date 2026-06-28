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

//Kick-start Program
func (p *Parser) ParseProgram() *Program {
	program := &Program{
		Statements: []Statement{},
	}
	
	for p.curToken.Type != EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	
	return program
}

//helper func to evaluate syntax errors
func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	} 
	return false
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

func (p *Parser) parseReturnStatement() *ReturnStatement {
	return nil
}