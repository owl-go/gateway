package mysql

import (
	"encoding/json"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

type User struct {
	Id       int    `json:"id",gorm:"id"`
	Account  string `gorm:"column:account;default:'1970-01-01''" json:"n"`
	Birthday string `json:"birthday" gorm:"column:birthday"`
}

//func (t Test) TableName() string {
//	return "test"
//}

type Account struct {
	Uid      int
	Nick     string
	Sex      uint8
	Age      uint8
	Province string //所在省份
	City     string //所在城市
}

func TestMysql(t *testing.T) {
	cfg := MysqlConfig{
		Host:     "39.108.87.16",
		Port:     "3306",
		Username: "root",
		Password: "Adinzx12.com",
		Database: "yip_com",
	}
	db := NewMysqlDriver(&cfg)

	jsonStr := `{"id":1000,"n":"hello","birthday":"2021-09-08"}`
	var te User
	//te.UserName = "haozy"

	json.Unmarshal([]byte(jsonStr), &te)
	fmt.Println(te)

	//fmt.Println(te)
	//db.NewRecord(te)
	err := db.DbCon.Create(&te).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(te.Id)
	//select

}
