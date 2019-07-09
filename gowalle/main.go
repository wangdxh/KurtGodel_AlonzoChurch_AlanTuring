package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func goclear() {
	err := os.RemoveAll("./db")
	if err != nil {
		log.Fatal(err)
		return
	}
}
func goformat(dir string) {
	golist := []string{}
	filepath.Walk("./db", func(path string, info os.FileInfo, err error) error {
		abspath, err := filepath.Abs(path)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if strings.HasSuffix(path, ".go") {
			golist = append(golist, abspath)
		}
		return nil
	})
	for _, path := range golist {
		fmt.Println(path)
		err := exec.Command("gofmt", "-l", "-w", path).Run()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
func pongo2filters() {
	joinlines := func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
		strin, bok := in.Interface().(string)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:gettablekey",
				OrigError: fmt.Errorf("Filter input argument must be of type 'string'."),
			}
		}
		strsep, bok := param.Interface().(string)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:gettablekey",
				OrigError: fmt.Errorf("Filter input argument must be of type 'string'."),
			}
		}
		strlines := strings.Split(strin, "\n")
		var linelist []string
		for _, value := range strlines {
			strtemp := strings.Trim(value, " \n\t\r")
			if len(strtemp) > 0 {
				linelist = append(linelist, strtemp)
			}
		}
		return pongo2.AsValue(strings.Join(linelist, strsep)), nil
	}
	pongo2.RegisterFilter("joinlines", joinlines)

	prefix := func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
		strin, bok := in.Interface().(string)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:prefix",
				OrigError: fmt.Errorf("Filter input argument must be of type 'string'."),
			}
		}
		strsub, bok := param.Interface().(string)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:prefix",
				OrigError: fmt.Errorf("Filter input argument must be of type 'string'."),
			}
		}
		return pongo2.AsValue(strings.HasPrefix(strin, strsub)), nil
	}
	pongo2.RegisterFilter("prefix", prefix)

}

type funcdef struct {
	FuncHeader  string `xml:"funcheader"`
	Structdef   string `xml:"structdefine"`
	QueryMethod string `xml:"querymethod"`
}
type tabinfo struct {
	TabName  string `xml:"tabname"`
	TabInfo  string `xml:"info"`
	Tabtype  string `xml:"type"`
	Tabgodef string `xml:"godef"`
}
type dbdefs struct {
	DbName    string    `xml:"dbname"`
	Funcdefs  []funcdef `xml:"funcdefs"`
	TableInfo []tabinfo `xml:"tabinfo"`
}
type SqlFile struct {
	Dbdefs []dbdefs `xml:"dbdefs"`
}

func main() {
	test()
	return

	goclear()
	pongo2filters()

	db, err := DbOpen("root", "kedacom", "172.16.64.92:3306",
		"", false)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer DbClose(db)

	dblist, err := GetAllDbandtables(db)
	if err != nil {
		fmt.Println(err)
	}

	for _, dbitem := range dblist {
		fmt.Println("dbname : ", dbitem.DbName)

		for _, tabinfo := range dbitem.TableInfoList {
			fmt.Println("\t : ", tabinfo.Name)
			for _, item := range tabinfo.Itemlist {
				fmt.Println("\t\t", item.Field, item.Type, item.Key, item.Null, item.GoType, item.Extra)
			}
		} //can be delete

		err = os.MkdirAll(fmt.Sprintf("./db/%s", dbitem.DbName), os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return
		}

		// generate the code
		tplstruct, err := pongo2.FromFile("./tpl/table.go.tpl")
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, tabinfo := range dbitem.TableInfoList {
			keyitemlist := []*TableItem{}
			for _, value := range tabinfo.Itemlist {
				if strings.Contains(value.Key, "PRI") {
					keyitemlist = append(keyitemlist, value)
				}
			}

			strdata, err := tplstruct.Execute(pongo2.Context{
				"package":     dbitem.DbName,
				"importmap":   tabinfo.Importmap,
				"TableInfo":   tabinfo,
				"keyitemlist": keyitemlist,
			})
			if err != nil {
				fmt.Println(err)
			}

			ioutil.WriteFile(fmt.Sprintf("./db/%s/%s.go", dbitem.DbName, tabinfo.Name),
				[]byte(strdata), 600)
		}

	}

	sql := SqlFile{}
	for _, dbitem := range dblist {
		dbtemp := dbdefs{DbName: dbitem.DbName}

		for _, tab := range dbitem.TableInfoList {
			fmt.Println("\t : ", tab.Name)
			tbtemp := tabinfo{TabName: tab.Name}

			for _, item := range tab.Itemlist {
				tbtemp.TabInfo += item.Field + ","
				tbtemp.Tabtype += item.Type + ","
				tbtemp.Tabgodef += item.Field + " " + item.GoType + ","
			}

			dbtemp.TableInfo = append(dbtemp.TableInfo, tbtemp)

		}
		dbtemp.Funcdefs = []funcdef{
			funcdef{},
		}
		sql.Dbdefs = append(sql.Dbdefs, dbtemp)
	}

	data, err := yaml.Marshal(sql) //xml.MarshalIndent(sql, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile("./config.sample.yaml", data, os.ModePerm)

	goformat("./db")
}
