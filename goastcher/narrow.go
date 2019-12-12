package goastcher

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"regexp"
	"strconv"
	"strings"
)

type (
	allOf struct {
		gs []Goastcher
		id string
	}

	anyOf struct {
		gs []Goastcher
		id string
	}

	unless struct {
		ger Goastcher
		id  string
	}

	asCode struct {
		code string
		id   string
	}

	matchCode struct {
		code *regexp.Regexp
		id   string
	}

	isSize struct {
		size int64
		id   string
	}

	hasOperator struct {
		op string
		id string
	}

	anything struct {
		id string
	}

	isType struct {
		typ string
		id  string
	}

	hasPrefix struct {
		prefix string
		id     string
	}

	equals struct {
		val interface{}
		id  string
	}

	hasSuffix struct {
		suffix string
		id     string
	}

	contains struct {
		substr string
		id     string
	}

	matchString struct {
		re *regexp.Regexp
		id string
	}

	isExported struct {
		id string
	}
)

// AllOf matches if all given matchers match
func AllOf(gs ...Goastcher) Goastcher {
	for _, g := range gs {
		if isErrorGer(g) {
			return g
		}
	}
	return &allOf{gs: gs}
}

// AnyOf matches if any of the given matchers matches
func AnyOf(gs ...Goastcher) Goastcher {
	for _, g := range gs {
		if isErrorGer(g) {
			return g
		}
	}
	return &anyOf{gs: gs}
}

// Unless matches if the provided matcher does not match
func Unless(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &unless{ger: g}
}

func (g *unless) Goastch(ctx *Context, node ast.Node) bool {
	return !g.ger.Goastch(ctx, node)
}

func (g *unless) String() string {
	return fmt.Sprintf("unless %s", g.ger)
}

func (g *unless) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *allOf) Goastch(ctx *Context, node ast.Node) bool {
	for _, g := range g.gs {
		if !g.Goastch(ctx, node) {
			return false
		}
	}
	return true
}

func (g *allOf) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *allOf) String() string {
	buf := bytes.NewBufferString("allOf [")
	buf.WriteString(fmt.Sprintf("%s", g.gs[0]))
	for _, ger := range g.gs[1:] {
		buf.WriteString(", ")
		buf.WriteString(fmt.Sprintf("%s", ger))
	}
	buf.WriteString("]")
	return buf.String()
}

func (g *anyOf) Goastch(ctx *Context, node ast.Node) bool {
	for _, g := range g.gs {
		if g.Goastch(ctx, node) {
			return true
		}
	}
	return false
}

func (g *anyOf) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *anyOf) String() string {
	buf := bytes.NewBufferString("anyOf [")
	buf.WriteString(fmt.Sprint(g.gs[0]))
	for _, g := range g.gs[1:] {
		buf.WriteString(", ")
		buf.WriteString(fmt.Sprint(g))
	}
	buf.WriteString("]")
	return buf.String()
}

// AsCode matches the code of any node that could formated by format.Node
func AsCode(code string) Goastcher {
	return &asCode{code: code}
}

func (g *asCode) Goastch(ctx *Context, node ast.Node) bool {
	var srcCode bytes.Buffer
	dstCode, _ := format.Source([]byte(g.code))
	fset := token.NewFileSet()
	_ = format.Node(&srcCode, fset, node)
	return bytes.Equal(srcCode.Bytes(), dstCode)
}

func (g *asCode) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *asCode) String() string {
	return fmt.Sprintf("asCode {%s}", g.code)
}

// MatchCode matches the code-patten of any node that could formated by
// format.Node
func MatchCode(patten string) Goastcher {
	re, err := regexp.CompilePOSIX(patten)
	if err != nil {
		return ErrorGer(err)
	}
	return &matchCode{code: re}
}

func (g *matchCode) Goastch(ctx *Context, node ast.Node) bool {
	var srcCode bytes.Buffer
	fset := token.NewFileSet()
	_ = format.Node(&srcCode, fset, node)
	return g.code.Match(srcCode.Bytes())
}

func (g *matchCode) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *matchCode) String() string {
	return fmt.Sprintf("asCode {%s}", g.code)
}

// IsSize matches nodes that is the specified size
func IsSize(size int64) Goastcher {
	if size < 0 {
		return ErrorGer(fmt.Errorf("required size greater than 0"))
	}
	return &isSize{size: size}
}

func (g *isSize) Goastch(ctx *Context, node ast.Node) bool {
	switch n := node.(type) {
	case *ast.CompositeLit:
		if len(n.Elts) != int(g.size) {
			return false
		}
	default:
		return false
	}
	return true
}

func (g *isSize) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *isSize) String() string {
	return fmt.Sprintf("isSize %d", g.size)
}

// HasOperator matches the operator of BinaryExpr, UnaryExpr, AssignStmt
func HasOperator(op string) Goastcher {
	return &hasOperator{op: op}
}

func (g *hasOperator) Goastch(ctx *Context, node ast.Node) bool {
	op := ""
	switch n := node.(type) {
	case *ast.BinaryExpr:
		op = n.Op.String()
	case *ast.UnaryExpr:
		op = n.Op.String()
	case *ast.AssignStmt:
		op = n.Tok.String()
	}
	return g.op == op
}

func (g *hasOperator) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *hasOperator) String() string {
	return fmt.Sprintf("hasOperator \"%s\"", g.op)
}

// Anything matches any node
func Anything() Goastcher {
	return &anything{}
}

func (g *anything) Goastch(ctx *Context, node ast.Node) bool {
	return true
}

func (g *anything) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *anything) String() string {
	return "anything"
}

// IsType matches the type of Field, Ident,ValueSpec
func IsType(typ string) Goastcher {
	return &isType{typ: typ}
}

func (g *isType) Goastch(ctx *Context, node ast.Node) bool {
	theTyp := ""
	switch n := node.(type) {
	case *ast.Field:
		theTyp = ctx.Info.Types[n.Type].Type.String()
	case *ast.Ident:
		def := ctx.Info.Defs[n]
		theTyp = def.Type().String()
	case *ast.ValueSpec:
		if n.Type != nil {
			theTyp = ctx.Info.Types[n.Type].Type.String()
		} else {
			theTyp = ctx.Info.Defs[n.Names[0]].Type().String()
		}
	}
	return theTyp == g.typ
}

func (g *isType) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *isType) String() string {
	return fmt.Sprintf("isType \"%s\"", g.typ)
}

// HasPrefix matches the prefix of Ident
func HasPrefix(prefix string) Goastcher {
	return &hasPrefix{prefix: prefix}
}

func (g *hasPrefix) Goastch(ctx *Context, node ast.Node) bool {
	str := ""
	switch n := node.(type) {
	case *ast.Ident:
		str = n.Name
	}
	return strings.HasPrefix(str, g.prefix)
}

func (g *hasPrefix) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *hasPrefix) String() string {
	return fmt.Sprintf("hasPrefix \"%s\"", g.prefix)
}

// Equals matches literals that are equal to the given value
func Equals(val interface{}) Goastcher {
	return &equals{val: val}
}

func (g *equals) Goastch(ctx *Context, node ast.Node) bool {
	var val interface{}
	switch n := node.(type) {
	case *ast.Ident:
		if n != nil {
			val = n.Name
		}
	case *ast.BasicLit:
		rawVal := n.Value
		switch n.Kind {
		case token.INT:
			i, err := strconv.ParseInt(rawVal, 10, 64)
			if err != nil {
				ctx.Err = err
				return false
			}
			val = i
		case token.FLOAT:
			f, err := strconv.ParseFloat(rawVal, 64)
			if err != nil {
				ctx.Err = err
				return false
			}
			val = f
		case token.STRING:
			val = strings.Trim(rawVal, "\"")
			// case token.CHAN:
			// case token.IMAG:
		}
	}
	return val == g.val
}

func (g *equals) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *equals) String() string {
	return fmt.Sprintf("equals \"%v\"", g.val)
}

// HasSuffix matches the suffix of Ident
func HasSuffix(suffix string) Goastcher {
	return &hasSuffix{suffix: suffix}
}

func (g *hasSuffix) Goastch(ctx *Context, node ast.Node) bool {
	str := ""
	switch n := node.(type) {
	case *ast.Ident:
		if n != nil {
			str = n.Name
		}
	}
	return strings.HasSuffix(str, g.suffix)
}

func (g *hasSuffix) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *hasSuffix) String() string {
	return fmt.Sprintf("hasSuffix \"%s\"", g.suffix)
}

// Contains matches Ident that contains the given substr
func Contains(substr string) Goastcher {
	return &contains{substr: substr}
}

func (g *contains) Goastch(ctx *Context, node ast.Node) bool {
	var str string
	switch n := node.(type) {
	case *ast.Ident:
		if n != nil {
			str = n.Name
		}
	}
	return strings.Contains(str, g.substr)
}

func (g *contains) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *contains) String() string {
	return fmt.Sprintf("contains \"%s\"", g.substr)
}

// MatchString matches Ident that match the given patten
func MatchString(patten string) Goastcher {
	re, err := regexp.CompilePOSIX(patten)
	if err != nil {
		return ErrorGer(err)
	}
	return &matchString{re: re}
}

func (g *matchString) Goastch(ctx *Context, node ast.Node) bool {
	var str string
	switch n := node.(type) {
	case *ast.Ident:
		if n != nil {
			str = n.Name
		}
	}
	return g.re.MatchString(str)
}

func (g *matchString) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *matchString) String() string {
	return fmt.Sprintf("match \"%s\"", g.re.String())
}

// IsExported matches Ident that is exported
func IsExported() Goastcher {
	return &isExported{}
}

func (g *isExported) Goastch(ctx *Context, node ast.Node) bool {
	var matched bool
	switch n := node.(type) {
	case *ast.Ident:
		matched = n.IsExported()
	}
	if matched {
		ctx.growBindings(g.id, node)
	}
	return matched
}

func (g *isExported) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *isExported) String() string {
	return "isExported"
}
