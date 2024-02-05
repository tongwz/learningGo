package main

import (
	"fmt"

	"github.com/gohouse/converter"
)

func main() {
	err := converter.NewTable2Struct().
		SavePath("./model/file.go").
		Dsn("root:password@tcp(xxx.xxx.xx.xx:3306)/sscf_company_module?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		Table("table_name").
		Run()
	fmt.Println(err)
}
