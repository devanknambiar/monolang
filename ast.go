package main

//Node interface
type Node interface {
	TokenLiteral() string
}

//Statement
type Statement interface {
	Node
	statementNode()	//dummy marker
}

//Expression
type Expression interface {
	Node
	expressionNode() //dummy marker
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}

type LetStatement struct {
	Token Token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {}

type ReturnStatement struct {
	Token Token
	Value Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}