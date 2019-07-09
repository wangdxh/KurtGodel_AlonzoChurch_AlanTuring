package db


type QueryBase struct {
	Offset int `form:"offset"
	Nums   int `form:"nums"
	Orderby string `form:"orderby"`
	Desc    bool   `form:"desc"`	
}