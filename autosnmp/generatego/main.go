package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const (
	NONE_ITEMTYPE     = "NONE"
	OBJECT_ITEMTYPE   = "OBJECT"
	SEQUENCE_ITEMTYPE = "SEQUENCE"
)

type ItemInfo struct {
	Name     string   `json:"name"`
	Oid      string   `json:"oid"`
	Strtype  string   `json:"strtype"`
	Itemtype string   `json:"itemtype"`
	Children []string `json:"children"`
}

func getrootitem() ItemInfo {
	return g_maptree["mgmt"]
}

var g_mapmibtypes = make(map[string]string)

func getallmibobjecttypes(item ItemInfo) {
	if item.Itemtype == OBJECT_ITEMTYPE && item.Children == nil {
		if _, bok := g_mapmibtypes[item.Strtype]; !bok {
			g_mapmibtypes[item.Strtype] = item.Strtype
		}
	}

	if item.Children != nil {
		for _, val := range item.Children {
			son := g_maptree[val]
			getallmibobjecttypes(son)
		}
	}
}

func recursiveprint(item ItemInfo, ndeep int) {
	strtab := strings.Repeat("  ", ndeep)
	fmt.Println(strtab, item.Name, ":", item.Oid, ":", item.Strtype)
	if item.Children != nil {
		for _, val := range item.Children {
			son := g_maptree[val]
			recursiveprint(son, ndeep+1)
		}
	}
}

var g_maptree map[string]ItemInfo

var g_mapmibtypestogotypes map[string]string = map[string]string{
	"AutonomousType":             "string",
	"BridgeId":                   "string",
	"Counter":                    "int",
	"Counter32":                  "int",
	"DateAndTime":                "string",
	"DisplayString":              "string",
	"Gauge":                      "int",
	"Gauge32":                    "int",
	"INTEGER":                    "int",
	"InterfaceIndex":             "int",
	"InterfaceIndexOrZero":       "int",
	"InternationalDisplayString": "string",
	"IpAddress":                  "string",
	"KBytes":                     "int",
	"MacAddress":                 "string",
	"NetworkAddress":             "string",
	"PhysAddress":                "string",
	"ProductID":                  "string",
	"String":                     "string",
	"TimeTicks":                  "int",
	"Timeout":                    "int",
	"TruthValue":                 "int",
}

func getgotype(mibtype string) string {
	ret, bok := g_mapmibtypestogotypes[mibtype]
	if !bok {
		log.Fatalf("mibtyps is not in the map ,can not to go code: ", mibtype)
	}
	return ret
}

func generatecode(item ItemInfo) {
	if item.Itemtype == SEQUENCE_ITEMTYPE && item.Children != nil && len(item.Children) == 1 {
		fmt.Printf("type %s struct  {\n", strings.Title(item.Name))

		itemson := g_maptree[item.Children[0]]
		for _, sunzi := range itemson.Children {

			itemsunzi := g_maptree[sunzi]
			strtag := fmt.Sprintf(" `snmp:\"%s\"`", itemsunzi.Strtype)
			fmt.Println("\t", strings.Title(itemsunzi.Name), " ", getgotype(itemsunzi.Strtype), strtag)
		}
		fmt.Println("}")
		return
	}
	if item.Itemtype == OBJECT_ITEMTYPE && item.Children == nil {
		fmt.Println(item.Name, ":", getgotype(item.Strtype))
	}

	if item.Children != nil {
		for _, sonname := range item.Children {

			son := g_maptree[sonname]
			generatecode(son)
		}
	}
}
func main() {

	data, err := ioutil.ReadFile("../snmpmibtoken.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &g_maptree)
	if err != nil {
		fmt.Println(err)
		return
	}
	recursiveprint(getrootitem(), 0)
	fmt.Println("-----------------------------------------types")
	getallmibobjecttypes(getrootitem())

	var strtypes []string
	for key, _ := range g_mapmibtypes {
		strtypes = append(strtypes, key)
	}
	sort.Strings(strtypes)
	for _, val := range strtypes {
		fmt.Printf(`"%s":"%s"`, val, val)
		fmt.Println()
	}
	fmt.Println("-----------------------------------------go code")
	generatecode(getrootitem())
}
