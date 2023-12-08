package dbconfig

import "fmt"

var username, password = "root", "1121"
var host, port, database = "luvx", 53306, "boot"
var Url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", username, password, host, port, database)
