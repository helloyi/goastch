package token

import (
	"strings"

	"github.com/helloyi/goastch/goastcher"
)

// Token ...
type Token int

const (
	// ILLEGAL ...
	ILLEGAL Token = iota
	EOF

	Node
	Narrow
	Travel

	Unless
	AnyOf
	AllOf

	Int
	Float
	Char
	String

	And
	End
	Bind
	From

	Unknow
)

var (
	buildinNodeObjs   map[string]object
	buildinTravelObjs map[string]object
	buildinNarrowObjs map[string]object
	buildinUnlessObjs map[string]object
	buildinAnyOfObjs  map[string]object
	buildinAllOfObjs  map[string]object

	andLits  map[string]object
	bindLits map[string]object
	endLits  map[string]object
	fromLits map[string]object
)

func (t Token) String() string {
	switch t {
	case ILLEGAL:
		return "illegal"
	case EOF:
		return "EOF"
	case Node:
		return "node"
	case Narrow:
		return "narrow"
	case Travel:
		return "travel"
	case Unless:
		return "unless"
	case AnyOf:
		return "anyof"
	case AllOf:
		return "allof"
	case Int:
		return "int"
	case Float:
		return "float"
	case Char:
		return "char"
	case String:
		return "string"
	case And:
		return "and"
	case End:
		return "end"
	case Bind:
		return "bind"
	case From:
		return "from"
	case Unknow:
		return "unknow"
	}
	return ""
}

// What ...
func What(lit string) Token {
	if isNodeLit(lit) {
		return Node
	} else if isNarrowLit(lit) {
		return Narrow
	} else if isTravelLit(lit) {
		return Travel
	} else if isUnlessLit(lit) {
		return Unless
	} else if isAllOfLit(lit) {
		return AllOf
	} else if isAnyOfLit(lit) {
		return AnyOf
	} else if isAndLit(lit) {
		return And
	} else if isEndLit(lit) {
		return End
	} else if isBindLit(lit) {
		return Bind
	} else if isFromLit(lit) {
		return From
	}
	return Unknow
}

func smappend(a map[string]string, b map[string]string) {
	for k, v := range b {
		a[k] = v
	}
}

// LogicGerDescs ...
func LogicGerDescs() map[string]string {
	descs := AllOfGerDescs()
	smappend(descs, AnyOfGerDescs())
	smappend(descs, UnlessGerDescs())
	return descs
}

// AllLitDescs ...
func AllLitDescs() map[string]string {
	descs := BindLitDescs()
	smappend(descs, AndLitDescs())
	smappend(descs, EndLitDescs())
	smappend(descs, FromLitDescs())
	return descs
}

// NodeGerDescs ...
func NodeGerDescs() map[string]string {
	descs := make(map[string]string)
	for ger, obj := range buildinNodeObjs {
		descs[ger] = obj.desc
	}
	return descs
}

// TravelGerDescs ...
func TravelGerDescs() map[string]string {
	descs := make(map[string]string)
	for ger, obj := range buildinTravelObjs {
		descs[ger] = obj.desc
	}
	return descs
}

// NarrowGerDescs ...
func NarrowGerDescs() map[string]string {
	descs := make(map[string]string)
	for ger, obj := range buildinNarrowObjs {
		descs[ger] = obj.desc
	}
	return descs
}

// AnyOfGerDescs ...
func AnyOfGerDescs() map[string]string {
	descs := make(map[string]string)
	for ger, obj := range buildinAnyOfObjs {
		descs[ger] = obj.desc
	}
	return descs
}

// AllOfGerDescs ...
func AllOfGerDescs() map[string]string {
	descs := make(map[string]string)
	for ger, obj := range buildinAllOfObjs {
		descs[ger] = obj.desc
	}
	return descs
}

// UnlessGerDescs ...
func UnlessGerDescs() map[string]string {
	descs := make(map[string]string)
	for ger, obj := range buildinUnlessObjs {
		descs[ger] = obj.desc
	}
	return descs
}

// AndLitDescs ...
func AndLitDescs() map[string]string {
	descs := make(map[string]string)
	for lit, obj := range andLits {
		descs[lit] = obj.desc
	}
	return descs
}

// EndLitDescs ...
func EndLitDescs() map[string]string {
	descs := make(map[string]string)
	for ger, obj := range endLits {
		descs[ger] = obj.desc
	}
	return descs
}

// BindLitDescs ...
func BindLitDescs() map[string]string {
	descs := make(map[string]string)
	for ger, obj := range bindLits {
		descs[ger] = obj.desc
	}
	return descs
}

// FromLitDescs ...
func FromLitDescs() map[string]string {
	descs := make(map[string]string)
	for ger, obj := range fromLits {
		descs[ger] = obj.desc
	}
	return descs
}

// Object ...
func Object(lit string) interface{} {
	l := strings.ToLower(lit)
	if obj := buildinNodeObjs[l].ger; obj != nil {
		return obj
	}
	if obj := buildinNarrowObjs[l].ger; obj != nil {
		return obj
	}
	if obj := buildinTravelObjs[l].ger; obj != nil {
		return obj
	}
	if obj := buildinUnlessObjs[l].ger; obj != nil {
		return obj
	}
	if obj := buildinAnyOfObjs[l].ger; obj != nil {
		return obj
	}
	return buildinAllOfObjs[l].ger
}

func isNodeLit(lit string) bool {
	_, ok := buildinNodeObjs[strings.ToLower(lit)]
	return ok
}

func isNarrowLit(lit string) bool {
	_, ok := buildinNarrowObjs[strings.ToLower(lit)]
	return ok
}

func isTravelLit(lit string) bool {
	_, ok := buildinTravelObjs[strings.ToLower(lit)]
	return ok
}

func isUnlessLit(lit string) bool {
	_, ok := buildinUnlessObjs[strings.ToLower(lit)]
	return ok
}

func isAnyOfLit(lit string) bool {
	_, ok := buildinAnyOfObjs[strings.ToLower(lit)]
	return ok
}

func isAllOfLit(lit string) bool {
	_, ok := buildinAllOfObjs[strings.ToLower(lit)]
	return ok
}

func isAndLit(lit string) bool {
	_, ok := andLits[strings.ToLower(lit)]
	return ok
}

func isEndLit(lit string) bool {
	_, ok := endLits[strings.ToLower(lit)]
	return ok
}

func isBindLit(lit string) bool {
	_, ok := bindLits[strings.ToLower(lit)]
	return ok
}

func isFromLit(lit string) bool {
	_, ok := fromLits[strings.ToLower(lit)]
	return ok
}

type object struct {
	ger  interface{}
	desc string
}

func init() {
	andLits = map[string]object{
		"and": {true, "connect goastchers of anyof/allof"},
	}

	endLits = map[string]object{
		",": {true, "end goastcher stmtment"},
	}

	bindLits = map[string]object{
		"@": {true, "binding goastcher"},
	}

	fromLits = map[string]object{
		"from": {true, "from \"go/package/path\""},
	}
	buildinNodeObjs = map[string]object{
		"arraytype": {
			goastcher.ArrayType,
			"match array node (slice and array decl)",
		},
		"assignstmt": {
			goastcher.AssignStmt,
			"match assign stmt",
		},
		"baddecl": {
			goastcher.BadDecl,
			"bad decl",
		},
		"badexpr": {
			goastcher.BadExpr,
			"bad expr",
		},
		"badstmt": {
			goastcher.BadStmt,
			"bad stmt"},
		"basiclit": {
			goastcher.BasicLit,
			"literal of basic type",
		},
		"binaryexpr": {
			goastcher.BinaryExpr,
			"binary expression",
		},
		"blockstmt": {
			goastcher.BlockStmt,
			"braced statement list",
		},
		"branchstmt": {
			goastcher.BranchStmt,
			"break, continue, goto, or fallthrough statement",
		},
		"callexpr": {
			goastcher.CallExpr,
			"",
		},
		"caseclause": {
			goastcher.CaseClause,
			"case of an expression or type switch statement",
		},
		"chantype": {
			goastcher.ChanType,
			"channel type",
		},
		"commclause": {
			goastcher.CommClause,
			"a case of a select statement",
		},
		"comment": {
			goastcher.Comment,
			"a single //-style or /*-style comment",
		},
		"commentgroup": {
			goastcher.CommentGroup,
			"comments with no other tokens and no empty lines between",
		},
		"compositelit": {
			goastcher.CompositeLit,
			"a composite literal",
		},
		"declstmt": {
			goastcher.DeclStmt,
			"a declaration in a statement list",
		},
		"deferstmt": {
			goastcher.DeferStmt,
			"a defer statement",
		},
		"ellipsis": {
			goastcher.Ellipsis,
			"the \"...\" type in a parameter list or the \"...\" length in an array type",
		},
		"emptystmt": {
			goastcher.EmptyStmt,
			"explicit or implicit semicolon",
		},
		"exprstmt": {
			goastcher.ExprStmt,
			"a (stand-alone) expression in a statement list",
		},
		"field": {
			goastcher.Field,
			"a struct field list, a interface method list, or a parameter/result",
		},
		"fieldlist": {
			goastcher.FieldList,
			"a list of fields, enclosed by parentheses or braces",
		},
		"file": {
			goastcher.File,
			"a go source file"},

		"forstmt": {
			goastcher.ForStmt,
			"a for statement"},

		"funcdecl": {
			goastcher.FuncDecl,
			"a function declaration"},

		"funclit": {
			goastcher.FuncLit,
			"a function literal"},

		"functype": {
			goastcher.FuncType,
			"a function type"},

		"gendecl": {
			goastcher.GenDecl,
			"an import, constant, type or variable declaration"},

		"gostmt": {
			goastcher.GoStmt,
			"a go statement"},

		"ident": {
			goastcher.Ident,
			"an identifier"},

		"ifstmt": {
			goastcher.IfStmt,
			"an if statement"},

		"importspec": {
			goastcher.ImportSpec,
			"a single package import"},
		"incdecstmt": {
			goastcher.IncDecStmt,
			"an increment or decrement statement"},
		"indexexpr": {
			goastcher.IndexExpr,
			"an expression followed by an index"},
		"interfacetype": {
			goastcher.InterfaceType,
			"an interface type",
		},
		"keyvalueexpr": {
			goastcher.KeyValueExpr,
			"(key : value) pairs in composite literals",
		},
		"labeledstmt": {
			goastcher.LabeledStmt,
			"a labeled statement",
		},
		"maptype": {
			goastcher.MapType,
			"a map type",
		},
		"package": {
			goastcher.Pkg,
			"a go package",
		},
		"parenexpr": {
			goastcher.ParenExpr,
			"a parenthesized expression",
		},
		"rangestmt": {
			goastcher.RangeStmt,
			"a for statement with a range clause",
		},
		"returnstmt": {
			goastcher.ReturnStmt,
			"a return statement",
		},
		"selectstmt": {
			goastcher.SelectStmt,
			"a select statement",
		},
		"selectorexpr": {
			goastcher.SelectorExpr,
			"an expression followed by a selector",
		},
		"sendstmt": {
			goastcher.SendStmt,
			"a send statement",
		},
		"sliceexpr": {
			goastcher.SliceExpr,
			"an expression followed by slice indices",
		},
		"starexpr": {
			goastcher.StarExpr,
			"\"*\" expression",
		},
		"structtype": {
			goastcher.StructType,
			"a struct type",
		},
		"switchstmt": {
			goastcher.SwitchStmt,
			"an expression switch statement",
		},
		"typeassertexpr": {
			goastcher.TypeAssertExpr,
			"an expression followed by a type assertion",
		},
		"typespec": {
			goastcher.TypeSpec,
			"a type declaration",
		},
		"typeswitchstmt": {
			goastcher.TypeSwitchStmt,
			"a type switch statement",
		},
		"unaryexpr": {
			goastcher.UnaryExpr,
			"a unary expression",
		},
		"valuespec": {
			goastcher.ValueSpec,
			"a constant or variable declaration",
		},

		"shortvardecl": {
			goastcher.ShortVarDecl,
			"a short variable declaration",
		},
		"slicetype": {
			goastcher.SliceType,
			"a slice type",
		},
		"intBasicLit": {
			goastcher.IntBasicLit,
			"an int literal",
		},
		"floatBasicLit": {
			goastcher.FloatBasicLit,
			"a float literal",
		},
		"imagBasicLit": {
			goastcher.ImagBasicLit,
			"an imag literal",
		},
		"charBasicLit": {
			goastcher.CharBasicLit,
			"a char literal ",
		},
		"stringBasicLit": {
			goastcher.StringBasicLit,
			"a string literal",
		},
	}
	buildinNarrowObjs = map[string]object{
		"ascode": {
			goastcher.AsCode,
			"matches if the matched node is represented by the given code",
		},
		"issize": {
			goastcher.IsSize,
			"matches nodes that are the specified size.",
		},
		"istype": {
			goastcher.IsType,
			"matches nodes that have the specified type.",
		},
		"equals": {
			goastcher.Equals,
			"matches literals that are equal to the given value",
		},
		"hasprefix": {
			goastcher.HasPrefix,
			"matches identifier that have the specified prefix",
		},
		"hassuffix": {
			goastcher.HasSuffix,
			"matches identifier that have the specified suffix",
		},
		"contains": {
			goastcher.Contains,
			"matches identifier that have the specified substr",
		},
		"matchstring": {
			goastcher.MatchString,
			"matches identifier that match the specified regexp",
		},
		"hasoperator": {
			goastcher.HasOperator,
			"matchs nodes that have the specified operator",
		},
		"anything": {
			goastcher.Anything,
			"matches any nodes",
		},
		"isexported": {
			goastcher.IsExported,
			"matches nodes that are exported",
		},
	}
	buildinTravelObjs = map[string]object{
		"has": {
			goastcher.Has,
			"matches nodes that have child nodes that match the provided goastcher.",
		},
		"hasdes": {
			goastcher.HasDescendant,
			"Matches nodes that have descendant nodes that match the provided goastcher",
		},
		"hasvalue": {
			goastcher.HasValue,
			"matches nodes that have Value node that match the specific goastcher",
		},
		"hasname": {
			goastcher.HasName,
			"matches nodes that have Name node that match the specific goastcher",
		},
		"fordecls": {
			goastcher.ForDecls,
			"matches each Decls by the given goastcher",
		},
		"forspecs": {
			goastcher.ForSpecs,
			"matches each Specs by the given goastcher",
		},
		"fornames": {
			goastcher.ForNames,
			"matches each Names by the given goastcher",
		},
		"forfields": {
			goastcher.ForFields,
			"matches each Field by the given goastcher",
		},
		"hasrecvname": {
			goastcher.HasRecvName,
			"",
		},
		"hasrhs": {
			goastcher.HasRhs,
			"",
		},
		"hasresults": {
			goastcher.HasResults,
			"matches Results by the given goastcher",
		},
		"hastype": {
			goastcher.HasType,
			"matches Types by the given goastcher",
		},
		"last": {
			goastcher.Last,
			"matches last field of composite node by the given goastcher",
		},
	}
	buildinUnlessObjs = map[string]object{
		"unless": {
			goastcher.Unless,
			"matches if the provided goastcher does not match",
		},
		"not": {
			goastcher.Unless,
			"matches if the provided goastcher does not match",
		},
	}

	buildinAllOfObjs = map[string]object{
		"allof": {
			goastcher.AllOf,
			"matches if all given matchers match",
		},
	}

	buildinAnyOfObjs = map[string]object{
		"anyof": {
			goastcher.AnyOf,
			"matches if any of the given matchers matches",
		},
	}
}
