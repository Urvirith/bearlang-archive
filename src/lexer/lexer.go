package lexer

import "github.com/Urvirith/bearlang/src/token"

// Structure defining the Lexer
type Lexer struct {
	in      string // Input data
	pos     int    // Current position in data - Current Character
	readPos int    // Current reading position in data - After Current Character
	ch      byte   // Current Char
}

// Create new instance and initialize the read position
func New(in string) *Lexer {
	lex := &Lexer{in: in}
	lex.readChar()
	return lex
}

// Fetch the next token
func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lex.ch {
	case '=':
		tok = newToken(token.ASSIGN, lex.ch)
	case '+':
		tok = newToken(token.ADD, lex.ch)
	case '-':
		tok = newToken(token.SUB, lex.ch)
	case '*':
		tok = newToken(token.MUL, lex.ch)
	case '/':
		tok = newToken(token.DIV, lex.ch)
	case '|':
		tok = newToken(token.OR, lex.ch)
	case '&':
		tok = newToken(token.AND, lex.ch)
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	case '[':
		tok = newToken(token.LBRACK, lex.ch)
	case ']':
		tok = newToken(token.RBRACK, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case ':':
		tok = newToken(token.COLON, lex.ch)
	case ';':
		tok = newToken(token.SCOLON, lex.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	lex.readChar()
	return tok
}

// Read the character of the input string
func (lex *Lexer) readChar() {
	// Read the Character or prevent overflow of read from the readPos
	if lex.readPos < len(lex.in) {
		lex.ch = lex.in[lex.readPos]
	} else {
		lex.ch = 0
	}
	lex.pos = lex.readPos
	lex.readPos += 1
}

// Return a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
