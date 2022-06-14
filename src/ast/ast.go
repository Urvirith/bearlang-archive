package ast

import "github.com/Urvirith/bearlang/src/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// PROGRAM SECTION
type Program struct {
	Statements []Statement
}

func (prg *Program) TokenLiteral() string {
	if len(prg.Statements) > 0 {
		return prg.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LET SECTION
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {
	// Placeholder
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// RETURN SECTION
type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) statementNode() {
	// Placeholder
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// IDENTIFIER SECTION
type Identifier struct {
	Token token.Token
	Value string
}

func (ind *Identifier) expressionNode() {
	// Placeholder
}

func (ind *Identifier) TokenLiteral() string {
	return ind.Token.Literal
}
