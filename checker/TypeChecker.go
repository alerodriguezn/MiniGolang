package checker

import (
	"MiniGolang/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type Checker struct {
	Table         *SymbolTable
	ErrorList     []string //ToDo change to error struct
	ErrorListener *parser.CustomErrorListener
	//My parser
	*parser.Baseexpr_parserVisitor
}

func NewChecker(errorListener *parser.CustomErrorListener) *Checker {
	return &Checker{
		Table:                  NewSymbolTable(),
		ErrorListener:          errorListener,
		ErrorList:              make([]string, 0),
		Baseexpr_parserVisitor: &parser.Baseexpr_parserVisitor{},
	}

}

func (v *Checker) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Checker) VisitChildren(node antlr.RuleNode) interface{} {
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

func (v *Checker) VisitRoot(ctx *parser.RootContext) interface{} {

	typs := ctx.TopDeclarationList().AllTypeDecl()

	for _, typ := range typs {
		v.Visit(typ)
	}

	functions := ctx.TopDeclarationList().AllFuncDecl()

	for _, function := range functions {
		rtype := DeclTypeInfo{
			Identifier: "",
			IsSlice:    false,
			IsStruct:   false,
			IsArray:    false,
		}
		//Get the return type of the function
		returnType := function.FuncFrontDecl().DeclType()
		if returnType != nil {
			rtype = v.Visit(returnType).(DeclTypeInfo)
		}

		// Save the arguments of the function
		var members []*Identifier
		if args := function.FuncFrontDecl().FuncArgDecls(); args != nil {
			for _, arg := range args.AllSingleVarDeclNoExps() {
				argType := v.Visit(arg.DeclType()).(DeclTypeInfo)

				ids := arg.IdentifierList().AllIDENTIFIER()
				for _, id := range ids {

					isValidType := v.Table.isValidType(argType.Identifier)
					if !isValidType {
						v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", argType.Identifier))
						v.ErrorListener.SyntaxError(nil, arg.GetStart(), arg.GetStart().GetLine(), arg.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", argType.Identifier), nil)
					}
					variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), argType.Identifier)

					//add to members
					members = append(members, variable)

				}

			}
		}
		// Add the function to the symbol table

		isValidFnType := v.Table.isValidType(rtype.Identifier)
		if !isValidFnType {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", rtype.Identifier))
			v.ErrorListener.SyntaxError(nil, function.GetStart(), function.GetStart().GetLine(), function.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", rtype.Identifier), nil)

		}
		fn := v.Table.NewFun(function.FuncFrontDecl().IDENTIFIER().GetSymbol(), function.FuncFrontDecl().IDENTIFIER().GetText(), rtype.Identifier, members)
		err := v.Table.AddElement(fn)
		if err != nil {
			v.ErrorList = append(v.ErrorList, err.Error())
			v.ErrorListener.SyntaxError(nil, function.GetStart(), function.GetStart().GetLine(), function.GetStart().GetColumn(), err.Error(), nil)

		}

	}

	vars := ctx.TopDeclarationList().AllTypeDecl()
	for _, variable := range vars {
		v.Visit(variable)

	}

	fns := ctx.TopDeclarationList().AllFuncDecl()
	for _, function := range fns {
		v.Visit(function)

	}

	return nil

	//return v.VisitChildren(ctx)
}

func (v *Checker) VisitSingleTypeDecl(ctx *parser.SingleTypeDeclContext) interface{} {

	id := ctx.IDENTIFIER()

	//Add the new type to the symbol table
	v.Table.AddType(id.GetText())

	_type := v.Visit(ctx.DeclType()).(DeclTypeInfo)

	// Add the type to the symbol table
	isValid := v.Table.isValidType(_type.Identifier)
	if !isValid {
		v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", _type.Identifier))
		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", _type.Identifier), nil)

	}
	typ := v.Table.NewType(id.GetSymbol(), id.GetText(), _type.Identifier, _type.IsSlice)
	err := v.Table.AddElement(typ)
	if err != nil {
		v.ErrorList = append(v.ErrorList, err.Error())
		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), nil)
	}

	return nil
}

func (v *Checker) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {

	funcName := ctx.FuncFrontDecl().IDENTIFIER().GetText()
	val := v.Table.search(funcName)
	if val == nil {
		v.ErrorList = append(v.ErrorList, fmt.Sprintf("function '%s' already defined in the current scope", funcName))
		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("function '%s' already defined in the current scope", funcName), nil)
	}

	v.Table.openScope()
	v.Visit(ctx.Block())
	v.Table.closeScope()

	return nil
}

func (v *Checker) VisitFuncFrontDecl(ctx *parser.FuncFrontDeclContext) interface{} {

	v.Table.openScope()
	args := ctx.FuncArgDecls()
	if args != nil {
		v.Visit(args)
	}

	return nil
}

func (v *Checker) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

type DeclTypeInfo struct {
	Identifier string
	IsSlice    bool
	IsStruct   bool
	IsArray    bool
}

func (v *Checker) VisitDeclType(ctx *parser.DeclTypeContext) interface{} {

	res := DeclTypeInfo{
		Identifier: "",
		IsSlice:    false,
		IsStruct:   false,
		IsArray:    false,
	}

	if parentheses := ctx.LPAREN(); parentheses != nil {
		//Get the DeclType in the middle of the parentheses
		child, exist := ctx.GetChild(1).(*parser.DeclTypeContext)
		if exist {
			return v.VisitDeclType(child)
		}
	}

	switch {
	case ctx.IDENTIFIER() != nil:
		res.Identifier = ctx.IDENTIFIER().GetText()
	case ctx.SliceDeclType() != nil:
		slice := ctx.SliceDeclType()
		sliceType := v.Visit(slice.DeclType()).(DeclTypeInfo)
		if ident := slice.DeclType().IDENTIFIER(); ident != nil {
			res.Identifier = sliceType.Identifier
			res.IsSlice = true
		}
	case ctx.ArrayDeclType() != nil:
		array := ctx.ArrayDeclType()
		if ident := array.DeclType().IDENTIFIER(); ident != nil {
			res.Identifier = ident.GetText()
			res.IsArray = true
		}
	case ctx.StructDeclType() != nil:

		structure := ctx.StructDeclType()
		structElement := v.Visit(structure).(*Identifier)
		err := v.Table.AddElement(structElement)
		if err != nil {
			v.ErrorList = append(v.ErrorList, err.Error())
			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), nil)

		}
		s := structElement.Value.(string)
		res.Identifier = s
		res.IsStruct = true

	}

	return res
}

func (v *Checker) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {

	// Get the type
	typ := v.Visit(ctx.DeclType()).(DeclTypeInfo)
	idList := ctx.IdentifierList().AllIDENTIFIER()
	// Add all var to the symbol table
	for _, id := range idList {

		isValid := v.Table.isValidType(typ.Identifier)
		if !isValid {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", typ.Identifier))
			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", typ.Identifier), nil)
		}
		variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), typ.Identifier)
		err := v.Table.AddElement(variable)
		if err != nil {
			v.ErrorList = append(v.ErrorList, err.Error())
			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), nil)

		}
	}

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitTopDeclarationList(ctx *parser.TopDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {

	// Check if SingleVarDecl is VarDeclNoExp or VarDeclWithType

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitMVarDecl(ctx *parser.MVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitInnerVarDecls(ctx *parser.InnerVarDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitVoidVarDecl(ctx *parser.VoidVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitVarDeclWithType(ctx *parser.VarDeclWithTypeContext) interface{} {
	// Get all Indentifiers

	identifiers := ctx.IdentifierList().AllIDENTIFIER()
	// Get the type of all identifiers
	idsType := v.Visit(ctx.DeclType()).(DeclTypeInfo)

	expressionList := ctx.ExpressionList().AllExpression()

	for i, id := range identifiers {
		exp := expressionList[i]
		expType := v.Visit(exp)
		if expType != idsType.Identifier {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("variable '%s' has type %s, expected %s", id.GetText(), expType, idsType.Identifier))
			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("variable '%s' has type %s, expected %s", id.GetText(), expType, idsType.Identifier), nil)
		}
		isValidType := v.Table.isValidType(idsType.Identifier)
		if !isValidType {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", idsType.Identifier))
			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("'%s' is not a valid type", idsType.Identifier), nil)
		}
		variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), idsType.Identifier)
		err := v.Table.AddElement(variable)
		if err != nil {
			v.ErrorList = append(v.ErrorList, err.Error())
			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), err.Error(), nil)

		}
	}
	return nil
}

func (v *Checker) VisitVarDeclWithoutType(ctx *parser.VarDeclWithoutTypeContext) interface{} {

	identifiers := ctx.IdentifierList().AllIDENTIFIER()
	// Get the type of all identifiers
	exprs := ctx.ExpressionList().AllExpression()

	for i, id := range identifiers {
		exp := exprs[i]
		expType := v.Visit(exp).(string)
		isValidType := v.Table.isValidType(expType)
		if !isValidType {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", expType))
			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("'%s' is not a valid type", expType), nil)
		}
		variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), expType)
		err := v.Table.AddElement(variable)
		if err != nil {
			v.ErrorList = append(v.ErrorList, err.Error())
			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), err.Error(), nil)

		}
	}

	return nil
}
func (v *Checker) VisitVarDeclNoExp(ctx *parser.VarDeclNoExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitTypeDec(ctx *parser.TypeDecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitMultiTypeDecl(ctx *parser.MultiTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitVoidTypeDecl(ctx *parser.VoidTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitInnerTypeDecls(ctx *parser.InnerTypeDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSliceDeclType(ctx *parser.SliceDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitArrayDeclType(ctx *parser.ArrayDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitStructDeclType(ctx *parser.StructDeclTypeContext) interface{} {
	newStruct := v.Table.NewStruct(ctx.GetStart(), "Str")

	members := ctx.StructMemDecls()
	if members != nil {
		varDecls := members.AllSingleVarDeclNoExps()
		for _, decl := range varDecls {
			ids := decl.IdentifierList().AllIDENTIFIER()
			_type := v.Visit(decl.DeclType()).(DeclTypeInfo)
			for _, id := range ids {
				isValid := v.Table.isValidType(_type.Identifier)
				if !isValid {
					v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", _type.Identifier))
					v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("'%s' is not a valid type", _type.Identifier), nil)
				}
				variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), _type.Identifier)
				newStruct.list = append(newStruct.list, variable)
			}

		}
	}

	return newStruct
}

func (v *Checker) VisitStructMemDecls(ctx *parser.StructMemDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpUnary(ctx *parser.ExpUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpPrimaryExp(ctx *parser.ExpPrimaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpBinary(ctx *parser.ExpBinaryContext) interface{} {

	if ctx.NOT_EQ() != nil || ctx.EQUALS() != nil || ctx.GREATER_THAN() != nil || ctx.LESS_THAN() != nil || ctx.LESS_THAN_OR_EQUALS() != nil || ctx.GREATER_THAN_OR_EQUALS() != nil {
		leftResult := v.Visit(ctx.Expression(0)).(*Identifier)  // Visita la expresión izquierda
		rightResult := v.Visit(ctx.Expression(1)).(*Identifier) // Visita la expresión derecha
		if leftResult.LitType != rightResult.LitType {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("Mismatch error: '%s' has type %s, expected %s", leftResult.Value, leftResult.LitType, rightResult.LitType))
			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("Mismatch error: '%s' has type %s, expected %s", leftResult.Value, leftResult.LitType, rightResult.LitType), nil)
		}
		return "bool"
	}

	if ctx.ADD() != nil || ctx.SUB() != nil || ctx.PIPE() != nil || ctx.CARET() != nil || ctx.AND() != nil || ctx.RSHIFT() != nil || ctx.LSHIFT() != nil || ctx.ANDNOT() != nil {
		leftResult := v.Visit(ctx.Expression(0)).(*Identifier)  // Visita la expresión izquierda
		rightResult := v.Visit(ctx.Expression(1)).(*Identifier) // Visita la expresión derecha
		if leftResult.LitType != rightResult.LitType {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("Mismatch error: '%s' has type %s, expected %s", leftResult.Value, leftResult.LitType, rightResult.LitType))
			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("Mismatch error: '%s' has type %s, expected %s", leftResult.Value, leftResult.LitType, rightResult.LitType), nil)
		}

		if leftResult == nil {
			return rightResult.LitType
		}
		return leftResult.LitType
	}

	if ctx.MULT() != nil || ctx.DIV() != nil || ctx.MOD() != nil {
		leftResult := v.Visit(ctx.Expression(0)).(*Identifier)  // Visita la expresión izquierda
		rightResult := v.Visit(ctx.Expression(1)).(*Identifier) // Visita la expresión derecha
		if leftResult.LitType != rightResult.LitType {
			//Display mismatch error message
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("Mismatch error: '%s' has type %s, expected %s", leftResult.Value, leftResult.LitType, rightResult.LitType))
			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("Mismatch error: '%s' has type %s, expected %s", leftResult.Value, leftResult.LitType, rightResult.LitType), nil)
		}

		if leftResult == nil {
			return rightResult.LitType
		}

		return leftResult.LitType
	}

	if ctx.LOG_OR() != nil || ctx.LOG_AND() != nil {
		leftResult := v.Visit(ctx.Expression(0)).(*Identifier)  // Visita la expresión izquierda
		rightResult := v.Visit(ctx.Expression(1)).(*Identifier) // Visita la expresión derecha
		if leftResult.LitType != rightResult.LitType {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("Mismatch error: '%s' has type %s, expected %s", leftResult.Value, leftResult.LitType, rightResult.LitType))
			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("Mismatch error: '%s' has type %s, expected %s", leftResult.Value, leftResult.LitType, rightResult.LitType), nil)
		}
		return "bool"

	}

	return nil

}

func (v *Checker) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {

	allExps := ctx.AllExpression()

	var list []*Identifier
	for _, exp := range allExps {
		typ := v.Visit(exp)
		if typ != nil {
			list = append(list, typ.(*Identifier))
		}
	}

	return list

}

func (v *Checker) VisitCapExp(ctx *parser.CapExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAppendExp(ctx *parser.AppendExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitLenExp(ctx *parser.LenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSelectorExp(ctx *parser.SelectorExpContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {

	function := v.Visit(ctx.PrimaryExpression()).(*Identifier)

	expression := ctx.Arguments().ExpressionList(0)

	arguments := v.Visit(expression).([]*Identifier)

	if len(arguments) != len(function.list) {
		v.ErrorList = append(v.ErrorList, fmt.Sprintf("function '%s' expects %d arguments, %d given", function.Value, len(function.list), len(arguments)))
		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("function '%s' expects %d arguments, %d given", function.Value, len(function.list), len(arguments)), nil)
	}

	if len(arguments) <= len(function.list) {
		for i, arg := range arguments {
			if arg.LitType != function.list[i].LitType {
				v.ErrorList = append(v.ErrorList, fmt.Sprintf("argument %d of function '%s' has type %d, expected %d", i, function.Value, arg.LitType, function.list[i].LitType))
				v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("argument %d of function '%s' has type %d, expected %d", i, function.Value, arg.LitType, function.list[i].LitType), nil)
			}
		}
	}

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitOpExp(ctx *parser.OpExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIndexExp(ctx *parser.IndexExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitOperand(ctx *parser.OperandContext) interface{} {

	if ctx.IDENTIFIER() != nil {
		id := ctx.IDENTIFIER()
		// Search the variable in the symbol table
		variable := v.Table.search(id.GetText())
		if variable == nil {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("variable '%s' not defined", id.GetText()))
			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("variable '%s' not defined", id.GetText()), nil)
		}
		return variable
	}

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIntLit(ctx *parser.IntLitContext) interface{} {
	return "int"
}

func (v *Checker) VisitFloatLit(ctx *parser.FloatLitContext) interface{} {
	return "float"
}

func (v *Checker) VisitRawStringLit(ctx *parser.RawStringLitContext) interface{} {
	return "string"
}

func (v *Checker) VisitInterpretedStringLit(ctx *parser.InterpretedStringLitContext) interface{} {
	return "string"
}

func (v *Checker) VisitRuneLit(ctx *parser.RuneLitContext) interface{} {
	return "char"
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

func (v *Checker) VisitPrintSt(ctx *parser.PrintStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitPrintlnSt(ctx *parser.PrintlnStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitReturnSt(ctx *parser.ReturnStContext) interface{} {

	//Get Current Function
	currentFunction := v.Table.getCurrentFunction()
	if currentFunction == nil {
		v.ErrorList = append(v.ErrorList, "return statement outside of function")
		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "return statement outside of function", nil)
		return nil
	}

	currFnType := currentFunction.LitType

	expr := ctx.Expression()
	exprType := v.Visit(expr).(*Identifier)

	if currFnType != exprType.LitType {
		v.ErrorList = append(v.ErrorList, fmt.Sprintf("function '%s' has return type %s, expected %s", currentFunction.Value, exprType.LitType, currFnType))
		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("function '%s' has return type %s, expected %s", currentFunction.Value, exprType.LitType, currFnType), nil)
	}

	return nil
}

func (v *Checker) VisitBreakSt(ctx *parser.BreakStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitContinueSt(ctx *parser.ContinueStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSimpleSt(ctx *parser.SimpleStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitVarDeclSt(ctx *parser.VarDeclStContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfStat(ctx *parser.IfStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitIfSt(ctx *parser.IfStContext) interface{} {

	v.Table.openScope()
	defer v.Table.closeScope()
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitShortDecSt(ctx *parser.ShortDecStContext) interface{} {

	leftExpression := ctx.ExpressionList(0).AllExpression()
	rightExpression := ctx.ExpressionList(1).AllExpression()

	//Check if the len of both expressions is the same
	if len(leftExpression) != len(rightExpression) {
		v.ErrorList = append(v.ErrorList, "The number of expressions in the left and right side of the short declaration must be the same")
		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "The number of expressions in the left and right side of the short declaration must be the same", nil)
		return nil
	}

	for i, leftExp := range leftExpression {
		exp := rightExpression[i]

		r := v.Visit(exp).(string)

		variable := v.Table.NewVar(leftExp.GetStart(), leftExp.GetText(), r)
		err := v.Table.AddElement(variable)
		if err != nil {
			v.ErrorList = append(v.ErrorList, err.Error())
			v.ErrorListener.SyntaxError(nil, leftExp.GetStart(), leftExp.GetStart().GetLine(), leftExp.GetStart().GetColumn(), err.Error(), nil)
		}
		if variable.LitType != r {
			v.ErrorList = append(v.ErrorList, fmt.Sprintf("variable '%s' has type %d, expected %d", leftExp.GetText(), variable.LitType, r))
			v.ErrorListener.SyntaxError(nil, leftExp.GetStart(), leftExp.GetStart().GetLine(), leftExp.GetStart().GetColumn(), fmt.Sprintf("variable '%s' has type %d, expected %d", leftExp.GetText(), variable.LitType, r), nil)
		}

	}
	return nil
}

func (v *Checker) VisitTypeDeclSt(ctx *parser.TypeDeclStContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignSt(ctx *parser.AssignStContext) interface{} {

	return v.VisitChildren(ctx)
}

func (v *Checker) VisitAssignStat(ctx *parser.AssignStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSwitchSt(ctx *parser.SwitchStContext) interface{} {
	v.Table.openScope()
	defer v.Table.closeScope()
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitSimpStSwitch(ctx *parser.SimpStSwitchContext) interface{} {

	st := ctx.SimpleStatement()
	if st != nil {
		v.Visit(st)
	}

	v.Visit(ctx.Expression())

	cases := ctx.ExpressionCaseClauseList()

	return v.Visit(cases)
}

func (v *Checker) VisitExpressionCaseClauseList(ctx *parser.ExpressionCaseClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Checker) VisitExpressionCaseClause(ctx *parser.ExpressionCaseClauseContext) interface{} {
	v.Visit(ctx.ExpressionSwitchCase())

	v.Table.openScope()
	defer v.Table.closeScope()
	return v.VisitChildren(ctx)

}

func (v *Checker) VisitCaseExp(ctx *parser.CaseExpContext) interface{} {

	//Get current Type
	//currentType := v.Table.getCurrentType()

	return nil
}

func (v *Checker) VisitDefaultExp(ctx *parser.DefaultExpContext) interface{} {
	return v.VisitChildren(ctx)
}

//func (v *Checker) VisitBlockSt(ctx *parser.BlockStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//

//
//func (v *Checker) VisitLoopSt(ctx *parser.LoopStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//

//
//func (v *Checker) VisitExpSt(ctx *parser.ExpStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//

//
//func (v *Checker) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//

//func (v *Checker) VisitIfElseIfSt(ctx *parser.IfElseIfStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitIfElseBlockSt(ctx *parser.IfElseBlockStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitIfSimpleSt(ctx *parser.IfSimpleStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitIfSimpleElseIfSt(ctx *parser.IfSimpleElseIfStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitIfSimpleElseBlockSt(ctx *parser.IfSimpleElseBlockStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitForSt(ctx *parser.ForStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitWhileExprSt(ctx *parser.WhileExprStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitForExprSt(ctx *parser.ForExprStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitOtherForSt(ctx *parser.OtherForStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitSwitch(ctx *parser.SwitchContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//

//
//func (v *Checker) VisitExpressionSwitchCase(ctx *parser.ExpressionSwitchCaseContext) interface{} {
//	return v.VisitChildren(ctx)
//}
