lexer grammar expr_lexer;

// LEXER RULES

// ---------- Keywords ------------------

PACKAGE   : 'package';
VAR       : 'var';
TYPE      : 'type';
FUNCTION  : 'func';
STRUCT    : 'struct';
APPEND    : 'append';
LENGHT    : 'len';
CAP       : 'cap';
PRINT     : 'print';
RETURN    : 'return';
BREAK     : 'break';
CONTINUE  : 'continue';
IF        : 'if';
ELSE      : 'else';
FOR       : 'for';
SWITH     : 'switch';
CASE      : 'case';
DEFAULT   : 'default';

// ---------- Symbols ------------------

ASSIGN    : '=' ;
SHORT_DEC : ':=';
SEMICOLON : ';' ;
COLON     : ':' ;
DOT       : '.' ;
LPAREN    : '(' ;
DPAREN    : ')' ;
RBRACK    : ']' ;
LBRACK    : '[' ;
LBRACE    : '{' ;
RBRACE    : '}' ;




// ---------- Operators  ------------------

ADD       : '+' ;
SUB       : '-' ;
MULT      : '*' ;
DIV       : '/' ;
MOD       : '%' ;
LSHIFT    : '<<';
RSHIFT    : '>>';
AND       : '&' ;
ANDNOT    : '&^';
PIPE      : '|';
CARET     : '^';
EQUALS    : '==';
NOT_EQ    : '!=';
LESS_THAN : '<';
LESS_THAN_OR_EQUALS: '<=';
GREATER_THAN: '>';
GREATER_THAN_OR_EQUALS: '>=';
LOG_AND: '&&';
LOG_OR: '||';
LOG_NOT: '!';
INCREMENT : '++';
DECREMENT : '--';
PLUS_ASSIGN: '+=';
AND_ASSIGN: '&=';
MINUS_ASSIGN: '-=';
OR_ASSIGN: '|=';
TIMES_ASSIGN: '*=';
XOR_ASSIGN: '^=';
LEFT_SHIFT_ASSIGN: '<<=';
RIGHT_SHIFT_ASSIGN: '>>=';
AND_NOT_ASSIGN: '&^=';
MOD_ASSIGN: '%=';
DIVIDE_ASSIGN: '/=';