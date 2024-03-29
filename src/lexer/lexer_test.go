package lexer

import (
	"testing"

	"github.com/Urvirith/bearlang/src/token"
)

func TestTokens(t *testing.T) {
	input := `= + - * / % | & ! ~ ^ += -= ++ -- |= &= ^= << >> == != > < >= <= || && => ( ) { } [ ] , : ; import fn let vol struct enum union const return if elif else match default for loop while true false i8 i16 i32 i64 i128 u8 u16 u32 u64 u128 f32 f64 bool`

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
		{token.NOT, "!"},
		{token.COMP, "~"},
		{token.XOR, "^"},
		{token.ADD_ASSIGN, "+="},
		{token.SUB_ASSIGN, "-="},
		{token.INC, "++"},
		{token.DEC, "--"},
		{token.OR_ASSIGN, "|="},
		{token.AND_ASSIGN, "&="},
		{token.XOR_ASSIGN, "^="},
		{token.LSHF, "<<"},
		{token.RSHF, ">>"},
		{token.EQU, "=="},
		{token.NEQ, "!="},
		{token.GRT, ">"},
		{token.LES, "<"},
		{token.GEQ, ">="},
		{token.LEQ, "<="},
		{token.COR, "||"},
		{token.CAND, "&&"},
		{token.MATCH_BRANCH, "=>"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.COMMA, ","},
		{token.COLON, ":"},
		{token.SCOLON, ";"},
		{token.IMPORT, "import"},
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
		{token.FOR, "for"},
		{token.LOOP, "loop"},
		{token.WHILE, "while"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.I8, "i8"},
		{token.I16, "i16"},
		{token.I32, "i32"},
		{token.I64, "i64"},
		{token.I128, "i128"},
		{token.U8, "u8"},
		{token.U16, "u16"},
		{token.U32, "u32"},
		{token.U64, "u64"},
		{token.U128, "u128"},
		{token.F32, "f32"},
		{token.F64, "f64"},
		{token.BOOL, "bool"},
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
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "x"},
		{token.ADD, "+"},
		{token.IDENTIFIER, "y"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
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

func TestNextToken(t *testing.T) {
	input := `let five: u32 = 5;
			  let ten: u32 = 10;
			  
			  let add = fn(x:u32, y:u32)(u32) {
				return x + y;
			  };
			  
			  let result = add(five, ten);
			  !-/*5;

			  if 5 <= 10 {
				return true;
			  } elif 10 > 5 {
				return true;
			  } else {
				return false;
			  }

			  10 == 10;
			  10 != 9;
			  let dogu864: f32 = 20.0;
			  `

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.COLON, ":"},
		{token.U32, "u32"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.COLON, ":"},
		{token.U32, "u32"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COLON, ":"},
		{token.U32, "u32"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.COLON, ":"},
		{token.U32, "u32"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.U32, "u32"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "x"},
		{token.ADD, "+"},
		{token.IDENTIFIER, "y"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SCOLON, ";"},
		{token.NOT, "!"},
		{token.SUB, "-"},
		{token.DIV, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SCOLON, ";"},
		{token.IF, "if"},
		{token.INT, "5"},
		{token.LEQ, "<="},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELIF, "elif"},
		{token.INT, "10"},
		{token.GRT, ">"},
		{token.INT, "5"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQU, "=="},
		{token.INT, "10"},
		{token.SCOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SCOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "dogu864"},
		{token.COLON, ":"},
		{token.F32, "f32"},
		{token.ASSIGN, "="},
		{token.INT, "20.0"},
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
