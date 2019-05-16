// Calc.g4
grammar Calc;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
PRINT: 'print';
LET: 'let';
VARNAME: [a-zA-Z]+;
WHITESPACE: (' ' | '\r' | '\n' | '\t')+ -> skip;

// Rules
start : expression | prints | variables EOF;

expression
   : expression op=('*'|'/') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | NUMBER                             # Number
   | VARNAME                            # VariableExp
   ;
prints
    : PRINT '(' expression ')' # PrintExp
    | PRINT '(' variables ')' # PrintVar
    ;
variables 
    : LET WHITESPACE? VARNAME  WHITESPACE? '=' WHITESPACE? (expression | VARNAME | NUMBER) # Variable
    ;