parser grammar GraffleParser;

options {
    tokenVocab=GraffleLexer;
}

file
    : (NEWLINE? function_declaration)* sequence EOF?
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
    : IF logical_expr ( ','? (THEN | DO | THEN DO) )? sequence else_stmnt? block_end
    ;
else_stmnt
    : NEWLINE? ELSE DO? sequence
    ;

if_is_stmnt
    : IF value case_stmnt* default_stmnt? block_end
    ;
case_stmnt
    : NEWLINE? IS value ( ','? (THEN | DO | THEN DO) )? sequence
    ;
default_stmnt
    : NEWLINE? DEFAULT DO? sequence
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
    | FOR variable IN RANGE? FROM? from=expr TO to=expr DO?                                  #ForRange
    ;
from_to_stmnt
    : FROM from=expr TO to=expr DO?
    ;

// Declarations:
// functions
function_declaration
    : one_line_function_declaration
    | mult_line_function_declaration
    | one_line_procedure_declaration
    | mult_line_procedure_declaration
    ;
one_line_function_declaration
    : function_declaration_head sequence_line
    ;
mult_line_function_declaration
    : function_declaration_head sequence block_end
    ;
function_declaration_head
    : ID '(' (variable (ARG_DELIM variable)*)? ')' ASSIGN return_val=value (','? WHERE)?
    ;
one_line_procedure_declaration
    : procedure_declaration_head sequence_line
    ;
mult_line_procedure_declaration
    : procedure_declaration_head NEWLINE sequence block_end
    ;
procedure_declaration_head
    : ID '(' (variable (ARG_DELIM variable)*)? ')' ASSIGN?
    ;

// vars
var_assign
    : vertice_assign
    | edge_assign
    | graph_assign
    | ID ASSIGN variable
    | ID ASSIGN value
    | ID arithm_assign_operator value
    | labeled_assign
    ;

vertice_assign
    : ID ASSIGN vertice_type
    ;
vertice_type
    : V_N '(' value ')'
    ;

edge_assign
    : ID ASSIGN edge_type
    ;
edge_type
    : value edge value
    | E_N '(' value ')'
    ;
edge
    : OR_EDGE_LR
    | or_w_edge_lr
    | OR_EDGE_RL
    | or_w_edge_rl
    | UNOR_EDGE
    | unor_w_edge
    ;
or_w_edge_lr
    : '-' '[' weight=NUMBER ']' OR_EDGE_LR
    ;
or_w_edge_rl
    : OR_EDGE_RL '[' weight=NUMBER ']' '-'
    ;
unor_w_edge
    : '-' '[' weight=NUMBER ']' '-'
    ;

graph_assign
    : ID ASSIGN graph_type
    | ID arithm_assign_operator value (ARG_DELIM value)+
    ;
graph_type
    : G_N '(' value (ARG_DELIM value)*')'
    | value (ARG_DELIM value)+
    ;

labeled_assign
    : edge_assign label?
    | graph_assign label?
    ;

// Expressions:
expr
    : variable
    | integral_expr
    | builtin_type
    | vertice_type
    | '(' edge_type ')'
    | '(' graph_type ')'
    ;

integral_expr
    : function_call
    | builtin_type
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
    : PRINTER variable
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
    ;
variable
    : '-'? ID
    ;

builtin_type
    : NUMBER
    | STRING
    | BOOL
    ;

block_end
    : NEWLINE? BLOCK_END ACT_DELIM* NEWLINE?
    ;
