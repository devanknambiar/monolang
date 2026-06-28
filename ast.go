package main

//Node interface
type Node interface {
	TokenLiteral() string
}

//Statement
type Statement interface {
	Node
	statementNode()
}

//Expression
type Expression interface {
	Node
	expressionNode()
}