package dbops

import (
_"github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
"common/config"
)

var (
	dbConn *gorm.DB
	err error
)

func init()  {
	dbHost := common.ConfMap["db_host"]
	dbPwd := common.ConfMap["db_password"]
	dbUserName := common.ConfMap["db_username"]
	dbDatabase := common.ConfMap["db_database"]
	dbPort := common.ConfMap["db_port"]
	dbConn, err = gorm.Open("mysql",  dbUserName+ ":"+ dbPwd + "@tcp("+ dbHost + ":" + dbPort + ")/"+ dbDatabase +"?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	dbConn.SingularTable(true)
}
