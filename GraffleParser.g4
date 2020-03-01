parser grammar GraffleParser;

options {
    tokenVocab=GraffleLexer;
    superClass=GraffleParserBase;
}

sequence
    : (sequence_element (ACT_DELIM sequence_element)* ACT_DELIM? (NEWLINE+ | EOF))*
    ;
sequence_element
    : action
    | one_line_stmnt
    ;

action
    : var_declaration
    | function_call
    | builtin_function_call
    ;

// Declarations:
var_declaration
    : ID ASSIGN ID
    | arc_declaration
    | vertice_declaration
    | graph_declaration
    | labeled_declaration
    ;

one_line_function_declaration
    : function_declaration ','? WHERE sequence
    ;
mult_line_function_declaration
    : function_declaration (','? WHERE)? NEWLINE
    ;
function_declaration
    : ID '(' (ID (ARG_DELIM ID)*)? ')' ASSIGN value
    ;
procedure_declaration
    : ID '(' (ID (ARG_DELIM ID)*)? ')' NEWLINE
    ;

arc_declaration
    : ID ASSIGN E_N '(' value ')'
    | ID ASSIGN value arc value
    ;
arc
    : OR_ARC_LR
    | OR_W_ARC_LR
    | OR_ARC_RL
    | OR_W_ARC_RL
    | UNOR_ARC
    | UNOR_W_ARC
    ;

vertice_declaration
    : ID ASSIGN V_N '(' value ')'
    | ID ASSIGN value
    | ID unar_arithm_operator value
    ;

graph_declaration
    : ID ASSIGN G_N '(' value ')'
    | ID ASSIGN
    ;

label
    : LABEL
    | ML_LABEL
    ;
labeled_declaration
    : vertice_declaration label?
    | arc_declaration label?
    | graph_declaration label?
    ;

value
    : expr
    | STRING
    ;

// Expressions:
expr
    : ID
    | integral_expr
    ;

logical_expr
    : expr bin_log_operator expr
    | un_log_operator expr
    | expr
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

arithm_expr
    : expr bin_arithm_operator expr
    |
    ;
bin_arithm_operator
    : MULT
    | DIV
    | ADD
    | SUB
    ;
unar_arithm_operator
    : ADD_ASSIGN
    | SUB_ASSIGN
    | MULT_ASSIGN
    | DIV_ASSIGN
    ;

integral_expr
    : NUMBER
    | function_call
    | BOOL
    | '(' logical_expr ')'
    | '(' arithm_expr ')'
    ;

// Function calls:
function_call
    : ID '(' (value (ARG_DELIM value)*)? ')'
    ;

// built-in functions
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

// Statements:
one_line_stmnt
    : stmnt sequence
    ;
mult_line_stmnt
    : stmnt NEWLINE
    ;

stmnt
    : if_stmnt
    | if_is_stmnt
    | case_stmnt
    | else_stmnt
    | while_stmnt
    | until_stmnt
    ;

if_stmnt
    : IF logical_expr ( ','? (THEN | DO | THEN DO) )?
    ;
else_stmnt
    : ELSE DO?
    ;

if_is_stmnt
    : IF value
    ;
case_stmnt
    : IS value ( ','? (THEN | DO | THEN DO) )?  #Case
    | DEFAULT DO?                               #Default
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
    | FOR ID IN RANGE? FROM? integral_expr TO integral_expr DO?
    ;
from_to_stmnt
    : FROM integral_expr TO integral_expr DO?
    ;
