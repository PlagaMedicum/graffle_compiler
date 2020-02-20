grammar graffle;

//
// LEXER
//

// types
INT     : '-'? DIGIT+ ;
NUMBER  : INT ('.'DIGIT+)? ;
STRING  : '"' ~[\\"]* '"' ;
TRUE    : 'true' | '1' ;
FALSE   : 'false' | '0' ;
BOOL    : TRUE | FALSE ;

// operators
OR_ARC_LR   : '->' ;
OR_ARC_RL   : '<-' ;
OR_W_ARC_LR : '-[' NUMBER ']' OR_ARC_LR ;
OR_W_ARC_RL : OR_ARC_RL '[' NUMBER ']-' ;
UNOR_ARC    : '--' ;
UNOR_W_ARC  : '-[' NUMBER ']-' ;

ASSIGN      : '=' ;
ADD_ASSIGN  : '+=' ;
SUB_ASSIGN  : '-=' ;

ADD         : '+' ;
SUB         : '-' ;
MULT        : '*' ;
DIV         : '/' ;

EQ          : '==' ;
NEQ         : '!=' ;
LESS_THAN   : '<' ;
GR_THAN     : '>' ;

NOT         : '!';

PRINTER     : '<<' ;
// COMMA : ',' ;

// type names
G_N : 'G' ; // graph
V_N : 'V' ; // vertice
E_N : 'E' ; // arc

// other
ID : [a-zA-Z_][a-zA-Z_0-9]* ;

// ignore
WS           : [ \t\n\r]+       -> skip ;
LINE_COMMENT : '#' ~[\r\n]*     -> skip ;

fragment DIGIT : [0-9] ;

//
// PARSER
//

function_param
    : STRING
    | NUMBER
    | ID
    | function_call
    | '(' logical_expr ')'
    ;

function_call
    : ID '(' (function_param (',' function_param)*)? ')'
    ;

function_declaration
    : ID '(' (ID (',' ID)*)? ')'
    ;

logical_operator
    : EQ
    | NEQ
    | LESS_THAN
    | GR_THAN
    ;

logical_expr
    : function_param logical_operator function_param
    | NOT function_param
    ;

typeAssignement
    : ID ASSIGN (G_N | V_N | E_N) '(' ID ')' // Example: graph = G(vertice)
    ;

nw_arc_declaration
    : ID ASSIGN ID (OR_ARC_LR | OR_ARC_RL | UNOR_ARC) ID // Example: arc = v1 -> v2
    ;
w_arc_declaration
    : ID ASSIGN ID (OR_W_ARC_LR | OR_W_ARC_RL | UNOR_W_ARC) ID // Example: w_arc = v1 <-[30]- v2
    ;

print
    : PRINTER ID
    | PRINTER NUMBER
    | PRINTER STRING
    ;
