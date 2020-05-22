// Code generated from /home/plagamedicum/Documents/yapis/graffle/GraffleParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // GraffleParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GraffleParserListener is a complete listener for a parse tree produced by GraffleParser.
type GraffleParserListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterSequence is called when entering the sequence production.
	EnterSequence(c *SequenceContext)

	// EnterSequence_element is called when entering the sequence_element production.
	EnterSequence_element(c *Sequence_elementContext)

	// EnterSequence_line is called when entering the sequence_line production.
	EnterSequence_line(c *Sequence_lineContext)

	// EnterOne_line_sequence_element is called when entering the one_line_sequence_element production.
	EnterOne_line_sequence_element(c *One_line_sequence_elementContext)

	// EnterAtom_action is called when entering the atom_action production.
	EnterAtom_action(c *Atom_actionContext)

	// EnterIf_stmnt is called when entering the if_stmnt production.
	EnterIf_stmnt(c *If_stmntContext)

	// EnterIf_is_stmnt is called when entering the if_is_stmnt production.
	EnterIf_is_stmnt(c *If_is_stmntContext)

	// EnterCase_stmnt is called when entering the case_stmnt production.
	EnterCase_stmnt(c *Case_stmntContext)

	// EnterOne_line_stmnt is called when entering the one_line_stmnt production.
	EnterOne_line_stmnt(c *One_line_stmntContext)

	// EnterMult_line_stmnt is called when entering the mult_line_stmnt production.
	EnterMult_line_stmnt(c *Mult_line_stmntContext)

	// EnterStmnt is called when entering the stmnt production.
	EnterStmnt(c *StmntContext)

	// EnterWhile_stmnt is called when entering the while_stmnt production.
	EnterWhile_stmnt(c *While_stmntContext)

	// EnterUntil_stmnt is called when entering the until_stmnt production.
	EnterUntil_stmnt(c *Until_stmntContext)

	// EnterForLogical is called when entering the ForLogical production.
	EnterForLogical(c *ForLogicalContext)

	// EnterForVar is called when entering the ForVar production.
	EnterForVar(c *ForVarContext)

	// EnterForRange is called when entering the ForRange production.
	EnterForRange(c *ForRangeContext)

	// EnterFrom_to_stmnt is called when entering the from_to_stmnt production.
	EnterFrom_to_stmnt(c *From_to_stmntContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterOne_line_function_declaration is called when entering the one_line_function_declaration production.
	EnterOne_line_function_declaration(c *One_line_function_declarationContext)

	// EnterMult_line_function_declaration is called when entering the mult_line_function_declaration production.
	EnterMult_line_function_declaration(c *Mult_line_function_declarationContext)

	// EnterFunction_declaration_head is called when entering the function_declaration_head production.
	EnterFunction_declaration_head(c *Function_declaration_headContext)

	// EnterOne_line_procedure_declaration is called when entering the one_line_procedure_declaration production.
	EnterOne_line_procedure_declaration(c *One_line_procedure_declarationContext)

	// EnterMult_line_procedure_declaration is called when entering the mult_line_procedure_declaration production.
	EnterMult_line_procedure_declaration(c *Mult_line_procedure_declarationContext)

	// EnterProcedure_declaration_head is called when entering the procedure_declaration_head production.
	EnterProcedure_declaration_head(c *Procedure_declaration_headContext)

	// EnterVar_declaration is called when entering the var_declaration production.
	EnterVar_declaration(c *Var_declarationContext)

	// EnterArc_declaration is called when entering the arc_declaration production.
	EnterArc_declaration(c *Arc_declarationContext)

	// EnterArc is called when entering the arc production.
	EnterArc(c *ArcContext)

	// EnterVertice_declaration is called when entering the vertice_declaration production.
	EnterVertice_declaration(c *Vertice_declarationContext)

	// EnterGraph_declaration is called when entering the graph_declaration production.
	EnterGraph_declaration(c *Graph_declarationContext)

	// EnterLabeled_declaration is called when entering the labeled_declaration production.
	EnterLabeled_declaration(c *Labeled_declarationContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterIntegral_expr is called when entering the integral_expr production.
	EnterIntegral_expr(c *Integral_exprContext)

	// EnterLogical_expr is called when entering the logical_expr production.
	EnterLogical_expr(c *Logical_exprContext)

	// EnterBin_log_operator is called when entering the bin_log_operator production.
	EnterBin_log_operator(c *Bin_log_operatorContext)

	// EnterUnar_log_operator is called when entering the unar_log_operator production.
	EnterUnar_log_operator(c *Unar_log_operatorContext)

	// EnterArithm_expr is called when entering the arithm_expr production.
	EnterArithm_expr(c *Arithm_exprContext)

	// EnterBin_arithm_operator is called when entering the bin_arithm_operator production.
	EnterBin_arithm_operator(c *Bin_arithm_operatorContext)

	// EnterArithm_assign_operator is called when entering the arithm_assign_operator production.
	EnterArithm_assign_operator(c *Arithm_assign_operatorContext)

	// EnterBuiltin_function_call is called when entering the builtin_function_call production.
	EnterBuiltin_function_call(c *Builtin_function_callContext)

	// EnterBuilt_func_print is called when entering the built_func_print production.
	EnterBuilt_func_print(c *Built_func_printContext)

	// EnterBuilt_func_input is called when entering the built_func_input production.
	EnterBuilt_func_input(c *Built_func_inputContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterBlock_end is called when entering the block_end production.
	EnterBlock_end(c *Block_endContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitSequence is called when exiting the sequence production.
	ExitSequence(c *SequenceContext)

	// ExitSequence_element is called when exiting the sequence_element production.
	ExitSequence_element(c *Sequence_elementContext)

	// ExitSequence_line is called when exiting the sequence_line production.
	ExitSequence_line(c *Sequence_lineContext)

	// ExitOne_line_sequence_element is called when exiting the one_line_sequence_element production.
	ExitOne_line_sequence_element(c *One_line_sequence_elementContext)

	// ExitAtom_action is called when exiting the atom_action production.
	ExitAtom_action(c *Atom_actionContext)

	// ExitIf_stmnt is called when exiting the if_stmnt production.
	ExitIf_stmnt(c *If_stmntContext)

	// ExitIf_is_stmnt is called when exiting the if_is_stmnt production.
	ExitIf_is_stmnt(c *If_is_stmntContext)

	// ExitCase_stmnt is called when exiting the case_stmnt production.
	ExitCase_stmnt(c *Case_stmntContext)

	// ExitOne_line_stmnt is called when exiting the one_line_stmnt production.
	ExitOne_line_stmnt(c *One_line_stmntContext)

	// ExitMult_line_stmnt is called when exiting the mult_line_stmnt production.
	ExitMult_line_stmnt(c *Mult_line_stmntContext)

	// ExitStmnt is called when exiting the stmnt production.
	ExitStmnt(c *StmntContext)

	// ExitWhile_stmnt is called when exiting the while_stmnt production.
	ExitWhile_stmnt(c *While_stmntContext)

	// ExitUntil_stmnt is called when exiting the until_stmnt production.
	ExitUntil_stmnt(c *Until_stmntContext)

	// ExitForLogical is called when exiting the ForLogical production.
	ExitForLogical(c *ForLogicalContext)

	// ExitForVar is called when exiting the ForVar production.
	ExitForVar(c *ForVarContext)

	// ExitForRange is called when exiting the ForRange production.
	ExitForRange(c *ForRangeContext)

	// ExitFrom_to_stmnt is called when exiting the from_to_stmnt production.
	ExitFrom_to_stmnt(c *From_to_stmntContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitOne_line_function_declaration is called when exiting the one_line_function_declaration production.
	ExitOne_line_function_declaration(c *One_line_function_declarationContext)

	// ExitMult_line_function_declaration is called when exiting the mult_line_function_declaration production.
	ExitMult_line_function_declaration(c *Mult_line_function_declarationContext)

	// ExitFunction_declaration_head is called when exiting the function_declaration_head production.
	ExitFunction_declaration_head(c *Function_declaration_headContext)

	// ExitOne_line_procedure_declaration is called when exiting the one_line_procedure_declaration production.
	ExitOne_line_procedure_declaration(c *One_line_procedure_declarationContext)

	// ExitMult_line_procedure_declaration is called when exiting the mult_line_procedure_declaration production.
	ExitMult_line_procedure_declaration(c *Mult_line_procedure_declarationContext)

	// ExitProcedure_declaration_head is called when exiting the procedure_declaration_head production.
	ExitProcedure_declaration_head(c *Procedure_declaration_headContext)

	// ExitVar_declaration is called when exiting the var_declaration production.
	ExitVar_declaration(c *Var_declarationContext)

	// ExitArc_declaration is called when exiting the arc_declaration production.
	ExitArc_declaration(c *Arc_declarationContext)

	// ExitArc is called when exiting the arc production.
	ExitArc(c *ArcContext)

	// ExitVertice_declaration is called when exiting the vertice_declaration production.
	ExitVertice_declaration(c *Vertice_declarationContext)

	// ExitGraph_declaration is called when exiting the graph_declaration production.
	ExitGraph_declaration(c *Graph_declarationContext)

	// ExitLabeled_declaration is called when exiting the labeled_declaration production.
	ExitLabeled_declaration(c *Labeled_declarationContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitIntegral_expr is called when exiting the integral_expr production.
	ExitIntegral_expr(c *Integral_exprContext)

	// ExitLogical_expr is called when exiting the logical_expr production.
	ExitLogical_expr(c *Logical_exprContext)

	// ExitBin_log_operator is called when exiting the bin_log_operator production.
	ExitBin_log_operator(c *Bin_log_operatorContext)

	// ExitUnar_log_operator is called when exiting the unar_log_operator production.
	ExitUnar_log_operator(c *Unar_log_operatorContext)

	// ExitArithm_expr is called when exiting the arithm_expr production.
	ExitArithm_expr(c *Arithm_exprContext)

	// ExitBin_arithm_operator is called when exiting the bin_arithm_operator production.
	ExitBin_arithm_operator(c *Bin_arithm_operatorContext)

	// ExitArithm_assign_operator is called when exiting the arithm_assign_operator production.
	ExitArithm_assign_operator(c *Arithm_assign_operatorContext)

	// ExitBuiltin_function_call is called when exiting the builtin_function_call production.
	ExitBuiltin_function_call(c *Builtin_function_callContext)

	// ExitBuilt_func_print is called when exiting the built_func_print production.
	ExitBuilt_func_print(c *Built_func_printContext)

	// ExitBuilt_func_input is called when exiting the built_func_input production.
	ExitBuilt_func_input(c *Built_func_inputContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitBlock_end is called when exiting the block_end production.
	ExitBlock_end(c *Block_endContext)
}
