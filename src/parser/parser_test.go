package parser

import (
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
