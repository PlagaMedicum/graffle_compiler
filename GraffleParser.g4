parser grammar GraffleParser;

options {
    tokenVocab=GraffleLexer;
    superClass=GraffleParserBase;
}

file
    : function_declaration+ sequence EOF?
    ;

sequence
    : NEWLINE? (sequence_element ((ACT_DELIM | NEWLINE)+ | EOF))+
    ;
sequence_element
    : if_stmnt
    | if_is_stmnt
    | mult_line_stmnt
    | sequence_line
    ;
sequence_line
    : one_line_sequence_element (ACT_DELIM one_line_sequence_element)* ACT_DELIM?
    ;
one_line_sequence_element
    : action
    | one_line_stmnt
    ;

action
    : var_declaration
    | function_call
    | builtin_function_call
    ;

// Statements:
one_line_stmnt
    : stmnt sequence_line
    ;
mult_line_stmnt
    : stmnt sequence block_end
    ;

stmnt
    : while_stmnt
    | until_stmnt
    | for_stmnt
    | from_to_stmnt
    ;

if_stmnt
    : IF cond=logical_expr ( ','? (THEN | DO | THEN DO) )? sequence+ NEWLINE? (ELSE DO? sequence+)? block_end
    ;

if_is_stmnt
    : IF value case_stmnt* NEWLINE? (DEFAULT DO? sequence+)? block_end
    ;
case_stmnt
    : NEWLINE? IS value ( ','? (THEN | DO | THEN DO) )? sequence+
    ;

while_stmnt
    : WHILE cond=logical_expr DO?
    ;
until_stmnt
    : UNTIL cond=logical_expr DO?
    ;
for_stmnt
    : FOR cond=logical_expr DO?                                  #ForLogical
    | FOR action ARG_DELIM logical_expr ARG_DELIM action DO?     #ForVar
    | FOR ID IN RANGE? FROM? integral_expr TO integral_expr DO?  #ForRange
    ;
from_to_stmnt
    : FROM integral_expr TO integral_expr DO?
    ;

// Declarations:
// functions
function_declaration
    : NEWLINE? one_line_function_declaration
    | NEWLINE? mult_line_function_declaration
    | NEWLINE? one_line_procedure_declaration
    | NEWLINE? mult_line_procedure_declaration
    ;
one_line_function_declaration
    : function_declaration_head sequence_line
    ;
mult_line_function_declaration
    : function_declaration_head sequence block_end
    ;
function_declaration_head
    : ID '(' (opd1=ID (ARG_DELIM opd2=ID)*)? ')' ASSIGN value (','? WHERE)?
    ;
one_line_procedure_declaration
    : procedure_declaration_head NEWLINE sequence block_end
    ;
mult_line_procedure_declaration
    : procedure_declaration_head sequence_line
    ;
procedure_declaration_head
    : ID '(' (opd1=ID (ARG_DELIM opd2=ID)*)? ')' ASSIGN?
    ;

// vars
var_declaration
    : var=ID ASSIGN val=ID
    | var=ID ASSIGN expr
    | arc_declaration
    | vertice_declaration
    | graph_declaration
    | labeled_declaration
    ;

arc_declaration
    : var=ID ASSIGN E_N '(' value ')'
    | var=ID ASSIGN value arc value
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

labeled_declaration
    : vertice_declaration label?
    | arc_declaration label?
    | graph_declaration label?
    ;

// Expressions:
expr
    : ID
    | integral_expr
    ;

integral_expr
    : NUMBER
    | function_call
    | BOOL
    | '(' logical_expr ')'
    | '(' arithm_expr ')'
    ;

logical_expr
    : left=expr bin_op=bin_log_operator right=expr
    | un_op=unar_log_operator expr
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
unar_log_operator
    : NOT
    ;

arithm_expr
    : expr bin_arithm_operator expr
    | expr unar_arithm_operator
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

// Function calls:
// built-in functions
builtin_function_call
    : print
    | input
    ;

print
    : PRINTER ID
    | PRINTER NUMBER
    | PRINTER STRING
    | PRINTER function_call
    ;
input
    : KEY_INPUT ID
    ;

function_call
    : ID '(' (value (ARG_DELIM value)*)? ')'
    ;

// Other:
label
    : LABEL
    | ML_LABEL
    ;
value
    : expr
    | STRING
    ;

block_end
    : NEWLINE? BLOCK_END ACT_DELIM?
    ;
