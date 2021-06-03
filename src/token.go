package src

import "fmt"

type TokenType int

const (
	Head TokenType = iota
	AND
	ASSING
	COMMA
	DIVISION
	ELSE
	EOF
	EQ
	FALSE
	FOR
	FUNCTION
	GT     // grather than
	GTOREQ // grater than or equeal
	IDENT
	IF
	ILLEGAL
	INT
	LBRACE
	LET
	LPAREN
	LT     // less than
	LTOREQ // less than or equal
	MINUS  // -
	NOT    // !
	NOT_EQ
	MOD
	OR
	PLUS
	RBRACE
	RETURN
	RPAREN
	SEMICOLON
	TIMES // *
	STRING
	TRUE
	WHILE
)

var tokens = [...]string{
	AND:       "AND",
	ASSING:    "ASSING",
	COMMA:     "COMMA",
	DIVISION:  "DIVISION",
	ELSE:      "ELSE",
	EOF:       "EOF",
	EQ:        "EQ",
	FALSE:     "FALSE",
	FUNCTION:  "FUNCTION",
	GT:        "GT",
	IDENT:     "IDENT",
	IF:        "IF",
	ILLEGAL:   "ILLEGAL",
	INT:       "INT",
	LBRACE:    "LBRACE",
	LET:       "LET",
	LPAREN:    "LPAREN",
	LT:        "LT",
	MINUS:     "MINNUS",
	NOT:       "NOT",
	NOT_EQ:    "NOT_EQ",
	MOD:       "MOD",
	OR:        "OR",
	PLUS:      "PLUS",
	RBRACE:    "RBRACE",
	RETURN:    "RETURN",
	RPAREN:    "RPAREN",
	SEMICOLON: "SEMICOLON",
	TIMES:     "TIMES",
	STRING:    "STRING",
	TRUE:      "TRUE",
	WHILE:     "MIENTRAS",
}

type Token struct {
	Token_type TokenType
	Literal    string
}

func (t *Token) PrintToken() string {
	return fmt.Sprintf("Token Type: %s, Literal: %s", tokens[t.Token_type], t.Literal)
}

// verify that given literal is a string
func LookUpTokenType(literal string) TokenType {
	keywords := map[string]TokenType{
		"falso":     FALSE,
		"funcion":   FUNCTION,
		"regresa":   RETURN,
		"si":        IF,
		"si_no":     ELSE,
		"var":       LET,
		"verdadero": TRUE,
		"mientras":  WHILE,
		"para":      FOR,
	}

	TokenType, exists := keywords[literal]
	if exists {
		return TokenType
	}

	return IDENT
}