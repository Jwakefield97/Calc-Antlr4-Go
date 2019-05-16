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
COMMENT: '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);

// Rules
start 
    : variables* expression* prints*  EOF 
    | variables* prints* expression*  EOF 
    | expression* prints* variables* EOF
    | expression* variables* prints* EOF 
    | prints* variables* expression* EOF
    | prints* expression* variables* EOF
    ;

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
    : LET (' ' | '\t')* VARNAME  (' ' | '\t')* '=' (' ' | '\t')* (expression) # Variable
    ;