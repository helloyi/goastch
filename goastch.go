// goastch ...
package goastch

import (
	"go/ast"
	"go/types"

	"github.com/helloyi/goastch/goastcher"
)

// Goastch ...
type Goastch struct {
	ctx *goastcher.Context
	ger goastcher.Goastcher
}

// New ...
func New(root ast.Node, info *types.Info, ger goastcher.Goastcher) (*Goastch, error) {
	if goastcher.IsError(ger) {
		return nil, ger.(error)
	}

	return &Goastch{
		ctx: &goastcher.Context{
			Root: root,
			Info: info,
		},
		ger: ger,
	}, nil
}

// Find ...
func (g *Goastch) Find(node ast.Node) (map[string][]ast.Node, error) {
	matched := g.ger.Goastch(g.ctx, node)
	if g.ctx.Err != nil {
		return nil, g.ctx.Err
	}
	if !matched {
		return nil, nil
	}
	return g.ctx.Bindings, nil
}

// Match ...
func (g *Goastch) Match(node ast.Node) (bool, error) {
	matched := g.ger.Goastch(g.ctx, node)
	if g.ctx.Err != nil {
		return false, g.ctx.Err
	}
	return matched, nil
}

// Match ...
func Match(n ast.Node, i *types.Info, g goastcher.Goastcher) (bool, error) {
	goastch, err := New(n, i, g)
	if err != nil {
		return false, err
	}
	return goastch.Match(n)
}

// Find ...
func Find(n ast.Node, i *types.Info, g goastcher.Goastcher) (map[string][]ast.Node, error) {
	goastch, err := New(n, i, g)
	if err != nil {
		return nil, err
	}
	return goastch.Find(n)
}
