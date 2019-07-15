grammar PREFIX;

//`(and, true, "a = 123", (atom, true, "b in [123, 354 ]"),  (or, ture, "", ""))`
expression :
    CSTRING  # expstring
    | '(' 'atom' ',' test ',' CSTRING ')' # expsignle
    | '(' operate ',' test ','  expression ',' expression (',' expression)* ')' # express
    ;

operate
    : 'and' | 'or' | 'atom'   
    ;
test 
    : 'true'| CSTRING 
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
