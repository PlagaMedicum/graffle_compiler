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

STRING      : '"'  ~[\\"]* '"'
            | '\'' ~[\\']* '\''
            ;
LABEL       : '@' ~[\r\n]* ;
ML_LABEL    : '@[' ~[\\[]* ']' ;

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
NEQ         : NOT '='
            | NOT EQUALS
            ;
EQUALS      : '=='
            | [Ee]'quals' To?
            ;
LESS_THAN   : '<'
            | [Ll]'ess' Than?
            ;
GR_THAN     : '>'
            | [Gg]'reater' Than?
            ;
LESS_THAN_E : '<='
            | [Ll]'ess' Than? OR EQUALS
            | NOT GR_THAN
            ;
GR_THAN_E   : '>='
            | [Gg]'reater' Than? OR EQUALS
            | NOT LESS_THAN
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
FROM    : [Ff]'rom'
        ;
TO      : [Tt]'o'
        ;
DO      : [Dd]'o'
        ;

ACT_DELIM : '.' ; // operations delimiter. Example: v1 = 1. v2 = 2
ARG_DELIM : ',' ;

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

// Other:
ID : [a-zA-Z_][a-zA-Z_0-9]* ;
// INDENT : ;

// Ignore:
WS      : ~[\n\r][ \t]+ -> skip ;
NEWLINE : [\n\r]+       -> skip ;

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

action_line
    : action (ACT_DELIM action)*
    ;

//sub_block
//    : ([ \t]+ action_line [\n\r])+
//    ;

function_param
    : STRING
    | NUMBER
    | ID
    | BOOL
    | function_call
    | '(' logical_expr ')'
    ;

integral_param
    : NUMBER
    | ID
    | function_call
    | BOOL
    | '(' logical_expr ')'
    ;

function_call
    : ID '(' (function_param (ARG_DELIM function_param)*)? ')'
    ;

bin_log_operator
    : EQUALS
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
    : arc_declaration
    | vertice_declaration
    | graph_declaration
    | ID ASSIGN ID
    ;

// Declarations:
function_declaration
    : ID '(' (ID (ARG_DELIM ID)*)? ')' ASSIGN
    ;

arc_declaration
    : ID ASSIGN E_N '(' ID ')'
    | ID ASSIGN ID OR_ARC_LR ID
    | ID ASSIGN ID OR_ARC_RL ID
    | ID ASSIGN ID UNOR_ARC ID
    | ID ASSIGN ID OR_W_ARC_LR ID
    | ID ASSIGN ID OR_W_ARC_RL ID
    | ID ASSIGN ID UNOR_W_ARC ID
    ;

vertice_declaration
    : ID ASSIGN V_N '(' ID 'ID'
    | ID ASSIGN (ID | vertice_value) // Example: v = "OSTIS" @ Best semantic technology ever
    ;

graph_declaration
    : ID ASSIGN G_N '(' ID ')'
    | ID ASSIGN
    ;

label
    : LABEL
    | ML_LABEL
    ;
labeled_expression
    : vertice_declaration label?
    | arc_declaration label?
    | graph_declaration label?
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

while_stmnt
    : WHILE logical_expr DO?
    ;
until_stmnt
    : UNTIL logical_expr DO?
    ;
for_stmnt
    : FOR logical_expr DO?
    | FOR action ARG_DELIM logical_expr ARG_DELIM action DO?
    ;
from_to_stmnt
    : FROM integral_param TO integral_param DO?
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
