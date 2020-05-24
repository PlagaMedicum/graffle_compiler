// Code generated from /home/plagamedicum/Documents/yapis/graffle/GraffleParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // GraffleParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGraffleParserListener is a complete listener for a parse tree produced by GraffleParser.
type BaseGraffleParserListener struct{}

var _ GraffleParserListener = &BaseGraffleParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGraffleParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGraffleParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGraffleParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGraffleParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseGraffleParserListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseGraffleParserListener) ExitFile(ctx *FileContext) {}

// EnterSequence is called when production sequence is entered.
func (s *BaseGraffleParserListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production sequence is exited.
func (s *BaseGraffleParserListener) ExitSequence(ctx *SequenceContext) {}

// EnterSequence_element is called when production sequence_element is entered.
func (s *BaseGraffleParserListener) EnterSequence_element(ctx *Sequence_elementContext) {}

// ExitSequence_element is called when production sequence_element is exited.
func (s *BaseGraffleParserListener) ExitSequence_element(ctx *Sequence_elementContext) {}

// EnterSequence_line is called when production sequence_line is entered.
func (s *BaseGraffleParserListener) EnterSequence_line(ctx *Sequence_lineContext) {}

// ExitSequence_line is called when production sequence_line is exited.
func (s *BaseGraffleParserListener) ExitSequence_line(ctx *Sequence_lineContext) {}

// EnterOne_line_sequence_element is called when production one_line_sequence_element is entered.
func (s *BaseGraffleParserListener) EnterOne_line_sequence_element(ctx *One_line_sequence_elementContext) {
}

// ExitOne_line_sequence_element is called when production one_line_sequence_element is exited.
func (s *BaseGraffleParserListener) ExitOne_line_sequence_element(ctx *One_line_sequence_elementContext) {
}

// EnterAtom_action is called when production atom_action is entered.
func (s *BaseGraffleParserListener) EnterAtom_action(ctx *Atom_actionContext) {}

// ExitAtom_action is called when production atom_action is exited.
func (s *BaseGraffleParserListener) ExitAtom_action(ctx *Atom_actionContext) {}

// EnterIf_stmnt is called when production if_stmnt is entered.
func (s *BaseGraffleParserListener) EnterIf_stmnt(ctx *If_stmntContext) {}

// ExitIf_stmnt is called when production if_stmnt is exited.
func (s *BaseGraffleParserListener) ExitIf_stmnt(ctx *If_stmntContext) {}

// EnterElse_stmnt is called when production else_stmnt is entered.
func (s *BaseGraffleParserListener) EnterElse_stmnt(ctx *Else_stmntContext) {}

// ExitElse_stmnt is called when production else_stmnt is exited.
func (s *BaseGraffleParserListener) ExitElse_stmnt(ctx *Else_stmntContext) {}

// EnterIf_is_stmnt is called when production if_is_stmnt is entered.
func (s *BaseGraffleParserListener) EnterIf_is_stmnt(ctx *If_is_stmntContext) {}

// ExitIf_is_stmnt is called when production if_is_stmnt is exited.
func (s *BaseGraffleParserListener) ExitIf_is_stmnt(ctx *If_is_stmntContext) {}

// EnterCase_stmnt is called when production case_stmnt is entered.
func (s *BaseGraffleParserListener) EnterCase_stmnt(ctx *Case_stmntContext) {}

// ExitCase_stmnt is called when production case_stmnt is exited.
func (s *BaseGraffleParserListener) ExitCase_stmnt(ctx *Case_stmntContext) {}

// EnterDefault_stmnt is called when production default_stmnt is entered.
func (s *BaseGraffleParserListener) EnterDefault_stmnt(ctx *Default_stmntContext) {}

// ExitDefault_stmnt is called when production default_stmnt is exited.
func (s *BaseGraffleParserListener) ExitDefault_stmnt(ctx *Default_stmntContext) {}

// EnterOne_line_stmnt is called when production one_line_stmnt is entered.
func (s *BaseGraffleParserListener) EnterOne_line_stmnt(ctx *One_line_stmntContext) {}

// ExitOne_line_stmnt is called when production one_line_stmnt is exited.
func (s *BaseGraffleParserListener) ExitOne_line_stmnt(ctx *One_line_stmntContext) {}

// EnterMult_line_stmnt is called when production mult_line_stmnt is entered.
func (s *BaseGraffleParserListener) EnterMult_line_stmnt(ctx *Mult_line_stmntContext) {}

// ExitMult_line_stmnt is called when production mult_line_stmnt is exited.
func (s *BaseGraffleParserListener) ExitMult_line_stmnt(ctx *Mult_line_stmntContext) {}

// EnterStmnt is called when production stmnt is entered.
func (s *BaseGraffleParserListener) EnterStmnt(ctx *StmntContext) {}

// ExitStmnt is called when production stmnt is exited.
func (s *BaseGraffleParserListener) ExitStmnt(ctx *StmntContext) {}

// EnterWhile_stmnt is called when production while_stmnt is entered.
func (s *BaseGraffleParserListener) EnterWhile_stmnt(ctx *While_stmntContext) {}

// ExitWhile_stmnt is called when production while_stmnt is exited.
func (s *BaseGraffleParserListener) ExitWhile_stmnt(ctx *While_stmntContext) {}

// EnterUntil_stmnt is called when production until_stmnt is entered.
func (s *BaseGraffleParserListener) EnterUntil_stmnt(ctx *Until_stmntContext) {}

// ExitUntil_stmnt is called when production until_stmnt is exited.
func (s *BaseGraffleParserListener) ExitUntil_stmnt(ctx *Until_stmntContext) {}

// EnterForLogical is called when production ForLogical is entered.
func (s *BaseGraffleParserListener) EnterForLogical(ctx *ForLogicalContext) {}

// ExitForLogical is called when production ForLogical is exited.
func (s *BaseGraffleParserListener) ExitForLogical(ctx *ForLogicalContext) {}

// EnterForVar is called when production ForVar is entered.
func (s *BaseGraffleParserListener) EnterForVar(ctx *ForVarContext) {}

// ExitForVar is called when production ForVar is exited.
func (s *BaseGraffleParserListener) ExitForVar(ctx *ForVarContext) {}

// EnterForRange is called when production ForRange is entered.
func (s *BaseGraffleParserListener) EnterForRange(ctx *ForRangeContext) {}

// ExitForRange is called when production ForRange is exited.
func (s *BaseGraffleParserListener) ExitForRange(ctx *ForRangeContext) {}

// EnterFrom_to_stmnt is called when production from_to_stmnt is entered.
func (s *BaseGraffleParserListener) EnterFrom_to_stmnt(ctx *From_to_stmntContext) {}

// ExitFrom_to_stmnt is called when production from_to_stmnt is exited.
func (s *BaseGraffleParserListener) ExitFrom_to_stmnt(ctx *From_to_stmntContext) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *BaseGraffleParserListener) EnterFunction_declaration(ctx *Function_declarationContext) {}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *BaseGraffleParserListener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterOne_line_function_declaration is called when production one_line_function_declaration is entered.
func (s *BaseGraffleParserListener) EnterOne_line_function_declaration(ctx *One_line_function_declarationContext) {
}

// ExitOne_line_function_declaration is called when production one_line_function_declaration is exited.
func (s *BaseGraffleParserListener) ExitOne_line_function_declaration(ctx *One_line_function_declarationContext) {
}

// EnterMult_line_function_declaration is called when production mult_line_function_declaration is entered.
func (s *BaseGraffleParserListener) EnterMult_line_function_declaration(ctx *Mult_line_function_declarationContext) {
}

// ExitMult_line_function_declaration is called when production mult_line_function_declaration is exited.
func (s *BaseGraffleParserListener) ExitMult_line_function_declaration(ctx *Mult_line_function_declarationContext) {
}

// EnterFunction_declaration_head is called when production function_declaration_head is entered.
func (s *BaseGraffleParserListener) EnterFunction_declaration_head(ctx *Function_declaration_headContext) {
}

// ExitFunction_declaration_head is called when production function_declaration_head is exited.
func (s *BaseGraffleParserListener) ExitFunction_declaration_head(ctx *Function_declaration_headContext) {
}

// EnterOne_line_procedure_declaration is called when production one_line_procedure_declaration is entered.
func (s *BaseGraffleParserListener) EnterOne_line_procedure_declaration(ctx *One_line_procedure_declarationContext) {
}

// ExitOne_line_procedure_declaration is called when production one_line_procedure_declaration is exited.
func (s *BaseGraffleParserListener) ExitOne_line_procedure_declaration(ctx *One_line_procedure_declarationContext) {
}

// EnterMult_line_procedure_declaration is called when production mult_line_procedure_declaration is entered.
func (s *BaseGraffleParserListener) EnterMult_line_procedure_declaration(ctx *Mult_line_procedure_declarationContext) {
}

// ExitMult_line_procedure_declaration is called when production mult_line_procedure_declaration is exited.
func (s *BaseGraffleParserListener) ExitMult_line_procedure_declaration(ctx *Mult_line_procedure_declarationContext) {
}

// EnterProcedure_declaration_head is called when production procedure_declaration_head is entered.
func (s *BaseGraffleParserListener) EnterProcedure_declaration_head(ctx *Procedure_declaration_headContext) {
}

// ExitProcedure_declaration_head is called when production procedure_declaration_head is exited.
func (s *BaseGraffleParserListener) ExitProcedure_declaration_head(ctx *Procedure_declaration_headContext) {
}

// EnterVar_assign is called when production var_assign is entered.
func (s *BaseGraffleParserListener) EnterVar_assign(ctx *Var_assignContext) {}

// ExitVar_assign is called when production var_assign is exited.
func (s *BaseGraffleParserListener) ExitVar_assign(ctx *Var_assignContext) {}

// EnterArc_assign is called when production arc_assign is entered.
func (s *BaseGraffleParserListener) EnterArc_assign(ctx *Arc_assignContext) {}

// ExitArc_assign is called when production arc_assign is exited.
func (s *BaseGraffleParserListener) ExitArc_assign(ctx *Arc_assignContext) {}

// EnterArc is called when production arc is entered.
func (s *BaseGraffleParserListener) EnterArc(ctx *ArcContext) {}

// ExitArc is called when production arc is exited.
func (s *BaseGraffleParserListener) ExitArc(ctx *ArcContext) {}

// EnterOr_w_arc_lr is called when production or_w_arc_lr is entered.
func (s *BaseGraffleParserListener) EnterOr_w_arc_lr(ctx *Or_w_arc_lrContext) {}

// ExitOr_w_arc_lr is called when production or_w_arc_lr is exited.
func (s *BaseGraffleParserListener) ExitOr_w_arc_lr(ctx *Or_w_arc_lrContext) {}

// EnterOr_w_arc_rl is called when production or_w_arc_rl is entered.
func (s *BaseGraffleParserListener) EnterOr_w_arc_rl(ctx *Or_w_arc_rlContext) {}

// ExitOr_w_arc_rl is called when production or_w_arc_rl is exited.
func (s *BaseGraffleParserListener) ExitOr_w_arc_rl(ctx *Or_w_arc_rlContext) {}

// EnterUnor_w_arc is called when production unor_w_arc is entered.
func (s *BaseGraffleParserListener) EnterUnor_w_arc(ctx *Unor_w_arcContext) {}

// ExitUnor_w_arc is called when production unor_w_arc is exited.
func (s *BaseGraffleParserListener) ExitUnor_w_arc(ctx *Unor_w_arcContext) {}

// EnterVertice_assign is called when production vertice_assign is entered.
func (s *BaseGraffleParserListener) EnterVertice_assign(ctx *Vertice_assignContext) {}

// ExitVertice_assign is called when production vertice_assign is exited.
func (s *BaseGraffleParserListener) ExitVertice_assign(ctx *Vertice_assignContext) {}

// EnterGraph_assign is called when production graph_assign is entered.
func (s *BaseGraffleParserListener) EnterGraph_assign(ctx *Graph_assignContext) {}

// ExitGraph_assign is called when production graph_assign is exited.
func (s *BaseGraffleParserListener) ExitGraph_assign(ctx *Graph_assignContext) {}

// EnterLabeled_assign is called when production labeled_assign is entered.
func (s *BaseGraffleParserListener) EnterLabeled_assign(ctx *Labeled_assignContext) {}

// ExitLabeled_assign is called when production labeled_assign is exited.
func (s *BaseGraffleParserListener) ExitLabeled_assign(ctx *Labeled_assignContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseGraffleParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseGraffleParserListener) ExitExpr(ctx *ExprContext) {}

// EnterIntegral_expr is called when production integral_expr is entered.
func (s *BaseGraffleParserListener) EnterIntegral_expr(ctx *Integral_exprContext) {}

// ExitIntegral_expr is called when production integral_expr is exited.
func (s *BaseGraffleParserListener) ExitIntegral_expr(ctx *Integral_exprContext) {}

// EnterLogical_expr is called when production logical_expr is entered.
func (s *BaseGraffleParserListener) EnterLogical_expr(ctx *Logical_exprContext) {}

// ExitLogical_expr is called when production logical_expr is exited.
func (s *BaseGraffleParserListener) ExitLogical_expr(ctx *Logical_exprContext) {}

// EnterLog_expr_operand is called when production log_expr_operand is entered.
func (s *BaseGraffleParserListener) EnterLog_expr_operand(ctx *Log_expr_operandContext) {}

// ExitLog_expr_operand is called when production log_expr_operand is exited.
func (s *BaseGraffleParserListener) ExitLog_expr_operand(ctx *Log_expr_operandContext) {}

// EnterBin_log_operator is called when production bin_log_operator is entered.
func (s *BaseGraffleParserListener) EnterBin_log_operator(ctx *Bin_log_operatorContext) {}

// ExitBin_log_operator is called when production bin_log_operator is exited.
func (s *BaseGraffleParserListener) ExitBin_log_operator(ctx *Bin_log_operatorContext) {}

// EnterUnar_log_operator is called when production unar_log_operator is entered.
func (s *BaseGraffleParserListener) EnterUnar_log_operator(ctx *Unar_log_operatorContext) {}

// ExitUnar_log_operator is called when production unar_log_operator is exited.
func (s *BaseGraffleParserListener) ExitUnar_log_operator(ctx *Unar_log_operatorContext) {}

// EnterArithm_expr is called when production arithm_expr is entered.
func (s *BaseGraffleParserListener) EnterArithm_expr(ctx *Arithm_exprContext) {}

// ExitArithm_expr is called when production arithm_expr is exited.
func (s *BaseGraffleParserListener) ExitArithm_expr(ctx *Arithm_exprContext) {}

// EnterArithm_expr_operand is called when production arithm_expr_operand is entered.
func (s *BaseGraffleParserListener) EnterArithm_expr_operand(ctx *Arithm_expr_operandContext) {}

// ExitArithm_expr_operand is called when production arithm_expr_operand is exited.
func (s *BaseGraffleParserListener) ExitArithm_expr_operand(ctx *Arithm_expr_operandContext) {}

// EnterBin_arithm_operator is called when production bin_arithm_operator is entered.
func (s *BaseGraffleParserListener) EnterBin_arithm_operator(ctx *Bin_arithm_operatorContext) {}

// ExitBin_arithm_operator is called when production bin_arithm_operator is exited.
func (s *BaseGraffleParserListener) ExitBin_arithm_operator(ctx *Bin_arithm_operatorContext) {}

// EnterArithm_assign_operator is called when production arithm_assign_operator is entered.
func (s *BaseGraffleParserListener) EnterArithm_assign_operator(ctx *Arithm_assign_operatorContext) {}

// ExitArithm_assign_operator is called when production arithm_assign_operator is exited.
func (s *BaseGraffleParserListener) ExitArithm_assign_operator(ctx *Arithm_assign_operatorContext) {}

// EnterBuiltin_function_call is called when production builtin_function_call is entered.
func (s *BaseGraffleParserListener) EnterBuiltin_function_call(ctx *Builtin_function_callContext) {}

// ExitBuiltin_function_call is called when production builtin_function_call is exited.
func (s *BaseGraffleParserListener) ExitBuiltin_function_call(ctx *Builtin_function_callContext) {}

// EnterBuilt_func_print is called when production built_func_print is entered.
func (s *BaseGraffleParserListener) EnterBuilt_func_print(ctx *Built_func_printContext) {}

// ExitBuilt_func_print is called when production built_func_print is exited.
func (s *BaseGraffleParserListener) ExitBuilt_func_print(ctx *Built_func_printContext) {}

// EnterBuilt_func_input is called when production built_func_input is entered.
func (s *BaseGraffleParserListener) EnterBuilt_func_input(ctx *Built_func_inputContext) {}

// ExitBuilt_func_input is called when production built_func_input is exited.
func (s *BaseGraffleParserListener) ExitBuilt_func_input(ctx *Built_func_inputContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseGraffleParserListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseGraffleParserListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseGraffleParserListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseGraffleParserListener) ExitLabel(ctx *LabelContext) {}

// EnterValue is called when production value is entered.
func (s *BaseGraffleParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseGraffleParserListener) ExitValue(ctx *ValueContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseGraffleParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseGraffleParserListener) ExitVariable(ctx *VariableContext) {}

// EnterBuiltin_type is called when production builtin_type is entered.
func (s *BaseGraffleParserListener) EnterBuiltin_type(ctx *Builtin_typeContext) {}

// ExitBuiltin_type is called when production builtin_type is exited.
func (s *BaseGraffleParserListener) ExitBuiltin_type(ctx *Builtin_typeContext) {}

// EnterBlock_end is called when production block_end is entered.
func (s *BaseGraffleParserListener) EnterBlock_end(ctx *Block_endContext) {}

// ExitBlock_end is called when production block_end is exited.
func (s *BaseGraffleParserListener) ExitBlock_end(ctx *Block_endContext) {}
