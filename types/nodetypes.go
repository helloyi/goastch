// types node type of go ast
package types

import (
	"go/ast"
)

// Typ node type of ast
type Typ ast.Node

var (
	ArrayType      *ast.ArrayType
	AssignStmt     *ast.AssignStmt
	BadDecl        *ast.BadDecl
	BadExpr        *ast.BadExpr
	BadStmt        *ast.BadStmt
	BasicLit       *ast.BasicLit
	BinaryExpr     *ast.BinaryExpr
	BlockStmt      *ast.BlockStmt
	BranchStmt     *ast.BranchStmt
	CallExpr       *ast.CallExpr
	CaseClause     *ast.CaseClause
	ChanType       *ast.ChanType
	CommClause     *ast.CommClause
	Comment        *ast.Comment
	CommentGroup   *ast.CommentGroup
	CompositeLit   *ast.CompositeLit
	DeclStmt       *ast.DeclStmt
	DeferStmt      *ast.DeferStmt
	Ellipsis       *ast.Ellipsis
	EmptyStmt      *ast.EmptyStmt
	ExprStmt       *ast.ExprStmt
	Field          *ast.Field
	FieldList      *ast.FieldList
	File           *ast.File
	ForStmt        *ast.ForStmt
	FuncDecl       *ast.FuncDecl
	FuncLit        *ast.FuncLit
	FuncType       *ast.FuncType
	GenDecl        *ast.GenDecl
	GoStmt         *ast.GoStmt
	Ident          *ast.Ident
	IfStmt         *ast.IfStmt
	ImportSpec     *ast.ImportSpec
	IncDecStmt     *ast.IncDecStmt
	IndexExpr      *ast.IndexExpr
	InterfaceType  *ast.InterfaceType
	KeyValueExpr   *ast.KeyValueExpr
	LabeledStmt    *ast.LabeledStmt
	MapType        *ast.MapType
	Package        *ast.Package
	ParenExpr      *ast.ParenExpr
	RangeStmt      *ast.RangeStmt
	ReturnStmt     *ast.ReturnStmt
	SelectStmt     *ast.SelectStmt
	SelectorExpr   *ast.SelectorExpr
	SendStmt       *ast.SendStmt
	SliceExpr      *ast.SliceExpr
	StarExpr       *ast.StarExpr
	StructType     *ast.StructType
	SwitchStmt     *ast.SwitchStmt
	TypeAssertExpr *ast.TypeAssertExpr
	TypeSpec       *ast.TypeSpec
	TypeSwitchStmt *ast.TypeSwitchStmt
	UnaryExpr      *ast.UnaryExpr
	ValueSpec      *ast.ValueSpec

	ntypeName = map[ast.Node]string{
		ArrayType:      "ArrayType",
		AssignStmt:     "AssignStmt",
		BadDecl:        "BadDecl",
		BadExpr:        "BadExpr",
		BadStmt:        "BadStmt",
		BasicLit:       "BasicLit",
		BinaryExpr:     "BinaryExpr",
		BlockStmt:      "BlockStmt",
		BranchStmt:     "BranchStmt",
		CallExpr:       "CallExpr",
		CaseClause:     "CaseClause",
		ChanType:       "ChanType",
		CommClause:     "CommClause",
		Comment:        "Comment",
		CommentGroup:   "CommentGroup",
		CompositeLit:   "CompositeLit",
		DeclStmt:       "DeclStmt",
		DeferStmt:      "DeferStmt",
		Ellipsis:       "Ellipsis",
		EmptyStmt:      "EmptyStmt",
		ExprStmt:       "ExprStmt",
		Field:          "Field",
		FieldList:      "FieldList",
		File:           "File",
		ForStmt:        "ForStmt",
		FuncDecl:       "FuncDecl",
		FuncLit:        "FuncLit",
		FuncType:       "FuncType",
		GenDecl:        "GenDecl",
		GoStmt:         "GoStmt",
		Ident:          "Ident",
		IfStmt:         "IfStmt",
		ImportSpec:     "ImportSpec",
		IncDecStmt:     "IncDecStmt",
		IndexExpr:      "IndexExpr",
		InterfaceType:  "InterfaceType",
		KeyValueExpr:   "KeyValueExpr",
		LabeledStmt:    "LabeledStmt",
		MapType:        "MapType",
		Package:        "Package",
		ParenExpr:      "ParenExpr",
		RangeStmt:      "RangeStmt",
		ReturnStmt:     "ReturnStmt",
		SelectStmt:     "SelectStmt",
		SelectorExpr:   "SelectorExpr",
		SendStmt:       "SendStmt",
		SliceExpr:      "SliceExpr",
		StarExpr:       "StarExpr",
		StructType:     "StructType",
		SwitchStmt:     "SwitchStmt",
		TypeAssertExpr: "TypeAssertExpr",
		TypeSpec:       "TypeSpec",
		TypeSwitchStmt: "TypeSwitchStmt",
		UnaryExpr:      "UnaryExpr",
		ValueSpec:      "ValueSpec",
	}
)

// TypeName ...
func TypeName(n ast.Node) string {
	if n == nil {
		return "Nil"
	}
	return ntypeName[Type(n)]
}

// Type return node type of 'n'
func Type(n ast.Node) Typ {
	switch n.(type) {
	case *ast.ArrayType:
		return ArrayType
	case *ast.AssignStmt:
		return AssignStmt
	case *ast.BadDecl:
		return BadDecl
	case *ast.BadExpr:
		return BadExpr
	case *ast.BadStmt:
		return BadStmt
	case *ast.BasicLit:
		return BasicLit
	case *ast.BinaryExpr:
		return BinaryExpr
	case *ast.BlockStmt:
		return BlockStmt
	case *ast.BranchStmt:
		return BranchStmt
	case *ast.CallExpr:
		return CallExpr
	case *ast.CaseClause:
		return CaseClause
	case *ast.ChanType:
		return ChanType
	case *ast.CommClause:
		return CommClause
	case *ast.Comment:
		return Comment
	case *ast.CommentGroup:
		return CommentGroup
	case *ast.CompositeLit:
		return CompositeLit
	case *ast.DeclStmt:
		return DeclStmt
	case *ast.DeferStmt:
		return DeferStmt
	case *ast.Ellipsis:
		return Ellipsis
	case *ast.EmptyStmt:
		return EmptyStmt
	case *ast.ExprStmt:
		return ExprStmt
	case *ast.Field:
		return Field
	case *ast.FieldList:
		return FieldList
	case *ast.File:
		return File
	case *ast.ForStmt:
		return ForStmt
	case *ast.FuncDecl:
		return FuncDecl
	case *ast.FuncLit:
		return FuncLit
	case *ast.FuncType:
		return FuncType
	case *ast.GenDecl:
		return GenDecl
	case *ast.GoStmt:
		return GoStmt
	case *ast.Ident:
		return Ident
	case *ast.IfStmt:
		return IfStmt
	case *ast.ImportSpec:
		return ImportSpec
	case *ast.IncDecStmt:
		return IncDecStmt
	case *ast.IndexExpr:
		return IndexExpr
	case *ast.InterfaceType:
		return InterfaceType
	case *ast.KeyValueExpr:
		return KeyValueExpr
	case *ast.LabeledStmt:
		return LabeledStmt
	case *ast.MapType:
		return MapType
	case *ast.Package:
		return Package
	case *ast.ParenExpr:
		return ParenExpr
	case *ast.RangeStmt:
		return RangeStmt
	case *ast.ReturnStmt:
		return ReturnStmt
	case *ast.SelectStmt:
		return SelectStmt
	case *ast.SelectorExpr:
		return SelectorExpr
	case *ast.SendStmt:
		return SendStmt
	case *ast.SliceExpr:
		return SliceExpr
	case *ast.StarExpr:
		return StarExpr
	case *ast.StructType:
		return StructType
	case *ast.SwitchStmt:
		return SwitchStmt
	case *ast.TypeAssertExpr:
		return TypeAssertExpr
	case *ast.TypeSpec:
		return TypeSpec
	case *ast.TypeSwitchStmt:
		return TypeSwitchStmt
	case *ast.UnaryExpr:
		return UnaryExpr
	case *ast.ValueSpec:
		return ValueSpec
	default:
		return nil
	}
}
