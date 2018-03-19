package scanner

import (
	"io"
	"text/scanner"

	"github.com/helloyi/goastch/galang/token"
)

// Scanner ...
type Scanner struct {
	scanner.Scanner
}

// Position ...
type Position struct {
	scanner.Position
}

// New ...
func New(src io.Reader) *Scanner {
	s := scanner.Scanner{
		Mode: scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats |
			scanner.ScanChars | scanner.ScanStrings | scanner.ScanRawStrings,
	}
	s.Init(src)
	return &Scanner{
		Scanner: s,
	}
}

// Scan ...
func (s *Scanner) Scan() token.Token {
	return s.token(s.Scanner.Scan())
}

// Pos ...
func (s *Scanner) Pos() Position {
	return Position{s.Scanner.Pos()}
}

func (s *Scanner) token(tok rune) token.Token {
	switch tok {
	case scanner.EOF:
		return token.EOF
	case scanner.Int:
		return token.Int
	case scanner.Float:
		return token.Float
	case scanner.Char:
		return token.Char
	case scanner.String:
		return token.String
	default:
		text := s.TokenText()
		tok := token.What(text)
		if tok == token.Unknow {
			return token.ILLEGAL
		}
		return tok
	}
}

// Next ...
func (s *Scanner) Next() token.Token {
	return s.token(s.Scanner.Next())
}

// TokenString ...
func (s *Scanner) TokenString() string {
	return s.TokenText()
}
