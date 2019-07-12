// goastcher impliments error goastcher
package goastcher

import (
	"fmt"
	"go/ast"
)

type errorGer struct {
	err error
}

// ErrorGer goastcher maker of error
func ErrorGer(err error) Goastcher {
	return &errorGer{err: err}
}

func (g *errorGer) Goastch(ctx *Context, n ast.Node) bool {
	ctx.Err = g.err
	return false
}

func (g *errorGer) String() string {
	return fmt.Sprintf("errorGer \"%s\"", g.err.Error())
}

func (g *errorGer) Error() string {
	return g.err.Error()
}

func (g *errorGer) Bind(id string) Goastcher {
	return g
}

// IsError return true if the given goastcher is ErrorGer
func IsError(g Goastcher) bool {
	return isErrorGer(g)
}

func isErrorGer(g Goastcher) bool {
	_, is := g.(*errorGer)
	return is
}
