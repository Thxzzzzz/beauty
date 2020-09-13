%{

package parser

import (
	"github.com/rushteam/beauty/tools/parser/ast"
)

const EOF =0

//@auth jwt
//@rest url
//@doc (key:val)
%}
// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	// empty struct{}
	token string
	str string
	// at_list []ast.At
	// at ast.At
	service ast.Service
	rpc_list []ast.RPC
	rpc ast.RPC
	route ast.Route
	route_list []ast.Route
	methods []string
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
// %type <val> expr number
%type <token> program
// %type <at> at
// %type <at_list> at_list
%type <service> service
%type <rpc> rpc
%type <rpc_list> rpc_list
%type <route> route
%type <route_list> route_list
%type <methods> methods


%start program
// same for terminals
// SERVICE RPC
%token ILLEGAL
%token Service Rpc Returns Route
%token <token> Comment Ident String
%token '(' ')' '|'

%%
program: {}
// | comment at_list comment service comment {}
// | comment at_list comment service {}
// | comment at_list service {}
// | at_list service {}
| service {
	if l, ok := yylex.(*Scanner); ok {
		stmt := ast.Stmt{}
		stmt.Services = append(stmt.Services,$1)
		l.Stmts = append(l.Stmts,stmt)
	}
}
;
// at_list: at_list at {
// 	$$ = append($1,$2)
// }
// | at {
// 	$$ = append($$,$1)
// };
// at: '@' Val Val {
// 	val := make(map[string]string,0)
// 	val["_"] = $3
// 	$$ = ast.At{
// 		Opt: $2,
// 		Val: val,
// 	}
// };

service: Service Ident '(' rpc_list ')' {
	$$ = ast.Service{
		Name: $2,
		Rpcs: $4,
	}
};
rpc_list: rpc_list rpc {
	$$ = append($1,$2)
} | rpc {
	$$ = append($$,$1)
};
rpc: route_list Rpc Ident '(' Ident ')' Returns '(' Ident ')' {
	$$ = ast.RPC{
		Routes: $1,
		Handler: $3,
		Request: $5,
		Response: $9,
	}
} | Rpc Ident '(' Ident ')' Returns '(' Ident ')' {
	$$ = ast.RPC{
		Routes: []ast.Route{},
		Handler: $2,
		Request: $4,
		Response: $8,
	}
};
route_list: route_list route {
	$$ = append($1,$2)
} | route {
	$$ = append($$,$1)
};
route: Route methods String {
	$$ = ast.Route{
		Methods: $2,
		URI: $3,
	}
};
methods: methods '|' Ident {
	$$ = append($1,$3)
} | Ident {
	$$ = append($$,$1)
};
%%
/*  start  of  programs  */
/*
service helloworld (
	@route GET|POST "/index/:id"
    rpc Index(getRequest) returns (getResponse)
    rpc Helloworld(getRequest) returns (getResponse)
)
*/