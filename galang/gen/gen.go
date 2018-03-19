package gen

import (
	"fmt"

	"github.com/helloyi/goastch/galang/ast"
	"github.com/helloyi/goastch/goastcher"
)

// Ger ...
func Ger(n ast.Ger) (goastcher.Goastcher, error) {
	ger := genGer(n)
	if goastcher.IsError(ger) {
		return nil, fmt.Errorf("%s", ger)
	}
	return ger, nil
}

// Stmt ...
type Stmt struct {
	Ger   goastcher.Goastcher
	Path  string
	IsPkg bool
}

// GenStmt ...
func GenStmt(n *ast.Stmt) (*Stmt, error) {
	ger, err := Ger(n.Ger)
	if err != nil {
		return nil, err
	}
	return &Stmt{
		Ger:  ger,
		Path: n.Path,
	}, nil
}

// genGer ...
func genGer(root ast.Ger) goastcher.Goastcher {
	switch n := root.(type) {
	case *ast.NodeGer:
		nextGer := genGer(n.Next)
		obj := n.Node.Object.(func(goastcher.Goastcher) goastcher.Goastcher)
		ger := obj(nextGer)
		if n.IsBind {
			ger = ger.Bind("")
		}
		return ger
	case *ast.TravelGer:
		nextGer := genGer(n.Next)
		obj := n.Travel.Object.(func(goastcher.Goastcher) goastcher.Goastcher)
		ger := obj(nextGer)
		if n.IsBind {
			ger = ger.Bind("")
		}
		return ger
	case *ast.LogicGer:
		switch g := n.Ger.(type) {
		case *ast.CompositeGer:
			gers := genCompositeGer(g)
			obj := n.Op.Object.(func(...goastcher.Goastcher) goastcher.Goastcher)
			return obj(gers...)
		default:
			ger := genGer(n.Ger)
			obj := n.Op.Object.(func(goastcher.Goastcher) goastcher.Goastcher)
			return obj(ger)
		}
	case *ast.NarrowGer:
		return genNarrowGer(n)
	case *ast.EmptyGer:
		return goastcher.ErrorGer(fmt.Errorf("%s: required goastcher", n.Pos))
	default:
		return nil
	}
}

func genNarrowGer(n *ast.NarrowGer) goastcher.Goastcher {
	if len(n.Args) > 1 {
		panic("just supported zero/one arg")
	}
	var arg interface{}
	if len(n.Args) != 0 {
		arg = n.Args[0].Value
	}
	switch obj := n.Narrow.Object.(type) {
	case func(int64) goastcher.Goastcher:
		return obj(arg.(int64))
	case func(float64) goastcher.Goastcher:
		return obj(arg.(float64))
	case func(string) goastcher.Goastcher:
		return obj(arg.(string))
	case func(interface{}) goastcher.Goastcher:
		return obj(arg)
	case func() goastcher.Goastcher:
		return obj()
	default:
		panic("unsupported narrow ger")
	}
}

func genCompositeGer(n *ast.CompositeGer) []goastcher.Goastcher {
	gers := make([]goastcher.Goastcher, 0)
	for _, ger := range n.Gers {
		gers = append(gers, genGer(ger))
	}
	return gers
}
