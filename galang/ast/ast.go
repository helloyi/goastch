package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/helloyi/goastch/galang/scanner"
	"github.com/helloyi/goastch/galang/token"
)

type (
	// Node ...
	Node interface{}

	// Stmt ...
	Stmt struct {
		Ger   Ger
		Path  string
		IsPkg bool
	}

	// Ger ...
	Ger interface {
		format(indent int) string
	}

	// NodeGer ...
	NodeGer struct {
		Node   *Ident
		Next   Ger
		IsBind bool
	}

	// NarrowGer ...
	NarrowGer struct {
		Narrow *Ident
		Args   []*ValueSpec
	}

	// ValueSpec ...
	ValueSpec struct {
		Value interface{}
	}

	// TravelGer ...
	TravelGer struct {
		Travel *Ident
		Next   Ger
		IsBind bool
	}

	// LogicGer ...
	LogicGer struct {
		Op  *Ident
		Ger Ger
	}

	// CompositeGer ...
	CompositeGer struct {
		Gers []Ger
	}

	// Ident ...
	Ident struct {
		Pos    scanner.Position
		Name   string
		Token  token.Token
		Object interface{}
	}

	// EmptyGer ...
	EmptyGer struct {
		Pos scanner.Position
	}
)

const (
	indentStr = "  "
)

func (n *NodeGer) String() string {
	return strings.TrimSpace(n.format(0))
}

func (n *NarrowGer) String() string {
	return strings.TrimSpace(n.format(0))
}

func (n *ValueSpec) String() string {
	return strings.TrimSpace(n.format(0))
}

func (n *TravelGer) String() string {
	return strings.TrimSpace(n.format(0))
}

func (n *LogicGer) String() string {
	return strings.TrimSpace(n.format(0))
}

func (n *CompositeGer) String() string {
	return strings.TrimSpace(n.format(0))
}

func (n *Ident) String() string {
	return strings.TrimSpace(n.format(0))
}

func (n *EmptyGer) String() string {
	return strings.TrimSpace(n.format(0))
}

func (n *NodeGer) format(indent int) string {
	buf := bytes.Buffer{}
	str := n.Node.format(indent)
	if n.IsBind {
		str = strings.TrimRight(str, "\n")
		buf.WriteString(str)
		buf.WriteString(", binded\n")
	} else {
		buf.WriteString(str)
	}
	buf.WriteString(n.Next.format(indent + 1))
	return buf.String()
}

func (n *NarrowGer) format(indent int) string {
	buf := bytes.Buffer{}
	buf.WriteString(n.Narrow.format(indent))
	if len(n.Args) == 0 {
		padding := strings.Repeat(indentStr, indent)
		buf.WriteString(fmt.Sprintf("%s%s(empty args)\n", padding, indentStr))
		return buf.String()
	}
	buf.WriteString(n.Args[0].format(indent + 1))
	for _, arg := range n.Args[1:] {
		buf.WriteString(arg.format(indent + 1))
	}
	return buf.String()
}

func (n *ValueSpec) format(indent int) string {
	padding := strings.Repeat(indentStr, indent)
	return fmt.Sprintf("%s%s\n", padding, n.Value)
}

func (n *TravelGer) format(indent int) string {
	buf := bytes.Buffer{}
	str := n.Travel.format(indent)
	if n.IsBind {
		str = strings.TrimRight(str, "\n")
		buf.WriteString(str)
		buf.WriteString(", binded\n")
	} else {
		buf.WriteString(str)
	}
	buf.WriteString(n.Next.format(indent + 1))
	return buf.String()
}

func (n *LogicGer) format(indent int) string {
	buf := bytes.Buffer{}
	buf.WriteString(n.Op.format(indent))
	buf.WriteString(n.Ger.format(indent + 1))
	return buf.String()
}

func (n *CompositeGer) format(indent int) string {
	buf := bytes.Buffer{}
	for _, ger := range n.Gers[:len(n.Gers)-1] {
		buf.WriteString(ger.format(indent))
	}
	buf.WriteString(n.Gers[len(n.Gers)-1].format(indent))
	return buf.String()
}

func (n *Ident) format(indent int) string {
	padding := strings.Repeat(indentStr, indent)
	return fmt.Sprintf("%s%s(%s)\n", padding, n.Name, n.Token)
}

func (n *EmptyGer) format(indent int) string {
	padding := strings.Repeat(indentStr, indent)
	return fmt.Sprintf("%semptyGer\n", padding)
}
