cd %cd%
go build
generatego.exe
cd ../snmplib
copy  snmp.go snmp.go.2
copy  getfunc.go getfunc.go.2
gofmt ./getfunc.go.2 > ./getfunc.go
gofmt ./snmp.go.2 > ./snmp.go
del  snmp.go.2
del  getfunc.go.2
