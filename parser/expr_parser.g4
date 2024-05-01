parser grammar expr_parser;

options {
  tokenVocab = expr_lexer;
}



root : PACKAGE IDENTIFIER SEMICOLON topDeclarationList;

topDeclarationList : (variableDecl | typeDecl | funcDecl)*;

//variableDecl : VAR (singleVarDecl | LPAREN innerVarDecls? RPAREN) SEMICOLON ;

variableDecl: VAR singleVarDecl SEMICOLON                #varDecl
            | VAR LPAREN innerVarDecls RPAREN SEMICOLON  #mVarDecl
            | VAR LPAREN RPAREN SEMICOLON                #voidVarDecl
            ;

innerVarDecls: singleVarDecl SEMICOLON (singleVarDecl SEMICOLON)*;


singleVarDecl : identifierList declType ASSIGN expressionList #varDeclWithType
			| identifierList ASSIGN expressionList #varDeclWithoutType
			| singleVarDeclNoExps #varDeclNoExp
			;

singleVarDeclNoExps: identifierList declType ;


typeDecl: TYPE singleTypeDecl SEMICOLON #typeDec
        | TYPE LPAREN innerTypeDecls RPAREN SEMICOLON #multiTypeDecl
        | TYPE LPAREN RPAREN SEMICOLON #voidTypeDecl
        ;

innerTypeDecls: singleTypeDecl SEMICOLON (singleTypeDecl ';')*;

singleTypeDecl : IDENTIFIER declType ;

funcDecl: funcFrontDecl block SEMICOLON;

//Alternative for function declaration
//funcFrontDecl : FUNCTION IDENTIFIER LPAREN funcArgDecls? RPAREN declType?;
funcFrontDecl : FUNCTION IDENTIFIER LPAREN funcArgDecls RPAREN declType
              | FUNCTION IDENTIFIER LPAREN funcArgDecls RPAREN
              | FUNCTION IDENTIFIER LPAREN RPAREN declType
              | FUNCTION IDENTIFIER LPAREN RPAREN
              | FUNCTION IDENTIFIER declType ;


funcArgDecls : (singleVarDeclNoExps COMMA)* singleVarDeclNoExps ;

declType : LPAREN declType RPAREN
         | IDENTIFIER 
         | sliceDeclType
         | arrayDeclType
         | structDeclType
         |  ;

sliceDeclType : LBRACK RBRACK declType ;
arrayDeclType : LBRACK INT_LIT RBRACK declType ;
structDeclType : STRUCT LBRACE structMemDecls? RBRACE ;

structMemDecls	: singleVarDeclNoExps SEMICOLON (singleVarDeclNoExps SEMICOLON)*;
identifierList	: IDENTIFIER (COMMA IDENTIFIER)* ;

expression : primaryExpression #expPrimaryExp
           | expression ( MULT | DIV | MOD | LSHIFT | RSHIFT | AND | ANDNOT | ADD | SUB | PIPE | CARET | EQUALS | NOT_EQ | LESS_THAN | LESS_THAN_OR_EQUALS | GREATER_THAN | GREATER_THAN_OR_EQUALS | LOG_AND | LOG_OR ) expression #expBinary
           | ( ADD | SUB | LOG_NOT | CARET ) expression  #expUnary
           ;

expressionList : expression (COMMA expression)* ;

primaryExpression : operand #opExp
			| primaryExpression selector #selectorExp
			| primaryExpression index #indexExp
			| primaryExpression arguments #funcCall
			| appendExpression  #appendExp
			| lengthExpression #lenExp
			| capExpression  #capExp
            ;
operand	: literal									
		| IDENTIFIER 
		| LPAREN expression RPAREN;

literal :
          INT_LIT #intLit
        | FLOAT_LIT #floatLit
        | RAW_STRING_LIT  #rawStringLit
        | INTERPRETED_STRING_LIT #interpretedStringLit
        | RUNE_LIT #runeLit
        ;

index : LBRACK expression RBRACK ;

arguments : LPAREN expressionList* RPAREN ;

selector : DOT IDENTIFIER ;

appendExpression : APPEND LPAREN expression COMMA expression RPAREN ;

lengthExpression : LENGHT LPAREN expression RPAREN ;

capExpression : CAP LPAREN expression RPAREN ;

statementList : statement* ;

block : LBRACE statementList RBRACE ;

statement : PRINT LPAREN expressionList? RPAREN SEMICOLON #printSt
          | PRINTLN LPAREN expressionList? RPAREN SEMICOLON #printlnSt
          | RETURN expression? SEMICOLON #returnSt
          | BREAK SEMICOLON    #breakSt
          | CONTINUE SEMICOLON #continueSt
          | simpleStatement SEMICOLON #simpleSt
          | block SEMICOLON #blockSt
          | switch SEMICOLON #switchSt
          | ifStatement SEMICOLON #ifStat
          | loop SEMICOLON #loopSt
          | typeDecl SEMICOLON #typeDeclSt
          | variableDecl  #varDeclSt
          ;


simpleStatement : expression (INCREMENT | DECREMENT)? #expSt
                | assignmentStatement #assignSt
                | expressionList SHORT_DEC expressionList #shortDecSt
                ;

assignmentStatement : expressionList ASSIGN expressionList #assignStat
                    | expression ( PLUS_ASSIGN | AND_ASSIGN | MINUS_ASSIGN | OR_ASSIGN | TIMES_ASSIGN| XOR_ASSIGN | LEFT_SHIFT_ASSIGN| RIGHT_SHIFT_ASSIGN | AND_NOT_ASSIGN | MOD_ASSIGN| DIVIDE_ASSIGN ) expression #otAssignSt
                    ;


ifStatement : IF expression block #ifSt
            | IF expression block ELSE ifStatement #ifElseIfSt
            | IF expression block ELSE block #ifElseBlockSt
            | IF simpleStatement SEMICOLON expression block #ifSimpleSt
            | IF simpleStatement SEMICOLON expression block ELSE ifStatement #ifSimpleElseIfSt
            | IF simpleStatement SEMICOLON expression block ELSE block #ifSimpleElseBlockSt
            ;

loop : FOR block #forSt
     | FOR expression block #whileExprSt
     | FOR simpleStatement SEMICOLON expression SEMICOLON simpleStatement block #forExprSt
     | FOR simpleStatement SEMICOLON SEMICOLON simpleStatement block #otherForSt
     ;


switch
    : SWITCH simpleStatement SEMICOLON expression LBRACE expressionCaseClauseList RBRACE #simpStSwitch
    | SWITCH expression LBRACE expressionCaseClauseList RBRACE #expSwitch
    | SWITCH simpleStatement SEMICOLON LBRACE expressionCaseClauseList RBRACE #simpStSwitchNoExp
    | SWITCH LBRACE expressionCaseClauseList RBRACE #switchDefault
    ;


expressionCaseClauseList: (expressionCaseClause expressionCaseClauseList)*
                        ;

expressionCaseClause : expressionSwitchCase COLON statementList ;

expressionSwitchCase : CASE expressionList #caseExp
                     | DEFAULT #defaultExp
                     ;