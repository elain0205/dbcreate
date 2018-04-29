// dbhelper
package dbhelper

import (
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

func Initdb(dbtype, ip, port, dbName, userName, pwd string) (bool, string) {

	switch dbtype {
	case "mysql":
		orm.RegisterDriver(dbtype, orm.DRMySQL)
		maxIdle := 30
		maxConn := 30
		ds := userName + ":" + pwd + "@" + ip + "/" + dbName + "?charset=utf8"
		orm.RegisterDataBase("default", dbtype, ds, maxIdle, maxConn)
	default:
		fmt.Println("数据库不支持")
		return false, "数据库不支持"
	}

	return true, "sucess"
}
