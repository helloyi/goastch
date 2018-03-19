// Package parser of goastch dsl
//
// syntax:
// ger       = nodeGer | narrowGer | travelGer | logicGer .
// nodeGer   = nodeObj (travelGer | narrowGer | logicGer) .
// travelGer = travelObj (nodeGer | logicGer) .
// narrowGer = narrowObj value .
// logicGer  = unaryGer | compositeGer .
// unaryGer  = unaryObj ger .
// compositeGer = compositeObj [unaryObj] object [ger] {'and' [unaryObj] object [ger]} .
//
// unaryObj  = 'unless'
// compositeObj = 'anyof' | 'allof'
// TODO
// 1. bind on logicger
// 2. binding key
package parser

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/helloyi/goastch/galang/ast"
	"github.com/helloyi/goastch/galang/scanner"
	"github.com/helloyi/goastch/galang/token"
)

// parser ...
type parser struct {
	scanner *scanner.Scanner

	pos scanner.Position
	tok token.Token
	lit string
}

func newParser(s *scanner.Scanner) *parser {
	p := &parser{
		scanner: s,
	}
	p.next()
	return p
}

func (p *parser) next() {
	p.tok = p.scanner.Scan()
	p.lit = p.scanner.TokenText()
	p.pos = p.scanner.Pos()
}

// ParseStmt ...
func ParseStmt(src string) (*ast.Stmt, error) {
	s := scanner.New(strings.NewReader(src))
	p := newParser(s)
	return p.parseStmt()
}

// ParseGer ...
func ParseGer(ger string) (ast.Ger, error) {
	return ParseGerFrom(strings.NewReader(ger))
}

// ParseGerFrom ...
func ParseGerFrom(reader io.Reader) (ast.Ger, error) {
	s := scanner.New(reader)
	p := newParser(s)
	return p.parseGer()
}

func (p *parser) parseStmt() (*ast.Stmt, error) {
	ger, err := p.parseGer()
	if err != nil {
		return nil, err
	}
	if !p.expected(token.From) {
		return nil, p.fmtErr("expected 'from'")
	}
	path, err := p.parsePath()
	if err != nil {
		return nil, err
	}

	var isPkg bool
	if file, _ := os.Stat(path); file.IsDir() {
		isPkg = true
	}

	return &ast.Stmt{
		Ger:   ger,
		Path:  path,
		IsPkg: isPkg,
	}, nil
}

func (p *parser) parseGer() (ast.Ger, error) {
	isBind := p.parseBind()
	switch p.tok {
	case token.Node:
		return p.parseNodeGer(isBind)
	case token.Narrow:
		return p.parseNarrowGer()
	case token.Travel:
		return p.parseTravelGer(isBind)
	case token.Unless:
		return p.parseUnlessGer()
	case token.AllOf, token.AnyOf:
		return p.parseLogicGer()
	default:
		return &ast.EmptyGer{Pos: p.pos}, nil
	}
}

func (p *parser) parseBind() bool {
	if p.tok != token.Bind {
		return false
	}
	p.next()
	return true
}

func (p *parser) parseNodeGer(isBind bool) (*ast.NodeGer, error) {
	node := p.parseIdent()
	next, err := p.parseGer()

	if err != nil {
		return nil, err
	}

	return &ast.NodeGer{
		Node:   node,
		IsBind: isBind,
		Next:   next,
	}, nil
}

func (p *parser) parseTravelGer(isBind bool) (*ast.TravelGer, error) {
	travel := p.parseIdent()
	next, err := p.parseGer()
	if err != nil {
		return nil, err
	}
	return &ast.TravelGer{
		Travel: travel,
		Next:   next,
		IsBind: isBind,
	}, nil
}

func (p *parser) parseNarrowGer() (*ast.NarrowGer, error) {
	narrow := p.parseIdent()
	args, err := p.parseArgs()
	if err != nil {
		return nil, err
	}
	return &ast.NarrowGer{
		Narrow: narrow,
		Args:   args,
	}, nil
}

func (p *parser) parseLogicGer() (*ast.LogicGer, error) {
	op := p.parseIdent()
	var ger ast.Ger
	var err error
	ger, err = p.parseCompositeGer()
	if err != nil {
		return nil, err
	}
	return &ast.LogicGer{
		Op:  op,
		Ger: ger,
	}, nil
}

func (p *parser) parseUnlessGer() (*ast.LogicGer, error) {
	unless := p.parseIdent()
	ger, err := p.parseGer()
	if err != nil {
		return nil, err
	}
	return &ast.LogicGer{
		Op:  unless,
		Ger: ger,
	}, nil
}

func (p *parser) parseCompositeGer() (*ast.CompositeGer, error) {
	gers := make([]ast.Ger, 0)
	ger, err := p.parseGer()
	if err != nil {
		return nil, err
	}
	gers = append(gers, ger)
	for {
		if !p.expected(token.And) {
			break
		}
		ger, err = p.parseGer()
		if err != nil {
			return nil, err
		}
		gers = append(gers, ger)
	}
	if err := p.resolve(gers); err != nil {
		return nil, err
	}
	if p.tok == token.End {
		p.next()
	}
	return &ast.CompositeGer{Gers: gers}, nil
}

func (p *parser) findLameGer(ger ast.Ger) ast.Ger {
	var next ast.Ger
	switch g := ger.(type) {
	case *ast.LogicGer:
		if g.Op.Token == token.Unless {
			return p.findLameGer(g.Ger)
		}
		next = g.Ger
	case *ast.NodeGer:
		next = g.Next
	case *ast.TravelGer:
		next = g.Next
	}
	if _, ok := next.(*ast.EmptyGer); ok {
		return ger
	}
	return nil
}

func (p *parser) findSaneGer(ger ast.Ger) ast.Ger {
	var next ast.Ger
	switch g := ger.(type) {
	case *ast.LogicGer:
		if g.Op.Token == token.Unless {
			return p.findSaneGer(g.Ger)
		}
		next = g.Ger
	case *ast.NodeGer:
		next = g.Next
	case *ast.TravelGer:
		next = g.Next
	}
	if _, ok := next.(*ast.EmptyGer); !ok {
		return ger
	}
	return nil
}

func (p *parser) resolve0(unresolved []ast.Ger, solution ast.Ger) {
	for _, ger := range unresolved {
		switch g := ger.(type) {
		case *ast.LogicGer:
			g.Ger = solution
		case *ast.NodeGer:
			g.Next = solution
		case *ast.TravelGer:
			g.Next = solution
		default:
			// fatal error (program bug)
			panic("invalid unresolved goastcher")
		}
	}
}

func (p *parser) resolve(gers []ast.Ger) error {
	unresolved := make([]ast.Ger, 0)
	for _, ger := range gers {
		if lameGer := p.findLameGer(ger); lameGer != nil {
			unresolved = append(unresolved, lameGer)
			continue
		}
		saneGer := p.findSaneGer(ger)
		p.resolve0(unresolved, saneGer)
		unresolved = unresolved[:0]
	}

	if len(unresolved) != 0 {
		return p.fmtErr("unresolved composite goastcher")
	}

	return nil
}

func (p *parser) expected(tok token.Token) bool {
	if p.tok != tok {
		return false
	}
	p.next()
	return true
}

func (p *parser) parseArgs() ([]*ast.ValueSpec, error) {
	args := make([]*ast.ValueSpec, 0)
	for {
		switch p.tok {
		case token.Int, token.Float, token.Char, token.String:
			value, err := p.ParseValueSpec()
			if err != nil {
				return nil, err
			}
			args = append(args, value)
		default:
			if len(args) == 0 {
				return nil, nil
			}
			return args, nil
		}
	}
}

func (p *parser) ParseValueSpec() (*ast.ValueSpec, error) {
	var value interface{}
	var err error
	switch p.tok {
	case token.Int:
		value, err = p.parseIntValue()
	case token.Float:
		value, err = p.parseFloatValue()
	case token.Char:
		value = p.parseCharValue()
	case token.String:
		value = p.parseStringValue()
	default:
		err = p.fmtErr("unsupported narrow argument")
	}
	return &ast.ValueSpec{Value: value}, err
}

func (p *parser) parseIntValue() (int64, error) {
	i, err := strconv.ParseInt(p.lit, 10, 64)
	if err != nil {
		return 0, p.fmtErr(err.Error())
	}
	p.next()
	return i, nil
}

func (p *parser) parseFloatValue() (float64, error) {
	f, err := strconv.ParseFloat(p.lit, 64)
	if err != nil {
		return 0, p.fmtErr(err.Error())
	}
	p.next()
	return f, nil
}

func (p *parser) parseCharValue() byte {
	c := strings.Trim(p.lit, "'")
	p.next()
	return c[0]
}

func (p *parser) parseStringValue() string {
	s := strings.Trim(p.lit, "\"")
	p.next()
	return s
}

func (p *parser) parseIdent() *ast.Ident {
	obj := token.Object(p.lit)
	if obj == nil {
		panic("invalid ident")
	}
	ident := &ast.Ident{
		Pos:    p.pos,
		Name:   p.lit,
		Token:  p.tok,
		Object: obj,
	}
	p.next()
	return ident
}

func (p *parser) parsePath() (string, error) {
	path := strings.Trim(p.lit, "\"'")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", p.fmtErr(path + " not exist")
	}
	return path, nil
}

func (p *parser) fmtErr(msg string) error {
	return fmt.Errorf("%s:%s", p.pos, msg)
}
