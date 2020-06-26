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
	// PLUS represents an addition operator
	PLUS = "+"
	// MINUS represents a substraction operator
	MINUS = "-"
	// BANG represents a !
	BANG = "!"
	// ASTERIX represents a *
	ASTERIX = "*"
	// SLASH represents a /
	SLASH = "/"
	// EQ represents a =
	EQ = "=="
	// NOT_EQ represents a !=
	NOT_EQ = "!="

	// LT represents a lesser than operator
	LT = "<"
	// GT represents a greater than operator
	GT = ">"

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

	// FUNCTION represents a function keyword
	FUNCTION = "FUNCTION"

	// LET represents a let keywotd
	LET = "LET"

	// TRUE represents a true boolean
	TRUE = "TRUE"

	// FALSE represents a false boolean
	FALSE = "FALSE"

	// IF represent a conditional keyword
	IF = "IF"

	// ELSE represents an else keyword
	ELSE = "ELSE"

	// RETURN represents a return keyword
	RETURN = "RETURN"
)

// Type represents a type of token
type Type string

// Token represents a lexed token
type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks if i is a keyword
func LookupIdent(i string) Type {
	if tok, ok := keywords[i]; ok {
		return tok
	}
	return IDEN
}
