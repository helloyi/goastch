package goastcher

import (
	"go/ast"
	"go/types"
)

// Goastcher ...
type Goastcher interface {
	Goastch(ctx *Context, node ast.Node) bool
	Bind(id string) Goastcher
}

// Context ...
type Context struct {
	Root ast.Node    // root node of ast
	Info *types.Info // type info

	Err      error
	Bindings map[string][]ast.Node // matched result
}

// IsPioneer ...
func IsPioneer(g Goastcher) bool {
	switch g.(type) {
	case *has, *hasDescendant, *hasName, *hasValue,
		*forDecls, *forSpecs, *forNames:
		return true
	default:
		return false
	}
}

func (ctx *Context) growBindings(id string, node ast.Node) {
	if id == "" {
		return
	}
	if ctx.Bindings == nil {
		ctx.Bindings = make(map[string][]ast.Node)
	}
	if ctx.Bindings[id] == nil {
		ctx.Bindings[id] = make([]ast.Node, 0)
	}
	ctx.Bindings[id] = append(ctx.Bindings[id], node)
}
