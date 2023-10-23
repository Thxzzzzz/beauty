// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2

import (
	"github.com/rushteam/beauty/tools/internal/parser/ast"
)

const EOF = 0

// go install golang.org/x/tools/cmd/goyacc@latest
// goyacc -o parser.go parser.go.y

//line parser.go.y:14
type yySymType struct {
	yys int
	// empty struct{}
	token         string
	str           string
	services      []*ast.Service
	service       *ast.Service
	rpcs          []*ast.RPC
	rpc           *ast.RPC
	route         *ast.Route
	routes        []*ast.Route
	route_methods []string
}

const ILLEGAL = 57346
const Service = 57347
const Rpc = 57348
const Returns = 57349
const Route = 57350
const Comment = 57351
const Ident = 57352
const String = 57353

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ILLEGAL",
	"Service",
	"Rpc",
	"Returns",
	"Route",
	"Comment",
	"Ident",
	"String",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"'|'",
	"'.'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:131

/*  start  of  programs  */

/*
service helloworld {
	@route GET|POST "/index/:id"
    rpc Index(getRequest) returns (getResponse)
    rpc Helloworld(getRequest) returns (getResponse)
}
*/
//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 60

var yyAct = [...]int8{
	28, 24, 25, 27, 23, 57, 52, 26, 56, 53,
	51, 44, 42, 37, 35, 12, 12, 14, 14, 50,
	45, 43, 15, 9, 13, 36, 34, 7, 55, 54,
	49, 47, 40, 38, 33, 32, 18, 31, 30, 29,
	22, 21, 19, 6, 17, 10, 14, 48, 46, 41,
	39, 4, 3, 20, 16, 5, 11, 8, 2, 1,
}

var yyPact = [...]int16{
	46, -1000, 46, -1000, 33, -1000, 15, 10, 9, -1000,
	-1000, 38, 32, -1000, 31, -1000, -1000, 30, -1000, -13,
	-9, -1000, -14, 29, 28, -1000, 27, 25, 24, 12,
	-1, -1000, 11, -2, 23, 43, 22, 42, -3, 7,
	-4, 6, 41, 21, 40, 20, 5, -5, -8, -6,
	19, -1000, 18, -1000, -7, -10, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 59, 58, 52, 45, 57, 24, 56, 53,
}

var yyR1 = [...]int8{
	0, 1, 1, 2, 2, 3, 3, 5, 5, 4,
	4, 4, 4, 7, 7, 6, 8, 8,
}

var yyR2 = [...]int8{
	0, 0, 1, 2, 1, 5, 4, 2, 1, 12,
	10, 11, 9, 2, 1, 3, 3, 1,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, 5, -3, 10, 12, -5, 13,
	-4, -7, 6, -6, 8, 13, -4, 6, -6, 10,
	-8, 10, 10, 17, 14, 11, 16, 17, 14, 10,
	10, 10, 10, 10, 14, 15, 14, 15, 10, 7,
	10, 7, 15, 14, 15, 14, 7, 10, 7, 10,
	14, 15, 14, 15, 10, 10, 15, 15,
}

var yyDef = [...]int8{
	1, -2, 2, 4, 0, 3, 0, 0, 0, 6,
	8, 0, 0, 14, 0, 5, 7, 0, 13, 0,
	0, 17, 0, 0, 0, 15, 0, 0, 0, 0,
	0, 16, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 12, 0, 10, 0, 0, 11, 9,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 15, 3, 3, 3, 3, 17, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 12, 16, 13,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:50
		{
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:51
		{
			if l, ok := yylex.(*Scanner); ok {
				stmt := &ast.Stmt{
					Services: yyDollar[1].services,
				}
				l.Stmts = append(l.Stmts, stmt)
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:60
		{
			yyVAL.services = append(yyDollar[1].services, yyDollar[2].service)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:62
		{
			yyVAL.services = append(yyVAL.services, yyDollar[1].service)
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:66
		{
			yyVAL.service = &ast.Service{
				Name: yyDollar[2].token,
				Rpcs: yyDollar[4].rpcs,
			}
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:71
		{
			yyVAL.service = &ast.Service{
				Name: yyDollar[2].token,
			}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:77
		{
			yyVAL.rpcs = append(yyDollar[1].rpcs, yyDollar[2].rpc)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:79
		{
			yyVAL.rpcs = append(yyVAL.rpcs, yyDollar[1].rpc)
		}
	case 9:
		yyDollar = yyS[yypt-12 : yypt+1]
//line parser.go.y:83
		{
			yyVAL.rpc = &ast.RPC{
				Routes:   yyDollar[1].routes,
				Handler:  yyDollar[3].token + "." + yyDollar[5].token,
				Request:  yyDollar[7].token,
				Response: yyDollar[11].token,
			}
		}
	case 10:
		yyDollar = yyS[yypt-10 : yypt+1]
//line parser.go.y:90
		{
			yyVAL.rpc = &ast.RPC{
				Routes:   yyDollar[1].routes,
				Handler:  yyDollar[3].token,
				Request:  yyDollar[5].token,
				Response: yyDollar[9].token,
			}
		}
	case 11:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:97
		{
			yyVAL.rpc = &ast.RPC{
				Routes:   []*ast.Route{},
				Handler:  yyDollar[2].token + "." + yyDollar[4].token,
				Request:  yyDollar[6].token,
				Response: yyDollar[10].token,
			}
		}
	case 12:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:104
		{
			yyVAL.rpc = &ast.RPC{
				Routes:   []*ast.Route{},
				Handler:  yyDollar[2].token,
				Request:  yyDollar[4].token,
				Response: yyDollar[8].token,
			}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:113
		{
			yyVAL.routes = append(yyDollar[1].routes, yyDollar[2].route)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:115
		{
			yyVAL.routes = append(yyVAL.routes, yyDollar[1].route)
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:119
		{
			yyVAL.route = &ast.Route{
				Methods: yyDollar[2].route_methods,
				URI:     yyDollar[3].token,
			}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:126
		{
			yyVAL.route_methods = append(yyDollar[1].route_methods, yyDollar[3].token)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:128
		{
			yyVAL.route_methods = append(yyVAL.route_methods, yyDollar[1].token)
		}
	}
	goto yystack /* stack new state and value */
}