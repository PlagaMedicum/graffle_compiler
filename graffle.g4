grammar graffle;

//
// LEXER
//

// Types:
NUMBER  : INT
        | FLOAT
        ;
INT     : '-'? Digit+ ;
FLOAT   : INT ('.' | ',') Digit+ ;

STRING  : '"'  ~[\\"]* '"'
        | '\'' ~[\\']* '\''
        ;
LABEL   : ':' ~[\\]* ;

BOOL    : True
        | False
        ;

// Operators:
// arcs
OR_ARC_LR   : '->' ;
OR_ARC_RL   : '<-' ;
OR_W_ARC_LR : '-[' NUMBER ']' OR_ARC_LR ;
OR_W_ARC_RL : OR_ARC_RL '[' NUMBER ']-' ;
UNOR_ARC    : '--' ;
UNOR_W_ARC  : '-[' NUMBER ']-' ;

// asiignements
ASSIGN      : '=' ;
ADD_ASSIGN  : '+=' ;
SUB_ASSIGN  : '-=' ;

// arithmetical operators
ADD         : '+' ;
SUB         : '-' ;
MULT        : '*' ;
DIV         : '/' ;

// logical
// - binary
NEQ         : '!='
            | [Nn]'ot' EQ
            ;
EQ          : '=='
            | [Ee]'quals' To?
            ;
LESS_THAN   : '<'
            | [Ll]'ess' Than?
            ;
GR_THAN     : '>'
            | [Gg]'reater' Than?
            ;
AND         : '&'
            | [Aa]'nd'
            | 'AND'
            ;
OR          : '|'
            | [Oo]'r'
            | 'OR'
            ;
XOR         : '^'
            | 'XOR'
            |[Xx]'or'
            ;
NOR         : NOT OR
            | [Nn]'or'
            | 'NOR'
            ;
NAND        : NOT AND
            | [Nn]'and'
            | 'NAND'
            ;
// - unary
NOT         : '!'
            | [Nn]'ot'
            | 'NOT'
            ;

// conditionals
IF      : [Ii]'f' ;
THEN    : [Tt]'hen' ;
ELSE    : [Ee]'lse' ;

// cycles
WHILE   : [Ww]'hile'
        ;
UNTIL   : [Uu]'ntil'
        ;
FOR     : [Ff]'or'
        ;
DO      : [Dd]'o'
        ;

DELIM : ',' ; // operations delimiter. Example: v1 = 1, v2 = 2

// Built-in functions:
PRINTER     : '<<<'
            | [Pp]'rint'
            ;
KEY_INPUT   : '>>>'
            | [Ii]'nput'
            ;

// Type names:
G_N : 'G'
    | 'Graph'
    ;
V_N : 'V'
    | 'Vertice'
    ;
E_N : 'E'
    | 'Arc'
    ;

ID : [a-zA-Z_][a-zA-Z_0-9]* ;

// Ignore:
WS      : [ \t]+    -> skip ;
NEWLINE : [\n\r]+   -> skip ;

LINE_COMMENT   : CommentSym{3} ~[\r\n]*      -> skip ;
M_LINE_COMMENT : CommentSym ~'`'* CommentSym -> skip ;

fragment CommentSym : '`' ;
fragment Digit      : [0-9] ;
fragment True       : [Tt]'rue' | '1' ;
fragment False      : [Ff]'alse' | '0' ;
fragment To         : [Tt]'o' ;
fragment Than       : [Tt]'han' ;

//
// PARSER
//

action
    : var_declaration
    | function_call
    | builtin_function_call
    ;

function_param
    : STRING
    | NUMBER
    | ID
    | BOOL
    | function_call
    | '(' logical_expr ')'
    ;

function_call
    : ID '(' (function_param (',' function_param)*)? ')'
    ;

bin_log_operator
    : EQ
    | NEQ
    | LESS_THAN
    | GR_THAN
    | AND
    | NAND
    | OR
    | NOR
    | XOR
    ;
un_log_operator
    : NOT
    ;

logical_expr
    : function_param bin_log_operator function_param
    | un_log_operator function_param
    | function_param
    ;

var_declaration
    : type_assignement
    | nw_arc_declaration
    | w_arc_declaration
    | vertice_declaration
    ;

type_assignement
    : ID ASSIGN (G_N | V_N | E_N) '(' ID ')' // Example: graph = G(vertice)
    ;

// Declarations:
function_declaration
    : ID '(' (ID (',' ID)*)? ')' ASSIGN
    ;

nw_arc_declaration
    : ID ASSIGN ID (OR_ARC_LR | OR_ARC_RL | UNOR_ARC) ID // Example: arc = v1 -> v2
    ;
w_arc_declaration
    : ID ASSIGN ID (OR_W_ARC_LR | OR_W_ARC_RL | UNOR_W_ARC) ID // Example: w_arc = v1 <-[30]- v2
    ;

vertice_declaration
    : ID ASSIGN (ID | vertice_value) LABEL? // Example: v = "OSTIS" {Best semantic technology}
    ;

vertice_value
    : NUMBER
    | STRING
    | BOOL
    ;

// Statements:
if_stmnt
    : IF logical_expr ','? THEN? DO?
    ;
else_stmnt
    : ELSE DO?
    ;

for_stmnt
    : FOR logical_expr DO?
    | FOR action DELIM logical_expr DELIM action DO?
    ;
while_stmnt
    : WHILE logical_expr DO?
    ;
until_stmnt
    : UNTIL logical_expr DO?
    ;

// Built-in functions
builtin_function_call
    : print
    | input
    ;

print
    : PRINTER ID
    | PRINTER NUMBER
    | PRINTER STRING
    ;
input
    : KEY_INPUT ID
    ;
