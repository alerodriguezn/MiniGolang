package llvm

import (
	"MiniGolang/checker"
	"MiniGolang/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"strconv"
)

type Encoder struct {
	ErrorListener *parser.CustomErrorListener
	symbolTable   *checker.SymbolTable
	modStorage    *moduleStorage
	//llvm ir
	blocks    *Stack[*ir.Block]
	functions *Stack[*Func]
	module    *ir.Module

	*parser.Baseexpr_parserVisitor
}

// get Module
func (v *Encoder) GetModule() *ir.Module {
	return v.module
}

// Str for prints
var strChar *ir.Global
var strInt *ir.Global
var strFloat *ir.Global
var strBool *ir.Global

// Predefined functions
var puts *ir.Func
var printf *ir.Func

func NewEncoder(listener *parser.CustomErrorListener) *Encoder {
	e := &Encoder{
		ErrorListener:          listener,
		module:                 ir.NewModule(),
		modStorage:             newModuleStorage(),
		blocks:                 CreateStack[*ir.Block](),
		functions:              CreateStack[*Func](),
		Baseexpr_parserVisitor: &parser.Baseexpr_parserVisitor{},
	}

	e.addPredefinedFunctions()
	e.addPredefinedGlobalDef()

	return e
}

func (v *Encoder) addPredefinedFunctions() {
	puts = v.module.NewFunc("puts", types.I32, ir.NewParam("", types.I8Ptr))
	printf = v.module.NewFunc("printf", types.I64, ir.NewParam("", types.I8Ptr), ir.NewParam("", types.I64))

	v.modStorage.addElement("puts", puts)
	v.modStorage.addElement("printf", printf)

}

func (v *Encoder) addPredefinedGlobalDef() {

	strInt = v.module.NewGlobalDef("", constant.NewCharArrayFromString("%lld\n\x00"))
	strFloat = v.module.NewGlobalDef("", constant.NewCharArrayFromString("%f\n\x00"))
	strChar = v.module.NewGlobalDef("", constant.NewCharArrayFromString("%c\n\x00"))
}

func (v *Encoder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Encoder) VisitChildren(node antlr.RuleNode) interface{} {
	var result any = nil
	for _, child := range node.GetChildren() {
		cpt, ok := child.(antlr.ParseTree)
		if !ok {
			panic("Not a ParseTree")
		}
		result = cpt.Accept(v)
	}
	return result
}

func (v *Encoder) VisitRoot(ctx *parser.RootContext) interface{} {

	allVarDecl := ctx.TopDeclarationList().AllVariableDecl()
	for _, varDecl := range allVarDecl {
		v.Visit(varDecl)
	}

	allFuncDecl := ctx.TopDeclarationList().AllFuncDecl()

	for _, funcDecl := range allFuncDecl {
		v.Visit(funcDecl)
	}

	return nil
}

func (v *Encoder) VisitSingleTypeDecl(ctx *parser.SingleTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

type Function struct {
	*ir.Func
	bodyBlock *ir.Block
}

func (v *Encoder) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {

	function := v.Visit(ctx.FuncFrontDecl()).(*Func)
	v.modStorage.addElement(ctx.FuncFrontDecl().IDENTIFIER().GetText(), function)

	//push the function to the stack
	v.functions.Push(function)
	//push the entry block to the stack
	v.blocks.Push(function.body)
	//enter scope
	v.modStorage.openScope()
	//visit the block
	v.Visit(ctx.Block())

	v.modStorage.closeScope()
	//pop the function from the stack
	_, _ = v.functions.Pop()
	//pop the entry block from the stack
	_, _ = v.blocks.Pop()

	return nil

}

func (v *Encoder) VisitFuncFrontDecl(ctx *parser.FuncFrontDeclContext) interface{} {

	retType := ctx.DeclType()
	var rt value.Value
	if retType != nil {
		rt = v.Visit(retType).(value.Value)
	}
	var params []*ir.Param
	if argum := ctx.FuncArgDecls(); argum != nil {

		for _, argument := range argum.AllSingleVarDeclNoExps() {
			varType := v.Visit(argument).(value.Value)
			for _, ident := range argument.IdentifierList().AllIDENTIFIER() {
				param := ir.NewParam(ident.GetText(), varType.Type())
				v.modStorage.addElement(argument.GetText(), varType)
				params = append(params, param)
			}

		}

	}

	var fn *ir.Func
	if rt == nil {
		fn = v.module.NewFunc(ctx.IDENTIFIER().GetText(), types.Void, params...)
	} else {
		fn = v.module.NewFunc(ctx.IDENTIFIER().GetText(), rt.Type(), params...)
	}

	//v.modStorage.addElement(ctx.IDENTIFIER().GetText(), fn)

	entryBlock := fn.NewBlock("")

	return &Func{fn, entryBlock}
}

func (v *Encoder) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitDeclType(ctx *parser.DeclTypeContext) interface{} {

	return v.VisitChildren(ctx)
}

//	func (v *Encoder) VisitNestedDeclType(ctx *parser.NestedDeclTypeContext) interface{} {
//		return v.VisitChildren(ctx)
//	}
func (v *Encoder) VisitIdDeclType(ctx *parser.IdDeclTypeContext) interface{} {
	name := ctx.IDENTIFIER().GetText()

	if name == "int" {
		a := &checker.Identifier{
			TypeId: 2,
			Value:  "int",
			LType:  types.I64,
		}
		return a
	}

	if name == "float" {
		a := &checker.Identifier{
			TypeId: 2,
			Value:  "float",
			LType:  types.Float,
		}
		return a
	}

	if name == "string" {
		a := &checker.Identifier{
			TypeId: 2,
			Value:  "string",
			LType:  types.I8Ptr,
		}
		return a
	}

	if name == "bool" {
		a := &checker.Identifier{
			TypeId: 2,
			Value:  "bool",
			LType:  types.I1,
		}
		return a
	}

	if name == "char" {
		a := &checker.Identifier{
			TypeId: 2,
			Value:  "char",
			LType:  types.I8,
		}
		return a
	}

	return nil
}

// VisitArrayDecType
//
//	func (v *Encoder) VisitSliceDecType(ctx *parser.SliceDecTypeContext) interface{} {
//		return v.VisitChildren(ctx)
//	}
func (v *Encoder) VisitArrayDecType(ctx *parser.ArrayDecTypeContext) interface{} {

	return v.VisitChildren(ctx)
}

//
//func (v *Encoder) VisitStructDecType(ctx *parser.StructDecTypeContext) interface{} {
//	return v.VisitChildren(ctx)
//}

func (v *Encoder) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {

	currFunction, _ := v.functions.Peek()

	expression := v.Visit(ctx.DeclType()).(types.Type)

	for _, identifier := range ctx.IdentifierList().AllIDENTIFIER() {
		idName := identifier.GetText()
		if currFunction == nil && currFunction.body == nil {

			global := v.module.NewGlobalDef("", constant.NewZeroInitializer(expression))
			v.modStorage.PutBack(idName, global)
		} else {
			newAlloca := currFunction.body.NewAlloca(expression)
			v.modStorage.addElement(idName, newAlloca)

		}
	}

	return expression
}

func (v *Encoder) VisitTopDeclarationList(ctx *parser.TopDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {

	return v.Visit(ctx.SingleVarDecl())
}

func (v *Encoder) VisitMVarDecl(ctx *parser.MVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitInnerVarDecls(ctx *parser.InnerVarDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVoidVarDecl(ctx *parser.VoidVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVarDeclWithType(ctx *parser.VarDeclWithTypeContext) interface{} {

	bl, _ := v.blocks.Peek()
	currFunc, _ := v.functions.Peek()

	for i, identifier := range ctx.IdentifierList().AllIDENTIFIER() {
		nameIdent := identifier.GetText()
		expr := v.Visit(ctx.ExpressionList().Expression(i)).(value.Value)

		switch e := expr.(type) {
		case constant.Constant:
			if bl != nil {
				newAlloc := currFunc.body.NewAlloca(e.Type())
				bl.NewStore(e, newAlloc)
				v.modStorage.addElement(nameIdent, newAlloc)
			} else {
				glob := v.module.NewGlobalDef("", e)
				v.modStorage.addElement(nameIdent, glob)
			}

		case *ir.InstAlloca:
			loader := bl.NewLoad(e.ElemType, expr)
			newAlloca := currFunc.body.NewAlloca(e.ElemType)
			bl.NewStore(loader, newAlloca)
			v.modStorage.addElement(nameIdent, newAlloca)

		}

	}

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVarDeclWithoutType(ctx *parser.VarDeclWithoutTypeContext) interface{} {

	var bl *ir.Block

	// get the current block if we are in a block
	if v.blocks != nil {
		bl, _ = v.blocks.Peek()
	}

	//get all identifiers
	identifiers := ctx.IdentifierList().AllIDENTIFIER()
	for i, identifier := range identifiers {
		ident := identifier.GetText()
		exp := v.Visit(ctx.ExpressionList().Expression(i)).(value.Value)

		switch val := exp.(type) {
		case constant.Constant:
			if bl != nil {
				//create a new alloca
				newAlloc := bl.NewAlloca(val.Type())
				bl.NewStore(val, newAlloc)
				v.modStorage.addElement(ident, newAlloc)
			} else {
				//create a new global if we are not in a block (Top level declaration)
				glob := v.module.NewGlobal("", val.Type())
				v.modStorage.addElement(ident, glob)
			}

		}

	}

	return nil
}
func (v *Encoder) VisitVarDeclNoExp(ctx *parser.VarDeclNoExpContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitTypeDec(ctx *parser.TypeDecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitMultiTypeDecl(ctx *parser.MultiTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVoidTypeDecl(ctx *parser.VoidTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitInnerTypeDecls(ctx *parser.InnerTypeDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSliceDeclType(ctx *parser.SliceDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitArrayDeclType(ctx *parser.ArrayDeclTypeContext) interface{} {
	arrayType := v.Visit(ctx.DeclType()).(value.Value)
	val, _ := strconv.ParseUint(ctx.INT_LIT().GetText(), 0, 64)
	return types.NewArray(val, arrayType.Type())
}

func (v *Encoder) VisitStructDeclType(ctx *parser.StructDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitStructMemDecls(ctx *parser.StructMemDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpUnary(ctx *parser.ExpUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpPrimaryExp(ctx *parser.ExpPrimaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpBinary(ctx *parser.ExpBinaryContext) interface{} {

	bl, e := v.blocks.Peek()
	if e != nil {
		panic("Error getting block")
	}

	leftExp := v.Visit(ctx.Expression(0)).(value.Value)
	rightExp := v.Visit(ctx.Expression(1)).(value.Value)

	if el, ok := rightExp.Type().(*types.PointerType); ok {
		//dereference the pointer
		rightExp = bl.NewLoad(el.ElemType, rightExp)

	}

	if el, ok := leftExp.Type().(*types.PointerType); ok {
		//dereference the pointer
		leftExp = bl.NewLoad(el.ElemType, leftExp)
	}

	switch {
	case ctx.MULT() != nil:
		return bl.NewMul(leftExp, rightExp)
	case ctx.DIV() != nil:
		return bl.NewSDiv(leftExp, rightExp)
	case ctx.MOD() != nil:
		return bl.NewSRem(leftExp, rightExp)
	case ctx.ADD() != nil:
		return bl.NewAdd(leftExp, rightExp)
	default:
		return nil
	}

}

func (v *Encoder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)

}

func (v *Encoder) VisitCapExp(ctx *parser.CapExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAppendExp(ctx *parser.AppendExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitLenExp(ctx *parser.LenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSelectorExp(ctx *parser.SelectorExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitOpExp(ctx *parser.OpExpContext) interface{} {
	return v.Visit(ctx.Operand())
}

func (v *Encoder) VisitIndexExp(ctx *parser.IndexExpContext) interface{} {

	bl, _ := v.blocks.Peek()

	indexVal := v.Visit(ctx.Index()).(value.Value)

	expression := v.Visit(ctx.PrimaryExpression()).(value.Value)

	e, _ := expression.Type().(*types.PointerType)

	if i, ok := indexVal.Type().(*types.PointerType); ok {
		bl.NewLoad(i.ElemType, indexVal)
	}

	res := bl.NewGetElementPtr(e.ElemType, expression, constant.NewInt(types.I64, 0), indexVal)

	return res
}

func (v *Encoder) VisitLiteralOp(ctx *parser.LiteralOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIdentifierOp(ctx *parser.IdentifierOpContext) interface{} {

	ident := ctx.IDENTIFIER().GetText()

	if val, ok := v.modStorage.getElement(ident); ok {
		return val.value
	}

	return nil
}

func (v *Encoder) VisitExpressionOp(ctx *parser.ExpressionOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIntLit(ctx *parser.IntLitContext) interface{} {

	num, err := strconv.ParseInt(ctx.INT_LIT().GetText(), 10, 64)
	if err != nil {
		panic("Error parsing int literal")
	}
	return constant.NewInt(types.I64, num)
}

func (v *Encoder) VisitFloatLit(ctx *parser.FloatLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitRawStringLit(ctx *parser.RawStringLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitInterpretedStringLit(ctx *parser.InterpretedStringLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitRuneLit(ctx *parser.RuneLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIndex(ctx *parser.IndexContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *Encoder) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSelector(ctx *parser.SelectorContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAppendExpression(ctx *parser.AppendExpressionContext) interface{} {

	return v.VisitChildren(ctx)

}

func (v *Encoder) VisitLengthExpression(ctx *parser.LengthExpressionContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitCapExpression(ctx *parser.CapExpressionContext) interface{} {
	return v.VisitChildren(ctx)

}

func (v *Encoder) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPrintSt(ctx *parser.PrintStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitPrintlnSt(ctx *parser.PrintlnStContext) interface{} {

	bl, _ := v.blocks.Peek()
	if expressionList := ctx.ExpressionList(); expressionList != nil {
		for _, expression := range expressionList.AllExpression() {
			exp := v.Visit(expression).(value.Value)

			switch e := exp.(type) {
			case *ir.InstAlloca:
				val := e.ElemType
				if _, ok := val.(*types.IntType); ok {
					loader := bl.NewLoad(e.ElemType, e)
					formatStrPtr := bl.NewBitCast(strInt, types.I8Ptr)
					bl.NewCall(printf, formatStrPtr, loader)
				}
			// Case when the expression is a constant
			case *constant.Int:
				formatStrPtr := bl.NewBitCast(strInt, types.I8Ptr)
				bl.NewCall(printf, formatStrPtr, exp)
			// Case when the expression is a result of an operation
			case *ir.InstFAdd, *ir.InstFSub, *ir.InstFMul, *ir.InstFDiv:
				formatStrPtr := bl.NewBitCast(strInt, types.I8Ptr)
				bl.NewCall(printf, formatStrPtr, exp)
			case *ir.InstAdd, *ir.InstSub, *ir.InstMul, *ir.InstSDiv, *ir.InstAnd, *ir.InstXor, *ir.InstOr:
				formatStrPtr := bl.NewBitCast(strInt, types.I8Ptr)
				bl.NewCall(printf, formatStrPtr, exp)

			case *ir.InstGetElementPtr:
				//array case
				if types.IsArray(e.ElemType) {
					loader := bl.NewLoad(e.ElemType, exp)
					formatStrPtr := bl.NewBitCast(strInt, types.I8Ptr)
					bl.NewCall(printf, formatStrPtr, loader)
				} else {
					bl.NewCall(puts, exp)
				}

			}

		}

	}

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitReturnSt(ctx *parser.ReturnStContext) interface{} {

	fmt.Println("Return statement")
	peekFn, _ := v.functions.Peek()
	peekBlock, _ := v.blocks.Peek()

	newBlock := peekFn.NewBlock("")
	peekBlock.NewBr(newBlock)

	k := peekFn.NewBlock("")
	k.NewBr(k)

	// Case when we are in the main function
	if peekFn.Name() == "main" {
		v.blocks.Push(k)
		peekBlock.NewRet(nil)
	}

	if expr := ctx.Expression(); expr != nil {
		expr := v.Visit(expr).(value.Value)
		//check if the expression is a pointer
		if types.IsPointer(expr.Type()) {
			v.blocks.Push(k)
			//load the value
			load := newBlock.NewLoad(peekFn.Sig.RetType, expr)
			//return the value
			return newBlock.NewRet(load)

		}

		v.blocks.Push(k)
		return newBlock.NewRet(expr)

	}

	return newBlock.NewRet(nil)
}

func (v *Encoder) VisitBreakSt(ctx *parser.BreakStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitContinueSt(ctx *parser.ContinueStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSimpleSt(ctx *parser.SimpleStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitVarDeclSt(ctx *parser.VarDeclStContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfStat(ctx *parser.IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfSt(ctx *parser.IfStContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitShortDecSt(ctx *parser.ShortDecStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitTypeDeclSt(ctx *parser.TypeDeclStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignSt(ctx *parser.AssignStContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignStat(ctx *parser.AssignStatContext) interface{} {

	bl, _ := v.blocks.Peek()

	//get expr right list
	exprListR := ctx.ExpressionList(1)

	for i, expr := range ctx.ExpressionList(0).AllExpression() {
		element := v.Visit(expr).(value.Value)
		expression := v.Visit(exprListR.Expression(i)).(value.Value)

		switch expression := expression.(type) {

		// Case when the expression is a constant
		case constant.Constant:
			bl.NewStore(expression, element)
		// Case when the expression is a pointer
		case *ir.InstAlloca:
			loader := bl.NewLoad(expression.ElemType, expression)
			bl.NewStore(loader, element)
		// Case when the expression is a result of an operation (Add, Sub, Mul, Div)
		case *ir.InstMul, *ir.InstFAdd, *ir.InstFSub, *ir.InstFMul, *ir.InstFDiv, *ir.InstSRem, *ir.InstAdd, *ir.InstSDiv, *ir.InstSub, *ir.InstFRem:
			if types.IsPointer(expression.Type()) {
				loader := bl.NewLoad(expression.Type(), expression)
				bl.NewStore(loader, element)
			} else {
				bl.NewStore(expression, element)
			}
		case *ir.InstGetElementPtr:
			switch exp := expression.Type().(type) {
			case *types.IntType:
				loader := bl.NewLoad(expression.ElemType, expression)
				bl.NewStore(loader, element)

			case *types.PointerType:
				switch i := exp.ElemType.(type) {
				case *types.IntType:
					loader := bl.NewLoad(i, expression)
					bl.NewStore(loader, element)
				case *types.ArrayType:
					loader := bl.NewLoad(i.ElemType, expression)
					bl.NewStore(loader, element)
				}

			default:
				bl.NewStore(expression, element)

			}

		default:
			fmt.Println("Error in AssignStat")
		}

	}

	return nil
}

func (v *Encoder) VisitSwitchSt(ctx *parser.SwitchStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitSimpStSwitch(ctx *parser.SimpStSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpressionCaseClauseList(ctx *parser.ExpressionCaseClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpressionCaseClause(ctx *parser.ExpressionCaseClauseContext) interface{} {
	return v.VisitChildren(ctx)

}

func (v *Encoder) VisitCaseExp(ctx *parser.CaseExpContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitDefaultExp(ctx *parser.DefaultExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitForSt(ctx *parser.ForStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitWhileExprSt(ctx *parser.WhileExprStContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitOtherForSt(ctx *parser.OtherForStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitBlockSt(ctx *parser.BlockStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitLoopSt(ctx *parser.LoopStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitExpSt(ctx *parser.ExpStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfElseIfSt(ctx *parser.IfElseIfStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfElseBlockSt(ctx *parser.IfElseBlockStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfSimpleSt(ctx *parser.IfSimpleStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfSimpleElseIfSt(ctx *parser.IfSimpleElseIfStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Encoder) VisitIfSimpleElseBlockSt(ctx *parser.IfSimpleElseBlockStContext) interface{} {
	return v.VisitChildren(ctx)
}
