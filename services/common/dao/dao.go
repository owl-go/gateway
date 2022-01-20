package dao

import (
	"server/infra/mysql"
)

type Dao struct {
	db *mysql.MysqlDriver
}

func NewDAO(database *mysql.MysqlDriver) *Dao {
	if database == nil {
		return nil
	}
	return &Dao{db: database}
}
