package ast

import (
	"bytes"

	"github.com/Urvirith/bearlang/src/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (prg *Program) String() string {
	var out bytes.Buffer

	for _, s := range prg.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
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

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")

	return out.String()
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

func (ind *Identifier) String() string {
	return ind.Value
}

// Expression Section
type ExpressionStatment struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatment) expressionNode() {
	// Placeholder
}

func (es *ExpressionStatment) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatment) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
