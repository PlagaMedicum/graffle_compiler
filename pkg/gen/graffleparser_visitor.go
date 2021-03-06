// Code generated from /home/plagamedicum/Documents/yapis/graffle/GraffleParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // GraffleParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GraffleParser.
type GraffleParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GraffleParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by GraffleParser#functions_block.
	VisitFunctions_block(ctx *Functions_blockContext) interface{}

	// Visit a parse tree produced by GraffleParser#sequence.
	VisitSequence(ctx *SequenceContext) interface{}

	// Visit a parse tree produced by GraffleParser#sequence_element.
	VisitSequence_element(ctx *Sequence_elementContext) interface{}

	// Visit a parse tree produced by GraffleParser#sequence_line.
	VisitSequence_line(ctx *Sequence_lineContext) interface{}

	// Visit a parse tree produced by GraffleParser#one_line_sequence_element.
	VisitOne_line_sequence_element(ctx *One_line_sequence_elementContext) interface{}

	// Visit a parse tree produced by GraffleParser#atom_action.
	VisitAtom_action(ctx *Atom_actionContext) interface{}

	// Visit a parse tree produced by GraffleParser#if_stmnt.
	VisitIf_stmnt(ctx *If_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#else_stmnt.
	VisitElse_stmnt(ctx *Else_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#elif_stmnt.
	VisitElif_stmnt(ctx *Elif_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#if_is_stmnt.
	VisitIf_is_stmnt(ctx *If_is_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#case_stmnt.
	VisitCase_stmnt(ctx *Case_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#default_stmnt.
	VisitDefault_stmnt(ctx *Default_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#one_line_stmnt.
	VisitOne_line_stmnt(ctx *One_line_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#mult_line_stmnt.
	VisitMult_line_stmnt(ctx *Mult_line_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#cycle_stmnt.
	VisitCycle_stmnt(ctx *Cycle_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#while_stmnt.
	VisitWhile_stmnt(ctx *While_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#until_stmnt.
	VisitUntil_stmnt(ctx *Until_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#ForLogical.
	VisitForLogical(ctx *ForLogicalContext) interface{}

	// Visit a parse tree produced by GraffleParser#ForVar.
	VisitForVar(ctx *ForVarContext) interface{}

	// Visit a parse tree produced by GraffleParser#ForRange.
	VisitForRange(ctx *ForRangeContext) interface{}

	// Visit a parse tree produced by GraffleParser#from_to_stmnt.
	VisitFrom_to_stmnt(ctx *From_to_stmntContext) interface{}

	// Visit a parse tree produced by GraffleParser#function_declaration.
	VisitFunction_declaration(ctx *Function_declarationContext) interface{}

	// Visit a parse tree produced by GraffleParser#one_line_function_declaration.
	VisitOne_line_function_declaration(ctx *One_line_function_declarationContext) interface{}

	// Visit a parse tree produced by GraffleParser#mult_line_function_declaration.
	VisitMult_line_function_declaration(ctx *Mult_line_function_declarationContext) interface{}

	// Visit a parse tree produced by GraffleParser#function_declaration_head.
	VisitFunction_declaration_head(ctx *Function_declaration_headContext) interface{}

	// Visit a parse tree produced by GraffleParser#one_line_procedure_declaration.
	VisitOne_line_procedure_declaration(ctx *One_line_procedure_declarationContext) interface{}

	// Visit a parse tree produced by GraffleParser#mult_line_procedure_declaration.
	VisitMult_line_procedure_declaration(ctx *Mult_line_procedure_declarationContext) interface{}

	// Visit a parse tree produced by GraffleParser#procedure_declaration_head.
	VisitProcedure_declaration_head(ctx *Procedure_declaration_headContext) interface{}

	// Visit a parse tree produced by GraffleParser#var_assign.
	VisitVar_assign(ctx *Var_assignContext) interface{}

	// Visit a parse tree produced by GraffleParser#vertice_assign.
	VisitVertice_assign(ctx *Vertice_assignContext) interface{}

	// Visit a parse tree produced by GraffleParser#vertice_type.
	VisitVertice_type(ctx *Vertice_typeContext) interface{}

	// Visit a parse tree produced by GraffleParser#edge_assign.
	VisitEdge_assign(ctx *Edge_assignContext) interface{}

	// Visit a parse tree produced by GraffleParser#edge_type.
	VisitEdge_type(ctx *Edge_typeContext) interface{}

	// Visit a parse tree produced by GraffleParser#edge.
	VisitEdge(ctx *EdgeContext) interface{}

	// Visit a parse tree produced by GraffleParser#or_w_edge_lr.
	VisitOr_w_edge_lr(ctx *Or_w_edge_lrContext) interface{}

	// Visit a parse tree produced by GraffleParser#or_w_edge_rl.
	VisitOr_w_edge_rl(ctx *Or_w_edge_rlContext) interface{}

	// Visit a parse tree produced by GraffleParser#unor_w_edge.
	VisitUnor_w_edge(ctx *Unor_w_edgeContext) interface{}

	// Visit a parse tree produced by GraffleParser#graph_assign.
	VisitGraph_assign(ctx *Graph_assignContext) interface{}

	// Visit a parse tree produced by GraffleParser#graph_type.
	VisitGraph_type(ctx *Graph_typeContext) interface{}

	// Visit a parse tree produced by GraffleParser#labeled_assign.
	VisitLabeled_assign(ctx *Labeled_assignContext) interface{}

	// Visit a parse tree produced by GraffleParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by GraffleParser#integral_expr.
	VisitIntegral_expr(ctx *Integral_exprContext) interface{}

	// Visit a parse tree produced by GraffleParser#logical_expr.
	VisitLogical_expr(ctx *Logical_exprContext) interface{}

	// Visit a parse tree produced by GraffleParser#log_expr_operand.
	VisitLog_expr_operand(ctx *Log_expr_operandContext) interface{}

	// Visit a parse tree produced by GraffleParser#bin_log_operator.
	VisitBin_log_operator(ctx *Bin_log_operatorContext) interface{}

	// Visit a parse tree produced by GraffleParser#unar_log_operator.
	VisitUnar_log_operator(ctx *Unar_log_operatorContext) interface{}

	// Visit a parse tree produced by GraffleParser#arithm_expr.
	VisitArithm_expr(ctx *Arithm_exprContext) interface{}

	// Visit a parse tree produced by GraffleParser#arithm_expr_operand.
	VisitArithm_expr_operand(ctx *Arithm_expr_operandContext) interface{}

	// Visit a parse tree produced by GraffleParser#bin_arithm_operator.
	VisitBin_arithm_operator(ctx *Bin_arithm_operatorContext) interface{}

	// Visit a parse tree produced by GraffleParser#arithm_assign_operator.
	VisitArithm_assign_operator(ctx *Arithm_assign_operatorContext) interface{}

	// Visit a parse tree produced by GraffleParser#builtin_function_call.
	VisitBuiltin_function_call(ctx *Builtin_function_callContext) interface{}

	// Visit a parse tree produced by GraffleParser#built_func_print.
	VisitBuilt_func_print(ctx *Built_func_printContext) interface{}

	// Visit a parse tree produced by GraffleParser#built_func_input.
	VisitBuilt_func_input(ctx *Built_func_inputContext) interface{}

	// Visit a parse tree produced by GraffleParser#function_call.
	VisitFunction_call(ctx *Function_callContext) interface{}

	// Visit a parse tree produced by GraffleParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by GraffleParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by GraffleParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by GraffleParser#builtin_type.
	VisitBuiltin_type(ctx *Builtin_typeContext) interface{}

	// Visit a parse tree produced by GraffleParser#number_type.
	VisitNumber_type(ctx *Number_typeContext) interface{}

	// Visit a parse tree produced by GraffleParser#block_end.
	VisitBlock_end(ctx *Block_endContext) interface{}
}
