package ast

import "github.com/Urvirith/bearlang/src/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statement []Statement
}

func (prog *Program) TokenLiteral() string {
	if len(prog.Statement) > 0 {
		return prog.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode() {
	// Placeholder
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ind *Identifier) expressionNode() {
	// Placeholder
}

func (ind *Identifier) TokenLiteral() string {
	return ind.Token.Literal
}
