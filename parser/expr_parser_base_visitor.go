// Code generated from C:/Users/navar/GolandProjects/MiniGolang/parser/expr_parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // expr_parser

import "github.com/antlr4-go/antlr/v4"

type Baseexpr_parserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *Baseexpr_parserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitTopDeclarationList(ctx *TopDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitMVarDecl(ctx *MVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitVoidVarDecl(ctx *VoidVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitVarDeclWithType(ctx *VarDeclWithTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitVarDeclWithoutType(ctx *VarDeclWithoutTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitVarDeclNoExp(ctx *VarDeclNoExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitTypeDec(ctx *TypeDecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitMultiTypeDecl(ctx *MultiTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitVoidTypeDecl(ctx *VoidTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitInnerTypeDecls(ctx *InnerTypeDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSingleTypeDecl(ctx *SingleTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitFuncDecl(ctx *FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitFuncFrontDecl(ctx *FuncFrontDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitFuncArgDecls(ctx *FuncArgDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitDeclType(ctx *DeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSliceDeclType(ctx *SliceDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitArrayDeclType(ctx *ArrayDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitStructDeclType(ctx *StructDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitStructMemDecls(ctx *StructMemDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitExpUnary(ctx *ExpUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitExpPrimaryExp(ctx *ExpPrimaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitExpBinary(ctx *ExpBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitCapExp(ctx *CapExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitAppendExp(ctx *AppendExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitLenExp(ctx *LenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSelectorExp(ctx *SelectorExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitOpExp(ctx *OpExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIndexExp(ctx *IndexExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitLiteralOp(ctx *LiteralOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIdentifierOp(ctx *IdentifierOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitExpressionOp(ctx *ExpressionOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIntLit(ctx *IntLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitFloatLit(ctx *FloatLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitRawStringLit(ctx *RawStringLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitInterpretedStringLit(ctx *InterpretedStringLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitRuneLit(ctx *RuneLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitAppendExpression(ctx *AppendExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitLengthExpression(ctx *LengthExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitCapExpression(ctx *CapExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitPrintSt(ctx *PrintStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitPrintlnSt(ctx *PrintlnStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitReturnSt(ctx *ReturnStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitBreakSt(ctx *BreakStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitContinueSt(ctx *ContinueStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSimpleSt(ctx *SimpleStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitBlockSt(ctx *BlockStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSwitchSt(ctx *SwitchStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIfStat(ctx *IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitLoopSt(ctx *LoopStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitTypeDeclSt(ctx *TypeDeclStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitVarDeclSt(ctx *VarDeclStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitExpSt(ctx *ExpStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitAssignSt(ctx *AssignStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitShortDecSt(ctx *ShortDecStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitAssignStat(ctx *AssignStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitOtAssignSt(ctx *OtAssignStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIfSt(ctx *IfStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIfElseIfSt(ctx *IfElseIfStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIfElseBlockSt(ctx *IfElseBlockStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIfSimpleSt(ctx *IfSimpleStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIfSimpleElseIfSt(ctx *IfSimpleElseIfStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitIfSimpleElseBlockSt(ctx *IfSimpleElseBlockStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitForSt(ctx *ForStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitWhileExprSt(ctx *WhileExprStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitOtherForSt(ctx *OtherForStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSimpStSwitch(ctx *SimpStSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitExpSwitch(ctx *ExpSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSimpStSwitchNoExp(ctx *SimpStSwitchNoExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitSwitchDefault(ctx *SwitchDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitCaseExp(ctx *CaseExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baseexpr_parserVisitor) VisitDefaultExp(ctx *DefaultExpContext) interface{} {
	return v.VisitChildren(ctx)
}
