package token

const (
	// ILLEGAL represents a token not known by the interpreter
	ILLEGAL = "ILLEGAL"
	// EOF represents an end of file
	EOF = "EOF"

	// IDEN represents an identifier token
	IDEN = "IDEN"
	// INT represents an integer number
	INT = "INT"

	// ASSIGN represents an assignment operator
	ASSIGN = "="
	// PLUS represents an addition operators
	PLUS = "+"

	// COMMA represents a comma
	COMMA = ","
	// SEMICOLON represents a comma
	SEMICOLON = ";"

	// LPAREN represents a left parenthesis
	LPAREN = "("
	// RPAREN represents a right parenthesis
	RPAREN = "("
	// LBRACE represents a left curly brace
	LBRACE = "{"
	// RBRACE represents a right curly brace
	RBRACE = "}"

	// FUNCTION represents a function token
	FUNCTION = "FUNCTION"
	// LET represents a declaration token
	LET = "LET"
)

// Type represents a type of token
type Type string

// Token represents a lexed token
type Token struct {
	Type    Type
	Literal string
}
