package parser

import (
	"fmt"
	"strconv"

	"github.com/Urvirith/bearlang/src/ast"
	"github.com/Urvirith/bearlang/src/lexer"
	"github.com/Urvirith/bearlang/src/token"
)

type Parser struct {
	lex            *lexer.Lexer
	curToken       token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
	errors         []string
}

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

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

var precedences = map[token.TokenType]int{
	token.EQU:      EQUALS,
	token.NEQ:      EQUALS,
	token.LES:      LESSGREATER,
	token.GRT:      LESSGREATER,
	token.ADD:      SUM,
	token.SUB:      SUM,
	token.DIV:      PRODUCT,
	token.ASTERISK: PRODUCT,
}

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

func New(lex *lexer.Lexer) *Parser {
	psr := &Parser{
		lex:    lex,
		errors: []string{},
	}

	// Read two tokens, curToken and peekToken are set
	psr.nextToken()
	psr.nextToken()

	psr.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	psr.infixParseFns = make(map[token.TokenType]infixParseFn)

	// REGISTER PREFIXES
	psr.registerPrefix(token.IDENTIFIER, psr.parseIdentifier)
	psr.registerPrefix(token.INT, psr.parseIntegerLiteral)
	psr.registerPrefix(token.SUB, psr.parsePrefixExpression)
	psr.registerPrefix(token.NOT, psr.parsePrefixExpression)
	psr.registerPrefix(token.TRUE, psr.parseBoolean)
	psr.registerPrefix(token.FALSE, psr.parseBoolean)
	psr.registerPrefix(token.LPAREN, psr.parseGroupExpression)

	psr.registerInfix(token.ADD, psr.parseInfixExpression)
	psr.registerInfix(token.SUB, psr.parseInfixExpression)
	psr.registerInfix(token.ASTERISK, psr.parseInfixExpression)
	psr.registerInfix(token.DIV, psr.parseInfixExpression)
	psr.registerInfix(token.EQU, psr.parseInfixExpression)
	psr.registerInfix(token.NEQ, psr.parseInfixExpression)
	psr.registerInfix(token.GRT, psr.parseInfixExpression)
	psr.registerInfix(token.LES, psr.parseInfixExpression)

	return psr
}

func (psr *Parser) ParseProgram() *ast.Program {
	prg := &ast.Program{}
	prg.Statements = []ast.Statement{}

	for psr.curToken.Type != token.EOF {
		stmt := psr.parseStatement()
		println(stmt.String())
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
		return psr.parseExpressionStatement()
		//return nil
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

// Parse the return statement
func (psr *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: psr.curToken}

	psr.nextToken()

	// TODO : Skimming expressions until semicolon
	for !psr.curTokenIs(token.SCOLON) {
		psr.nextToken()
	}

	return stmt
}

// Parse Expression Statements
func (psr *Parser) parseExpressionStatement() *ast.ExpressionStatment {
	stmt := &ast.ExpressionStatment{
		Token: psr.curToken,
	}

	stmt.Expression = psr.parseExpression(LOWEST)

	if psr.peekTokenIs(token.SCOLON) {
		psr.nextToken()
	}

	return stmt
}

// Parse Expression
func (psr *Parser) parseExpression(precedence int) ast.Expression {
	prefix := psr.prefixParseFns[psr.curToken.Type]

	if prefix == nil {
		psr.noPrefixParseFnError(psr.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !psr.peekTokenIs(token.SCOLON) && precedence < psr.peekPrecedence() {
		infix := psr.infixParseFns[psr.peekToken.Type]

		if infix == nil {
			return leftExp
		}

		psr.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (psr *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: psr.curToken,
		Value: psr.curToken.Literal,
	}
}

func (psr *Parser) parseIntegerLiteral() ast.Expression {
	literal := &ast.IntegerLiteral{
		Token: psr.curToken,
	}

	value, err := strconv.ParseInt(psr.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", psr.curToken.Literal)
		psr.errors = append(psr.errors, msg)
		return nil
	}

	literal.Value = value

	return literal
}

func (psr *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    psr.curToken,
		Operator: psr.curToken.Literal,
	}

	psr.nextToken()

	exp.Right = psr.parseExpression(PREFIX)

	return exp
}

func (psr *Parser) noPrefixParseFnError(tokenType token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function found for %s found", tokenType)
	psr.errors = append(psr.errors, msg)
}

// Infix Expressions
func (psr *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token:    psr.curToken,
		Operator: psr.curToken.Literal,
		Left:     left,
	}

	prec := psr.curPrecedence()
	psr.nextToken()
	expr.Right = psr.parseExpression(prec)

	return expr
}

func (psr *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: psr.curToken,
		Value: psr.curTokenIs(token.TRUE),
	}
}

func (psr *Parser) parseGroupExpression() ast.Expression {
	psr.nextToken()

	exp := psr.parseExpression(LOWEST)

	if !psr.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
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

// Peek Precedence
func (psr *Parser) peekPrecedence() int {
	if p, ok := precedences[psr.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

// Current Precedence
func (psr *Parser) curPrecedence() int {
	if p, ok := precedences[psr.curToken.Type]; ok {
		return p
	}

	return LOWEST
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

// Register a prefix for an expression
func (psr *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	psr.prefixParseFns[tokenType] = fn
}

// Register a infix for an expression
func (psr *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	psr.infixParseFns[tokenType] = fn
}
