package main

import (
	_ "TantanDemo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	// set db driver
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=lixiaodong password=pingan10000 dbname=lixiaodong host=127.0.0.1 port=5432 sslmode=disable")
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	beego.Run()
}
