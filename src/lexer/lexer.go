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

	lex.consumeWhitespace()

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
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readID()
			tok.Type = token.LookupID(tok.Literal)
			return tok
		} else if isDigit(lex.ch) {
			tok.Type = token.I32
			tok.Literal = lex.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.ch)
		}
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

// Read the identifier of the input string
func (lex *Lexer) readID() string {
	// Read the Identifer or prevent overflow of read from the readPos
	pos := lex.pos
	for isLetter(lex.ch) {
		lex.readChar()
	}
	return lex.in[pos:lex.pos]
}

// Read the digits of the input string
func (lex *Lexer) readNumber() string {
	// Read the Identifer or prevent overflow of read from the readPos
	pos := lex.pos
	for isDigit(lex.ch) {
		lex.readChar()
	}
	return lex.in[pos:lex.pos]
}

// Consume whitespace as it serves no purpose
func (lex *Lexer) consumeWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

// Return a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
