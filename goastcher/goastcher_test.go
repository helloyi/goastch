package goastcher_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"

	t "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/helloyi/goastch"
	. "github.com/helloyi/goastch/goastcher"
)

var _ = t.Describe("Node Goastcher", func() {
	cases := map[ast.Node]Goastcher{
		&ast.ArrayType{
			Len: &ast.BadExpr{},
		}: ArrayType(Anything()),

		new(ast.AssignStmt):     AssignStmt(Anything()),
		new(ast.BadDecl):        BadDecl(Anything()),
		new(ast.BadExpr):        BadExpr(Anything()),
		new(ast.BadStmt):        BadStmt(Anything()),
		new(ast.BasicLit):       BasicLit(Anything()),
		new(ast.BinaryExpr):     BinaryExpr(Anything()),
		new(ast.BlockStmt):      BlockStmt(Anything()),
		new(ast.BranchStmt):     BranchStmt(Anything()),
		new(ast.CallExpr):       CallExpr(Anything()),
		new(ast.CaseClause):     CaseClause(Anything()),
		new(ast.ChanType):       ChanType(Anything()),
		new(ast.CommClause):     CommClause(Anything()),
		new(ast.Comment):        Comment(Anything()),
		new(ast.CommentGroup):   CommentGroup(Anything()),
		new(ast.CompositeLit):   CompositeLit(Anything()),
		new(ast.DeclStmt):       DeclStmt(Anything()),
		new(ast.DeferStmt):      DeferStmt(Anything()),
		new(ast.Ellipsis):       Ellipsis(Anything()),
		new(ast.EmptyStmt):      EmptyStmt(Anything()),
		new(ast.ExprStmt):       ExprStmt(Anything()),
		new(ast.Field):          Field(Anything()),
		new(ast.FieldList):      FieldList(Anything()),
		new(ast.File):           File(Anything()),
		new(ast.ForStmt):        ForStmt(Anything()),
		new(ast.FuncDecl):       FuncDecl(Anything()),
		new(ast.FuncLit):        FuncLit(Anything()),
		new(ast.FuncType):       FuncType(Anything()),
		new(ast.GenDecl):        GenDecl(Anything()),
		new(ast.GoStmt):         GoStmt(Anything()),
		new(ast.Ident):          Ident(Anything()),
		new(ast.IfStmt):         IfStmt(Anything()),
		new(ast.ImportSpec):     ImportSpec(Anything()),
		new(ast.IncDecStmt):     IncDecStmt(Anything()),
		new(ast.IndexExpr):      IndexExpr(Anything()),
		new(ast.InterfaceType):  InterfaceType(Anything()),
		new(ast.KeyValueExpr):   KeyValueExpr(Anything()),
		new(ast.LabeledStmt):    LabeledStmt(Anything()),
		new(ast.MapType):        MapType(Anything()),
		new(ast.Package):        Pkg(Anything()),
		new(ast.ParenExpr):      ParenExpr(Anything()),
		new(ast.RangeStmt):      RangeStmt(Anything()),
		new(ast.ReturnStmt):     ReturnStmt(Anything()),
		new(ast.SelectStmt):     SelectStmt(Anything()),
		new(ast.SelectorExpr):   SelectorExpr(Anything()),
		new(ast.SendStmt):       SendStmt(Anything()),
		new(ast.SliceExpr):      SliceExpr(Anything()),
		new(ast.StarExpr):       StarExpr(Anything()),
		new(ast.StructType):     StructType(Anything()),
		new(ast.SwitchStmt):     SwitchStmt(Anything()),
		new(ast.TypeAssertExpr): TypeAssertExpr(Anything()),
		new(ast.TypeSpec):       TypeSpec(Anything()),
		new(ast.TypeSwitchStmt): TypeSwitchStmt(Anything()),
		new(ast.UnaryExpr):      UnaryExpr(Anything()),
		new(ast.ValueSpec):      ValueSpec(Anything()),

		&ast.AssignStmt{Tok: token.DEFINE}: ShortVarDecl(Anything()),
		&ast.ArrayType{Len: nil}:           SliceType(Anything()),
		&ast.BasicLit{Kind: token.INT}:     IntBasicLit(Anything()),
		&ast.BasicLit{Kind: token.FLOAT}:   FloatBasicLit(Anything()),
		&ast.BasicLit{Kind: token.IMAG}:    ImagBasicLit(Anything()),
		&ast.BasicLit{Kind: token.CHAR}:    CharBasicLit(Anything()),
		&ast.BasicLit{Kind: token.STRING}:  StringBasicLit(Anything()),
	}
	t.It("invalid", func() {
		for node, ger := range cases {
			bindings, err := goastch.Find(node, nil, ger.Bind("x"))
			Expect(err).To(BeNil())
			Expect(bindings).Should(Not(BeNil()))
			Expect(len(bindings["x"])).Should(Equal(1))
			Expect(bindings["x"][0]).To(Equal(node))
		}
	})
})

var _ = t.Describe("Travel Goastcher", func() {
	t.XIt("hasDescendant", func() {
		file := parse(`package test; var a []string`)
		bindings, err := goastch.Find(file, nil, HasDescendant(Anything()).Bind("x"))
		Expect(err).To(BeNil())
		Expect(len(bindings["x"])).To(Equal(counts(file)))
	})

	t.Describe("hasName", func() {
		test := func(node ast.Node, name *ast.Ident) {
			b, err := goastch.Find(node, nil, HasName(Anything()).Bind("x"))
			Expect(err).To(BeNil())
			Expect(b).Should(Not(BeNil()))
			Expect(len(b["x"])).Should(Equal(1))
			Expect(b["x"][0]).To(Equal(name))
		}
		t.It("is ast.ImportSpec", func() {
			node := &ast.ImportSpec{Name: new(ast.Ident)}
			test(node, node.Name)
		})
		t.It("is ast.File", func() {
			node := &ast.File{Name: new(ast.Ident)}
			test(node, node.Name)
		})
	})

	hasTest := func(node ast.Node, it ast.Node, g Goastcher) {
		b, err := goastch.Find(node, nil, g.Bind("x"))
		Expect(err).To(BeNil())
		Expect(b).Should(Not(BeNil()))
		Expect(len(b["x"])).Should(Equal(1))
		Expect(b["x"][0]).To(Equal(it))
	}
	t.Describe("hasName", func() {
		t.It("is ImportSpec", func() {
			node := &ast.ImportSpec{Name: new(ast.Ident)}
			hasTest(node, node.Name, HasName(Anything()))
		})
		t.It("is File", func() {
			node := &ast.File{Name: new(ast.Ident)}
			hasTest(node, node.Name, HasName(Anything()))
		})
	})

	t.Describe("hasValue", func() {
		t.It("is RangeStmt", func() {
			node := &ast.RangeStmt{Value: new(ast.BadExpr)}
			hasTest(node, node.Value, HasValue(Anything()))
		})
	})

	t.Describe("hasRecvName", func() {
		t.It("is FuncDecl", func() {
			node := &ast.FuncDecl{}
			node.Recv = &ast.FieldList{}
			node.Recv.List = []*ast.Field{
				&ast.Field{Names: []*ast.Ident{new(ast.Ident)}},
			}
			hasTest(node, node.Recv.List[0].Names[0], HasRecvName(Anything()))
		})
	})

	t.Describe("hasRhs", func() {
		t.It("is AssignStmt", func() {
			node := &ast.AssignStmt{
				Rhs: []ast.Expr{new(ast.BadExpr)},
			}
			hasTest(node, node.Rhs[0], HasRhs(Anything()))
		})
	})

	t.Describe("hasResults", func() {
		t.It("is FuncType", func() {
			node := &ast.FuncType{
				Results: new(ast.FieldList),
			}
			hasTest(node, node.Results, HasResults(Anything()))
		})
	})

	t.Describe("hasType", func() {
		ger := HasType(Anything())
		t.It("is CompositeLit", func() {
			node := &ast.CompositeLit{Type: new(ast.BadExpr)}
			hasTest(node, node.Type, ger)
		})
		t.It("is Field", func() {
			node := &ast.Field{Type: new(ast.BadExpr)}
			hasTest(node, node.Type, ger)
		})
		t.It("is FuncDecl", func() {
			node := &ast.FuncDecl{Type: new(ast.FuncType)}
			hasTest(node, node.Type, ger)
		})
		t.It("is FuncLit", func() {
			node := &ast.FuncLit{Type: new(ast.FuncType)}
			hasTest(node, node.Type, ger)
		})
		t.It("is TypeAssertExpr", func() {
			node := &ast.TypeAssertExpr{Type: new(ast.BadExpr)}
			hasTest(node, node.Type, ger)
		})
		t.It("is TypeSpec", func() {
			node := &ast.TypeSpec{Type: new(ast.BadExpr)}
			hasTest(node, node.Type, ger)
		})
		t.It("is ValueSpec", func() {
			node := &ast.ValueSpec{Type: new(ast.BadExpr)}
			hasTest(node, node.Type, ger)
		})
	})

	t.Describe("hasCond", func() {
		ger := HasCond(Anything())
		t.It("is IfStmt", func() {
			node := &ast.IfStmt{
				Cond: new(ast.BadExpr),
			}
			hasTest(node, node.Cond, ger)
		})
		t.It("is ForStmt", func() {
			node := &ast.ForStmt{
				Cond: new(ast.BadExpr),
			}
			hasTest(node, node.Cond, ger)
		})
	})

	t.Describe("hasDecl", func() {
		ger := HasDecl(Anything())
		t.It("is DeclStmt", func() {
			node := &ast.DeclStmt{
				Decl: new(ast.BadDecl),
			}
			hasTest(node, node.Decl, ger)
		})
	})

	t.Describe("hasLen", func() {
		ger := HasLen(Anything())
		t.It("is ArrayType", func() {
			node := &ast.ArrayType{
				Len: new(ast.BadExpr),
			}
			hasTest(node, node.Len, ger)
		})
	})

	t.Describe("hasElement", func() {
		ger := HasElement(Anything())
		t.It("is ArrayType", func() {
			node := &ast.ArrayType{
				Elt: new(ast.BadExpr),
			}
			hasTest(node, node.Elt, ger)
		})
		t.It("is Ellipsis", func() {
			node := &ast.Ellipsis{
				Elt: new(ast.BadExpr),
			}
			hasTest(node, node.Elt, ger)
		})
	})

	t.Describe("hasLabel", func() {
		ger := HasLabel(Anything())
		t.It("is BranchStmt", func() {
			node := &ast.BranchStmt{
				Label: new(ast.Ident),
			}
			hasTest(node, node.Label, ger)
		})
		t.It("is LabeledStmt", func() {
			node := &ast.LabeledStmt{
				Label: new(ast.Ident),
			}
			hasTest(node, node.Label, ger)
		})
	})
})

var _ = t.Describe("Logic Goastcher", func() {
	var (
		node     = new(ast.BadExpr)
		falseGer = Unless(Anything())
		trueGer  = Anything()
	)
	t.Describe("Unless", func() {
		t.It("is true", func() {
			matched, err := goastch.Match(node, nil, Unless(falseGer))
			Expect(err).To(BeNil())
			Expect(matched).To(BeTrue())

		})
		t.It("is false", func() {
			matched, err := goastch.Match(node, nil, Unless(trueGer))
			Expect(err).To(BeNil())
			Expect(matched).To(BeFalse())
		})
	})

	t.Describe("AllOf", func() {
		t.It("is true", func() {
			matched, err := goastch.Match(new(ast.BadExpr), nil,
				AllOf(trueGer, trueGer, trueGer))
			Expect(err).To(BeNil())
			Expect(matched).To(BeTrue())
		})
		t.It("is false", func() {
			matched, err := goastch.Match(new(ast.BadExpr), nil,
				AllOf(trueGer, falseGer, trueGer))
			Expect(err).To(BeNil())
			Expect(matched).To(BeFalse())

			matched, err = goastch.Match(new(ast.BadExpr), nil,
				AllOf(falseGer, falseGer, falseGer))
			Expect(err).To(BeNil())
			Expect(matched).To(BeFalse())
		})
	})

	t.Describe("AnyOf", func() {
		t.It("is true", func() {
			matched, err := goastch.Match(new(ast.BadExpr), nil,
				AnyOf(trueGer, trueGer, trueGer))
			Expect(err).To(BeNil())
			Expect(matched).To(BeTrue())

			matched, err = goastch.Match(new(ast.BadExpr), nil,
				AnyOf(falseGer, trueGer, falseGer))
			Expect(err).To(BeNil())
			Expect(matched).To(BeTrue())
		})
		t.It("is false", func() {
			matched, err := goastch.Match(new(ast.BadExpr), nil,
				AnyOf(falseGer, falseGer, falseGer))
			Expect(err).To(BeNil())
			Expect(matched).To(BeFalse())
		})
	})
})

var _ = t.Describe("Attr Goastcher", func() {
})

func counts(root ast.Node) (count int) {
	ast.Inspect(root, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		count++
		return true
	})
	return
}

func parse(src string) *ast.File {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test", src, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}
	return f
}

func load(src string) (*ast.File, *types.Info) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test", src, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}

	info := types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}
	var conf types.Config
	_, err = conf.Check("fib", fset, []*ast.File{f}, &info)
	if err != nil {
		panic(err)
	}

	return f, &info
}
