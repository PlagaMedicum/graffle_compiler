grammar graffle;

//
// LEXER
//

fragment DIGIT : [0-9] ;

// types
INT : '-'? DIGIT+ ;
NUMBER : INT ('.'DIGIT+)? ;
STRING : '"' ~[\\"]* '"' ;

// operators
OR_ARC : '->' ;
OR_W_ARC : '-[' NUMBER ']' OR_ARC ;
UNOR_ARC : '--' ;
UNOR_W_ARC : '-[' NUMBER ']-' ;
ASSIGN : '=' ;
ADD_ASSIGN : '+=' ;
SUB_ASSIGN : '-=' ;
ADD : '+' ;
SUB : '-' ;
MULT : '*' ;
DIV : '/' ;
EQ : '==' ;
NEQ : '!=' ;
// COMMA : ',' ;

// type names
V_N : 'V' ; // vertice
E_N : 'E' ; // arc
G_N : 'G' ; // graph

// other
ID : [a-zA-Z_][a-zA-Z_0-9]* ;

// ignore
WS
    : [ \t\n\r]+
        -> skip
    ;
LINE_COMMENT
    : '#' ~[\r\n]*
        -> skip
    ;

//
// PARSER
//

type_assignement : ID ASSIGN (V_N | E_N | G_N) '(' ID ')' ; // graph = G(vertice)
