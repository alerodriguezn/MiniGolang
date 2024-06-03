package checker
//
//import (
//	"MiniGolang/parser"
//	"fmt"
//	"github.com/antlr4-go/antlr/v4"
//	"strconv"
//)
//
//type Checker struct {
//	Table         *SymbolTable
//	ErrorList     []string //ToDo change to error struct
//	ErrorListener *parser.CustomErrorListener
//	//My parser
//	*parser.Baseexpr_parserVisitor
//}
//
//func NewChecker(errorListener *parser.CustomErrorListener) *Checker {
//	return &Checker{
//		Table:                  NewSymbolTable(),
//		ErrorListener:          errorListener,
//		ErrorList:              make([]string, 0),
//		Baseexpr_parserVisitor: &parser.Baseexpr_parserVisitor{},
//	}
//
//}
//
//func (v *Checker) Visit(tree antlr.ParseTree) interface{} {
//	return tree.Accept(v)
//}
//
//func (v *Checker) VisitChildren(node antlr.RuleNode) interface{} {
//	var result any = nil
//	for _, child := range node.GetChildren() {
//		cpt, ok := child.(antlr.ParseTree)
//		if !ok {
//			panic("Not a ParseTree")
//		}
//		result = cpt.Accept(v)
//	}
//	return result
//}
//
//func (v *Checker) VisitRoot(ctx *parser.RootContext) interface{} {
//
//	typs := ctx.TopDeclarationList().AllTypeDecl()
//
//	for _, typ := range typs {
//		v.Visit(typ)
//	}
//
//	//functions := ctx.TopDeclarationList().AllFuncDecl()
//	//
//	//for _, function := range functions {
//	//	rtype := DeclTypeInfo{
//	//		Identifier: "",
//	//		IsSlice:    false,
//	//		IsStruct:   false,
//	//		IsArray:    false,
//	//	}
//	//	//Get the return type of the function
//	//	returnType := function.FuncFrontDecl().DeclType()
//	//	if returnType != nil {
//	//		rtype = v.Visit(returnType).(DeclTypeInfo)
//	//	}
//	//
//	//	// Save the arguments of the function
//	//	var members []*Identifier
//	//	if args := function.FuncFrontDecl().FuncArgDecls(); args != nil {
//	//		for _, arg := range args.AllSingleVarDeclNoExps() {
//	//			argType := v.Visit(arg.DeclType()).(DeclTypeInfo)
//	//
//	//			ids := arg.IdentifierList().AllIDENTIFIER()
//	//			for _, id := range ids {
//	//
//	//				isValidType := v.Table.isValidType(argType.Identifier)
//	//				if !isValidType {
//	//					v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", argType.Identifier))
//	//					v.ErrorListener.SyntaxError(nil, arg.GetStart(), arg.GetStart().GetLine(), arg.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", argType.Identifier), nil)
//	//				}
//	//				variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), argType.Identifier)
//	//
//	//				//add to members
//	//				members = append(members, variable)
//	//
//	//			}
//	//
//	//		}
//	//	}
//	//	// Add the function to the symbol table
//	//
//	//	isValidFnType := v.Table.isValidType(rtype.Identifier)
//	//	if !isValidFnType {
//	//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", rtype.Identifier))
//	//		v.ErrorListener.SyntaxError(nil, function.GetStart(), function.GetStart().GetLine(), function.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", rtype.Identifier), nil)
//	//
//	//	}
//	//	fn := v.Table.NewFun(function.FuncFrontDecl().IDENTIFIER().GetSymbol(), function.FuncFrontDecl().IDENTIFIER().GetText(), rtype.Identifier, members)
//	//	err := v.Table.AddElement(fn)
//	//	if err != nil {
//	//		v.ErrorList = append(v.ErrorList, err.Error())
//	//		v.ErrorListener.SyntaxError(nil, function.GetStart(), function.GetStart().GetLine(), function.GetStart().GetColumn(), err.Error(), nil)
//	//
//	//	}
//	//
//	//}
//
//	vars := ctx.TopDeclarationList().AllVariableDecl()
//	for _, variable := range vars {
//		v.Visit(variable)
//
//	}
//
//	fns := ctx.TopDeclarationList().AllFuncDecl()
//	for _, function := range fns {
//		v.Visit(function)
//
//	}
//
//	return nil
//
//	//return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitSingleTypeDecl(ctx *parser.SingleTypeDeclContext) interface{} {
//
//	id := ctx.IDENTIFIER()
//
//	//Add the new type to the symbol table
//	v.Table.AddType(id.GetText())
//
//	_type := v.Visit(ctx.DeclType()).(DeclTypeInfo)
//
//	// Add the type to the symbol table
//	isValid := v.Table.isValidType(_type.Identifier)
//	if !isValid {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", _type.Identifier))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", _type.Identifier), nil)
//
//	}
//	typ := v.Table.NewType(id.GetSymbol(), id.GetText(), _type.Identifier, _type.IsSlice)
//	err := v.Table.AddElement(typ)
//	if err != nil {
//		v.ErrorList = append(v.ErrorList, err.Error())
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), nil)
//	}
//
//	return nil
//}
//
//func (v *Checker) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
//
//	funcFrontDecl := ctx.FuncFrontDecl()
//	rtype := DeclTypeInfo{
//		Identifier: "",
//		IsSlice:    false,
//		IsStruct:   false,
//		IsArray:    false,
//	}
//	//Get the return type of the function
//	returnType := funcFrontDecl.DeclType()
//	if returnType != nil {
//		rtype = v.Visit(returnType).(DeclTypeInfo)
//	}
//
//	// Save the arguments of the function
//	var members []*Identifier
//	if args := funcFrontDecl.FuncArgDecls(); args != nil {
//		for _, arg := range args.AllSingleVarDeclNoExps() {
//			argType := v.Visit(arg.DeclType()).(DeclTypeInfo)
//
//			ids := arg.IdentifierList().AllIDENTIFIER()
//			for _, id := range ids {
//
//				isValidType := v.Table.isValidType(argType.Identifier)
//				if !isValidType {
//					v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", argType.Identifier))
//					v.ErrorListener.SyntaxError(nil, arg.GetStart(), arg.GetStart().GetLine(), arg.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", argType.Identifier), nil)
//				}
//				variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), argType.Identifier)
//
//				//add to members
//				members = append(members, variable)
//
//			}
//
//		}
//
//	}
//
//	isValidFnType := v.Table.isValidType(rtype.Identifier)
//	if !isValidFnType {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", rtype.Identifier))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", rtype.Identifier), nil)
//
//	}
//	fn := v.Table.NewFun(funcFrontDecl.IDENTIFIER().GetSymbol(), funcFrontDecl.IDENTIFIER().GetText(), rtype.Identifier, members)
//	err := v.Table.AddElement(fn)
//	if err != nil {
//		v.ErrorList = append(v.ErrorList, err.Error())
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), nil)
//	}
//
//	// Visit the block
//	v.Table.openScope()
//	v.Visit(ctx.Block())
//	v.Table.closeScope()
//
//	return nil
//}
//func ComputeFac(x int) {
//}
//
//func (v *Checker) VisitFuncFrontDecl(ctx *parser.FuncFrontDeclContext) interface{} {
//	// ----
//	rtype := DeclTypeInfo{
//		Identifier: "",
//		IsSlice:    false,
//		IsStruct:   false,
//		IsArray:    false,
//	}
//	returnType := ctx.DeclType()
//	if returnType != nil {
//		rtype = v.Visit(returnType).(DeclTypeInfo)
//	}
//	// -----
//	v.Table.openScope()
//	var members []*Identifier
//	if args := ctx.FuncArgDecls(); args != nil {
//		for _, arg := range args.AllSingleVarDeclNoExps() {
//			argType := v.Visit(arg.DeclType()).(DeclTypeInfo)
//
//			ids := arg.IdentifierList().AllIDENTIFIER()
//			for _, id := range ids {
//
//				isValidType := v.Table.isValidType(argType.Identifier)
//				if !isValidType {
//					v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", argType.Identifier))
//					v.ErrorListener.SyntaxError(nil, arg.GetStart(), arg.GetStart().GetLine(), arg.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", argType.Identifier), nil)
//				}
//				variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), argType.Identifier)
//
//				//add to members
//				members = append(members, variable)
//
//			}
//
//		}
//	}
//	// Add the function to the symbol table
//
//	isValidFnType := v.Table.isValidType(rtype.Identifier)
//	if !isValidFnType {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", rtype.Identifier))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", rtype.Identifier), nil)
//	}
//
//	//args := ctx.FuncArgDecls()
//	//if args != nil {
//	//	v.Visit(args)
//	//}
//
//	return rtype
//}
//
//func (v *Checker) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//type DeclTypeInfo struct {
//	Identifier string
//	IsSlice    bool
//	IsStruct   bool
//	IsArray    bool
//}
//
//func (v *Checker) VisitDeclType(ctx *parser.DeclTypeContext) interface{} {
//
//	res := DeclTypeInfo{
//		Identifier: "",
//		IsSlice:    false,
//		IsStruct:   false,
//		IsArray:    false,
//	}
//
//	if parentheses := ctx.LPAREN(); parentheses != nil {
//		//Get the DeclType in the middle of the parentheses
//		child, exist := ctx.GetChild(1).(*parser.DeclTypeContext)
//		if exist {
//			return v.VisitDeclType(child)
//		}
//	}
//
//	switch {
//	case ctx.IDENTIFIER() != nil:
//		res.Identifier = ctx.IDENTIFIER().GetText()
//	case ctx.SliceDeclType() != nil:
//		slice := ctx.SliceDeclType()
//		sliceType := v.Visit(slice.DeclType()).(DeclTypeInfo)
//		if ident := slice.DeclType().IDENTIFIER(); ident != nil {
//			res.Identifier = sliceType.Identifier
//			res.IsSlice = true
//		}
//	case ctx.ArrayDeclType() != nil:
//		array := v.Visit(ctx.ArrayDeclType()).(*Identifier)
//		res.Identifier = array.LitType
//		res.IsArray = true
//
//	case ctx.StructDeclType() != nil:
//
//		structure := ctx.StructDeclType()
//		structElement := v.Visit(structure).(*Identifier)
//		err := v.Table.AddElement(structElement)
//		if err != nil {
//			v.ErrorList = append(v.ErrorList, err.Error())
//			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), nil)
//
//		}
//		s := structElement.Value.(string)
//		res.Identifier = s
//		res.IsStruct = true
//
//	}
//
//	return res
//}
//
//func (v *Checker) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
//
//	// Get the type
//	typ := v.Visit(ctx.DeclType()).(DeclTypeInfo)
//	idList := ctx.IdentifierList().AllIDENTIFIER()
//	// Add all var to the symbol table
//	for _, id := range idList {
//
//		isValid := v.Table.isValidType(typ.Identifier)
//		if !isValid {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", typ.Identifier))
//			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("'%s' is not a valid type", typ.Identifier), nil)
//		}
//		variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), typ.Identifier)
//		err := v.Table.AddElement(variable)
//		if err != nil {
//			v.ErrorList = append(v.ErrorList, err.Error())
//			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), nil)
//
//		}
//	}
//
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitTopDeclarationList(ctx *parser.TopDeclarationListContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitVarDecl(ctx *parser.VarDeclContext) interface{} {
//
//	// Check if SingleVarDecl is VarDeclNoExp or VarDeclWithType
//
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitMVarDecl(ctx *parser.MVarDeclContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitInnerVarDecls(ctx *parser.InnerVarDeclsContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitVoidVarDecl(ctx *parser.VoidVarDeclContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitVarDeclWithType(ctx *parser.VarDeclWithTypeContext) interface{} {
//	// Get all Indentifiers
//
//	identifiers := ctx.IdentifierList().AllIDENTIFIER()
//	// Get the type of all identifiers
//	idsType := v.Visit(ctx.DeclType()).(DeclTypeInfo)
//
//	expressionList := ctx.ExpressionList().AllExpression()
//
//	for i, id := range identifiers {
//		exp := expressionList[i]
//		expType := v.Visit(exp).(*Identifier)
//		if expType.LitType != idsType.Identifier {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("variable '%s' has type %s, expected %s", id.GetText(), expType.LitType, idsType.Identifier))
//			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("variable '%s' has type %s, expected %s", id.GetText(), expType, idsType.Identifier), nil)
//		}
//		isValidType := v.Table.isValidType(idsType.Identifier)
//		if !isValidType {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", idsType.Identifier))
//			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("'%s' is not a valid type", idsType.Identifier), nil)
//		}
//		variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), idsType.Identifier)
//		err := v.Table.AddElement(variable)
//		if err != nil {
//			v.ErrorList = append(v.ErrorList, err.Error())
//			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), err.Error(), nil)
//
//		}
//	}
//	return nil
//}
//
//func (v *Checker) VisitVarDeclWithoutType(ctx *parser.VarDeclWithoutTypeContext) interface{} {
//
//	identifiers := ctx.IdentifierList().AllIDENTIFIER()
//	// Get the type of all identifiers
//	exprs := ctx.ExpressionList().AllExpression()
//
//	for i, id := range identifiers {
//		exp := exprs[i]
//		expType := v.Visit(exp).(string)
//		isValidType := v.Table.isValidType(expType)
//		if !isValidType {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", expType))
//			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("'%s' is not a valid type", expType), nil)
//		}
//		variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), expType)
//		err := v.Table.AddElement(variable)
//		if err != nil {
//			v.ErrorList = append(v.ErrorList, err.Error())
//			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), err.Error(), nil)
//
//		}
//	}
//
//	return nil
//}
//func (v *Checker) VisitVarDeclNoExp(ctx *parser.VarDeclNoExpContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitTypeDec(ctx *parser.TypeDecContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitMultiTypeDecl(ctx *parser.MultiTypeDeclContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitVoidTypeDecl(ctx *parser.VoidTypeDeclContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitInnerTypeDecls(ctx *parser.InnerTypeDeclsContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitSliceDeclType(ctx *parser.SliceDeclTypeContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitArrayDeclType(ctx *parser.ArrayDeclTypeContext) interface{} {
//
//	id := v.Visit(ctx.DeclType()).(DeclTypeInfo)
//	//convert arraySize to int
//	arraySize, _ := strconv.Atoi(ctx.INT_LIT().GetText())
//
//	array := &Identifier{
//		Value:   id.Identifier,
//		Level:   v.Table.CurrentLevel,
//		LitType: id.Identifier,
//		Size:    arraySize,
//		list:    nil,
//	}
//
//	return array
//}
//
//func (v *Checker) VisitStructDeclType(ctx *parser.StructDeclTypeContext) interface{} {
//
//	newStruct := v.Table.NewStruct(ctx.GetStart(), "struct")
//
//	members := ctx.StructMemDecls()
//	if members != nil {
//		varDecls := members.AllSingleVarDeclNoExps()
//		for _, decl := range varDecls {
//			ids := decl.IdentifierList().AllIDENTIFIER()
//			_type := v.Visit(decl.DeclType()).(DeclTypeInfo)
//			for _, id := range ids {
//				isValid := v.Table.isValidType(_type.Identifier)
//				if !isValid {
//					v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' is not a valid type", _type.Identifier))
//					v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("'%s' is not a valid type", _type.Identifier), nil)
//				}
//				variable := v.Table.NewVar(id.GetSymbol(), id.GetText(), _type.Identifier)
//				newStruct.list = append(newStruct.list, variable)
//			}
//
//		}
//	}
//
//	return newStruct
//}
//
//func (v *Checker) VisitStructMemDecls(ctx *parser.StructMemDeclsContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitExpUnary(ctx *parser.ExpUnaryContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitExpPrimaryExp(ctx *parser.ExpPrimaryExpContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitExpBinary(ctx *parser.ExpBinaryContext) interface{} {
//
//	if ctx.NOT_EQ() != nil || ctx.EQUALS() != nil || ctx.GREATER_THAN() != nil || ctx.LESS_THAN() != nil || ctx.LESS_THAN_OR_EQUALS() != nil || ctx.GREATER_THAN_OR_EQUALS() != nil {
//		leftResult := v.Visit(ctx.Expression(0)).(*Identifier)  // Visita la expresión izquierda
//		rightResult := v.Visit(ctx.Expression(1)).(*Identifier) // Visita la expresión derecha
//		if leftResult.LitType != rightResult.LitType {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("Mismatch error: '%s' has type %s, giving %s", leftResult.Value, leftResult.LitType, rightResult.LitType))
//			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("Mismatch error: '%s' has type %s, giving %s", leftResult.Value, leftResult.LitType, rightResult.LitType), nil)
//		}
//		return "bool"
//	}
//
//	if ctx.ADD() != nil || ctx.SUB() != nil || ctx.PIPE() != nil || ctx.CARET() != nil || ctx.AND() != nil || ctx.RSHIFT() != nil || ctx.LSHIFT() != nil || ctx.ANDNOT() != nil {
//		leftResult := v.Visit(ctx.Expression(0)).(*Identifier)  // Visita la expresión izquierda
//		rightResult := v.Visit(ctx.Expression(1)).(*Identifier) // Visita la expresión derecha
//
//		if leftResult.LitType != rightResult.LitType {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("Mismatch error: '%s' has type %s, giving %s", leftResult.Value, leftResult.LitType, rightResult.LitType))
//			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("Mismatch error: '%s' has type %s, and %s has type %s", ctx.Expression(0).GetText(), leftResult.LitType, ctx.Expression(1).GetText(), rightResult.LitType), nil)
//		}
//
//		if leftResult == nil {
//			return rightResult
//		}
//		return leftResult
//	}
//
//	if ctx.MULT() != nil || ctx.DIV() != nil || ctx.MOD() != nil {
//		leftResult := v.Visit(ctx.Expression(0)).(*Identifier)  // Visita la expresión izquierda
//		rightResult := v.Visit(ctx.Expression(1)).(*Identifier) // Visita la expresión derecha
//		if leftResult.LitType != rightResult.LitType {
//			//Display mismatch error message
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("Mismatch error: '%s' has type %s, giving %s", leftResult.Value, leftResult.LitType, rightResult.LitType))
//			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("Mismatch error: '%s' has type %s, giving %s", leftResult.Value, leftResult.LitType, rightResult.LitType), nil)
//		}
//
//		if leftResult == nil {
//			return rightResult
//		}
//
//		return leftResult
//	}
//
//	if ctx.LOG_OR() != nil || ctx.LOG_AND() != nil {
//		leftResult := v.Visit(ctx.Expression(0)).(*Identifier)  // Visita la expresión izquierda
//		rightResult := v.Visit(ctx.Expression(1)).(*Identifier) // Visita la expresión derecha
//		if leftResult.LitType != rightResult.LitType {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("Mismatch error: '%s' has type %s, expected %s", leftResult.Value, leftResult.LitType, rightResult.LitType))
//			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("Mismatch error: '%s' has type %s, giving %s", leftResult.Value, leftResult.LitType, rightResult.LitType), nil)
//		}
//		return "bool"
//
//	}
//
//	return nil
//
//}
//
//func (v *Checker) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
//
//	allExps := ctx.AllExpression()
//
//	var list []*Identifier
//	for _, exp := range allExps {
//		typ := v.Visit(exp)
//		if typ != nil {
//			list = append(list, typ.(*Identifier))
//		}
//	}
//
//	return list
//
//}
//
//func (v *Checker) VisitCapExp(ctx *parser.CapExpContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitAppendExp(ctx *parser.AppendExpContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitLenExp(ctx *parser.LenExpContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitSelectorExp(ctx *parser.SelectorExpContext) interface{} {
//
//	fmt.Println("Entrooooooo")
//	// Get the struct
//	structure := v.Visit(ctx.PrimaryExpression()).(*Identifier)
//	fmt.Println(structure.Value)
//
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
//
//	function := v.Visit(ctx.PrimaryExpression()).(*Identifier)
//	arguments := v.Visit(ctx.Arguments().ExpressionList(0)).([]*Identifier)
//
//	//Check if the function exists
//	if function == nil {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("function '%s' not defined", ctx.PrimaryExpression().GetText()))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("function '%s' not defined", ctx.PrimaryExpression().GetText()), nil)
//		return nil
//	}
//
//	if len(arguments) != len(function.list) {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("function '%s' expects %d arguments, %d given", function.Value, len(function.list), len(arguments)))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("function '%s' expects %d arguments, %d given", function.Value, len(function.list), len(arguments)), nil)
//	}
//
//	if len(arguments) <= len(function.list) {
//		for i, arg := range arguments {
//			if arg.LitType != function.list[i].LitType {
//				v.ErrorList = append(v.ErrorList, fmt.Sprintf("argument %d of function '%s' has type %s, expected %s", i, function.Value, arg.LitType, function.list[i].LitType))
//				v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("argument %d of function '%s' has type %s, expected %s", i, function.Value, arg.LitType, function.list[i].LitType), nil)
//
//			}
//		}
//	}
//
//	return function
//}
//
//func (v *Checker) VisitOpExp(ctx *parser.OpExpContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitIndexExp(ctx *parser.IndexExpContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitLiteralOp(ctx *parser.LiteralOpContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitIdentifierOp(ctx *parser.IdentifierOpContext) interface{} {
//	if ctx.IDENTIFIER() != nil {
//		id := ctx.IDENTIFIER()
//		// Search the variable in the symbol table
//		variable := v.Table.search(id.GetText())
//		if variable == nil {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("variable '%s' not defined", id.GetText()))
//			v.ErrorListener.SyntaxError(nil, id.GetSymbol(), id.GetSymbol().GetLine(), id.GetSymbol().GetColumn(), fmt.Sprintf("variable '%s' not defined", id.GetText()), nil)
//		}
//		return variable
//	}
//
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitExpressionOp(ctx *parser.ExpressionOpContext) interface{} {
//
//	res := v.Visit(ctx.Expression())
//	return res
//}
//
//func (v *Checker) VisitIntLit(ctx *parser.IntLitContext) interface{} {
//
//	return &Identifier{
//		LitType: "int",
//	}
//}
//
//func (v *Checker) VisitFloatLit(ctx *parser.FloatLitContext) interface{} {
//	return &Identifier{
//		LitType: "float",
//	}
//}
//
//func (v *Checker) VisitRawStringLit(ctx *parser.RawStringLitContext) interface{} {
//	return &Identifier{
//		LitType: "string",
//	}
//}
//
//func (v *Checker) VisitInterpretedStringLit(ctx *parser.InterpretedStringLitContext) interface{} {
//	return &Identifier{
//		LitType: "string",
//	}
//}
//
//func (v *Checker) VisitRuneLit(ctx *parser.RuneLitContext) interface{} {
//	return &Identifier{
//		LitType: "char",
//	}
//}
//
//func (v *Checker) VisitIndex(ctx *parser.IndexContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitSelector(ctx *parser.SelectorContext) interface{} {
//
//	// Get the struct
//
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitAppendExpression(ctx *parser.AppendExpressionContext) interface{} {
//
//	sliceExpression := ctx.Expression(0)
//	valueExpression := ctx.Expression(1)
//
//	slice := v.Visit(sliceExpression).(*Identifier)
//	value := v.Visit(valueExpression).(*Identifier)
//
//	if slice.LitType != value.LitType {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("append expects %s, %s given", slice.LitType, value.LitType))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("append expects %s, %s given", slice.LitType, value.LitType), nil)
//	}
//
//	return slice
//
//}
//
//func (v *Checker) VisitLengthExpression(ctx *parser.LengthExpressionContext) interface{} {
//
//	fmt.Println("Entrando a length")
//	id := v.Visit(ctx.Expression()).(*Identifier)
//
//	if id.LitType != "string" && id.LitType != "array" && id.LitType != "slice" {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("len expects string, array or slice, %s given", id.LitType))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("len expects string, array or slice, %s given", id.LitType), nil)
//	}
//
//	return &Identifier{
//		LitType: "int",
//	}
//}
//
//func (v *Checker) VisitCapExpression(ctx *parser.CapExpressionContext) interface{} {
//	expression := ctx.Expression()
//	id := v.Visit(expression).(*Identifier)
//
//	//Check if the expression is a slice or an array
//	if id.LitType != "slice" && id.LitType != "array" {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("cap expects slice or array, %s given", id.LitType))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("cap expects slice or array, %s given", id.LitType), nil)
//		return nil
//	}
//
//	return &Identifier{
//		LitType: "int",
//	}
//
//}
//
//func (v *Checker) VisitStatementList(ctx *parser.StatementListContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitBlock(ctx *parser.BlockContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitPrintSt(ctx *parser.PrintStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitPrintlnSt(ctx *parser.PrintlnStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitReturnSt(ctx *parser.ReturnStContext) interface{} {
//
//	//Get Current Function
//	currentFunction := v.Table.getCurrentFunction()
//	if currentFunction == nil {
//		v.ErrorList = append(v.ErrorList, "return statement outside of function")
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "return statement outside of function", nil)
//		return nil
//	}
//
//	currFnType := currentFunction.LitType
//
//	expr := ctx.Expression()
//	exprType := v.Visit(expr).(*Identifier)
//
//	if currFnType != exprType.LitType {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("function '%s' has return type %s, expected %s", currentFunction.Value, exprType.LitType, currFnType))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("function '%s' has return type %s, expected %s", currentFunction.Value, exprType.LitType, currFnType), nil)
//	}
//
//	return nil
//}
//
//func (v *Checker) VisitBreakSt(ctx *parser.BreakStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitContinueSt(ctx *parser.ContinueStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitSimpleSt(ctx *parser.SimpleStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitVarDeclSt(ctx *parser.VarDeclStContext) interface{} {
//
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitIfStat(ctx *parser.IfStatContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitIfSt(ctx *parser.IfStContext) interface{} {
//
//	v.Table.openScope()
//	defer v.Table.closeScope()
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitShortDecSt(ctx *parser.ShortDecStContext) interface{} {
//
//	leftExpression := ctx.ExpressionList(0).AllExpression()
//	rightExpression := ctx.ExpressionList(1).AllExpression()
//
//	//Check if the len of both expressions is the same
//	if len(leftExpression) != len(rightExpression) {
//		v.ErrorList = append(v.ErrorList, "The number of expressions in the left and right side of the short declaration must be the same")
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "The number of expressions in the left and right side of the short declaration must be the same", nil)
//		return nil
//	}
//
//	for i, leftExp := range leftExpression {
//		exp := rightExpression[i]
//
//		r := v.Visit(exp).(*Identifier)
//
//		variable := v.Table.NewVar(leftExp.GetStart(), leftExp.GetText(), r.LitType)
//		err := v.Table.AddElement(variable)
//		if err != nil {
//			v.ErrorList = append(v.ErrorList, err.Error())
//			v.ErrorListener.SyntaxError(nil, leftExp.GetStart(), leftExp.GetStart().GetLine(), leftExp.GetStart().GetColumn(), err.Error(), nil)
//		}
//		if variable.LitType != r.LitType {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("variable '%s' has type %d, expected %d", leftExp.GetText(), variable.LitType, r))
//			v.ErrorListener.SyntaxError(nil, leftExp.GetStart(), leftExp.GetStart().GetLine(), leftExp.GetStart().GetColumn(), fmt.Sprintf("variable '%s' has type %d, expected %d", leftExp.GetText(), variable.LitType, r), nil)
//		}
//
//	}
//	return nil
//}
//
//func (v *Checker) VisitTypeDeclSt(ctx *parser.TypeDeclStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitAssignSt(ctx *parser.AssignStContext) interface{} {
//
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitAssignStat(ctx *parser.AssignStatContext) interface{} {
//
//	r := ctx.ExpressionList(1).AllExpression()
//	fmt.Printf("Right: %v\n", r)
//	l := ctx.ExpressionList(0).AllExpression()
//
//	if len(r) != len(l) {
//		v.ErrorList = append(v.ErrorList, "The number of expressions in the left and right side of the assignment must be the same")
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "The number of expressions in the left and right side of the assignment must be the same", nil)
//		return nil
//	}
//
//	for i, leftExp := range l {
//		exp := r[i]
//
//		s := v.Visit(exp).(*Identifier)
//		variable := v.Table.search(leftExp.GetText())
//
//		if variable == nil {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("variable '%s' not defined", leftExp.GetText()))
//			v.ErrorListener.SyntaxError(nil, leftExp.GetStart(), leftExp.GetStart().GetLine(), leftExp.GetStart().GetColumn(), fmt.Sprintf("variable '%s' not defined", leftExp.GetText()), nil)
//			return nil
//		}
//		if variable.LitType != s.LitType {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("variable '%s' has type %s, expected %s", leftExp.GetText(), variable.LitType, s.LitType))
//			v.ErrorListener.SyntaxError(nil, leftExp.GetStart(), leftExp.GetStart().GetLine(), leftExp.GetStart().GetColumn(), fmt.Sprintf("variable '%s' has type %s, expected %s", leftExp.GetText(), variable.LitType, s.LitType), nil)
//		}
//	}
//
//	return nil
//}
//
//func (v *Checker) VisitSwitchSt(ctx *parser.SwitchStContext) interface{} {
//	v.Table.openScope()
//	defer v.Table.closeScope()
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitSimpStSwitch(ctx *parser.SimpStSwitchContext) interface{} {
//
//	st := ctx.SimpleStatement()
//	if st != nil {
//		v.Visit(st)
//	}
//
//	v.Visit(ctx.Expression())
//
//	cases := ctx.ExpressionCaseClauseList()
//
//	return v.Visit(cases)
//}
//
//func (v *Checker) VisitExpressionCaseClauseList(ctx *parser.ExpressionCaseClauseListContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitExpressionCaseClause(ctx *parser.ExpressionCaseClauseContext) interface{} {
//	v.Visit(ctx.ExpressionSwitchCase())
//
//	v.Table.openScope()
//	defer v.Table.closeScope()
//	return v.VisitChildren(ctx)
//
//}
//
//func (v *Checker) VisitCaseExp(ctx *parser.CaseExpContext) interface{} {
//
//	//Get current Type
//	//currentType := v.Table.getCurrentType()
//
//	return nil
//}
//
//func (v *Checker) VisitDefaultExp(ctx *parser.DefaultExpContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitForSt(ctx *parser.ForStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitWhileExprSt(ctx *parser.WhileExprStContext) interface{} {
//
//	expr := v.Visit(ctx.Expression()).(string)
//
//	if expr != "bool" {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("while expects bool, %s given", expr))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("while expects bool, %s given", expr), nil)
//	}
//	block := ctx.Block()
//	return v.Visit(block)
//}
//
//func (v *Checker) VisitOtherForSt(ctx *parser.OtherForStContext) interface{} {
//
//	// Get Both Simple Statements
//	simpleSt1 := ctx.SimpleStatement(0)
//	simpleSt2 := ctx.SimpleStatement(1)
//	// Get Expr
//	expr := ctx.Expression()
//
//	//Visit Both Simple Statements
//	v.Visit(simpleSt1)
//	v.Visit(simpleSt2)
//
//	//Visit Expression for get the type
//	t := v.Visit(expr).(string)
//	fmt.Printf("Type: %s\n", t)
//	if t != "bool" {
//		v.ErrorList = append(v.ErrorList, fmt.Sprintf("for expects bool, %s given", t))
//		v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("for expects bool, %s given", t), nil)
//	}
//
//	block := ctx.Block()
//
//	return v.Visit(block)
//}
//
//func (v *Checker) VisitBlockSt(ctx *parser.BlockStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitLoopSt(ctx *parser.LoopStContext) interface{} {
//	return v.VisitChildren(ctx)
//}
//
//func (v *Checker) VisitExpSt(ctx *parser.ExpStContext) interface{} {
//
//	if ctx.DECREMENT() != nil || ctx.INCREMENT() != nil {
//		expr := ctx.Expression()
//		id := v.Visit(expr).(*Identifier)
//		if id.LitType != "int" {
//			v.ErrorList = append(v.ErrorList, fmt.Sprintf("'%s' has type %s, expected int", id.Value, id.LitType))
//			v.ErrorListener.SyntaxError(nil, ctx.GetStart(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("'%s' has type %s, expected int", id.Value, id.LitType), nil)
//
//		}
//		return nil
//	}
//
//	return v.VisitChildren(ctx)
//}
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
