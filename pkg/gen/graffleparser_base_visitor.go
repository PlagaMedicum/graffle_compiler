// Code generated from /home/plagamedicum/Documents/yapis/graffle/GraffleParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // GraffleParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGraffleParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGraffleParserVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitSequence(ctx *SequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitSequence_element(ctx *Sequence_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitSequence_line(ctx *Sequence_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitOne_line_sequence_element(ctx *One_line_sequence_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitAtom_action(ctx *Atom_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitIf_stmnt(ctx *If_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitElse_stmnt(ctx *Else_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitIf_is_stmnt(ctx *If_is_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitCase_stmnt(ctx *Case_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitDefault_stmnt(ctx *Default_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitOne_line_stmnt(ctx *One_line_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitMult_line_stmnt(ctx *Mult_line_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitStmnt(ctx *StmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitWhile_stmnt(ctx *While_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitUntil_stmnt(ctx *Until_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitForLogical(ctx *ForLogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitForVar(ctx *ForVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitForRange(ctx *ForRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitFrom_to_stmnt(ctx *From_to_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitFunction_declaration(ctx *Function_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitOne_line_function_declaration(ctx *One_line_function_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitMult_line_function_declaration(ctx *Mult_line_function_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitFunction_declaration_head(ctx *Function_declaration_headContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitOne_line_procedure_declaration(ctx *One_line_procedure_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitMult_line_procedure_declaration(ctx *Mult_line_procedure_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitProcedure_declaration_head(ctx *Procedure_declaration_headContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitVar_assign(ctx *Var_assignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitVertice_assign(ctx *Vertice_assignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitVertice_type(ctx *Vertice_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitEdge_assign(ctx *Edge_assignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitEdge_type(ctx *Edge_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitEdge(ctx *EdgeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitOr_w_edge_lr(ctx *Or_w_edge_lrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitOr_w_edge_rl(ctx *Or_w_edge_rlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitUnor_w_edge(ctx *Unor_w_edgeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitGraph_assign(ctx *Graph_assignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitGraph_type(ctx *Graph_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitLabeled_assign(ctx *Labeled_assignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitIntegral_expr(ctx *Integral_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitLogical_expr(ctx *Logical_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitLog_expr_operand(ctx *Log_expr_operandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitBin_log_operator(ctx *Bin_log_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitUnar_log_operator(ctx *Unar_log_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitArithm_expr(ctx *Arithm_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitArithm_expr_operand(ctx *Arithm_expr_operandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitBin_arithm_operator(ctx *Bin_arithm_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitArithm_assign_operator(ctx *Arithm_assign_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitBuiltin_function_call(ctx *Builtin_function_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitBuilt_func_print(ctx *Built_func_printContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitBuilt_func_input(ctx *Built_func_inputContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitFunction_call(ctx *Function_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitBuiltin_type(ctx *Builtin_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitBlock_end(ctx *Block_endContext) interface{} {
	return v.VisitChildren(ctx)
}
