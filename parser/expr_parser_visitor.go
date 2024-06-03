// Code generated from C:/Users/navar/GolandProjects/MiniGolang/parser/expr_parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // expr_parser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by expr_parser.
type expr_parserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by expr_parser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by expr_parser#topDeclarationList.
	VisitTopDeclarationList(ctx *TopDeclarationListContext) interface{}

	// Visit a parse tree produced by expr_parser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by expr_parser#mVarDecl.
	VisitMVarDecl(ctx *MVarDeclContext) interface{}

	// Visit a parse tree produced by expr_parser#voidVarDecl.
	VisitVoidVarDecl(ctx *VoidVarDeclContext) interface{}

	// Visit a parse tree produced by expr_parser#innerVarDecls.
	VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{}

	// Visit a parse tree produced by expr_parser#varDeclWithoutType.
	VisitVarDeclWithoutType(ctx *VarDeclWithoutTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#varDeclWithType.
	VisitVarDeclWithType(ctx *VarDeclWithTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#varDeclNoExp.
	VisitVarDeclNoExp(ctx *VarDeclNoExpContext) interface{}

	// Visit a parse tree produced by expr_parser#singleVarDeclNoExps.
	VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{}

	// Visit a parse tree produced by expr_parser#typeDec.
	VisitTypeDec(ctx *TypeDecContext) interface{}

	// Visit a parse tree produced by expr_parser#multiTypeDecl.
	VisitMultiTypeDecl(ctx *MultiTypeDeclContext) interface{}

	// Visit a parse tree produced by expr_parser#voidTypeDecl.
	VisitVoidTypeDecl(ctx *VoidTypeDeclContext) interface{}

	// Visit a parse tree produced by expr_parser#innerTypeDecls.
	VisitInnerTypeDecls(ctx *InnerTypeDeclsContext) interface{}

	// Visit a parse tree produced by expr_parser#singleTypeDecl.
	VisitSingleTypeDecl(ctx *SingleTypeDeclContext) interface{}

	// Visit a parse tree produced by expr_parser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by expr_parser#funcFrontDecl.
	VisitFuncFrontDecl(ctx *FuncFrontDeclContext) interface{}

	// Visit a parse tree produced by expr_parser#funcArgDecls.
	VisitFuncArgDecls(ctx *FuncArgDeclsContext) interface{}

	// Visit a parse tree produced by expr_parser#nestedDeclType.
	VisitNestedDeclType(ctx *NestedDeclTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#idDeclType.
	VisitIdDeclType(ctx *IdDeclTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#sliceDecType.
	VisitSliceDecType(ctx *SliceDecTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#arrayDecType.
	VisitArrayDecType(ctx *ArrayDecTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#structDecType.
	VisitStructDecType(ctx *StructDecTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#sliceDeclType.
	VisitSliceDeclType(ctx *SliceDeclTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#arrayDeclType.
	VisitArrayDeclType(ctx *ArrayDeclTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#structDeclType.
	VisitStructDeclType(ctx *StructDeclTypeContext) interface{}

	// Visit a parse tree produced by expr_parser#structMemDecls.
	VisitStructMemDecls(ctx *StructMemDeclsContext) interface{}

	// Visit a parse tree produced by expr_parser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by expr_parser#expUnary.
	VisitExpUnary(ctx *ExpUnaryContext) interface{}

	// Visit a parse tree produced by expr_parser#expPrimaryExp.
	VisitExpPrimaryExp(ctx *ExpPrimaryExpContext) interface{}

	// Visit a parse tree produced by expr_parser#expBinary.
	VisitExpBinary(ctx *ExpBinaryContext) interface{}

	// Visit a parse tree produced by expr_parser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by expr_parser#capExp.
	VisitCapExp(ctx *CapExpContext) interface{}

	// Visit a parse tree produced by expr_parser#appendExp.
	VisitAppendExp(ctx *AppendExpContext) interface{}

	// Visit a parse tree produced by expr_parser#lenExp.
	VisitLenExp(ctx *LenExpContext) interface{}

	// Visit a parse tree produced by expr_parser#selectorExp.
	VisitSelectorExp(ctx *SelectorExpContext) interface{}

	// Visit a parse tree produced by expr_parser#funcCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by expr_parser#opExp.
	VisitOpExp(ctx *OpExpContext) interface{}

	// Visit a parse tree produced by expr_parser#indexExp.
	VisitIndexExp(ctx *IndexExpContext) interface{}

	// Visit a parse tree produced by expr_parser#literalOp.
	VisitLiteralOp(ctx *LiteralOpContext) interface{}

	// Visit a parse tree produced by expr_parser#identifierOp.
	VisitIdentifierOp(ctx *IdentifierOpContext) interface{}

	// Visit a parse tree produced by expr_parser#expressionOp.
	VisitExpressionOp(ctx *ExpressionOpContext) interface{}

	// Visit a parse tree produced by expr_parser#intLit.
	VisitIntLit(ctx *IntLitContext) interface{}

	// Visit a parse tree produced by expr_parser#floatLit.
	VisitFloatLit(ctx *FloatLitContext) interface{}

	// Visit a parse tree produced by expr_parser#rawStringLit.
	VisitRawStringLit(ctx *RawStringLitContext) interface{}

	// Visit a parse tree produced by expr_parser#interpretedStringLit.
	VisitInterpretedStringLit(ctx *InterpretedStringLitContext) interface{}

	// Visit a parse tree produced by expr_parser#runeLit.
	VisitRuneLit(ctx *RuneLitContext) interface{}

	// Visit a parse tree produced by expr_parser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by expr_parser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by expr_parser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by expr_parser#appendExpression.
	VisitAppendExpression(ctx *AppendExpressionContext) interface{}

	// Visit a parse tree produced by expr_parser#lengthExpression.
	VisitLengthExpression(ctx *LengthExpressionContext) interface{}

	// Visit a parse tree produced by expr_parser#capExpression.
	VisitCapExpression(ctx *CapExpressionContext) interface{}

	// Visit a parse tree produced by expr_parser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by expr_parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by expr_parser#printSt.
	VisitPrintSt(ctx *PrintStContext) interface{}

	// Visit a parse tree produced by expr_parser#printlnSt.
	VisitPrintlnSt(ctx *PrintlnStContext) interface{}

	// Visit a parse tree produced by expr_parser#returnSt.
	VisitReturnSt(ctx *ReturnStContext) interface{}

	// Visit a parse tree produced by expr_parser#breakSt.
	VisitBreakSt(ctx *BreakStContext) interface{}

	// Visit a parse tree produced by expr_parser#continueSt.
	VisitContinueSt(ctx *ContinueStContext) interface{}

	// Visit a parse tree produced by expr_parser#simpleSt.
	VisitSimpleSt(ctx *SimpleStContext) interface{}

	// Visit a parse tree produced by expr_parser#blockSt.
	VisitBlockSt(ctx *BlockStContext) interface{}

	// Visit a parse tree produced by expr_parser#switchSt.
	VisitSwitchSt(ctx *SwitchStContext) interface{}

	// Visit a parse tree produced by expr_parser#ifStat.
	VisitIfStat(ctx *IfStatContext) interface{}

	// Visit a parse tree produced by expr_parser#loopSt.
	VisitLoopSt(ctx *LoopStContext) interface{}

	// Visit a parse tree produced by expr_parser#typeDeclSt.
	VisitTypeDeclSt(ctx *TypeDeclStContext) interface{}

	// Visit a parse tree produced by expr_parser#varDeclSt.
	VisitVarDeclSt(ctx *VarDeclStContext) interface{}

	// Visit a parse tree produced by expr_parser#expSt.
	VisitExpSt(ctx *ExpStContext) interface{}

	// Visit a parse tree produced by expr_parser#assignSt.
	VisitAssignSt(ctx *AssignStContext) interface{}

	// Visit a parse tree produced by expr_parser#shortDecSt.
	VisitShortDecSt(ctx *ShortDecStContext) interface{}

	// Visit a parse tree produced by expr_parser#assignStat.
	VisitAssignStat(ctx *AssignStatContext) interface{}

	// Visit a parse tree produced by expr_parser#otAssignSt.
	VisitOtAssignSt(ctx *OtAssignStContext) interface{}

	// Visit a parse tree produced by expr_parser#ifSt.
	VisitIfSt(ctx *IfStContext) interface{}

	// Visit a parse tree produced by expr_parser#ifElseIfSt.
	VisitIfElseIfSt(ctx *IfElseIfStContext) interface{}

	// Visit a parse tree produced by expr_parser#ifElseBlockSt.
	VisitIfElseBlockSt(ctx *IfElseBlockStContext) interface{}

	// Visit a parse tree produced by expr_parser#ifSimpleSt.
	VisitIfSimpleSt(ctx *IfSimpleStContext) interface{}

	// Visit a parse tree produced by expr_parser#ifSimpleElseIfSt.
	VisitIfSimpleElseIfSt(ctx *IfSimpleElseIfStContext) interface{}

	// Visit a parse tree produced by expr_parser#ifSimpleElseBlockSt.
	VisitIfSimpleElseBlockSt(ctx *IfSimpleElseBlockStContext) interface{}

	// Visit a parse tree produced by expr_parser#forSt.
	VisitForSt(ctx *ForStContext) interface{}

	// Visit a parse tree produced by expr_parser#whileExprSt.
	VisitWhileExprSt(ctx *WhileExprStContext) interface{}

	// Visit a parse tree produced by expr_parser#otherForSt.
	VisitOtherForSt(ctx *OtherForStContext) interface{}

	// Visit a parse tree produced by expr_parser#simpStSwitch.
	VisitSimpStSwitch(ctx *SimpStSwitchContext) interface{}

	// Visit a parse tree produced by expr_parser#expSwitch.
	VisitExpSwitch(ctx *ExpSwitchContext) interface{}

	// Visit a parse tree produced by expr_parser#simpStSwitchNoExp.
	VisitSimpStSwitchNoExp(ctx *SimpStSwitchNoExpContext) interface{}

	// Visit a parse tree produced by expr_parser#switchDefault.
	VisitSwitchDefault(ctx *SwitchDefaultContext) interface{}

	// Visit a parse tree produced by expr_parser#expressionCaseClauseList.
	VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{}

	// Visit a parse tree produced by expr_parser#expressionCaseClause.
	VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{}

	// Visit a parse tree produced by expr_parser#caseExp.
	VisitCaseExp(ctx *CaseExpContext) interface{}

	// Visit a parse tree produced by expr_parser#defaultExp.
	VisitDefaultExp(ctx *DefaultExpContext) interface{}
}
