// goastcher impliments node goastcher
package goastcher

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/helloyi/goastch/types"
)

type (
	node struct {
		next Goastcher
		id   string
	}

	// purebred node matchers
	arrayType      struct{ node }
	assignStmt     struct{ node }
	badDecl        struct{ node }
	badExpr        struct{ node }
	badStmt        struct{ node }
	basicLit       struct{ node }
	binaryExpr     struct{ node }
	blockStmt      struct{ node }
	branchStmt     struct{ node }
	callExpr       struct{ node }
	caseClause     struct{ node }
	chanType       struct{ node }
	commClause     struct{ node }
	comment        struct{ node }
	commentGroup   struct{ node }
	compositeLit   struct{ node }
	declStmt       struct{ node }
	deferStmt      struct{ node }
	ellipsis       struct{ node }
	emptyStmt      struct{ node }
	exprStmt       struct{ node }
	field          struct{ node }
	fieldList      struct{ node }
	file           struct{ node }
	forStmt        struct{ node }
	funcDecl       struct{ node }
	funcLit        struct{ node }
	funcType       struct{ node }
	genDecl        struct{ node }
	goStmt         struct{ node }
	ident          struct{ node }
	ifStmt         struct{ node }
	importSpec     struct{ node }
	incDecStmt     struct{ node }
	indexExpr      struct{ node }
	interfaceType  struct{ node }
	keyValueExpr   struct{ node }
	labeledStmt    struct{ node }
	mapType        struct{ node }
	pkg            struct{ node }
	parenExpr      struct{ node }
	rangeStmt      struct{ node }
	returnStmt     struct{ node }
	selectStmt     struct{ node }
	selectorExpr   struct{ node }
	sendStmt       struct{ node }
	sliceExpr      struct{ node }
	starExpr       struct{ node }
	structType     struct{ node }
	switchStmt     struct{ node }
	typeAssertExpr struct{ node }
	typeSpec       struct{ node }
	typeSwitchStmt struct{ node }
	unaryExpr      struct{ node }
	valueSpec      struct{ node }

	// varietal node matchers
	shortVarDecl   struct{ node }
	sliceType      struct{ node }
	intBasicLit    struct{ node }
	floatBasicLit  struct{ node }
	imagBasicLit   struct{ node }
	charBasicLit   struct{ node }
	stringBasicLit struct{ node }
)

// StringBasicLit matches string literals
func StringBasicLit(g Goastcher) Goastcher {
	if _, is := g.(*errorGer); is {
		return g
	}
	return &stringBasicLit{node: newNode(g)}
}

func (g *stringBasicLit) Goastch(ctx *Context, node ast.Node) bool {
	n, ok := node.(*ast.BasicLit)
	if !ok {
		return false
	}
	if n.Kind != token.STRING {
		return false
	}

	if !g.next.Goastch(ctx, node) {
		return false
	}

	ctx.growBindings(g.id, node)
	return true
}

func (g *stringBasicLit) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *stringBasicLit) String() string {
	return fmt.Sprintf("stringBasicLit %s", g.next)
}

// CharBasicLit matches char literals
func CharBasicLit(g Goastcher) Goastcher {
	if _, is := g.(*errorGer); is {
		return g
	}
	return &charBasicLit{node: newNode(g)}
}

func (g *charBasicLit) Goastch(ctx *Context, node ast.Node) bool {
	n, ok := node.(*ast.BasicLit)
	if !ok {
		return false
	}
	if n.Kind != token.CHAR {
		return false
	}

	if !g.next.Goastch(ctx, node) {
		return false
	}

	ctx.growBindings(g.id, node)
	return true
}

func (g *charBasicLit) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *charBasicLit) String() string {
	return fmt.Sprintf("charBasicLit %s", g.next)
}

// ImagBasicLit matches imag literals
func ImagBasicLit(g Goastcher) Goastcher {
	if _, is := g.(*errorGer); is {
		return g
	}
	return &imagBasicLit{node: newNode(g)}
}

func (g *imagBasicLit) Goastch(ctx *Context, node ast.Node) bool {
	n, ok := node.(*ast.BasicLit)
	if !ok {
		return false
	}
	if n.Kind != token.IMAG {
		return false
	}

	if !g.next.Goastch(ctx, node) {
		return false
	}

	ctx.growBindings(g.id, node)
	return true
}

func (g *imagBasicLit) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *imagBasicLit) String() string {
	return fmt.Sprintf("imagBasicLit %s", g.next)
}

// FloatBasicLit matches float literals
func FloatBasicLit(g Goastcher) Goastcher {
	if _, is := g.(*errorGer); is {
		return g
	}
	return &floatBasicLit{node: newNode(g)}
}

func (g *floatBasicLit) Goastch(ctx *Context, node ast.Node) bool {
	n, ok := node.(*ast.BasicLit)
	if !ok {
		return false
	}
	if n.Kind != token.FLOAT {
		return false
	}

	if !g.next.Goastch(ctx, node) {
		return false
	}

	ctx.growBindings(g.id, node)
	return true
}

func (g *floatBasicLit) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *floatBasicLit) String() string {
	return fmt.Sprintf("floatBasicLit %s", g.next)
}

// IntBasicLit matches int literals
func IntBasicLit(g Goastcher) Goastcher {
	if _, is := g.(*errorGer); is {
		return g
	}
	return &intBasicLit{node: newNode(g)}
}

func (g *intBasicLit) Goastch(ctx *Context, node ast.Node) bool {
	n, ok := node.(*ast.BasicLit)
	if !ok {
		return false
	}
	if n.Kind != token.INT {
		return false
	}

	if !g.next.Goastch(ctx, node) {
		return false
	}

	ctx.growBindings(g.id, node)
	return true
}

func (g *intBasicLit) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *intBasicLit) String() string {
	return fmt.Sprintf("intBasicLit %s", g.next)
}

// SliceType matches slice type node
func SliceType(g Goastcher) Goastcher {
	if _, is := g.(*errorGer); is {
		return g
	}
	return &sliceType{node: newNode(g)}
}

func (g *sliceType) Goastch(ctx *Context, node ast.Node) bool {
	if types.Type(node) != types.ArrayType {
		return false
	}
	n := node.(*ast.ArrayType)
	if n.Len != nil {
		return false
	}

	if !g.next.Goastch(ctx, node) {
		return false
	}

	ctx.growBindings(g.id, node)
	return true
}

func (g *sliceType) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *sliceType) String() string {
	return fmt.Sprintf("sliceType %s", g.next)
}

// ShortVarDecl matches short variable declare
func ShortVarDecl(g Goastcher) Goastcher {
	if _, is := g.(*errorGer); is {
		return g
	}
	return &shortVarDecl{node: newNode(g)}
}

func (g *shortVarDecl) Goastch(ctx *Context, node ast.Node) bool {
	if types.Type(node) != types.AssignStmt {
		return false
	}
	n := node.(*ast.AssignStmt)
	if n.Tok.String() != ":=" {
		return false
	}

	if !g.next.Goastch(ctx, node) {
		return false
	}

	ctx.growBindings(g.id, node)
	return true
}

func (g *shortVarDecl) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *shortVarDecl) String() string {
	return fmt.Sprintf("shortVarDecl %s", g.next)
}

// newNode ...
func newNode(g Goastcher) node {
	return node{next: g}
}

// match ...
func (g *node) match(ctx *Context, n ast.Node, typ types.Typ) bool {
	if types.Type(n) != typ {
		return false
	}
	if !g.next.Goastch(ctx, n) {
		return false
	}

	ctx.growBindings(g.id, n)
	return true
}

func (g *node) bind(id string, ger Goastcher) Goastcher {
	if isErrorGer(ger) {
		return ger
	}

	if id == "" {
		g.id = fmt.Sprintf("%p", ger)
	} else {
		g.id = id
	}

	return ger
}

// ArrayType matches array type node
//
// []Type
// [N]Type
// [...]Type
func ArrayType(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &arrayType{node: newNode(next)}
}

// AssignStmt matches assignment/sort-variable-declaration node
func AssignStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &assignStmt{node: newNode(next)}
}

// BadDecl matches bad declaration node
func BadDecl(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &badDecl{node: newNode(next)}
}

// BadExpr matches bad expression node
func BadExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &badExpr{node: newNode(next)}
}

// BadStmt matches bad statement node
func BadStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &badStmt{node: newNode(next)}
}

// BasicLit matches literals for basic type
func BasicLit(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &basicLit{node: newNode(next)}
}

// BinaryExpr matches binary expressions
func BinaryExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &binaryExpr{node: newNode(next)}
}

// BlockStmt matches braced statement list
func BlockStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &blockStmt{node: newNode(next)}
}

// BranchStmt matches break, continue, goto, or fallthrough statement
func BranchStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &branchStmt{node: newNode(next)}
}

// CallExpr matches expression followed by an argument list
func CallExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &callExpr{node: newNode(next)}
}

// CaseClause matches case of an expression or type switch statement
func CaseClause(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &caseClause{node: newNode(next)}
}

// ChanType matches channel type
func ChanType(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &chanType{node: newNode(next)}
}

// CommClause matches case of a select statement
func CommClause(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &commClause{node: newNode(next)}
}

// Comment matches single //-style or /*-style comment
func Comment(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &comment{node: newNode(next)}
}

// CommentGroup matches sequence of comments with no other tokens and no empty
// lines between
func CommentGroup(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &commentGroup{node: newNode(next)}
}

// CompositeLit matches composite literal
func CompositeLit(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &compositeLit{node: newNode(next)}
}

// DeclStmt matches declaration in a statement list
func DeclStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &declStmt{node: newNode(next)}
}

// DeferStmt matches defer statement
func DeferStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &deferStmt{node: newNode(next)}
}

// Ellipsis matches node for the "..." type in a parameter list or the "..."
// length in an array type
func Ellipsis(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &ellipsis{node: newNode(next)}
}

// EmptyStmt matches explicit or implicit semicolon
func EmptyStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &emptyStmt{node: newNode(next)}
}

// ExprStmt matches stand-alone expression in a statement list
func ExprStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &exprStmt{node: newNode(next)}
}

// Field matches field declaration list in a struct type, method list in an
// interface type, or a parameter/result declaration in a signature
func Field(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &field{node: newNode(next)}
}

// FieldList matches fields list, enclosed by parentheses or braces
func FieldList(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &fieldList{node: newNode(next)}
}

// File matches node that represents a Go source file
func File(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &file{node: newNode(next)}
}

// ForStmt matches for statement
func ForStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &forStmt{node: newNode(next)}
}

// FuncDecl matches function declaration
func FuncDecl(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &funcDecl{node: newNode(next)}
}

// FuncLit matches function literal
func FuncLit(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &funcLit{node: newNode(next)}
}

// FuncType matches function type
func FuncType(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &funcType{node: newNode(next)}
}

// GenDecl matches import, constant, type or variable declaration
func GenDecl(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &genDecl{node: newNode(next)}
}

// GoStmt matches go statement
func GoStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &goStmt{node: newNode(next)}
}

// Ident matches identifier
func Ident(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &ident{node: newNode(next)}
}

// IfStmt matches if stamtement
func IfStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &ifStmt{node: newNode(next)}
}

// ImportSpec matches single package import specification
func ImportSpec(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &importSpec{node: newNode(next)}
}

// IncDecStmt matches increment or decrement statement
func IncDecStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &incDecStmt{node: newNode(next)}
}

// IndexExpr matches expression followed by an index
func IndexExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &indexExpr{node: newNode(next)}
}

// InterfaceType matches interface type
func InterfaceType(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &interfaceType{node: newNode(next)}
}

// KeyValueExpr matches (key : value) pairs in composite literals
func KeyValueExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &keyValueExpr{node: newNode(next)}
}

// LabeledStmt matches labeled statement
func LabeledStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &labeledStmt{node: newNode(next)}
}

// MapType matches map type
func MapType(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &mapType{node: newNode(next)}
}

// Pkg matches Go package
func Pkg(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &pkg{node: newNode(next)}
}

// ParenExpr matches parenthesized expression
func ParenExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &parenExpr{node: newNode(next)}
}

// RangeStmt matches for statement with a range clause
func RangeStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &rangeStmt{node: newNode(next)}
}

// ReturnStmt matches return statement
func ReturnStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &returnStmt{node: newNode(next)}
}

// SelectStmt matches select statement
func SelectStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &selectStmt{node: newNode(next)}
}

// SelectorExpr matches expression followed by a selector
func SelectorExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &selectorExpr{node: newNode(next)}
}

// SendStmt matches send statement
func SendStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &sendStmt{node: newNode(next)}
}

// SliceExpr matches expression followed by slice indices
func SliceExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &sliceExpr{node: newNode(next)}
}

// StarExpr matches expression of the form "*" Expression. Semantically it could
// be a unary "*" expression, or a pointer type
func StarExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &starExpr{node: newNode(next)}
}

// StructType matches struct type
func StructType(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &structType{node: newNode(next)}
}

// SwitchStmt matches switch statement
func SwitchStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &switchStmt{node: newNode(next)}
}

// TypeAssertExpr matches expression followed by a type assertion
func TypeAssertExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &typeAssertExpr{node: newNode(next)}
}

// TypeSpec matches type declaration
func TypeSpec(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &typeSpec{node: newNode(next)}
}

// TypeSwitchStmt matches type switch statement
func TypeSwitchStmt(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &typeSwitchStmt{node: newNode(next)}
}

// UnaryExpr unary expression. Unary "*" expressions are matched via StarExpr
func UnaryExpr(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &unaryExpr{node: newNode(next)}
}

// ValueSpec constant or variable declaration
func ValueSpec(next Goastcher) Goastcher {
	if isErrorGer(next) {
		return next
	}
	return &valueSpec{node: newNode(next)}
}

func (g *arrayType) Goastch(ctx *Context, node ast.Node) bool {
	if types.Type(node) != types.ArrayType {
		return false
	}
	n := node.(*ast.ArrayType)
	if n.Len == nil {
		return false
	}
	if !g.next.Goastch(ctx, node) {
		return false
	}

	ctx.growBindings(g.id, n)
	return true
}

func (g *assignStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.AssignStmt)
}

func (g *badDecl) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.BadDecl)
}

func (g *badExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.BadExpr)
}

func (g *badStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.BadStmt)
}

func (g *basicLit) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.BasicLit)
}

func (g *binaryExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.BinaryExpr)
}

func (g *blockStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.BlockStmt)
}

func (g *branchStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.BranchStmt)
}

func (g *callExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.CallExpr)
}

func (g *caseClause) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.CaseClause)
}

func (g *chanType) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.ChanType)
}

func (g *commClause) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.CommClause)
}

func (g *comment) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.Comment)
}

func (g *commentGroup) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.CommentGroup)
}

func (g *compositeLit) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.CompositeLit)
}

func (g *declStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.DeclStmt)
}

func (g *deferStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.DeferStmt)
}

func (g *ellipsis) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.Ellipsis)
}

func (g *emptyStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.EmptyStmt)
}

func (g *exprStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.ExprStmt)
}

func (g *field) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.Field)
}

func (g *fieldList) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.FieldList)
}

func (g *file) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.File)
}

func (g *forStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.ForStmt)
}

func (g *funcDecl) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.FuncDecl)
}

func (g *funcLit) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.FuncLit)
}

func (g *funcType) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.FuncType)
}

func (g *genDecl) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.GenDecl)
}

func (g *goStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.GoStmt)
}

func (g *ident) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.Ident)
}

func (g *ifStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.IfStmt)
}

func (g *importSpec) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.ImportSpec)
}

func (g *incDecStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.IncDecStmt)
}

func (g *indexExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.IndexExpr)
}

func (g *interfaceType) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.InterfaceType)
}

func (g *keyValueExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.KeyValueExpr)
}

func (g *labeledStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.LabeledStmt)
}

func (g *mapType) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.MapType)
}

func (g *pkg) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.Package)
}

func (g *parenExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.ParenExpr)
}

func (g *rangeStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.RangeStmt)
}

func (g *returnStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.ReturnStmt)
}

func (g *selectStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.SelectStmt)
}

func (g *selectorExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.SelectorExpr)
}

func (g *sendStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.SendStmt)
}

func (g *sliceExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.SliceExpr)
}

func (g *starExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.StarExpr)
}

func (g *structType) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.StructType)
}

func (g *switchStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.SwitchStmt)
}

func (g *typeAssertExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.TypeAssertExpr)
}

func (g *typeSpec) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.TypeSpec)
}

func (g *typeSwitchStmt) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.TypeSwitchStmt)
}

func (g *unaryExpr) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.UnaryExpr)
}

func (g *valueSpec) Goastch(ctx *Context, n ast.Node) bool {
	return g.match(ctx, n, types.ValueSpec)
}

// Stringer

func (g *arrayType) String() string {
	return fmt.Sprintf("arrayType %s", g.next)
}

func (g *assignStmt) String() string {
	return fmt.Sprintf("assignStmt %s", g.next)
}

func (g *badDecl) String() string {
	return fmt.Sprintf("badDecl %s", g.next)
}

func (g *badExpr) String() string {
	return fmt.Sprintf("badExpr %s", g.next)
}

func (g *badStmt) String() string {
	return fmt.Sprintf("badStmt %s", g.next)
}

func (g *basicLit) String() string {
	return fmt.Sprintf("basicLit %s", g.next)
}

func (g *binaryExpr) String() string {
	return fmt.Sprintf("binaryExpr %s", g.next)
}

func (g *blockStmt) String() string {
	return fmt.Sprintf("blockStmt %s", g.next)
}

func (g *branchStmt) String() string {
	return fmt.Sprintf("branchStmt %s", g.next)
}

func (g *callExpr) String() string {
	return fmt.Sprintf("callExpr %s", g.next)
}

func (g *caseClause) String() string {
	return fmt.Sprintf("caseClause %s", g.next)
}

func (g *chanType) String() string {
	return fmt.Sprintf("chanType %s", g.next)
}

func (g *commClause) String() string {
	return fmt.Sprintf("commClause %s", g.next)
}

func (g *comment) String() string {
	return fmt.Sprintf("comment %s", g.next)
}

func (g *commentGroup) String() string {
	return fmt.Sprintf("commentGroup %s", g.next)
}

func (g *compositeLit) String() string {
	return fmt.Sprintf("compositeLit %s", g.next)
}

func (g *declStmt) String() string {
	return fmt.Sprintf("declStmt %s", g.next)
}

func (g *deferStmt) String() string {
	return fmt.Sprintf("deferStmt %s", g.next)
}

func (g *ellipsis) String() string {
	return fmt.Sprintf("ellipsis %s", g.next)
}

func (g *emptyStmt) String() string {
	return fmt.Sprintf("emptyStmt %s", g.next)
}

func (g *exprStmt) String() string {
	return fmt.Sprintf("exprStmt %s", g.next)
}

func (g *field) String() string {
	return fmt.Sprintf("field %s", g.next)
}

func (g *fieldList) String() string {
	return fmt.Sprintf("fieldList %s", g.next)
}

func (g *file) String() string {
	return fmt.Sprintf("file %s", g.next)
}

func (g *forStmt) String() string {
	return fmt.Sprintf("forStmt %s", g.next)
}

func (g *funcDecl) String() string {
	return fmt.Sprintf("funcDecl %s", g.next)
}

func (g *funcLit) String() string {
	return fmt.Sprintf(" funcLit %s", g.next)
}

func (g *funcType) String() string {
	return fmt.Sprintf(" funcType %s", g.next)
}

func (g *genDecl) String() string {
	return fmt.Sprintf("genDecl %s", g.next)
}

func (g *goStmt) String() string {
	return fmt.Sprintf("goStmt %s", g.next)
}

func (g *ident) String() string {
	return fmt.Sprintf("ident %s", g.next)
}

func (g *ifStmt) String() string {
	return fmt.Sprintf("ifStmt %s", g.next)
}

func (g *importSpec) String() string {
	return fmt.Sprintf("importSpec %s", g.next)
}

func (g *incDecStmt) String() string {
	return fmt.Sprintf("incDecStmt %s", g.next)
}

func (g *indexExpr) String() string {
	return fmt.Sprintf("indexExpr %s", g.next)
}

func (g *interfaceType) String() string {
	return fmt.Sprintf("interfaceType %s", g.next)
}

func (g *keyValueExpr) String() string {
	return fmt.Sprintf("keyValueExpr %s", g.next)
}

func (g *labeledStmt) String() string {
	return fmt.Sprintf("labeledStmt %s", g.next)
}

func (g *mapType) String() string {
	return fmt.Sprintf("mapType %s", g.next)
}

func (g *pkg) String() string {
	return fmt.Sprintf("package %s", g.next)
}

func (g *parenExpr) String() string {
	return fmt.Sprintf("parenExpr %s", g.next)
}

func (g *rangeStmt) String() string {
	return fmt.Sprintf("rangeStmt %s", g.next)
}

func (g *returnStmt) String() string {
	return fmt.Sprintf("returnStmt %s", g.next)
}

func (g *selectStmt) String() string {
	return fmt.Sprintf("selectStmt %s", g.next)
}

func (g *selectorExpr) String() string {
	return fmt.Sprintf("selectorExpr %s", g.next)
}

func (g *sendStmt) String() string {
	return fmt.Sprintf("sendStmt %s", g.next)
}

func (g *sliceExpr) String() string {
	return fmt.Sprintf("sliceExpr %s", g.next)
}

func (g *starExpr) String() string {
	return fmt.Sprintf("starExpr %s", g.next)
}

func (g *structType) String() string {
	return fmt.Sprintf("structType %s", g.next)
}

func (g *switchStmt) String() string {
	return fmt.Sprintf("switchStmt %s", g.next)
}

func (g *typeAssertExpr) String() string {
	return fmt.Sprintf("typeAssertExpr %s", g.next)
}

func (g *typeSpec) String() string {
	return fmt.Sprintf("typeSpec %s", g.next)
}

func (g *typeSwitchStmt) String() string {
	return fmt.Sprintf("typeSwitchStmt %s", g.next)
}

func (g *unaryExpr) String() string {
	return fmt.Sprintf("unaryExpr %s", g.next)
}

func (g *valueSpec) String() string {
	return fmt.Sprintf("valueSpec %s", g.next)
}

// Bind

func (g *arrayType) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *assignStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *badDecl) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *badExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *badStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *basicLit) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *binaryExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *blockStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *branchStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *callExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *caseClause) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *chanType) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *commClause) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *comment) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *commentGroup) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *compositeLit) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *declStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *deferStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *ellipsis) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *emptyStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *exprStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *field) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *fieldList) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *file) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *forStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *funcDecl) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *funcLit) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *funcType) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *genDecl) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *goStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *ident) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *ifStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *importSpec) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *incDecStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *indexExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *interfaceType) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *keyValueExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *labeledStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *mapType) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *pkg) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *parenExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *rangeStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *returnStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *selectStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *selectorExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *sendStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *sliceExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *starExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *structType) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *switchStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *typeAssertExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *typeSpec) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *typeSwitchStmt) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *unaryExpr) Bind(id string) Goastcher {
	return g.bind(id, g)
}

func (g *valueSpec) Bind(id string) Goastcher {
	return g.bind(id, g)
}
