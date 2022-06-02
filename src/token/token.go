package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FN,
	"let":    LET,
	"return": RETURN,
}

// Constants For The Types Of Tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers & Literals
	ID   = "ID"   // add, x, y, etc...
	I8   = "I8"   // Signed Integer 8 Bit
	I16  = "I16"  // Signed Integer 16 Bit
	I32  = "I32"  // Signed Integer 32 Bit
	I64  = "I64"  // Signed Integer 64 Bit
	I128 = "I128" // Signed Integer 128 Bit
	U8   = "U8"   // Unsigned Integer 8 Bit
	U16  = "U16"  // Unsigned Integer 16 Bit
	U32  = "U32"  // Unsigned Integer 32 Bit
	U64  = "U64"  // Unsigned Integer 64 Bit
	U128 = "U128" // Unsigned Integer 128 Bit

	// Operators
	ASSIGN = "="
	ADD    = "+"
	SUB    = "-"
	MUL    = "*"
	DIV    = "/"
	OR     = "|"
	AND    = "&"

	// Comparators
	EQU  = "=="
	NEQ  = "!="
	GRT  = ">"
	LES  = "<"
	GEQ  = ">="
	LEQ  = "<="
	COR  = "||"
	CAND = "&&"

	// Delimiters
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"
	COMMA  = ","
	COLON  = ":"
	SCOLON = ";"

	// Keywords
	FN     = "FN"     // Function
	LET    = "LET"    // Let (Variable Declare)
	VOL    = "VOL"    // Volitile
	STRUCT = "STRUCT" // Structure
	ENUM   = "ENUM"   // Enumeration
	UNION  = "UNION"  // Union
	CONST  = "CONST"  // Constant
	RETURN = "RETURN" // Return
)

func LookupID(id string) TokenType {
	if tok, ok := keywords[id]; ok {
		return tok
	}
	return ID
}
