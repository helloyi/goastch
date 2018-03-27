package goastcher

import (
	"go/ast"
	"go/types"
)

// Goastcher go ast matcher
type Goastcher interface {
	// Goastch matcher interface
	Goastch(ctx *Context, node ast.Node) bool
	// Bind binding 'id' to a goastcher, the matched nodes will bind to 'id'.
	// Call goatch.Find get the result.
	Bind(id string) Goastcher
}

// Context context of goastcher links
type Context struct {
	Root     ast.Node              // root node of ast
	Info     *types.Info           // type info
	Err      error                 // error when matching
	Bindings map[string][]ast.Node // matched result
}

// growBindings grow Context.Bindings
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
