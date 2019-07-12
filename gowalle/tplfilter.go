package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"strings"
)

func init() {

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

	isallatomfilter := func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
		where, bok := in.Interface().(Where)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:isallatomfilter",
				OrigError: fmt.Errorf("Filter input argument must be of type 'where'."),
			}
		}
		return pongo2.AsValue(isallatom(where)), nil
	}
	pongo2.RegisterFilter("isallatom", isallatomfilter)

	generateatomfilter := func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
		where, bok := in.Interface().(Where)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:generateatomfilter",
				OrigError: fmt.Errorf("Filter input argument must be of type 'where'."),
			}
		}
		return pongo2.AsValue(generateatom(where)), nil
	}
	pongo2.RegisterFilter("generateatom", generateatomfilter)

	trim := func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
		strin, bok := in.Interface().(string)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:trim",
				OrigError: fmt.Errorf("Filter input argument must be of type 'string'."),
			}
		}
		/*strsub, bok := param.Interface().(string)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:trim",
				OrigError: fmt.Errorf("Filter input argument must be of type 'string'."),
			}
		}*/
		return pongo2.AsValue(strings.Trim(strin, `"`)), nil
	}
	pongo2.RegisterFilter("trim", trim)

	sql := func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
		strin, bok := in.Interface().(string)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:sql",
				OrigError: fmt.Errorf("Filter input argument must be of type 'string'."),
			}
		}
		/*strsub, bok := param.Interface().(string)
		if !bok {
			return nil, &pongo2.Error{
				Sender:    "filter:sql",
				OrigError: fmt.Errorf("Filter input argument must be of type 'string'."),
			}
		}*/
		return pongo2.AsValue("(" + strin + ")"), nil
	}
	pongo2.RegisterFilter("sql", sql)

}
