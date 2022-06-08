package parser

import (
	"testing"

	"github.com/Urvirith/bearlang/src/ast"
	"github.com/Urvirith/bearlang/src/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	lex := lexer.New(input)
	psr := New(lex)
	prg := psr.ParseProgram()

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
