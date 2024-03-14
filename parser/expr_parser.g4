parser grammar expr_parser;

options {
  tokenVocab = expr_lexer;
}



root : PACKAGE IDENTIFIER SEMICOLON topDeclarationList;


topDeclarationList : (variableDecl | typeDecl | funcDecl)*;

variableDecl : VAR (singleVarDecl | LPAREN innerVarDecls? RPAREN) SEMICOLON ;

innerVarDecls : (singleVarDecl SEMICOLON)+ ;


singleVarDecl : identifierList declType ASSIGN expressionList
			| identifierList ASSIGN expressionList
			| singleVarDeclNoExps;

singleVarDeclNoExps: identifierList declType ;

typeDecl : TYPE (singleTypeDecl | LPAREN innerTypeDecls? RPAREN) SEMICOLON ;

innerTypeDecls: singleTypeDecl SEMICOLON (singleTypeDecl ';')*;

singleTypeDecl : IDENTIFIER declType ;

funcDecl: funcFrontDecl block SEMICOLON;

//Alternative for function declaration
funcFrontDecl : FUNCTION IDENTIFIER LPAREN funcArgDecls RPAREN declType
              | FUNCTION IDENTIFIER LPAREN funcArgDecls RPAREN
              | FUNCTION IDENTIFIER LPAREN RPAREN declType
              | FUNCTION IDENTIFIER LPAREN RPAREN ;


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

expression : primaryExpression
           | expression ( MULT | DIV | MOD | LSHIFT | RSHIFT | AND | ANDNOT | ADD | SUB | PIPE | CARET | EQUALS | NOT_EQ | LESS_THAN | LESS_THAN_OR_EQUALS | GREATER_THAN | GREATER_THAN_OR_EQUALS | LOG_AND | LOG_OR ) expression
           | ( ADD | SUB | LOG_NOT | CARET ) expression ;

expressionList : expression (COMMA expression)* ;

primaryExpression : operand								
			| primaryExpression selector 
			| primaryExpression index 
			| primaryExpression arguments 
			| appendExpression 
			| lengthExpression
			| capExpression ;

operand	: literal									
		| IDENTIFIER 
		| LPAREN expression RPAREN;

literal : INT_LIT | FLOAT_LIT |  RAW_STRING_LIT | INTERPRETED_STRING_LIT;

index : LBRACK expression RBRACK ;

arguments : LPAREN expressionList? RPAREN ;

selector : DOT IDENTIFIER ;

appendExpression : APPEND LPAREN expression COMMA expression RPAREN ;

lengthExpression : LENGHT LPAREN expression RPAREN ;

capExpression : CAP LPAREN expression RPAREN ;

statementList : statement* ;

block : LBRACE statementList RBRACE ;

statement : PRINT LPAREN expressionList? RPAREN COMMA 
          | PRINTLN LPAREN expressionList? RPAREN COMMA 
          | RETURN expression? COMMA
          | BREAK COMMA
          | CONTINUE COMMA
          | simpleStatement COMMA
          | block COMMA
          | switch COMMA
          | ifStatement COMMA
          | loop COMMA
          | typeDecl
          | variableDecl ;


simpleStatement : expression (INCREMENT | DECREMENT)? 
                | assignmentStatement 
                | expressionList SHORT_DEC expressionList ;

assignmentStatement : expressionList ASSIGN expressionList 
                    | expression ( PLUS_ASSIGN | AND_ASSIGN | MINUS_ASSIGN | OR_ASSIGN | TIMES_ASSIGN| XOR_ASSIGN | LEFT_SHIFT_ASSIGN| RIGHT_SHIFT_ASSIGN | AND_NOT_ASSIGN | MOD_ASSIGN| DIVIDE_ASSIGN ) expression ;


ifStatement : IF expression block 
            | IF expression block ELSE ifStatement 
            | IF expression block ELSE block 
            | IF simpleStatement SEMICOLON expression block 
            | IF simpleStatement SEMICOLON expression block ELSE ifStatement
            | IF simpleStatement SEMICOLON expression block ELSE block ;

loop : FOR block 
     | FOR expression block 
     | FOR simpleStatement SEMICOLON expression SEMICOLON simpleStatement block
     | FOR simpleStatement SEMICOLON SEMICOLON simpleStatement block ;

switch : SWITCH simpleStatement? SEMICOLON? expression? LBRACE expressionCaseClauseList RBRACE ;


expressionCaseClauseList : expressionCaseClause* ;

expressionCaseClause : expressionSwitchCase ':' statementList ;

expressionSwitchCase : CASE expressionList 
                     | DEFAULT ;