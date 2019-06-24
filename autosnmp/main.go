package main

import (
	"github.com/soniah/gosnmp"
	"time"
	"fmt"
	"./snmplib"
)

// gosnmp https://blog.csdn.net/wbchen2330/article/details/48864293
func Dosnmp(strip string, port uint16) {
	params := &gosnmp.GoSNMP{
		Target:    strip,
		Port:      port,
		Community: "public",
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(5) * time.Second,
		Retries:   2,
		//Logger:    log.New(os.Stdout, "", 0),
	}
	err := params.Connect()
	if err != nil {
		fmt.Printf("%s Connect() err: %v\n", strip, err)
		return
	}
	defer params.Conn.Close()

	tab, err :=snmplib.GetifTable(params)
	if err !=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("length: ", len(tab))
	for _, val := range tab {
		fmt.Println(val)
	}
	ret, err := snmplib.GetsysDescr(params)
	fmt.Println(ret, err)
	ret2, err := snmplib.GetsysUpTime(params)
	fmt.Println(ret2, err)
}
func main()  {
	Dosnmp("127.0.0.1", 161)
}