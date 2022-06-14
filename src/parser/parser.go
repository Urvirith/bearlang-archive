package parser

import (
	"fmt"

	"github.com/Urvirith/bearlang/src/ast"
	"github.com/Urvirith/bearlang/src/lexer"
	"github.com/Urvirith/bearlang/src/token"
)

type Parser struct {
	lex       *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

var datatypes = []token.TokenType{
	token.I8,
	token.I16,
	token.I32,
	token.I64,
	token.I128,
	token.U8,
	token.U16,
	token.U32,
	token.U64,
	token.U128,
	token.F32,
	token.F64,
	token.BOOL,
}

func New(lex *lexer.Lexer) *Parser {
	psr := &Parser{
		lex:    lex,
		errors: []string{},
	}

	psr.nextToken()

	return psr
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
	case token.RETURN:
		return psr.parseReturnStatement()
	default:
		return nil
	}
}

// Parse the let statement
func (psr *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: psr.curToken}

	// Let is not followed by Identifer (Variable)
	if !psr.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: psr.curToken, Value: psr.curToken.Literal}

	// Identifer Is Not Followed By :
	if !psr.expectPeek(token.COLON) {
		return nil
	}

	// Identifer Is Not Followed DataType :
	if !psr.expectPeekDataType() {
		return nil
	}

	// Identifer is not followed by an =
	if !psr.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO : Skimming expressioins until semicolon
	for !psr.curTokenIs(token.SCOLON) {
		psr.nextToken()
	}

	return stmt
}

func (psr *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: psr.curToken}

	psr.nextToken()

	// TODO : Skimming expressions until semicolon
	for !psr.curTokenIs(token.SCOLON) {
		psr.nextToken()
	}

	return stmt
}

// COMMON FUNCTIONS
// Verify if the token is as expected
func (psr *Parser) curTokenIs(tok token.TokenType) bool {
	return psr.curToken.Type == tok
}

// Verify if the next token is as expected
func (psr *Parser) peekTokenIs(tok token.TokenType) bool {
	return psr.peekToken.Type == tok
}

// Verify if the next token is as expected if so move on, else return error
func (psr *Parser) expectPeek(tok token.TokenType) bool {
	if psr.peekTokenIs(tok) {
		psr.nextToken()
		return true
	} else {
		psr.peekError(tok)
		return false
	}
}

// Verify all allowed types for all data
func (psr *Parser) expectPeekDataType() bool {
	for i := range datatypes {
		if psr.peekTokenIs(datatypes[i]) {
			psr.nextToken()
			return true
		}
	}

	psr.peekDataError()
	return false
}

// Return errors from data structure
func (psr *Parser) Errors() []string {
	return psr.errors
}

// Add an error for the expected error
func (psr *Parser) peekError(tok token.TokenType) {
	msg := fmt.Sprintf("expected next rune to be %s, got %s instead", tok, psr.peekToken.Type)
	psr.errors = append(psr.errors, msg)
}

// Add an error if the data type is not found
func (psr *Parser) peekDataError() {
	msg := fmt.Sprintf("expected next rune to be %v, got %s instead", datatypes, psr.peekToken.Type)
	psr.errors = append(psr.errors, msg)
}

// Move on to the next token, and peek ahead the following token
func (psr *Parser) nextToken() {
	psr.curToken = psr.peekToken
	psr.peekToken = psr.lex.NextToken()
}
