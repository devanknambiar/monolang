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