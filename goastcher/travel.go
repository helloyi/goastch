// goastcher impliments traversal goastcher
package goastcher

import (
	"fmt"
	"go/ast"
)

type (
	hasDescendant struct {
		ger Goastcher
		id  string
	}

	has struct {
		ger Goastcher
		id  string
	}

	hasName struct {
		ger Goastcher
		id  string
	}

	hasValue struct {
		ger Goastcher
		id  string
	}

	forDecls struct {
		ger Goastcher
		id  string
	}

	forSpecs struct {
		ger Goastcher
		id  string
	}

	forNames struct {
		ger Goastcher
		id  string
	}

	forFields struct {
		ger Goastcher
		id  string
	}

	hasRecvName struct {
		ger Goastcher
		id  string
	}

	hasRhs struct {
		ger Goastcher
		id  string
	}

	hasResults struct {
		ger Goastcher
		id  string
	}

	hasType struct {
		ger Goastcher
		id  string
	}

	last struct {
		ger Goastcher
		id  string
	}

	hasCond struct {
		ger Goastcher
		id  string
	}

	hasDecl struct {
		ger Goastcher
		id  string
	}

	hasLen struct {
		ger Goastcher
		id  string
	}

	hasElement struct {
		ger Goastcher
		id  string
	}

	forLHS struct {
		ger Goastcher
		id  string
	}

	forRHS struct {
		ger Goastcher
		id  string
	}

	hasLabel struct {
		ger Goastcher
		id  string
	}

	forArgs struct {
		ger Goastcher
		id  string
	}

	hasFun struct {
		ger Goastcher
		id  string
	}

	// TODO CaseClause
	// CommClause
)

// HasFun matches Fun of CallExpr
func HasFun(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasFun{ger: g}
}

func (g *hasFun) String() string {
	return fmt.Sprintf("hasFun %s", g.ger)
}

func (g *hasFun) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *hasFun) Goastch(ctx *Context, node ast.Node) bool {
	if n, ok := node.(*ast.CallExpr); ok {
		if g.ger.Goastch(ctx, n.Fun) {
			ctx.growBindings(g.id, n.Fun)
			return true
		}
	}
	return false
}

// ForArgs matches all arguments
func ForArgs(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &forArgs{ger: g}
}

func (g *forArgs) String() string {
	return fmt.Sprintf("forArgs %s", g.ger)
}

func (g *forArgs) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *forArgs) Goastch(ctx *Context, node ast.Node) bool {
	var args []ast.Expr
	switch n := node.(type) {
	case *ast.CallExpr:
		args = n.Args
	}

	matched := false
	for _, arg := range args {
		if g.ger.Goastch(ctx, arg) {
			ctx.growBindings(g.id, arg)
			matched = true
		}
	}
	return matched
}

// HasLabel matches Label of BranchStmt, LabeledStmt
func HasLabel(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasLabel{ger: g}
}

func (g *hasLabel) String() string {
	return fmt.Sprintf("hasLabel %s", g.ger)
}

func (g *hasLabel) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *hasLabel) Goastch(ctx *Context, node ast.Node) bool {
	var label *ast.Ident
	switch n := node.(type) {
	case *ast.BranchStmt:
		label = n.Label
	case *ast.LabeledStmt:
		label = n.Label
	}
	if !g.ger.Goastch(ctx, label) {
		return false
	}
	ctx.growBindings(g.id, label)
	return true
}

// ForRHS matches all RHS expr of AssignStmt
func ForRHS(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &forRHS{ger: g}
}

func (g *forRHS) String() string {
	return fmt.Sprintf("forRHS %s", g.ger)
}

func (g *forRHS) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *forRHS) Goastch(ctx *Context, node ast.Node) bool {
	matched := false
	switch n := node.(type) {
	case *ast.AssignStmt:
		for _, expr := range n.Rhs {
			if g.ger.Goastch(ctx, expr) {
				matched = true
				ctx.growBindings(g.id, expr)
			}
		}
	}
	return matched
}

// ForLHS matches all LHS expr of AssignStmt
func ForLHS(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &forLHS{ger: g}
}

// String ...
func (g *forLHS) String() string {
	return fmt.Sprintf("forLHS %s", g.ger)
}

func (g *forLHS) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *forLHS) Goastch(ctx *Context, node ast.Node) bool {
	matched := false
	switch n := node.(type) {
	case *ast.AssignStmt:
		for _, expr := range n.Lhs {
			if g.ger.Goastch(ctx, expr) {
				matched = true
				ctx.growBindings(g.id, expr)
			}
		}
	}
	return matched
}

// HasElement matches Elt node of ArrayType, Ellipsis
func HasElement(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasElement{ger: g}
}

func (g *hasElement) String() string {
	return fmt.Sprintf("hasElement %s", g.ger)
}

func (g *hasElement) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *hasElement) Goastch(ctx *Context, node ast.Node) bool {
	var elt ast.Node
	switch n := node.(type) {
	case *ast.ArrayType:
		elt = n.Elt
	case *ast.Ellipsis:
		elt = n.Elt
	}
	if !g.ger.Goastch(ctx, elt) {
		return false
	}
	ctx.growBindings(g.id, elt)
	return true
}

// HasLen matches Len node of ArrayType
func HasLen(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasLen{ger: g}
}

func (g *hasLen) String() string {
	return fmt.Sprintf("hasLen %s", g.ger)
}

func (g *hasLen) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *hasLen) Goastch(ctx *Context, node ast.Node) bool {
	var len ast.Node
	switch n := node.(type) {
	case *ast.ArrayType:
		len = n.Len
	}
	if !g.ger.Goastch(ctx, len) {
		return false
	}
	ctx.growBindings(g.id, len)
	return true
}

// HasDecl matches Decl node of DeclStmt
func HasDecl(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasDecl{ger: g}
}

func (g *hasDecl) String() string {
	return fmt.Sprintf("hasDecl %s", g.ger)
}

func (g *hasDecl) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *hasDecl) Goastch(ctx *Context, node ast.Node) bool {
	var decl ast.Node
	switch n := node.(type) {
	case *ast.DeclStmt:
		decl = n.Decl
	}
	if !g.ger.Goastch(ctx, decl) {
		return false
	}
	ctx.growBindings(g.id, decl)
	return true
}

// HasCond matches Cond node of ForStmt, IfStmt
func HasCond(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasCond{ger: g}
}

func (g *hasCond) String() string {
	return fmt.Sprintf("hasCond %s", g.ger)
}

func (g *hasCond) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *hasCond) Goastch(ctx *Context, node ast.Node) bool {
	var cond ast.Node
	switch n := node.(type) {
	case *ast.ForStmt:
		cond = n.Cond
	case *ast.IfStmt:
		cond = n.Cond
	}
	if !g.ger.Goastch(ctx, cond) {
		return false
	}
	ctx.growBindings(g.id, cond)
	return true
}

// Last matches last node of some List-Type nodes, eg. FieldList, BlockStmt,
// CommentGroup
func Last(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &last{ger: g}
}

func (g *last) String() string {
	return fmt.Sprintf("last %s", g.ger)
}

func (g *last) Goastch(ctx *Context, node ast.Node) bool {
	var lastNode ast.Node
	switch n := node.(type) {
	case *ast.FieldList:
		if n == nil {
			return false
		}
		if len(n.List) == 0 {
			return false
		}
		lastNode = n.List[len(n.List)-1]
	case *ast.BlockStmt:
		if len(n.List) == 0 {
			return false
		}
		lastNode = n.List[len(n.List)-1]
	case *ast.CommentGroup:
		if len(n.List) == 0 {
			return false
		}
		lastNode = n.List[len(n.List)-1]
	}
	if !g.ger.Goastch(ctx, lastNode) {
		return false
	}
	ctx.growBindings(g.id, lastNode)
	return true
}

func (g *last) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// ForFields matches all Fields of FieldList, StructType
func ForFields(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &forFields{ger: g}
}

func (g *forFields) String() string {
	return fmt.Sprintf("forFields %s", g.ger)
}

func (g *forFields) Goastch(ctx *Context, node ast.Node) bool {
	var fields []*ast.Field
	switch n := node.(type) {
	case *ast.FieldList:
		fields = n.List
	case *ast.StructType:
		fields = n.Fields.List
	}

	matched := false
	for _, field := range fields {
		if g.ger.Goastch(ctx, field) {
			ctx.growBindings(g.id, field)
			matched = true
		}
	}

	return matched
}

func (g *forFields) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// HasType matches Type node
func HasType(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasType{ger: g}
}

func (g *hasType) String() string {
	return fmt.Sprintf("hasType %s", g.ger)
}

func (g *hasType) Goastch(ctx *Context, node ast.Node) bool {
	var typeNode ast.Node
	switch n := node.(type) {
	case *ast.CompositeLit:
		typeNode = n.Type
	case *ast.Field:
		typeNode = n.Type
	case *ast.FuncDecl:
		typeNode = n.Type
	case *ast.FuncLit:
		typeNode = n.Type
	case *ast.TypeAssertExpr:
		typeNode = n.Type
	case *ast.TypeSpec:
		typeNode = n.Type
	case *ast.ValueSpec:
		typeNode = n.Type
	}

	if !g.ger.Goastch(ctx, typeNode) {
		return false
	}
	ctx.growBindings(g.id, typeNode)
	return true
}

func (g *hasType) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// HasResults matches Results node of FuncType
func HasResults(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasResults{ger: g}
}

func (g *hasResults) String() string {
	return fmt.Sprintf("hasResults %s", g.ger)
}

func (g *hasResults) Goastch(ctx *Context, node ast.Node) bool {
	switch n := node.(type) {
	case *ast.FuncType:
		if !g.ger.Goastch(ctx, n.Results) {
			return false
		}
		ctx.growBindings(g.id, n.Results)
	}
	return true
}

func (g *hasResults) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// HasRhs ...
func HasRhs(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasRhs{ger: g}
}

func (g *hasRhs) String() string {
	return fmt.Sprintf("hasRhs %s", g.ger)
}

func (g *hasRhs) Goastch(ctx *Context, node ast.Node) bool {
	switch as := node.(type) {
	case *ast.AssignStmt:
		if len(as.Rhs) != 1 {
			return false
		}
		if !g.ger.Goastch(ctx, as.Rhs[0]) {
			return false
		}
		ctx.growBindings(g.id, as.Rhs[0])
	}

	return true
}

func (g *hasRhs) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// HasRecvName matches RecvName node of FuncDecl
func HasRecvName(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasRecvName{ger: g}
}

func (g *hasRecvName) String() string {
	return fmt.Sprintf("hasRecvName %s", g.ger)
}

func (g *hasRecvName) Goastch(ctx *Context, node ast.Node) bool {
	switch fn := node.(type) {
	case *ast.FuncDecl:
		if fn.Recv == nil || len(fn.Recv.List) == 0 {
			return false
		}
		names := fn.Recv.List[0].Names
		if len(names) < 1 {
			return false
		}
		name := names[0]
		if !g.ger.Goastch(ctx, name) {
			return false
		}
		ctx.growBindings(g.id, name)
	}

	return true
}

func (g *hasRecvName) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// ForDecls matches all Decls of File
func ForDecls(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &forDecls{ger: g}
}

func (g *forDecls) String() string {
	return fmt.Sprintf("forDecls %s", g.ger)
}

func (g *forDecls) Goastch(ctx *Context, node ast.Node) bool {
	matched := false
	switch n := node.(type) {
	case *ast.File:
		for _, decl := range n.Decls {
			if g.ger.Goastch(ctx, decl) {
				ctx.growBindings(g.id, decl)
				matched = true
			}
		}
	}

	return matched
}

func (g *forDecls) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// ForSpecs matches all Specs of GenDecl
func ForSpecs(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &forSpecs{ger: g}
}

func (g *forSpecs) String() string {
	return fmt.Sprintf("forSpecs %s", g.ger)
}

func (g *forSpecs) Goastch(ctx *Context, node ast.Node) bool {
	matched := false
	switch n := node.(type) {
	case *ast.GenDecl:
		for _, spec := range n.Specs {
			if g.ger.Goastch(ctx, spec) {
				ctx.growBindings(g.id, spec)
				matched = true
			}
		}
	}

	return matched
}
func (g *forSpecs) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// ForNames matches all Names of ValueSpec
func ForNames(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &forNames{ger: g}
}

func (g *forNames) String() string {
	return fmt.Sprintf("forNames %s", g.ger)
}

func (g *forNames) Goastch(ctx *Context, node ast.Node) bool {
	matched := false
	switch n := node.(type) {
	case *ast.ValueSpec:
		for _, name := range n.Names {
			if g.ger.Goastch(ctx, name) {
				ctx.growBindings(g.id, name)
				matched = true
			}
		}
	}

	return matched
}

func (g *forNames) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// HasValue matches Value node of RangStmt
func HasValue(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasValue{ger: g}
}

func (g *hasValue) String() string {
	return fmt.Sprintf("hasValue %s", g.ger)
}

func (g *hasValue) Goastch(ctx *Context, node ast.Node) bool {
	var valueNode ast.Node
	switch n := node.(type) {
	case *ast.RangeStmt:
		valueNode = n.Value
	}
	if !g.ger.Goastch(ctx, valueNode) {
		return false
	}
	ctx.growBindings(g.id, valueNode)
	return true
}

func (g *hasValue) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// HasName matches Name node of ImportSpec, File, FuncDecl
func HasName(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasName{ger: g}
}

func (g *hasName) String() string {
	return fmt.Sprintf("hasName %s", g.ger)
}

func (g *hasName) Goastch(ctx *Context, node ast.Node) bool {
	var nameNode ast.Node
	switch n := node.(type) {
	case *ast.ImportSpec:
		nameNode = n.Name
	case *ast.File:
		nameNode = n.Name
	case *ast.FuncDecl:
		nameNode = n.Name
	}

	if !g.ger.Goastch(ctx, nameNode) {
		return false
	}
	ctx.growBindings(g.id, nameNode)
	return true
}

func (g *hasName) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

type inspector func(ast.Node)

func (f inspector) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	f(node)
	return f
}

// HasDescendant matches all descendant AST nodes
func HasDescendant(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &hasDescendant{ger: g}
}

func (g *hasDescendant) Goastch(ctx *Context, node ast.Node) bool {
	matched := false
	ast.Walk(inspector(func(n ast.Node) {
		if g.ger.Goastch(ctx, n) {
			ctx.growBindings(g.id, n)
			matched = true
		}
	}), node)
	return matched
}

func (g *hasDescendant) String() string {
	return fmt.Sprintf("hasDescendant %s", g.ger)
}

func (g *hasDescendant) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

// Has matches all child AST nodes
func Has(g Goastcher) Goastcher {
	if isErrorGer(g) {
		return g
	}
	return &has{ger: g}
}

func (g *has) String() string {
	return fmt.Sprintf("has %s", g.ger)
}

func (g *has) Bind(id string) Goastcher {
	if id == "" {
		g.id = fmt.Sprintf("%p", g)
	} else {
		g.id = id
	}
	return g
}

func (g *has) Goastch(ctx *Context, node ast.Node) bool {
	switch n := node.(type) {
	case *ast.ArrayType:
		return g.arrayType(ctx, n)
	case *ast.AssignStmt:
		return g.assignStmt(ctx, n)
	case *ast.BadDecl:
		return g.badDecl(ctx, n)
	case *ast.BadExpr:
		return g.badExpr(ctx, n)
	case *ast.BadStmt:
		return g.badStmt(ctx, n)
	case *ast.BasicLit:
		return g.basicLit(ctx, n)
	case *ast.BinaryExpr:
		return g.binaryExpr(ctx, n)
	case *ast.BlockStmt:
		return g.blockStmt(ctx, n)
	case *ast.BranchStmt:
		return g.branchStmt(ctx, n)
	case *ast.CallExpr:
		return g.callExpr(ctx, n)
	case *ast.CaseClause:
		return g.caseClause(ctx, n)
	case *ast.ChanType:
		return g.chanType(ctx, n)
	case *ast.CommClause:
		return g.commClause(ctx, n)
	case *ast.Comment:
		return g.comment(ctx, n)
	case *ast.CommentGroup:
		return g.commentGroup(ctx, n)
	case *ast.CompositeLit:
		return g.compositeLit(ctx, n)
	case *ast.DeclStmt:
		return g.declStmt(ctx, n)
	case *ast.DeferStmt:
		return g.deferStmt(ctx, n)
	case *ast.Ellipsis:
		return g.ellipsis(ctx, n)
	case *ast.EmptyStmt:
		return g.emptyStmt(ctx, n)
	case *ast.ExprStmt:
		return g.exprStmt(ctx, n)
	case *ast.Field:
		return g.field(ctx, n)
	case *ast.FieldList:
		return g.fieldList(ctx, n)
	case *ast.File:
		return g.file(ctx, n)
	case *ast.ForStmt:
		return g.forStmt(ctx, n)
	case *ast.FuncDecl:
		return g.funcDecl(ctx, n)
	case *ast.FuncLit:
		return g.funcLit(ctx, n)
	case *ast.FuncType:
		return g.funcType(ctx, n)
	case *ast.GenDecl:
		return g.genDecl(ctx, n)
	case *ast.GoStmt:
		return g.goStmt(ctx, n)
	case *ast.Ident:
		return g.ident(ctx, n)
	case *ast.IfStmt:
		return g.ifStmt(ctx, n)
	case *ast.ImportSpec:
		return g.importSpec(ctx, n)
	case *ast.IncDecStmt:
		return g.incDecStmt(ctx, n)
	case *ast.IndexExpr:
		return g.indexExpr(ctx, n)
	case *ast.InterfaceType:
		return g.interfaceType(ctx, n)
	case *ast.KeyValueExpr:
		return g.keyValueExpr(ctx, n)
	case *ast.LabeledStmt:
		return g.labeledStmt(ctx, n)
	case *ast.MapType:
		return g.mapType(ctx, n)
	case *ast.Package:
		return g.pkg(ctx, n)
	case *ast.ParenExpr:
		return g.parenExpr(ctx, n)
	case *ast.RangeStmt:
		return g.rangeStmt(ctx, n)
	case *ast.ReturnStmt:
		return g.returnStmt(ctx, n)
	case *ast.SelectStmt:
		return g.selectStmt(ctx, n)
	case *ast.SelectorExpr:
		return g.selectorExpr(ctx, n)
	case *ast.SendStmt:
		return g.sendStmt(ctx, n)
	case *ast.SliceExpr:
		return g.sliceExpr(ctx, n)
	case *ast.StarExpr:
		return g.starExpr(ctx, n)
	case *ast.StructType:
		return g.structType(ctx, n)
	case *ast.SwitchStmt:
		return g.switchStmt(ctx, n)
	case *ast.TypeAssertExpr:
		return g.typeAssertExpr(ctx, n)
	case *ast.TypeSpec:
		return g.typeSpec(ctx, n)
	case *ast.TypeSwitchStmt:
		return g.typeSwitchStmt(ctx, n)
	case *ast.UnaryExpr:
		return g.unaryExpr(ctx, n)
	case *ast.ValueSpec:
		return g.valueSpec(ctx, n)
	default:
		return false
	}
}

func (g *has) arrayType(ctx *Context, node *ast.ArrayType) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Len) {
		ctx.growBindings(g.id, node.Len)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Elt) {
		ctx.growBindings(g.id, node.Elt)
		matched = true
	}

	return matched
}

func (g *has) assignStmt(ctx *Context, node *ast.AssignStmt) bool {
	matched := false
	for _, expr := range node.Lhs {
		if g.ger.Goastch(ctx, expr) {
			ctx.growBindings(g.id, expr)
			matched = true
		}
	}
	for _, expr := range node.Rhs {
		if g.ger.Goastch(ctx, expr) {
			ctx.growBindings(g.id, expr)
			matched = true
		}
	}

	return matched
}

func (g *has) badDecl(ctx *Context, node *ast.BadDecl) bool {
	return false
}
func (g *has) badExpr(ctx *Context, node *ast.BadExpr) bool {
	return false
}
func (g *has) badStmt(ctx *Context, node *ast.BadStmt) bool {
	return false
}

func (g *has) basicLit(ctx *Context, node *ast.BasicLit) bool {
	return false
}
func (g *has) binaryExpr(ctx *Context, node *ast.BinaryExpr) bool {
	matched := false
	if g.ger.Goastch(ctx, node.X) {
		ctx.growBindings(g.id, node.X)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Y) {
		ctx.growBindings(g.id, node.Y)
		matched = true
	}

	return matched
}

func (g *has) blockStmt(ctx *Context, node *ast.BlockStmt) bool {
	matched := false
	for _, stmt := range node.List {
		if g.ger.Goastch(ctx, stmt) {
			ctx.growBindings(g.id, stmt)
			matched = true
		}
	}
	return matched
}

func (g *has) branchStmt(ctx *Context, node *ast.BranchStmt) bool {
	if !g.ger.Goastch(ctx, node.Label) {
		return false
	}
	ctx.growBindings(g.id, node.Label)
	return true
}

func (g *has) callExpr(ctx *Context, node *ast.CallExpr) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Fun) {
		ctx.growBindings(g.id, node.Fun)
		matched = true
	}
	for _, expr := range node.Args {
		if g.ger.Goastch(ctx, expr) {
			ctx.growBindings(g.id, expr)
			matched = true
		}
	}
	return matched
}

func (g *has) caseClause(ctx *Context, node *ast.CaseClause) bool {
	matched := false
	for _, expr := range node.List {
		if g.ger.Goastch(ctx, expr) {
			ctx.growBindings(g.id, expr)
			matched = true
		}
	}
	for _, stmt := range node.Body {
		if g.ger.Goastch(ctx, stmt) {
			ctx.growBindings(g.id, stmt)
			matched = true
		}
	}
	return matched
}

func (g *has) chanType(ctx *Context, node *ast.ChanType) bool {
	if !g.ger.Goastch(ctx, node.Value) {
		return false
	}
	ctx.growBindings(g.id, node.Value)
	return true
}

func (g *has) commClause(ctx *Context, node *ast.CommClause) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Comm) {
		ctx.growBindings(g.id, node.Comm)
		matched = true
	}
	for _, stmt := range node.Body {
		if g.ger.Goastch(ctx, stmt) {
			ctx.growBindings(g.id, stmt)
			matched = true
		}
	}
	return matched
}

func (g *has) comment(ctx *Context, node *ast.Comment) bool {
	return false
}

func (g *has) commentGroup(ctx *Context, node *ast.CommentGroup) bool {
	matched := false
	for _, comment := range node.List {
		if g.ger.Goastch(ctx, comment) {
			ctx.growBindings(g.id, comment)
			matched = true
		}
	}
	return matched
}

func (g *has) compositeLit(ctx *Context, node *ast.CompositeLit) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Type) {
		ctx.growBindings(g.id, node.Type)
		matched = true
	}
	for _, expr := range node.Elts {
		if g.ger.Goastch(ctx, expr) {
			ctx.growBindings(g.id, expr)
			matched = true
		}
	}
	return matched
}

func (g *has) declStmt(ctx *Context, node *ast.DeclStmt) bool {
	if !g.ger.Goastch(ctx, node.Decl) {
		return false
	}
	ctx.growBindings(g.id, node.Decl)
	return true
}

func (g *has) deferStmt(ctx *Context, node *ast.DeferStmt) bool {
	if !g.ger.Goastch(ctx, node.Call) {
		return false
	}
	ctx.growBindings(g.id, node.Call)
	return true
}

func (g *has) ellipsis(ctx *Context, node *ast.Ellipsis) bool {
	if !g.ger.Goastch(ctx, node.Elt) {
		return false
	}
	ctx.growBindings(g.id, node.Elt)
	return true
}

func (g *has) emptyStmt(ctx *Context, node *ast.EmptyStmt) bool {
	return false
}

func (g *has) exprStmt(ctx *Context, node *ast.ExprStmt) bool {
	if !g.ger.Goastch(ctx, node.X) {
		return false
	}
	ctx.growBindings(g.id, node.X)
	return true
}

func (g *has) field(ctx *Context, node *ast.Field) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Doc) {
		ctx.growBindings(g.id, node.Doc)
		matched = true
	}
	for _, name := range node.Names {
		if g.ger.Goastch(ctx, name) {
			ctx.growBindings(g.id, name)
			matched = true
		}
	}
	if g.ger.Goastch(ctx, node.Type) {
		ctx.growBindings(g.id, node.Type)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Tag) {
		ctx.growBindings(g.id, node.Tag)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Comment) {
		ctx.growBindings(g.id, node.Comment)
		matched = true
	}
	return matched
}

func (g *has) fieldList(ctx *Context, node *ast.FieldList) bool {
	matched := false
	for _, field := range node.List {
		if g.ger.Goastch(ctx, field) {
			ctx.growBindings(g.id, field)
			matched = true
		}
	}
	return matched
}

func (g *has) file(ctx *Context, node *ast.File) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Doc) {
		ctx.growBindings(g.id, node.Doc)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Name) {
		ctx.growBindings(g.id, node.Name)
		matched = true
	}
	for _, decl := range node.Decls {
		if g.ger.Goastch(ctx, decl) {
			ctx.growBindings(g.id, decl)
			matched = true
		}
	}
	for _, imp := range node.Imports {
		if g.ger.Goastch(ctx, imp) {
			ctx.growBindings(g.id, imp)
			matched = true
		}
	}
	for _, ident := range node.Unresolved {
		if g.ger.Goastch(ctx, ident) {
			ctx.growBindings(g.id, ident)
			matched = true
		}
	}
	for _, comment := range node.Comments {
		if g.ger.Goastch(ctx, comment) {
			ctx.growBindings(g.id, comment)
			matched = true
		}
	}
	return matched
}

func (g *has) forStmt(ctx *Context, node *ast.ForStmt) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Init) {
		ctx.growBindings(g.id, node.Init)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Cond) {
		ctx.growBindings(g.id, node.Cond)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Post) {
		ctx.growBindings(g.id, node.Post)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Body) {
		ctx.growBindings(g.id, node.Body)
		matched = true
	}
	return matched
}

func (g *has) funcDecl(ctx *Context, node *ast.FuncDecl) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Doc) {
		ctx.growBindings(g.id, node.Doc)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Recv) {
		ctx.growBindings(g.id, node.Recv)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Name) {
		ctx.growBindings(g.id, node.Name)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Type) {
		ctx.growBindings(g.id, node.Type)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Body) {
		ctx.growBindings(g.id, node.Body)
		matched = true
	}
	return matched
}

func (g *has) funcLit(ctx *Context, node *ast.FuncLit) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Type) {
		ctx.growBindings(g.id, node.Type)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Body) {
		ctx.growBindings(g.id, node.Body)
		matched = true
	}
	return matched
}

func (g *has) funcType(ctx *Context, node *ast.FuncType) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Params) {
		ctx.growBindings(g.id, node.Params)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Results) {
		ctx.growBindings(g.id, node.Results)
		matched = true
	}
	return matched
}

func (g *has) genDecl(ctx *Context, node *ast.GenDecl) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Doc) {
		ctx.growBindings(g.id, node.Doc)
		matched = true
	}
	for _, spec := range node.Specs {
		if g.ger.Goastch(ctx, spec) {
			ctx.growBindings(g.id, spec)
			matched = true
		}
	}
	return matched
}

func (g *has) goStmt(ctx *Context, node *ast.GoStmt) bool {
	if !g.ger.Goastch(ctx, node.Call) {
		return false
	}
	ctx.growBindings(g.id, node.Call)
	return true
}

func (g *has) ident(ctx *Context, node *ast.Ident) bool {
	return false
}

func (g *has) ifStmt(ctx *Context, node *ast.IfStmt) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Init) {
		ctx.growBindings(g.id, node.Init)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Cond) {
		ctx.growBindings(g.id, node.Cond)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Body) {
		ctx.growBindings(g.id, node.Body)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Else) {
		ctx.growBindings(g.id, node.Else)
		matched = true
	}

	return matched
}

func (g *has) importSpec(ctx *Context, node *ast.ImportSpec) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Doc) {
		ctx.growBindings(g.id, node.Doc)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Name) {
		ctx.growBindings(g.id, node.Name)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Path) {
		ctx.growBindings(g.id, node.Path)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Comment) {
		ctx.growBindings(g.id, node.Comment)
		matched = true
	}
	return matched
}

func (g *has) incDecStmt(ctx *Context, node *ast.IncDecStmt) bool {
	if !g.ger.Goastch(ctx, node.X) {
		return false
	}
	ctx.growBindings(g.id, node.X)
	return true

}

func (g *has) indexExpr(ctx *Context, node *ast.IndexExpr) bool {
	matched := false
	if g.ger.Goastch(ctx, node.X) {
		ctx.growBindings(g.id, node.X)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Index) {
		ctx.growBindings(g.id, node.Index)
		matched = true
	}
	return matched
}

func (g *has) interfaceType(ctx *Context, node *ast.InterfaceType) bool {
	if !g.ger.Goastch(ctx, node.Methods) {
		return false
	}
	ctx.growBindings(g.id, node.Methods)
	return true
}

func (g *has) keyValueExpr(ctx *Context, node *ast.KeyValueExpr) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Key) {
		ctx.growBindings(g.id, node.Key)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Value) {
		ctx.growBindings(g.id, node.Value)
		matched = true
	}
	return matched
}

func (g *has) labeledStmt(ctx *Context, node *ast.LabeledStmt) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Label) {
		ctx.growBindings(g.id, node.Label)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Stmt) {
		ctx.growBindings(g.id, node.Stmt)
		matched = true
	}
	return matched
}

func (g *has) mapType(ctx *Context, node *ast.MapType) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Key) {
		ctx.growBindings(g.id, node.Key)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Value) {
		ctx.growBindings(g.id, node.Value)
		matched = true
	}
	return matched
}

func (g *has) pkg(ctx *Context, node *ast.Package) bool {
	matched := false
	for _, f := range node.Files {
		if g.ger.Goastch(ctx, f) {
			ctx.growBindings(g.id, f)
			matched = true
		}
	}
	return matched
}

func (g *has) parenExpr(ctx *Context, node *ast.ParenExpr) bool {
	if !g.ger.Goastch(ctx, node.X) {
		return false
	}
	ctx.growBindings(g.id, node.X)
	return true
}

func (g *has) rangeStmt(ctx *Context, node *ast.RangeStmt) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Key) {
		ctx.growBindings(g.id, node.Key)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Value) {
		ctx.growBindings(g.id, node.Value)
		matched = true
	}
	if g.ger.Goastch(ctx, node.X) {
		ctx.growBindings(g.id, node.X)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Body) {
		ctx.growBindings(g.id, node.Body)
		matched = true
	}
	return matched
}

func (g *has) returnStmt(ctx *Context, node *ast.ReturnStmt) bool {
	matched := false
	for _, res := range node.Results {
		if g.ger.Goastch(ctx, res) {
			ctx.growBindings(g.id, res)
			matched = true
		}
	}
	return matched
}

func (g *has) selectStmt(ctx *Context, node *ast.SelectStmt) bool {
	if !g.ger.Goastch(ctx, node.Body) {
		return false
	}
	ctx.growBindings(g.id, node.Body)
	return true
}

func (g *has) selectorExpr(ctx *Context, node *ast.SelectorExpr) bool {
	matched := false
	if g.ger.Goastch(ctx, node.X) {
		ctx.growBindings(g.id, node.X)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Sel) {
		ctx.growBindings(g.id, node.Sel)
		matched = true
	}
	return matched
}

func (g *has) sendStmt(ctx *Context, node *ast.SendStmt) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Chan) {
		ctx.growBindings(g.id, node.Chan)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Value) {
		ctx.growBindings(g.id, node.Value)
		matched = true
	}
	return matched
}

func (g *has) sliceExpr(ctx *Context, node *ast.SliceExpr) bool {
	matched := false
	if g.ger.Goastch(ctx, node.X) {
		ctx.growBindings(g.id, node.X)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Low) {
		ctx.growBindings(g.id, node.Low)
		matched = true
	}
	if g.ger.Goastch(ctx, node.High) {
		ctx.growBindings(g.id, node.High)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Max) {
		ctx.growBindings(g.id, node.Max)
		matched = true
	}
	return matched
}

func (g *has) starExpr(ctx *Context, node *ast.StarExpr) bool {
	if !g.ger.Goastch(ctx, node.X) {
		return false
	}
	ctx.growBindings(g.id, node.X)
	return true
}

func (g *has) structType(ctx *Context, node *ast.StructType) bool {
	if !g.ger.Goastch(ctx, node.Fields) {
		return false
	}
	ctx.growBindings(g.id, node.Fields)
	return true
}

func (g *has) switchStmt(ctx *Context, node *ast.SwitchStmt) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Init) {
		ctx.growBindings(g.id, node.Init)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Tag) {
		ctx.growBindings(g.id, node.Tag)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Body) {
		ctx.growBindings(g.id, node.Body)
		matched = true
	}
	return matched
}

func (g *has) typeAssertExpr(ctx *Context, node *ast.TypeAssertExpr) bool {
	matched := false
	if g.ger.Goastch(ctx, node.X) {
		ctx.growBindings(g.id, node.X)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Type) {
		ctx.growBindings(g.id, node.Type)
		matched = true
	}
	return matched
}

func (g *has) typeSpec(ctx *Context, node *ast.TypeSpec) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Doc) {
		ctx.growBindings(g.id, node.Doc)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Name) {
		ctx.growBindings(g.id, node.Name)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Type) {
		ctx.growBindings(g.id, node.Type)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Comment) {
		ctx.growBindings(g.id, node.Comment)
		matched = true
	}
	return matched
}

func (g *has) typeSwitchStmt(ctx *Context, node *ast.TypeSwitchStmt) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Init) {
		ctx.growBindings(g.id, node.Init)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Assign) {
		ctx.growBindings(g.id, node.Assign)
		matched = true
	}
	if g.ger.Goastch(ctx, node.Body) {
		ctx.growBindings(g.id, node.Body)
		matched = true
	}
	return matched
}

func (g *has) unaryExpr(ctx *Context, node *ast.UnaryExpr) bool {
	if !g.ger.Goastch(ctx, node.X) {
		return false
	}
	ctx.growBindings(g.id, node.X)
	return true
}

func (g *has) valueSpec(ctx *Context, node *ast.ValueSpec) bool {
	matched := false
	if g.ger.Goastch(ctx, node.Doc) {
		ctx.growBindings(g.id, node.Doc)
		matched = true
	}
	for _, name := range node.Names {
		if g.ger.Goastch(ctx, name) {
			ctx.growBindings(g.id, name)
			matched = true
		}
	}
	if g.ger.Goastch(ctx, node.Type) {
		ctx.growBindings(g.id, node.Type)
		matched = true
	}
	for _, val := range node.Values {
		if g.ger.Goastch(ctx, val) {
			ctx.growBindings(g.id, val)
			matched = true
		}
	}
	if g.ger.Goastch(ctx, node.Comment) {
		ctx.growBindings(g.id, node.Comment)
		matched = true
	}
	return matched
}
