package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

/**
	docker容器内访问宿主机MySql，修改监听地址为0.0.0.0  /etc/mysql/mysql.conf.d/mysqld.cnf
 */
func NewMySqlConnetcion(dbHost string, dbName string, dbUser string, dbPassword string) *gorm.DB{
	var err error

	dbType 		:= "mysql"
	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbName))

	if err != nil {
		panic("MySql Connect Failed !" + err.Error())
	}

	DB.LogMode(true)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	//DB.SetLogger(loger.Loger)

	fmt.Println("MySql Connect Success.", dbHost)
	return DB
}


