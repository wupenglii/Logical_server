package sysinit

import (
	//模块
	_ "Logical_server/System/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//数据库连接
	_ "github.com/go-sql-driver/mysql"
)

func dbinit(aliases ...string) {
	isDev := (beego.AppConfig.String("runmode") == "dev")

	if len(aliases) > 0 {
		for _, alias := range aliases {
			registerDatabase(alias)
			if "w" == alias {
				orm.RunSyncdb("default", false, isDev)
			}
		}

	} else {
		registerDatabase("w")
		orm.RunSyncdb("default", false, isDev) //自动建表
	}

	if isDev {
		orm.Debug = isDev
	}
}

func registerDatabase(alias string) {
	if len(alias) <= 0 {
		return
	}
	dbAlias := alias
	if "w" == alias || "default" == alias {
		dbAlias = "default"
		alias = "w"
	}
	//数据库名称
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	//数据库用户名
	dbuser := beego.AppConfig.String("db_" + alias + "_username")
	//数据库密码
	dbPwd := beego.AppConfig.String("db_" + alias + "_password")
	//数据库IP
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	//数据库端口
	dbPost := beego.AppConfig.String("db_" + alias + "_port")

	orm.RegisterDataBase(dbAlias, "mysql", dbuser+":"+dbPwd+"@tcp("+dbHost+":"+dbPost+")/"+dbName+"?charset=utf8", 30)

}
