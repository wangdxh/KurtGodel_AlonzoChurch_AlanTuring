package main

import (
	"fmt"
	"github.com/flosch/pongo2"
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

func main() {
	goclear()

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
				fmt.Println("\t\t", item.Field, item.Type, item.Key, item.Null, item.GoType)
			}
		} //can be delete

		err = os.MkdirAll(fmt.Sprintf("./db/%s", dbitem.DbName), os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return
		}
		// generate the code
		tplstruct, err := pongo2.FromFile("./table.go.tpl")
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, tabinfo := range dbitem.TableInfoList {
			strdata, _ := tplstruct.Execute(pongo2.Context{
				"package":   dbitem.DbName,
				"importmap": tabinfo.Importmap,
				"TableInfo": tabinfo,
			})

			ioutil.WriteFile(fmt.Sprintf("./db/%s/%s.go", dbitem.DbName, tabinfo.Name),
				[]byte(strdata), 600)
		}

	}

	goformat("./db")
}
