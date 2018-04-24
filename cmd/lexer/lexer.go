package lexer

import (
	"bytes"
	"fmt"
	oLexer "github.com/alecthomas/participle/lexer"
	"io"
	"strconv"
	"strings"
	"text/scanner"
	"unicode/utf8"
)

var MatrixDefinition = &matrixDefinition{}

type matrixDefinition struct{}

func (d *matrixDefinition) Lex(r io.Reader) oLexer.Lexer {
	return Lex(r)
}

func (d *matrixDefinition) Symbols() map[string]rune {
	return map[string]rune{
		"EOF":       scanner.EOF,
		"Char":      scanner.Char,
		"Ident":     scanner.Ident,
		"Int":       scanner.Int,
		"Float":     scanner.Float,
		"String":    scanner.String,
		"RawString": scanner.RawString,
		"Comment":   scanner.Comment,
	}
}

// textScannerLexer is a Lexer based on text/scanner.Scanner
type textScannerLexer struct {
	scanner  *scanner.Scanner
	filename string
}

// Lex an io.Reader with text/scanner.Scanner.
//
// This provides very fast lexing of source code compatibile with Go tokens.
//
// Note that this differs from text/scanner.Scanner in that string tokens will be unquoted.
func Lex(r io.Reader) oLexer.Lexer {
	lexer := lexWithScanner(r, &scanner.Scanner{})
	lexer.scanner.Error = func(s *scanner.Scanner, msg string) {
		// This is to support single quoted strings. Hacky.
		if msg != "illegal char literal" {
			oLexer.Panic(oLexer.Position(lexer.scanner.Pos()), msg)
		}
	}
	return oLexer.Upgrade(lexer)
}

// LexWithScanner creates a Lexer from a user-provided scanner.Scanner.
//
// Useful if you need to customise the Scanner.
//
// Note that if this function is used, single-quoted strings are not supported. See the source for
// Lex() for how to achieve this.

func LexWithScanner(r io.Reader, scan *scanner.Scanner) oLexer.Lexer {
	return oLexer.Upgrade(lexWithScanner(r, scan))
}

func lexWithScanner(r io.Reader, scan *scanner.Scanner) *textScannerLexer {
	lexer := &textScannerLexer{
		filename: oLexer.NameOfReader(r),
		scanner:  scan,
	}
	lexer.scanner.Init(r)
	return lexer
}

// LexBytes returns a new default lexer over bytes.
func LexBytes(b []byte) oLexer.Lexer {
	return Lex(bytes.NewReader(b))
}

// LexString returns a new default lexer over a string.
func LexString(s string) oLexer.Lexer {
	return Lex(strings.NewReader(s))
}

func (t *textScannerLexer) Next() oLexer.Token {
	typ := t.scanner.Scan()
	text := t.scanner.TokenText()
	pos := oLexer.Position(t.scanner.Position)
	pos.Filename = t.filename
	out := oLexer.Token{
		Type:  typ,
		Value: text,
		Pos:   pos,
	}
	out.Pos.Filename = t.filename
	// Unquote strings.
	switch out.Type {
	case scanner.Char:
		// FIXME(alec): This is pretty hacky...we convert a single quoted char into a double
		// quoted string in order to support single quoted strings.
		out.Value = fmt.Sprintf("\"%s\"", out.Value[1:len(out.Value)-1])
		fallthrough
	case scanner.String:
		s, err := strconv.Unquote(out.Value)
		if err != nil {
			oLexer.Panic(out.Pos, err.Error())
		}
		out.Value = s
		if out.Type == scanner.Char && utf8.RuneCountInString(s) > 1 {
			out.Type = scanner.String
		}
	case scanner.RawString:
		out.Value = out.Value[1 : len(out.Value)-1]
	}
	return out
}
