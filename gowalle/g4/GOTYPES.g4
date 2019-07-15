grammar GOTYPES;

prog: expression+;

expression
    : structdef NEWLINE?
    | funcdef NEWLINE?
    | NEWLINE
    ;
// fucntion define

funcdef
    : 'func' IDENTIFIER '(' args? (',' args)* ')' IDENTIFIER
    ;

args
    : IDENTIFIER BRACKETS? POINTER? IDENTIFIER?
    ;

// struct define
structdef
    : 'type' IDENTIFIER 'struct' '{' NEWLINE
        (structitem NEWLINE*)+
        '}'
    ;

structitem
    :
        IDENTIFIER? BRACKETS? POINTER? IDENTIFIER tag?
    ;

tag
    : '`' (tagitem)+ '`'
    ;

tagitem
    : IDENTIFIER ':' CSTRING
    ;

POINTER
    : '*'
    ;

BRACKETS
    :'[' ']'
    ;

WS  
    :  (' '|'\r'|'\t'|'\u000C'|'\n') -> skip
    ;

CSTRING
    :  '"' ( EscapeSequence | ~('\\'|'"') )* '"'
    ;

APOSTROPHE
	:	'\''
	;

fragment
EscapeSequence
    :   '\\' ('b'|'t'|'n'|'f'|'r'|'"'|APOSTROPHE|'\\')
    ;

fragment
LETTER
    :  '\u0024' |
    	'\u002d' |
     '\u0041'..'\u005a' |
       '\u005f' |
       '\u0061'..'\u007a' |
       '\u00c0'..'\u00d6' |
       '\u00d8'..'\u00f6' |
       '\u00f8'..'\u00ff' |
       '\u0100'..'\u1fff' |
       '\u3040'..'\u318f' |
       '\u3300'..'\u337f' |
       '\u3400'..'\u3d2d' |
       '\u4e00'..'\u9fff' |
       '\uf900'..'\ufaff'
    ;

fragment
JavaIDDigit
    :  '-'|
	   '\u0030'..'\u0039' |
       '\u0660'..'\u0669' |
       '\u06f0'..'\u06f9' |
       '\u0966'..'\u096f' |
       '\u09e6'..'\u09ef' |
       '\u0a66'..'\u0a6f' |
       '\u0ae6'..'\u0aef' |
       '\u0b66'..'\u0b6f' |
       '\u0be7'..'\u0bef' |
       '\u0c66'..'\u0c6f' |
       '\u0ce6'..'\u0cef' |
       '\u0d66'..'\u0d6f' |
       '\u0e50'..'\u0e59' |
       '\u0ed0'..'\u0ed9' |
       '\u1040'..'\u1049'
   ;

//OBJECTCLASSREFERENCE
//	: UPPER (UPPER | LOWER | '-')
//	;
IDENTIFIER
    :   LETTER (LETTER | JavaIDDigit)*
    ;

NEWLINE:'\r'? '\n';

/*LINE_COMMENT
    : '//' ~('\n'|'\r')* '\r'? '\n' -> skip
    ;*/
COMMENT : '/*' .*? '*/' -> skip;
