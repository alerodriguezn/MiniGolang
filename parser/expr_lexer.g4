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
fragment UNDERSCORE : '_';
fragment DECIMAL_DIGIT : '0'..'9' ;
fragment BINARY_DIGIT  : '0' | '1' ;
fragment OCTAL_DIGIT   : '0'..'7';
fragment HEX_DIGIT : [0-9a-fA-F];

fragment DECIMAL_DIGITS : DECIMAL_DIGIT (UNDERSCORE? DECIMAL_DIGIT)* ;
fragment BINARY_DIGITS  : BINARY_DIGIT (UNDERSCORE? BINARY_DIGIT)* ;
fragment OCTAL_DIGITS   : OCTAL_DIGIT (UNDERSCORE? OCTAL_DIGIT)* ;
fragment HEX_DIGITS     : HEX_DIGIT (UNDERSCORE? HEX_DIGIT)* ;

INT_LIT     : DECIMAL_LIT | BINARY_LIT | OCTAL_LIT | HEX_LIT ;
fragment DECIMAL_LIT : '0' | ('1' .. '9') '_'? DECIMAL_DIGITS?;
fragment BINARY_LIT  : '0' ('B' | 'b') '_'? BINARY_DIGITS;
fragment OCTAL_LIT   : '0' ('O' | 'o') '_'? OCTAL_DIGITS;
fragment HEX_LIT     : '0' ('X' | 'x') '_'? HEX_DIGITS;


// ---------- Identifier   ------------------

fragment UNICODE_LETTER : [\p{L}];
fragment UNICODE_DIGIT : [\p{Nd}];

fragment LETTER : UNICODE_LETTER | UNDERSCORE ;
IDENTIFIER : LETTER (LETTER | UNICODE_DIGIT)* ;

// ---------- Float Literal     -----------------------

fragment DECIMAL_EXPONENT : [eE] [+\-]? DECIMAL_DIGITS ;
fragment DECIMAL_FLOAT_LIT : DECIMAL_DIGITS '.' (DECIMAL_DIGITS? DECIMAL_EXPONENT? | DECIMAL_EXPONENT)
    | DECIMAL_DIGITS DECIMAL_EXPONENT
    | '.' DECIMAL_DIGITS (DECIMAL_EXPONENT)?
    ;
fragment HEX_FLOAT_LIT
    : '0' [xX] HEX_MANTISSA HEX_EXPONENT
    ;
fragment HEX_MANTISSA
    : UNDERSCORE? HEX_DIGITS '.' (HEX_DIGITS? | UNDERSCORE)
    | UNDERSCORE? HEX_DIGITS
    | '.' HEX_DIGITS
    ;
fragment HEX_EXPONENT
    : [pP] [+\-]? DECIMAL_DIGITS
    ;

FLOAT_LIT : DECIMAL_FLOAT_LIT | HEX_FLOAT_LIT ;

// ---------- Char Literal (RUNELITERAL)    -----------------------
fragment UNICODE_CHAR : . ;
fragment LITTLE_U_VALUE : '\\' 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT;
fragment BIG_U_VALUE : '\\U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT;
fragment ESCAPED_CHAR : '\\' ( 'a' | 'b' | 'f' | 'n' | 'r' | 't' | 'v' | '\\' | '\'' | '"' ) ;
fragment UNICODE_VALUE : UNICODE_CHAR | LITTLE_U_VALUE | BIG_U_VALUE | ESCAPED_CHAR;
fragment OCTAL_BYTE_VALUE : '\\' OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT ;
fragment HEX_BYTE_VALUE : '\\' 'x' HEX_DIGIT HEX_DIGIT ;
fragment BYTE_VALUE : OCTAL_BYTE_VALUE | HEX_BYTE_VALUE;

RUNE_LIT : '\'' ( UNICODE_VALUE | BYTE_VALUE ) '\'';

// ---------- Strings Literal (RAWSTRINGLITERAL,INTERPRETEDSTRINGLITERAL )   -----------------------

fragment NEWLINE: '\n';
fragment RAW_STRING_LIT: '`' (UNICODE_CHAR | NEWLINE)* '`';
fragment INTERPRETED_STRING_LIT: '"' (UNICODE_VALUE | BYTE_VALUE)* '"';
STRING_LIT : RAW_STRING_LIT | INTERPRETED_STRING_LIT;


// ---------- Others   -----------------------

WS  :   [ \t\n\r]+ -> skip ;