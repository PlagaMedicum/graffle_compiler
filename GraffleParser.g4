parser grammar GraffleParser;

options {
    tokenVocab=GraffleLexer;
}

file
    : NEWLINE? function_declaration* main_body=sequence EOF?
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
    : atom_action
    | one_line_stmnt
    ;

atom_action
    : var_assign
    | function_call
    | builtin_function_call
    ;

// Statements:
if_stmnt
    : IF cond=logical_expr ( ','? (THEN | DO | THEN DO) )? sequence+ NEWLINE? (ELSE DO? sequence+)? block_end
    ;

if_is_stmnt
    : IF value case_stmnt* NEWLINE? (DEFAULT DO? sequence+)? block_end
    ;
case_stmnt
    : NEWLINE? IS value ( ','? (THEN | DO | THEN DO) )? sequence+
    ;

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

while_stmnt
    : WHILE cond=logical_expr DO?
    ;
until_stmnt
    : UNTIL cond=logical_expr DO?
    ;
for_stmnt
    : FOR cond=logical_expr DO?                                                              #ForLogical
    | FOR pre_act=atom_action ARG_DELIM cond=logical_expr ARG_DELIM post_act=atom_action DO? #ForVar
    | FOR VAR IN RANGE? FROM? from=expr TO to=expr DO?                                       #ForRange
    ;
from_to_stmnt
    : FROM from=expr TO to=expr DO?
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
    : ID '(' (opd1=VAR (ARG_DELIM opd2=VAR)*)? ')' ASSIGN return_val=value (','? WHERE)?
    ;
one_line_procedure_declaration
    : procedure_declaration_head sequence_line
    ;
mult_line_procedure_declaration
    : procedure_declaration_head NEWLINE sequence block_end
    ;
procedure_declaration_head
    : ID '(' (opd1=VAR (ARG_DELIM opd2=VAR)*)? ')' ASSIGN?
    ;

// vars
var_assign
    : variable=ID ASSIGN val=VAR
    | variable=ID ASSIGN expr
    | arc_assign
    | vertice_assign
    | graph_assign
    | labeled_assign
    ;

arc_assign
    : variable=ID ASSIGN E_N '(' value ')'
    | variable=ID ASSIGN value arc value
    ;
arc
    : OR_ARC_LR
    | or_w_arc_lr
    | OR_ARC_RL
    | or_w_arc_rl
    | UNOR_ARC
    | unor_w_arc
    ;
or_w_arc_lr
    : '-' '[' weight=NUMBER ']' OR_ARC_LR
    ;
or_w_arc_rl
    : OR_ARC_RL '[' weight=NUMBER ']' '-'
    ;
unor_w_arc
    : '-' '[' weight=NUMBER ']' '-'
    ;

vertice_assign
    : ID ASSIGN '(' V_N ')' value
    | ID ASSIGN value
    | ID arithm_assign_operator value
    ;

graph_assign
    : ID ASSIGN G_N '(' value ')'
    | ID ASSIGN value (ARG_DELIM value)*
    | ID arithm_assign_operator value
    ;

labeled_assign
    : vertice_assign label?
    | arc_assign label?
    | graph_assign label?
    ;

// Expressions:
expr
    : VAR
    | integral_expr
    ;

integral_expr
    : function_call
    | builtin
    | '(' logical_expr ')'
    | '(' arithm_expr ')'
    ;

logical_expr
    : left=expr bin_op=bin_log_operator right=log_expr_operand
    | un_op=unar_log_operator log_expr_operand
    | arithm_expr
    | expr
    ;
log_expr_operand
    : expr
    | logical_expr
    ;
bin_log_operator
    : EQUALS
    | NEQ
    | LESS_THAN
    | GR_THAN
    | LESS_THAN_E
    | GR_THAN_E
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
    : left=expr op=bin_arithm_operator right=arithm_expr_operand
    ;
arithm_expr_operand
    : expr
    | arithm_expr
    ;
bin_arithm_operator
    : MULT
    | DIV
    | ADD
    | SUB
    ;
arithm_assign_operator
    : ADD_ASSIGN
    | SUB_ASSIGN
    | MULT_ASSIGN
    | DIV_ASSIGN
    ;

// Function calls:
// built-in functions
builtin_function_call
    : built_func_print
    | built_func_input
    ;

built_func_print
    : PRINTER VAR
    | PRINTER value
    | PRINTER function_call
    ;
built_func_input
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
    : logical_expr
    | arithm_expr
    | expr
    | builtin
    ;

builtin
    : NUMBER
    | STRING
    | BOOL
    ;

block_end
    : NEWLINE? BLOCK_END NEWLINE?
    ;
