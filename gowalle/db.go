package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"regexp"
	"strings"
)

func DbOpen(username, password, ipaddr, databasename string, unsafe bool) (db *sqlx.DB, err error) {
	strurl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Asia%%2fShanghai", username, password, ipaddr, databasename)
	if len(databasename) == 0 {
		strurl = fmt.Sprintf("%s:%s@tcp(%s)/", username, password, ipaddr)
	}
	//db, err = sql.Open("mysql", strurl)
	db, err = sqlx.Open("mysql", strurl)
	if err != nil {
		fmt.Println(db == nil, err)
		return
	}
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(8)

	if unsafe {
		// 牛逼的源泉
		db = db.Unsafe()
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		db = nil
		fmt.Println(err)
		return
	}
	return
}

func DbClose(db *sqlx.DB) {
	if db != nil {
		db.Close()
	}
}

var g_dbfilter = []string{"mysql", "information_schema", "performance_schema"}

func dbnameisok(dbname string) bool {
	for _, val := range g_dbfilter {
		if val == dbname {
			return false
		}
	}
	return true
}

type TableItem struct {
	Field      string         `db:"Field"`
	Type       string         `db:"Type"`
	Collation  sql.NullString `db:"Collation"`
	Null       string         `db:"Null"`
	Key        string         `db:"Key"`
	Default    sql.NullString `db:"Default"`
	Extra      string         `db:"Extra"`
	Privileges string         `db:"Privileges"`
	Comment    string         `db:"Comment"`
	GoType     string
}

type TableInfo struct {
	Name      string
	Itemlist  []*TableItem
	Importmap map[string]string
}

type DbInfo struct {
	DbName        string
	TableInfoList []*TableInfo
}

func GetAllDbandtables(db *sqlx.DB) (resultdblist []*DbInfo, err error) {
	var dblist []string
	err = db.Select(&dblist, "show databases;")
	if err != nil {
		return
	}

	for _, dbname := range dblist {
		if dbnameisok(dbname) {
			fmt.Println("db", dbname)

			_, err = db.Exec(fmt.Sprintf("use %s;", dbname))
			if err != nil {
				return
			}
			var tablelist []string

			err = db.Select(&tablelist, "show tables;")
			if err != nil {
				return
			}

			var tabinfolist []*TableInfo
			for _, tabname := range tablelist {
				var importmap = map[string]string{}
				tabinfo := &TableInfo{Name: tabname, Importmap: importmap}
				err = db.Select(&tabinfo.Itemlist, fmt.Sprintf("SHOW FULL COLUMNS FROM  %s;", tabname))
				if err != nil {
					return
				}
				for _, item := range tabinfo.Itemlist {
					item.GoType = sqltypetogo(item.Type, item.Null)

					if strings.Contains(item.GoType, "time") {
						importmap["time"] = "time"
					} else if strings.Contains(item.GoType, "sql") {
						importmap["database/sql"] = "database/sql"
					}
				}

				tabinfolist = append(tabinfolist, tabinfo)
			}
			dbinfo := &DbInfo{DbName: dbname, TableInfoList: tabinfolist}
			resultdblist = append(resultdblist, dbinfo)
		}
	}

	return
}

func IsNormalSqlWord(wordlist ...string) error {
	for _, val := range wordlist {
		if len(val) == 0 {
			continue
		}
		matched, err := regexp.MatchString(`^([0-9]|[a-z]|[A-Z]|_)+$`, val)
		if err != nil {
			return err
		} else if matched == false {
			return fmt.Errorf("%s is not normal sql word", val)
		}
	}
	return nil
}

var g_mapsqltogotype = map[string]string{
	"TINYINT":   "int8",
	"SMALLINT":  "int16",
	"MEDIUMINT": "int",
	"INT":       "int",
	"INTEGER":   "int",
	"BIGINT":    "int64",

	"CHAR":       "string",
	"VARCHAR":    "string",
	"TINYTEXT":   "string",
	"TEXT":       "string",
	"MEDIUMTEXT": "string",
	"LONGTEXT":   "string",

	"TINYBLOB":   "[]byte",
	"BLOB":       "[]byte",
	"MEDIUMBLOB": "[]byte",
	"LOGNGBLOB":  "[]byte",

	"FLOAT":   "float32",
	"DOUBLE":  "float64",
	"DECIMAL": "float64",

	"DATE":      "time.Time",
	"TIME":      "time.Time",
	"DATETIME":  "time.Time",
	"TIMESTAMP": "time.Time",
	"YEAR":      "int",
}

//https://www.cnblogs.com/-xlp/p/8617760.html
func sqltypetogo(sqltype, cannull string) string {
	for key, value := range g_mapsqltogotype {
		keylower := strings.ToLower(key)
		sqltypelower := strings.ToLower(sqltype)
		if strings.HasPrefix(sqltypelower, keylower) {
			if cannull == "YES" {
				if strings.Contains(value, "int") {
					return "sql.NullInt64"
				} else if strings.Contains(value, "float") {
					return "sql.NullFloat64"
				} else if strings.Contains(value, "string") {
					return "sql.NullString"
				}
				return value
			}
			if strings.Contains(sqltypelower, "unsigned") && strings.Contains(value, "int") {
				return "u" + value
			}
			return value
		}
	}
	panic("do not find the mysql type")
	return ""
}
