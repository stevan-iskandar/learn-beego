package database

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

type Orm struct {
	CreatedORM orm.Ormer
}

func init() {
	sqlconn, _ := web.AppConfig.String("sqlconn")
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", sqlconn)
}
