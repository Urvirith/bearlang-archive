package parser

import (
	"fmt"
	"testing"

	"github.com/Urvirith/bearlang/src/ast"
	"github.com/Urvirith/bearlang/src/lexer"
)

func TestLetStatementsPass(t *testing.T) {
	input := `
	let x: u16 = 5;
	let y: u32 = 10;
	let foobar: f32 = 838383;
	`
	lex := lexer.New(input)
	psr := New(lex)
	prg := psr.ParseProgram()
	checkParserErrors(t, psr)

	if prg == nil {
		t.Fatalf("ParseProgram() returned null")
	}

	if len(prg.Statements) != 3 {
		t.Fatalf("program.Statements does not have 3 statements. got=%d", len(prg.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := prg.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestLetStatementsFail(t *testing.T) {
	input := `
	let x: u16 5;
	let : u32 = 10;
	let 838383;
	`
	lex := lexer.New(input)
	psr := New(lex)
	prg := psr.ParseProgram()
	checkParserErrors(t, psr)

	if prg == nil {
		t.Fatalf("ParseProgram() returned null")
	}

	if len(prg.Statements) != 3 {
		t.Fatalf("program.Statements does not have 3 statements. got=%d", len(prg.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := prg.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral not 'let', got: %q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)

	if !ok {
		t.Errorf("letStmt.Name.Value not: '%s', got: %s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name not: '%s', got: %s", name, letStmt.Name)
		return false
	}

	return true
}

func TestReturnStatementsPass(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 903030;
	`
	lex := lexer.New(input)
	psr := New(lex)
	prg := psr.ParseProgram()
	checkParserErrors(t, psr)

	if prg == nil {
		t.Fatalf("ParseProgram() returned null")
	}

	if len(prg.Statements) != 3 {
		t.Fatalf("program.Statements does not have 3 statements. got=%d", len(prg.Statements))
	}

	for _, stmt := range prg.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, psr *Parser) {
	errors := psr.errors

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error; %q", msg)
	}
	t.FailNow()
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	lex := lexer.New(input)
	psr := New(lex)

	program := psr.ParseProgram()
	checkParserErrors(t, psr)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statments. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatment)

	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatment. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("expression not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	lex := lexer.New(input)
	psr := New(lex)

	program := psr.ParseProgram()
	checkParserErrors(t, psr)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statments. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatment)

	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatment. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expression not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Fatalf("ident.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Fatalf("ident.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		lex := lexer.New(tt.input)
		psr := New(lex)

		program := psr.ParseProgram()
		checkParserErrors(t, psr)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statments. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatment)

		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatment. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf("expression not *ast.Identifier. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator not %s. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integer.Value != value {
		t.Errorf("integer.Value not %d. got=%d", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral not %d. got=%s", value, integer.TokenLiteral())
		return false
	}

	return true
}
