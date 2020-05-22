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

func (v *BaseGraffleParserVisitor) VisitIf_is_stmnt(ctx *If_is_stmntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitCase_stmnt(ctx *Case_stmntContext) interface{} {
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

func (v *BaseGraffleParserVisitor) VisitVar_declaration(ctx *Var_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitArc_declaration(ctx *Arc_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitArc(ctx *ArcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitVertice_declaration(ctx *Vertice_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitGraph_declaration(ctx *Graph_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitLabeled_declaration(ctx *Labeled_declarationContext) interface{} {
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

func (v *BaseGraffleParserVisitor) VisitBin_log_operator(ctx *Bin_log_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitUnar_log_operator(ctx *Unar_log_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGraffleParserVisitor) VisitArithm_expr(ctx *Arithm_exprContext) interface{} {
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

func (v *BaseGraffleParserVisitor) VisitBlock_end(ctx *Block_endContext) interface{} {
	return v.VisitChildren(ctx)
}
