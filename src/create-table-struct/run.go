package main

import (
	"fmt"

	"github.com/gohouse/converter"
)

func main() {
	err := converter.NewTable2Struct().
		SavePath("./model/cmf_live_record.go").
		Dsn("root:password@tcp(xxx.xxx.xx.xx:3306)/sscf_company_module?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		Table("cmf_live_record").
		Run()
	fmt.Println(err)
}
