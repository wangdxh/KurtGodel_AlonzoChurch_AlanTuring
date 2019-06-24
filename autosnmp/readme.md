#### parsemib
parsemib use antlr4 to parse the snmp mib file, and generate the snmpmibtoken.txt    
it is a json file contains the mib info    

#### generatego
the generatego read the snmpmibtoken file, and use template to generate the snmplib go files    
this could be done just by parsemib java project, but the java template i am not good at ......   

#### snmplib  
the snmplib contains the generated go files, contains snmp apis in snmp.go       


#### need to do 
now the snmplib api return the data type is just int and string   
in the generatego/main.go  file:  g_mapmibtypestogotypes define the translate type    
     
```
var g_mapmibtypestogotypes map[string]string = map[string]string{    
	"AutonomousType":             "string",    
	"BridgeId":                   "string",
	"Counter":                    "int",
	"Counter32":                  "int",
	"DateAndTime":                "string",
	......
}
```
