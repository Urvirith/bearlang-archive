package parser

import (
	"github.com/Urvirith/bearlang/src/ast"
	"github.com/Urvirith/bearlang/src/lexer"
	"github.com/Urvirith/bearlang/src/token"
)

type Parser struct {
	lex       *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(lex *lexer.Lexer) *Parser {
	psr := &Parser{lex: lex}

	psr.nextToken()

	return psr
}

func (psr *Parser) nextToken() {
	psr.curToken = psr.peekToken
	psr.peekToken = psr.lex.NextToken()
}

func (psr *Parser) ParseProgram() *ast.Program {
	prg := &ast.Program{}
	prg.Statements = []ast.Statement{}

	for psr.curToken.Type != token.EOF {
		stmt := psr.parseStatement()
		if stmt != nil {
			prg.Statements = append(prg.Statements, stmt)
		}
		psr.nextToken()
	}

	return prg
}

func (psr *Parser) parseStatement() ast.Statement {
	switch psr.curToken.Type {
	case token.LET:
		return psr.parseLetStatement()
	default:
		return nil
	}
}

func (psr *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: psr.curToken}

	// Let is not followed by Identifer (Variable)
	if !psr.expectPeek(token.ID) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: psr.curToken, Value: psr.curToken.Literal}

	// Identifer is not followed by an Assign
	if !psr.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO : Skimming expressioins until semicolon
	for !psr.curTokenIs(token.SCOLON) {
		psr.nextToken()
	}

	return stmt
}

func (psr *Parser) curTokenIs(tok token.TokenType) bool {
	return psr.curToken.Type == tok
}

func (psr *Parser) peekTokenIs(tok token.TokenType) bool {
	return psr.peekToken.Type == tok
}

func (psr *Parser) expectPeek(tok token.TokenType) bool {
	if psr.peekTokenIs(tok) {
		psr.nextToken()
		return true
	}
	return false
}
