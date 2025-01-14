// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
package token

// here

type TokenType uint
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota
	EOF
	// Identifiers + literals
	IDENT // add, foobar, x, y, ...
	INT   // 1343456
	// Operators
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	LT
	GT
	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	keyword_beg
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
	keyword_end

	EQ
	NOT_EQ
)

var keywords map[string]TokenType
var Tokens = [...]string{
	// Identifiers + literals
	IDENT: "IDENT", // add
	INT:   "INT",   // 1343456
	// Operators
	ASSIGN:   "=",
	PLUS:     "+",
	MINUS:    "-",
	BANG:     "!",
	ASTERISK: "*",
	SLASH:    "/",
	LT:       "<",
	GT:       ">",
	// Delimiters
	COMMA:     ",",
	SEMICOLON: ";",
	LPAREN:    "(",
	RPAREN:    ")",
	LBRACE:    "{",
	RBRACE:    "}",
	// Keywords
	FUNCTION: "fn",
	LET:      "let",
	TRUE:     "true",
	FALSE:    "false",
	IF:       "if",
	ELSE:     "else",
	RETURN:   "return",
	EQ:       "==",
	NOT_EQ:   "!=",
}
