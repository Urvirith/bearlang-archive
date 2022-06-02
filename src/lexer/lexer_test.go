package lexer

import (
	"testing"

	"github.com/Urvirith/bearlang/src/token"
)

func TestTokens(t *testing.T) {
	input := `= + - * / % | & ! ~ ^ ( ) { } [ ] , : ; fn let vol struct enum union const return if elif else match default true false`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.ASSIGN, "="},
		{token.ADD, "+"},
		{token.SUB, "-"},
		{token.ASTERISK, "*"},
		{token.DIV, "/"},
		{token.MOD, "%"},
		{token.OR, "|"},
		{token.AND, "&"},
		{token.BANG, "!"},
		{token.NOT, "~"},
		{token.XOR, "^"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.COMMA, ","},
		{token.COLON, ":"},
		{token.SCOLON, ";"},
		{token.FUNCTION, "fn"},
		{token.LET, "let"},
		{token.VOLITILE, "vol"},
		{token.STRUCT, "struct"},
		{token.ENUM, "enum"},
		{token.UNION, "union"},
		{token.CONST, "const"},
		{token.RETURN, "return"},
		{token.IF, "if"},
		{token.ELIF, "elif"},
		{token.ELSE, "else"},
		{token.MATCH, "match"},
		{token.DEFAULT, "default"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectType {
			t.Fatalf("tests[%d] - tokentype wrong. expected: %q, got: %q", i, tt.expectType, tok.Type)
		}

		if tok.Literal != tt.expectLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected: %q, got: %q", i, tt.expectLiteral, tok.Literal)
		}
	}
}

func TestCode(t *testing.T) {
	input := `let five = 5;
			  let ten = 10;
			  
			  let add = fn(x, y) {
				return x + y;
			  };
			  
			  let result = add(five, ten);
			  `

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.LET, "let"},
		{token.ID, "five"},
		{token.ASSIGN, "="},
		{token.I32, "5"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.ID, "ten"},
		{token.ASSIGN, "="},
		{token.I32, "10"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.ID, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.ID, "x"},
		{token.COMMA, ","},
		{token.ID, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.ID, "x"},
		{token.ADD, "+"},
		{token.ID, "y"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.ID, "result"},
		{token.ASSIGN, "="},
		{token.ID, "add"},
		{token.LPAREN, "("},
		{token.ID, "five"},
		{token.COMMA, ","},
		{token.ID, "ten"},
		{token.RPAREN, ")"},
		{token.SCOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectType {
			t.Fatalf("tests[%d] - tokentype wrong. expected: %q, got: %q", i, tt.expectType, tok.Type)
		}

		if tok.Literal != tt.expectLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected: %q, got: %q", i, tt.expectLiteral, tok.Literal)
		}
	}
}
