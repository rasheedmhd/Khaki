package token

type TokenType string 

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"

	// Identifies and Literals
	IDENT   = "IDENT"
	INT 	= "INT"

	// Operators
	ASSIGN  = "="
	PLUS	= "+"

	// Delimiters
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE	 = "{"
	RBRACE	 = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

)