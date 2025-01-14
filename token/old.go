// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
package token

import (
	"strconv"
	"unicode"
	"unicode/utf8"
)

// TokenInt is the set of lexical tokens of the Go programming language.
type TokenInt int

// The list of tokens.
const (
	// Special tokens
	ILLEGAL__ TokenInt = iota
	EOF__
	COMMENT__

	literal_beg__
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT__  // main
	INT__    // 12345
	FLOAT__  // 123.45
	IMAG__   // 123.45i
	CHAR__   // 'a'
	STRING__ // "abc"
	literal_end__

	operator_beg__
	// Operators and delimiters
	ADD__ // +
	SUB__ // -
	MUL__ // *
	QUO__ // /
	REM__ // %

	AND__     // &
	OR__      // |
	XOR__     // ^
	SHL__     // <<
	SHR__     // >>
	AND_NOT__ // &^

	ADD_ASSIGN__ // +=
	SUB_ASSIGN__ // -=
	MUL_ASSIGN__ // *=
	QUO_ASSIGN__ // /=
	REM_ASSIGN__ // %=

	AND_ASSIGN__     // &=
	OR_ASSIGN__      // |=
	XOR_ASSIGN__     // ^=
	SHL_ASSIGN__     // <<=
	SHR_ASSIGN__     // >>=
	AND_NOT_ASSIGN__ // &^=

	LAND__  // &&
	LOR__   // ||
	ARROW__ // <-
	INC__   // ++
	DEC__   // --

	EQL__    // ==
	LSS__    // <
	GTR__    // >
	ASSIGN__ // =
	NOT__    // !

	NEQ__      // !=
	LEQ__      // <=
	GEQ__      // >=
	DEFINE__   // :=
	ELLIPSIS__ // ...

	LPAREN__ // (
	LBRACK   // [
	LBRACE__ // {
	COMMA__  // ,
	PERIOD   // .

	RPAREN__    // )
	RBRACK__    // ]
	RBRACE__    // }
	SEMICOLON__ // ;
	COLON__     // :
	operator_end__

	keyword_beg__
	// Keywords
	BREAK__
	CASE__
	CHAN__
	CONST__
	CONTINUE__

	DEFAULT__
	DEFER__
	ELSE__
	FALLTHROUGH__
	FOR__

	FUNC__
	GO__
	GOTO__
	IF__
	IMPORT__

	INTERFACE__
	MAP__
	PACKAGE__
	RANGE__
	RETURN__

	SELECT__
	STRUCT__
	SWITCH__
	TYPE__
	VAR__
	keyword_end__

	additional_beg__
	// additional tokens, handled in an ad-hoc manner
	TILDE__
	additional_end__
)

var tokens__ = [...]string{
	ILLEGAL__: "ILLEGAL",

	EOF__:     "EOF",
	COMMENT__: "COMMENT",

	IDENT__:  "IDENT",
	INT__:    "INT",
	FLOAT__:  "FLOAT",
	IMAG__:   "IMAG",
	CHAR__:   "CHAR",
	STRING__: "STRING",

	ADD__: "+",
	SUB__: "-",
	MUL__: "*",
	QUO__: "/",
	REM__: "%",

	AND__:     "&",
	OR__:      "|",
	XOR__:     "^",
	SHL__:     "<<",
	SHR__:     ">>",
	AND_NOT__: "&^",

	ADD_ASSIGN__: "+=",
	SUB_ASSIGN__: "-=",
	MUL_ASSIGN__: "*=",
	QUO_ASSIGN__: "/=",
	REM_ASSIGN__: "%=",

	AND_ASSIGN__:     "&=",
	OR_ASSIGN__:      "|=",
	XOR_ASSIGN__:     "^=",
	SHL_ASSIGN__:     "<<=",
	SHR_ASSIGN__:     ">>=",
	AND_NOT_ASSIGN__: "&^=",

	LAND__:  "&&",
	LOR__:   "||",
	ARROW__: "<-",
	INC__:   "++",
	DEC__:   "--",

	EQL__:    "==",
	LSS__:    "<",
	GTR__:    ">",
	ASSIGN__: "=",
	NOT__:    "!",

	NEQ__:      "!=",
	LEQ__:      "<=",
	GEQ__:      ">=",
	DEFINE__:   ":=",
	ELLIPSIS__: "...",

	LPAREN__: "(",
	LBRACK:   "[",
	LBRACE__: "{",
	COMMA__:  ",",
	PERIOD:   ".",

	RPAREN__:    ")",
	RBRACK__:    "]",
	RBRACE__:    "}",
	SEMICOLON__: ";",
	COLON__:     ":",

	BREAK__:    "break",
	CASE__:     "case",
	CHAN__:     "chan",
	CONST__:    "const",
	CONTINUE__: "continue",

	DEFAULT__:     "default",
	DEFER__:       "defer",
	ELSE__:        "else",
	FALLTHROUGH__: "fallthrough",
	FOR__:         "for",

	FUNC__:   "func",
	GO__:     "go",
	GOTO__:   "goto",
	IF__:     "if",
	IMPORT__: "import",

	INTERFACE__: "interface",
	MAP__:       "map",
	PACKAGE__:   "package",
	RANGE__:     "range",
	RETURN__:    "return",

	SELECT__: "select",
	STRUCT__: "struct",
	SWITCH__: "switch",
	TYPE__:   "type",
	VAR__:    "var",

	TILDE__: "~",
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token [ADD__], the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token [IDENT__], the string is "IDENT").
func (tok TokenInt) String() string {
	s := ""
	if 0 <= tok && tok < TokenInt(len(tokens__)) {
		s = tokens__[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
func (op TokenInt) Precedence() int {
	switch op {
	case LOR__:
		return 1
	case LAND__:
		return 2
	case EQL__, NEQ__, LSS__, LEQ__, GTR__, GEQ__:
		return 3
	case ADD__, SUB__, OR__, XOR__:
		return 4
	case MUL__, QUO__, REM__, SHL__, SHR__, AND__, AND_NOT__:
		return 5
	}
	return LowestPrec
}

var keywords__ map[string]TokenInt

func init() {
	keywords__ = make(map[string]TokenInt, keyword_end__-(keyword_beg__+1))
	for i := keyword_beg__ + 1; i < keyword_end__; i++ {
		keywords__[tokens__[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or [IDENT__] (if not a keyword).
func Lookup(ident string) TokenInt {
	if tok, is_keyword := keywords__[ident]; is_keyword {
		return tok
	}
	return IDENT__
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
func (tok TokenInt) IsLiteral() bool { return literal_beg__ < tok && tok < literal_end__ }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
func (tok TokenInt) IsOperator() bool {
	return (operator_beg__ < tok && tok < operator_end__) || tok == TILDE__
}

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
func (tok TokenInt) IsKeyword() bool { return keyword_beg__ < tok && tok < keyword_end__ }

// IsExported reports whether name starts with an upper-case letter.
func IsExported(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}

// IsKeyword reports whether name is a Go keyword, such as "func" or "return".
func IsKeyword(name string) bool {
	// TODO: opt: use a perfect hash function instead of a global map.
	_, ok := keywords__[name]
	return ok
}

// IsIdentifier reports whether name is a Go identifier, that is, a non-empty
// string made up of letters, digits, and underscores, where the first character
// is not a digit. Keywords are not identifiers.
func IsIdentifier(name string) bool {
	if name == "" || IsKeyword(name) {
		return false
	}
	for i, c := range name {
		if !unicode.IsLetter(c) && c != '_' && (i == 0 || !unicode.IsDigit(c)) {
			return false
		}
	}
	return true
}
