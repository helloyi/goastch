// main ...
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/printer"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/helloyi/goastch"
	"github.com/helloyi/goastch/galang/gen"
	"github.com/helloyi/goastch/galang/parser"
	gatoken "github.com/helloyi/goastch/galang/token"
	"github.com/c-bata/go-prompt"
	"golang.org/x/tools/go/buildutil"
	"golang.org/x/tools/go/loader"
)

// Pkg package info
type Pkg struct {
	ast.Package
	fset *token.FileSet
	info types.Info
}

// Interpreter interpreter
type Interpreter struct {
	loader  *loader.Config
	program *loader.Program
	writer  prompt.ConsoleWriter
}

var itpt *Interpreter

func main() {
	itpt = &Interpreter{
		loader: &loader.Config{Build: &build.Default},
		writer: prompt.NewStdoutWriter(),
	}

	p := prompt.New(func(src string) { itpt.exec(src) }, completer, prompt.OptionPrefix("ga> "))
	p.Run()
}

func (i *Interpreter) exec(src string) {
	stmt, err := i.genStmt(src)
	if err != nil {
		i.printErr(err)
		return
	}

	pkg := i.findPkg(stmt.Path)
	if pkg == nil {
		err = i.loadGo(stmt.Path)
		if err != nil {
			i.printErr(err)
			return
		}
		pkg = i.findPkg(stmt.Path)
	}

	binds, err := goastch.Find(&pkg.Package, &pkg.info, stmt.Ger)
	if err != nil {
		i.printErr(err)
		return
	}
	i.printBinds(binds, pkg.fset)
}

func (i *Interpreter) genStmt(src string) (*gen.Stmt, error) {
	node, err := parser.ParseStmt(src)
	if err != nil {
		return nil, err
	}
	return gen.GenStmt(node)
}

func (i *Interpreter) loadGo(path string) error {
	fnames, err := i.getFnames(path)
	if err != nil {
		return err
	}
	i.loader.CreateFromFilenames(path, fnames...)

	program, err := i.loader.Load()
	if err != nil {
		return err
	}
	i.program = program
	return nil
}

func (i *Interpreter) findPkg(path string) *Pkg {
	if i.program == nil {
		return nil
	}
	prog := i.program
	pkgInfo := prog.Package(path)
	if pkgInfo == nil {
		return nil
	}
	fset := prog.Fset
	files := make(map[string]*ast.File)
	for _, file := range pkgInfo.Files {
		fname := fset.File(file.Package).Name()
		files[fname] = file
	}

	pkg := &Pkg{}
	pkg.Name = pkgInfo.Pkg.Name()
	pkg.Files = files
	pkg.fset = fset
	pkg.info = pkgInfo.Info
	return pkg
}

func (i *Interpreter) printErr(err error) {
	i.writer.SetColor(prompt.Red, prompt.DefaultColor, false)
	i.writer.WriteStr(err.Error())
	i.writer.WriteStr("\n")
	_ = i.writer.Flush()
}

func (i *Interpreter) printBinds(binds map[string][]ast.Node, fset *token.FileSet) {
	var buf bytes.Buffer
	writer := i.writer
	for _, nodes := range binds {
		for _, node := range nodes {
			err := printer.Fprint(&buf, fset, node)
			if err != nil {
				i.printErr(err)
			} else {
				writer.SetColor(prompt.Green, prompt.DefaultColor, false)
				pos := fset.Position(node.Pos())
				_, fname := filepath.Split(pos.Filename)
				pos.Filename = fname
				writer.WriteStr(pos.String())
				writer.WriteStr(": ")

				writer.SetColor(prompt.Blue, prompt.DefaultColor, false)
				writer.Write(buf.Bytes())
				writer.WriteStr("\n")
				buf.Reset()

				_ = writer.Flush()
			}
		}
	}
}

func (i *Interpreter) getFnames(p string) ([]string, error) {
	file, err := os.Stat(p)
	if err != nil {
		return nil, err
	}
	if !file.IsDir() {
		return []string{p}, nil
	}

	files, err := buildutil.ReadDir(i.loader.Build, p)
	if err != nil {
		return nil, err
	}

	fnames := make([]string, 0)
	for _, file := range files {
		fname := file.Name()
		if strings.HasPrefix(fname, ".") {
			continue
		}
		if filepath.Ext(fname) != ".go" {
			continue
		}
		if strings.HasSuffix(fname, "_test.go") {
			continue
		}
		if file.IsDir() {
			continue
		}
		path := filepath.Join(p, fname)
		fnames = append(fnames, path)
	}
	return fnames, nil
}

var (
	travelGerSuggest []prompt.Suggest
	unodeGerSuggest  []prompt.Suggest
	allSuggest       []prompt.Suggest
)

func init() {
	tgds := gatoken.TravelGerDescs()
	ngds := gatoken.NodeGerDescs()
	wgds := gatoken.NarrowGerDescs()
	lgds := gatoken.LogicGerDescs()
	alds := gatoken.AllLitDescs()

	aglen := len(tgds) + len(ngds) + len(wgds) + len(lgds) + len(alds)

	allSuggest = make([]prompt.Suggest, 0, aglen)
	travelGerSuggest = make([]prompt.Suggest, 0, len(tgds))
	unodeGerSuggest = make([]prompt.Suggest, 0, aglen-len(ngds))

	allSuggest = sgappend(allSuggest, tgds, ngds, wgds, lgds, alds)
	travelGerSuggest = sgappend(travelGerSuggest, tgds)
	unodeGerSuggest = sgappend(unodeGerSuggest, tgds, wgds, lgds, alds)
}

func sgappend(s []prompt.Suggest, manyDescs ...map[string]string) []prompt.Suggest {
	for _, descs := range manyDescs {
		for name, desc := range descs {
			s = append(s, prompt.Suggest{
				Text:        name,
				Description: desc,
			})
		}
	}
	return s
}

func completer(d prompt.Document) []prompt.Suggest {
	fields := strings.Fields(d.CurrentLineBeforeCursor())
	if len(fields) == 0 {
		return []prompt.Suggest{}
	}

	var sub string
	var preToken string

	lastWord := d.GetWordBeforeCursor()
	if gatoken.What(strings.TrimLeft(lastWord, "@")) == gatoken.Unknow {
		i := len(fields) - 2
		if i < 0 {
			i = 0
		}
		preToken = fields[i]
		sub = lastWord
	} else {
		preToken = lastWord
	}
	preToken = strings.TrimLeft(preToken, "@")

	switch gatoken.What(preToken) {
	case gatoken.Node:
		return gerCompleter(unodeGerSuggest, sub)
	case gatoken.Narrow:
		return []prompt.Suggest{}
	case gatoken.From:
		return pathCompleter(sub)
	case gatoken.Unknow:
		return gerCompleter(travelGerSuggest, sub)
	default:
		return gerCompleter(allSuggest, sub)
	}
}

func gerCompleter(completions []prompt.Suggest, sub string) []prompt.Suggest {
	if sub == "" {
		return completions
	}

	var ret []prompt.Suggest
	var hasBind bool
	if strings.HasPrefix(sub, "@") {
		hasBind = true
		sub = strings.TrimLeft(sub, "@")
	}
	sub = strings.ToLower(sub)

	for _, c := range completions {
		if !strings.Contains(c.Text, sub) {
			continue
		}
		if hasBind {
			c.Text = "@" + c.Text
		}
		ret = append(ret, c)
	}
	return ret
}

func pathCompleter(prefix string) []prompt.Suggest {
	var s []prompt.Suggest
	if strings.HasSuffix(prefix, "\"") {
		return s
	}
	prefix = strings.Trim(prefix, "\"")
	dir, fnamePrefix := filepath.Split(prefix)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return s
	}

	s = make([]prompt.Suggest, 0, len(files))
	for _, file := range files {
		fname := file.Name()
		if strings.HasPrefix(fname, fnamePrefix) {
			text := filepath.Join(dir, fname)
			text = fmt.Sprintf("\"%s", text)
			s = append(s, prompt.Suggest{Text: text})
		}
	}
	return s
}
