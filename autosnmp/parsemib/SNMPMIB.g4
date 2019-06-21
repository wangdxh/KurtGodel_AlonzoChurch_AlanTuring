

grammar SNMPMIB;

moduleDefinition :  IDENTIFIER ('{' (IDENTIFIER '(' NUMBER ')')* '}')?     
	 'DEFINITIONS'
     tagDefault
     extensionDefault     
	 '::='
     'BEGIN'
     moduleBody
     'END'
;

tagDefault : ( ('EXPLICIT'|'IMPLICIT'|'AUTOMATIC') 'TAGS' )?
;

extensionDefault :
   ('EXTENSIBILITY' 'IMPLIED')?
;

moduleBody :  (exports imports assignmentList) ?
;
//-------------------------exports and imports 实际对我们来说没啥用---
exports :   ('EXPORTS' symbolsExported ';' 
			| 'EXPORTS' 'ALL' ';' )?
;

symbolsExported : ( symbolList )?
;

imports :   ('IMPORTS' symbolsImported ';' )?
;

symbolsImported : (symbolsFromModuleList )?
;

symbolsFromModuleList :
     (symbolsFromModule) (symbolsFromModule)*
;

symbolsFromModule : symbolList 'FROM' globalModuleReference
;

globalModuleReference : IDENTIFIER
;


symbolList   : (symbol) (',' symbol)*
;

symbol  : (IDENTIFIER | valuetype | asnType)
;

//-------------------------------------------- ok 正式进入赋值语句的定义
assignmentList :  (assignment) (assignment)*
;

assignment :
	(  valueAssignment
	 | typeAssignment
	)
;

typeAssignment : IDENTIFIER  '::='  asnType
;

valueAssignment : IDENTIFIER valuetype
    (valuedesctype)*
	  '::=' '{'  IDENTIFIER NUMBER '}'
;

asnType : (builtinType | IDENTIFIER)  (constraint | sizeConstraint)*
;
builtinType :
   octetStringType
 | integerType
 | objectidentifiertype
 | sequencetype  
 ;

valuetype:
	objectidentifiervaluetype 
	| objectvalueType
	| modulevalueType
;
modulevalueType : 'OBJECT-TYPE';
objectvalueType :'MODULE-IDENTITY';
objectidentifiervaluetype :  'OBJECT' 'IDENTIFIER';

valuedesctype:
	syntaxdesctype
	| indexdesctype
	| unknowndesctype
;
syntaxdesctype:'SYNTAX' (asnType | ('SEQUENCE' 'OF' IDENTIFIER));
//descdesctype: 'DESCRIPTION' CSTRING;
//accessdesctype: 'ACCESS' IDENTIFIER;
//statusdesctype:'STATUS' IDENTIFIER;
indexdesctype: 'INDEX' '{' IDENTIFIER (',' IDENTIFIER)* '}';
unknowndesctype: IDENTIFIER (CSTRING|IDENTIFIER);

octetStringType  :  'OCTET' 'STRING';
objectidentifiertype  :  'OBJECT' 'IDENTIFIER';
sequencetype: 'SEQUENCE' '{' typedefine (',' typedefine)* '}'
;
typedefine : IDENTIFIER asnType
;

integerType:'INTEGER'  ('{' namedNumberList '}')?
;
namedNumberList : (namedNumber) (',' namedNumber)*
;
namedNumber :   IDENTIFIER '(' signedNumber ')'
;
signedNumber :  ('-')? NUMBER
;


sizeConstraint : '(' 'SIZE' constraint ')'
;
constraint :'(' rangeconstraint ')'
;
rangeconstraint: signedNumber '..' signedNumber
;

//////////////////////////////////////////////////////

COMMENT
    :	'--'
    ;

fragment DIGIT
    : '0'..'9'
    ;

fragment UPPER
    : ('A'..'Z')
    ;

fragment LOWER
    : ('a'..'z')
    ;

NUMBER
    : DIGIT+
    ;

WS
    :  (' '|'\r'|'\t'|'\u000C'|'\n') -> skip
    ;

fragment Exponent
    : ('e'|'E') ('+'|'-')? NUMBER
    ;

LINE_COMMENT
    : '--' ~('\n'|'\r')* '\r'? '\n' ->skip
    ;

BSTRING
    : APOSTROPHE ('0'..'1')* '\'B'
    ;

fragment HEXDIGIT
    : (DIGIT|'a'..'f'|'A'..'F')
    ;

HSTRING
    : APOSTROPHE HEXDIGIT* '\'H'
    ;




//IDENTIFIER : ('a'..'z'|'A'..'Z') ('0'..'9'|'a'..'z'|'A'..'Z')* ;
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



//fragment

/**I found this char range in JavaCC's grammar, but Letter and Digit overlap.
   Still works, but...
 */
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

/*IDENTIFIER
   :
   NONDIGIT (NONDIGIT | DIGIT)*
   ;

fragment Identifiernondigit
   : NONDIGIT   
   ;

fragment NONDIGIT
   : [a-zA-Z_\-]
   ;*/

/*fragment DIGIT
   : [0-9]
   ;*/