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

// ---------- INT, DECIMAL, BIN, HEX LITERAL   ------------------
UNDERSCORE : '_';
DECIMAL_DIGIT : '0'..'9' ;
BINARY_DIGIT  : '0' | '1' ;
OCTAL_DIGIT   : '0'..'7';
HEX_DIGIT     : '0'..'9' | 'A'..'F' | 'a'..'f';

DECIMAL_DIGITS : DECIMAL_DIGIT (UNDERSCORE? DECIMAL_DIGIT)* ;
BINARY_DIGITS  : BINARY_DIGIT (UNDERSCORE? BINARY_DIGIT)* ;
OCTAL_DIGITS   : OCTAL_DIGIT (UNDERSCORE? OCTAL_DIGIT)* ;
HEX_DIGITS     : HEX_DIGIT (UNDERSCORE? HEX_DIGIT)* ;

INT_LIT     : DECIMAL_LIT | BINARY_LIT | OCTAL_LIT | HEX_LIT ;
DECIMAL_LIT : '0' | ('1' .. '9') '_'? DECIMAL_DIGITS?;
BINARY_LIT  : '0' ('B' | 'b') '_'? BINARY_DIGITS;
OCTAL_LIT   : '0' ('O' | 'o') '_'? OCTAL_DIGITS;
HEX_LIT     : '0' ('X' | 'x') '_'? HEX_DIGITS;


// ---------- Identifier   ------------------

UNICODE_LETTER : [\p{L}];
UNICODE_DIGIT : [\p{Nd}];

LETTER : UNICODE_LETTER | UNDERSCORE ;
IDENTIFIER : LETTER (LETTER | UNICODE_DIGIT)* ;

// ---------- Float Literal     -----------------------

DECIMAL_EXPONENT : [eE] [+\-]? DECIMAL_DIGITS ;
DECIMAL_FLOAT_LIT : DECIMAL_DIGITS '.' (DECIMAL_DIGITS? DECIMAL_EXPONENT? | DECIMAL_EXPONENT)
    | DECIMAL_DIGITS DECIMAL_EXPONENT
    | '.' DECIMAL_DIGITS (DECIMAL_EXPONENT)?
    ;
HEX_FLOAT_LIT
    : '0' [xX] HEX_MANTISSA HEX_EXPONENT
    ;
HEX_MANTISSA
    : UNDERSCORE? HEX_DIGITS '.' (HEX_DIGITS? | UNDERSCORE)
    | UNDERSCORE? HEX_DIGITS
    | '.' HEX_DIGITS
    ;
HEX_EXPONENT
    : [pP] [+\-]? DECIMAL_DIGITS
    ;
FLOAT_LIT : DECIMAL_FLOAT_LIT | HEX_FLOAT_LIT ;

// ---------- Char Literal (RUNELITERAL)    -----------------------


// ---------- Strings Literal (RAWSTRINGLITERAL,INTERPRETEDSTRINGLITERAL )   -----------------------