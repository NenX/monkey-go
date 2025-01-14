// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
package token

func Get_keywords() map[string]TokenType {

	return keywords
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
func init() {
	println("token package init")
	keywords = make(map[string]TokenType, keyword_end-keyword_beg)

	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[Tokens[i]] = i
	}
}
