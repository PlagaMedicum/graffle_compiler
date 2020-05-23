package listener

import (
	"bytes"
	"fmt"
	"github.com/PlagaMedicum/graffle/pkg/gen"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"strings"
)

type GraffleListener struct {
	*parser.BaseGraffleParserListener
	Buffer bytes.Buffer
	paramStack []string
}

func (s *GraffleListener) pushParam(p string) {
	s.paramStack = append(s.paramStack, p)
}

func (s *GraffleListener) pushParamf(format string, a ...interface{}) {
	s.paramStack = append(s.paramStack, fmt.Sprintf(format, a...))
}

func (s *GraffleListener) popParam() string {
	if len(s.paramStack) < 1 {
		log.Fatal("Parsing error! Getting param from empty stack!")
	}

	last := len(s.paramStack) - 1
	res := s.paramStack[last]
	s.paramStack = s.paramStack[:last]

	return res
}

func (s *GraffleListener) writeBuf(format string, a ...interface{}) {
	s.Buffer.WriteString(fmt.Sprintf(format, a...))
}

// VisitTerminal is called when a terminal node is visited.
func (s *GraffleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *GraffleListener) VisitErrorNode(node antlr.ErrorNode) {
	log.Fatalf("Error parsing: %s", node.GetText())
}

// EnterEveryRule is called when any rule is entered.
func (s *GraffleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *GraffleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *GraffleListener) EnterFile(ctx *parser.FileContext) {
	s.writeBuf(`package main

import . "github.com/PlagaMedicum/graffle/pkg/runtime"

func main() {`)
}

// ExitFile is called when production file is exited.
func (s *GraffleListener) ExitFile(ctx *parser.FileContext) {
	s.writeBuf("\n}")
}

// EnterSequence is called when production sequence is entered.
func (s *GraffleListener) EnterSequence(ctx *parser.SequenceContext) {}

// ExitSequence is called when production sequence is exited.
func (s *GraffleListener) ExitSequence(ctx *parser.SequenceContext) {}

// EnterSequence_element is called when production sequence_element is entered.
func (s *GraffleListener) EnterSequence_element(ctx *parser.Sequence_elementContext) {}

// ExitSequence_element is called when production sequence_element is exited.
func (s *GraffleListener) ExitSequence_element(ctx *parser.Sequence_elementContext) {}

// EnterSequence_line is called when production sequence_line is entered.
func (s *GraffleListener) EnterSequence_line(ctx *parser.Sequence_lineContext) {}

// ExitSequence_line is called when production sequence_line is exited.
func (s *GraffleListener) ExitSequence_line(ctx *parser.Sequence_lineContext) {}

// EnterOne_line_sequence_element is called when production one_line_sequence_element is entered.
func (s *GraffleListener) EnterOne_line_sequence_element(ctx *parser.One_line_sequence_elementContext) {
}

// ExitOne_line_sequence_element is called when production one_line_sequence_element is exited.
func (s *GraffleListener) ExitOne_line_sequence_element(ctx *parser.One_line_sequence_elementContext) {
}

// EnterAtom_action is called when production atom_action is entered.
func (s *GraffleListener) EnterAtom_action(ctx *parser.Atom_actionContext) {}

// ExitAtom_action is called when production atom_action is exited.
func (s *GraffleListener) ExitAtom_action(ctx *parser.Atom_actionContext) {}

// EnterIf_stmnt is called when production if_stmnt is entered.
func (s *GraffleListener) EnterIf_stmnt(ctx *parser.If_stmntContext) {}

// ExitIf_stmnt is called when production if_stmnt is exited.
func (s *GraffleListener) ExitIf_stmnt(ctx *parser.If_stmntContext) {}

// EnterIf_is_stmnt is called when production if_is_stmnt is entered.
func (s *GraffleListener) EnterIf_is_stmnt(ctx *parser.If_is_stmntContext) {}

// ExitIf_is_stmnt is called when production if_is_stmnt is exited.
func (s *GraffleListener) ExitIf_is_stmnt(ctx *parser.If_is_stmntContext) {}

// EnterCase_stmnt is called when production case_stmnt is entered.
func (s *GraffleListener) EnterCase_stmnt(ctx *parser.Case_stmntContext) {}

// ExitCase_stmnt is called when production case_stmnt is exited.
func (s *GraffleListener) ExitCase_stmnt(ctx *parser.Case_stmntContext) {}

// EnterOne_line_stmnt is called when production one_line_stmnt is entered.
func (s *GraffleListener) EnterOne_line_stmnt(ctx *parser.One_line_stmntContext) {}

// ExitOne_line_stmnt is called when production one_line_stmnt is exited.
func (s *GraffleListener) ExitOne_line_stmnt(ctx *parser.One_line_stmntContext) {}

// EnterMult_line_stmnt is called when production mult_line_stmnt is entered.
func (s *GraffleListener) EnterMult_line_stmnt(ctx *parser.Mult_line_stmntContext) {}

// ExitMult_line_stmnt is called when production mult_line_stmnt is exited.
func (s *GraffleListener) ExitMult_line_stmnt(ctx *parser.Mult_line_stmntContext) {}

// EnterStmnt is called when production stmnt is entered.
func (s *GraffleListener) EnterStmnt(ctx *parser.StmntContext) {}

// ExitStmnt is called when production stmnt is exited.
func (s *GraffleListener) ExitStmnt(ctx *parser.StmntContext) {}

// EnterWhile_stmnt is called when production while_stmnt is entered.
func (s *GraffleListener) EnterWhile_stmnt(ctx *parser.While_stmntContext) {}

// ExitWhile_stmnt is called when production while_stmnt is exited.
func (s *GraffleListener) ExitWhile_stmnt(ctx *parser.While_stmntContext) {}

// EnterUntil_stmnt is called when production until_stmnt is entered.
func (s *GraffleListener) EnterUntil_stmnt(ctx *parser.Until_stmntContext) {}

// ExitUntil_stmnt is called when production until_stmnt is exited.
func (s *GraffleListener) ExitUntil_stmnt(ctx *parser.Until_stmntContext) {}

// EnterForLogical is called when production ForLogical is entered.
func (s *GraffleListener) EnterForLogical(ctx *parser.ForLogicalContext) {}

// ExitForLogical is called when production ForLogical is exited.
func (s *GraffleListener) ExitForLogical(ctx *parser.ForLogicalContext) {}

// EnterForVar is called when production ForVar is entered.
func (s *GraffleListener) EnterForVar(ctx *parser.ForVarContext) {}

// ExitForVar is called when production ForVar is exited.
func (s *GraffleListener) ExitForVar(ctx *parser.ForVarContext) {}

// EnterForRange is called when production ForRange is entered.
func (s *GraffleListener) EnterForRange(ctx *parser.ForRangeContext) {}

// ExitForRange is called when production ForRange is exited.
func (s *GraffleListener) ExitForRange(ctx *parser.ForRangeContext) {}

// EnterFrom_to_stmnt is called when production from_to_stmnt is entered.
func (s *GraffleListener) EnterFrom_to_stmnt(ctx *parser.From_to_stmntContext) {}

// ExitFrom_to_stmnt is called when production from_to_stmnt is exited.
func (s *GraffleListener) ExitFrom_to_stmnt(ctx *parser.From_to_stmntContext) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *GraffleListener) EnterFunction_declaration(ctx *parser.Function_declarationContext) {}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *GraffleListener) ExitFunction_declaration(ctx *parser.Function_declarationContext) {}

// EnterOne_line_function_declaration is called when production one_line_function_declaration is entered.
func (s *GraffleListener) EnterOne_line_function_declaration(ctx *parser.One_line_function_declarationContext) {
}

// ExitOne_line_function_declaration is called when production one_line_function_declaration is exited.
func (s *GraffleListener) ExitOne_line_function_declaration(ctx *parser.One_line_function_declarationContext) {
}

// EnterMult_line_function_declaration is called when production mult_line_function_declaration is entered.
func (s *GraffleListener) EnterMult_line_function_declaration(ctx *parser.Mult_line_function_declarationContext) {
}

// ExitMult_line_function_declaration is called when production mult_line_function_declaration is exited.
func (s *GraffleListener) ExitMult_line_function_declaration(ctx *parser.Mult_line_function_declarationContext) {
}

// EnterFunction_declaration_head is called when production function_declaration_head is entered.
func (s *GraffleListener) EnterFunction_declaration_head(ctx *parser.Function_declaration_headContext) {
}

// ExitFunction_declaration_head is called when production function_declaration_head is exited.
func (s *GraffleListener) ExitFunction_declaration_head(ctx *parser.Function_declaration_headContext) {
}

// EnterOne_line_procedure_declaration is called when production one_line_procedure_declaration is entered.
func (s *GraffleListener) EnterOne_line_procedure_declaration(ctx *parser.One_line_procedure_declarationContext) {
}

// ExitOne_line_procedure_declaration is called when production one_line_procedure_declaration is exited.
func (s *GraffleListener) ExitOne_line_procedure_declaration(ctx *parser.One_line_procedure_declarationContext) {
}

// EnterMult_line_procedure_declaration is called when production mult_line_procedure_declaration is entered.
func (s *GraffleListener) EnterMult_line_procedure_declaration(ctx *parser.Mult_line_procedure_declarationContext) {
}

// ExitMult_line_procedure_declaration is called when production mult_line_procedure_declaration is exited.
func (s *GraffleListener) ExitMult_line_procedure_declaration(ctx *parser.Mult_line_procedure_declarationContext) {
}

// EnterProcedure_declaration_head is called when production procedure_declaration_head is entered.
func (s *GraffleListener) EnterProcedure_declaration_head(ctx *parser.Procedure_declaration_headContext) {
}

// ExitProcedure_declaration_head is called when production procedure_declaration_head is exited.
func (s *GraffleListener) ExitProcedure_declaration_head(ctx *parser.Procedure_declaration_headContext) {
}

// EnterVar_declaration is called when production var_declaration is entered.
func (s *GraffleListener) EnterVar_declaration(ctx *parser.Var_declarationContext) {

}

// ExitVar_declaration is called when production var_declaration is exited.
func (s *GraffleListener) ExitVar_declaration(ctx *parser.Var_declarationContext) {}

// EnterArc_declaration is called when production arc_declaration is entered.
func (s *GraffleListener) EnterArc_declaration(ctx *parser.Arc_declarationContext) {}

// ExitArc_declaration is called when production arc_declaration is exited.
func (s *GraffleListener) ExitArc_declaration(ctx *parser.Arc_declarationContext) {}

// EnterArc is called when production arc is entered.
func (s *GraffleListener) EnterArc(ctx *parser.ArcContext) {}

// ExitArc is called when production arc is exited.
func (s *GraffleListener) ExitArc(ctx *parser.ArcContext) {}

// EnterVertice_declaration is called when production vertice_declaration is entered.
func (s *GraffleListener) EnterVertice_declaration(ctx *parser.Vertice_declarationContext) {}

// ExitVertice_declaration is called when production vertice_declaration is exited.
func (s *GraffleListener) ExitVertice_declaration(ctx *parser.Vertice_declarationContext) {}

// EnterGraph_declaration is called when production graph_declaration is entered.
func (s *GraffleListener) EnterGraph_declaration(ctx *parser.Graph_declarationContext) {}

// ExitGraph_declaration is called when production graph_declaration is exited.
func (s *GraffleListener) ExitGraph_declaration(ctx *parser.Graph_declarationContext) {}

// EnterLabeled_declaration is called when production labeled_declaration is entered.
func (s *GraffleListener) EnterLabeled_declaration(ctx *parser.Labeled_declarationContext) {}

// ExitLabeled_declaration is called when production labeled_declaration is exited.
func (s *GraffleListener) ExitLabeled_declaration(ctx *parser.Labeled_declarationContext) {}

// EnterExpr is called when production expr is entered.
func (s *GraffleListener) EnterExpr(ctx *parser.ExprContext) {
	v := ctx.GetTokens(parser.GraffleParserVAR)
	if len(v) > 0 {
		s.writeBuf(v[0].GetText())
	}
}

// ExitExpr is called when production expr is exited.
func (s *GraffleListener) ExitExpr(ctx *parser.ExprContext) {}

// EnterIntegral_expr is called when production integral_expr is entered.
func (s *GraffleListener) EnterIntegral_expr(ctx *parser.Integral_exprContext) {
	ctx.GetToken(parser.GraffleLexerNUMBER, 0)
}

// ExitIntegral_expr is called when production integral_expr is exited.
func (s *GraffleListener) ExitIntegral_expr(ctx *parser.Integral_exprContext) {}

// EnterLogical_expr is called when production logical_expr is entered.
func (s *GraffleListener) EnterLogical_expr(ctx *parser.Logical_exprContext) {}

// ExitLogical_expr is called when production logical_expr is exited.
func (s *GraffleListener) ExitLogical_expr(ctx *parser.Logical_exprContext) {}

// EnterBin_log_operator is called when production bin_log_operator is entered.
func (s *GraffleListener) EnterBin_log_operator(ctx *parser.Bin_log_operatorContext) {}

// ExitBin_log_operator is called when production bin_log_operator is exited.
func (s *GraffleListener) ExitBin_log_operator(ctx *parser.Bin_log_operatorContext) {}

// EnterUnar_log_operator is called when production unar_log_operator is entered.
func (s *GraffleListener) EnterUnar_log_operator(ctx *parser.Unar_log_operatorContext) {}

// ExitUnar_log_operator is called when production unar_log_operator is exited.
func (s *GraffleListener) ExitUnar_log_operator(ctx *parser.Unar_log_operatorContext) {}

// EnterArithm_expr is called when production arithm_expr is entered.
func (s *GraffleListener) EnterArithm_expr(ctx *parser.Arithm_exprContext) {
}

// ExitArithm_expr is called when production arithm_expr is exited.
func (s *GraffleListener) ExitArithm_expr(ctx *parser.Arithm_exprContext) {
	r := s.popParam()
	l := s.popParam()
	switch ctx.GetOp().GetStart().GetTokenType() {
	case parser.GraffleParserADD:
		s.pushParamf("Add(%s, %s)", l, r)
	case parser.GraffleParserSUB:
		s.pushParamf("Subtract(%s, %s)", l, r)
	case parser.GraffleParserMULT:
		s.pushParamf("Multiply(%s, %s)", l, r)
	case parser.GraffleParserDIV:
		s.pushParamf("Divide(%s, %s)", l, r)
	}
}

// EnterArithm_expr_operand is called when production arithm_expr_operand is entered.
func (s *GraffleListener) EnterArithm_expr_operand(ctx *parser.Arithm_expr_operandContext) {
}

// ExitArithm_expr_operand is called when production arithm_expr_operand is exited.
func (s *GraffleListener) ExitArithm_expr_operand(ctx *parser.Arithm_expr_operandContext) {}

// EnterBin_arithm_operator is called when production bin_arithm_operator is entered.
func (s *GraffleListener) EnterBin_arithm_operator(ctx *parser.Bin_arithm_operatorContext) {}

// ExitBin_arithm_operator is called when production bin_arithm_operator is exited.
func (s *GraffleListener) ExitBin_arithm_operator(ctx *parser.Bin_arithm_operatorContext) {}

// EnterArithm_assign_operator is called when production arithm_assign_operator is entered.
func (s *GraffleListener) EnterArithm_assign_operator(ctx *parser.Arithm_assign_operatorContext) {}

// ExitArithm_assign_operator is called when production arithm_assign_operator is exited.
func (s *GraffleListener) ExitArithm_assign_operator(ctx *parser.Arithm_assign_operatorContext) {}

// EnterBuiltin_function_call is called when production builtin_function_call is entered.
func (s *GraffleListener) EnterBuiltin_function_call(ctx *parser.Builtin_function_callContext) {}

// ExitBuiltin_function_call is called when production builtin_function_call is exited.
func (s *GraffleListener) ExitBuiltin_function_call(ctx *parser.Builtin_function_callContext) {}

// EnterBuilt_func_print is called when production built_func_print is entered.
func (s *GraffleListener) EnterBuilt_func_print(ctx *parser.Built_func_printContext) {}

// ExitBuilt_func_print is called when production built_func_print is exited.
func (s *GraffleListener) ExitBuilt_func_print(ctx *parser.Built_func_printContext) {
	s.writeBuf("\nPrint(%s);", s.popParam())
}

// EnterBuilt_func_input is called when production built_func_input is entered.
func (s *GraffleListener) EnterBuilt_func_input(ctx *parser.Built_func_inputContext) {}

// ExitBuilt_func_input is called when production built_func_input is exited.
func (s *GraffleListener) ExitBuilt_func_input(ctx *parser.Built_func_inputContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *GraffleListener) EnterFunction_call(ctx *parser.Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *GraffleListener) ExitFunction_call(ctx *parser.Function_callContext) {}

// EnterLabel is called when production label is entered.
func (s *GraffleListener) EnterLabel(ctx *parser.LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *GraffleListener) ExitLabel(ctx *parser.LabelContext) {}

// EnterValue is called when production value is entered.
func (s *GraffleListener) EnterValue(ctx *parser.ValueContext) {}

// ExitValue is called when production value is exited.
func (s *GraffleListener) ExitValue(ctx *parser.ValueContext) {}

// EnterBuiltin is called when production builtin is entered.
func (s *GraffleListener) EnterBuiltin(ctx *parser.BuiltinContext) {
	switch ctx.GetStart().GetTokenType() {
	case parser.GraffleParserNUMBER:
		v := strings.ReplaceAll(ctx.GetText(), ",", ".")
		s.pushParamf("NewNumber(%s)", v)
	case parser.GraffleLexerSTRING:
		s.pushParamf("NewString(%s)", ctx.GetText())
	case parser.GraffleParserBOOL:
		v := strings.ToLower(ctx.GetText())
		s.pushParamf("NewBool(%s)", v)
	}
}

// ExitBuiltin is called when production builtin is exited.
func (s *GraffleListener) ExitBuiltin(ctx *parser.BuiltinContext) {}

// EnterBlock_end is called when production block_end is entered.
func (s *GraffleListener) EnterBlock_end(ctx *parser.Block_endContext) {}

// ExitBlock_end is called when production block_end is exited.
func (s *GraffleListener) ExitBlock_end(ctx *parser.Block_endContext) {}
