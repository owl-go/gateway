package mysql

import (
	"fmt"
	"log"

	//_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type MysqlDriver struct {
	*MysqlConfig
	DbCon *gorm.DB
}

func IsRecordNotFoundError(err error) bool {
	if err != nil && err.Error() == "sql: no rows in result set" {
		return true
	} else {
		return false
	}
}

func NewMysqlDriver(c *MysqlConfig) *MysqlDriver {
	dbDriver := &MysqlDriver{
		MysqlConfig: c,
	}
	dbDriver.Open()
	return dbDriver
}

//Open 连接数据库
func (dbDriver *MysqlDriver) Open() {
	dbConn, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbDriver.Username, dbDriver.Password, dbDriver.Host, dbDriver.Port, dbDriver.Database))
	if err != nil {
		log.Fatal(err)
	}
	dbDriver.DbCon = dbConn
}

//关闭数据库
func (dbDriver *MysqlDriver) Close() {
	dbDriver.DbCon.Close()
}
