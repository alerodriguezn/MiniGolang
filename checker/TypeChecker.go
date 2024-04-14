package checker

import (
	"MiniGolang/parser"
	"github.com/antlr4-go/antlr/v4"
)

type Checker struct {
	Table     *SymbolTable
	ErrorList []string //ToDo change to error struct
	*antlr.BaseParseTreeVisitor
	parser antlr.Parser
}

func (v *Checker) VisitRoot(ctx *parser.RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitTopDeclarationList(ctx *parser.TopDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitInnerVarDecls(ctx *parser.InnerVarDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSingleVarDecl(ctx *parser.SingleVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitTypeDecl(ctx *parser.TypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitInnerTypeDecls(ctx *parser.InnerTypeDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSingleTypeDecl(ctx *parser.SingleTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitFuncFrontDecl(ctx *parser.FuncFrontDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitDeclType(ctx *parser.DeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSliceDeclType(ctx *parser.SliceDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitArrayDeclType(ctx *parser.ArrayDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStructDeclType(ctx *parser.StructDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStructMemDecls(ctx *parser.StructMemDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitOperand(ctx *parser.OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIndex(ctx *parser.IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSelector(ctx *parser.SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAppendExpression(ctx *parser.AppendExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLengthExpression(ctx *parser.LengthExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitCapExpression(ctx *parser.CapExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStatement(ctx *parser.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSimpleStatement(ctx *parser.SimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLoop(ctx *parser.LoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSwitch(ctx *parser.SwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionCaseClauseList(ctx *parser.ExpressionCaseClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionCaseClause(ctx *parser.ExpressionCaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionSwitchCase(ctx *parser.ExpressionSwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}
