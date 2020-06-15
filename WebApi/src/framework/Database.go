package framework

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"fmt"
)

const (
	username ="root"
	password ="123456"
	ip ="120.78.205.70"
	port="3306"
	dbName="DesktopApp"
	driverName="mysql"
)
var DB *sql.DB
func InitDB(){
	path :=strings.Join([]string{username,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8&parseTime=True&loc=Local"},"")
	DB, _=sql.Open(driverName,path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err :=DB.Ping();err!=nil {
		fmt.Println("database connect error")
	}
	fmt.Println("database connect success")
}

func CreateTable(){
	userTable :="Create Table if not exists `user(`"+
		"`id` int Unsigned auto_increment, "+
		"`username` varchar(20) not null,"+
		"`password` varchar(40) not null,"+
		"`create_time` DateTime,"+
		"PRIMARY KEY (`id`)"+
		")ENGINE=InnoDB DEFAULT CHARSET=utf8;"
	_,err:=DB.Exec(userTable)
	if err!=nil {
		fmt.Println("CreateTable error")
	}
	fmt.Println("CreateTable success")
}