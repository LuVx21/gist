package db

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
    "luvx/config"
)

var MySQLClient *gorm.DB

func init() {
    fmt.Println("初始化数据库连接...")
    var err error
    mysqlConfig := config.AppConfig.MySQL
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Dbname)
    MySQLClient, err = gorm.Open(mysql.New(mysql.Config{
        DSN: dsn,
    }), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            //TablePrefix:   "t_", // 表名前缀
            SingularTable: true, // 使用单数表名
        },
    })

    if err != nil {
        panic(err)
    }
}
