package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "luvx/api/common"
    "luvx/mysql/dbconfig"
    "net/http"
)

var db *sql.DB

func init() {
    db, _ = sql.Open("mysql", dbconfig.Url)
    db.SetMaxOpenConns(2000)
    db.SetMaxIdleConns(1000)
    db.Ping()
}

func pool(w http.ResponseWriter, r *http.Request) {
    //ddl()
    //insert()
    _select()
    fmt.Fprintln(w, "finish")
}

func ddl() {
    sqlText := `
        CREATE TABLE IF NOT EXISTS user (
          id bigint(20) NOT NULL AUTO_INCREMENT,
          user_name varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
          password varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
          age int(11) DEFAULT NULL,
          PRIMARY KEY (id)
        ) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8;
    `
    _, err := db.Exec(sqlText)
    if err != nil {
        fmt.Println(err)
    }
}

func insert() {
    sqlText := `INSERT INTO user(user_name, password, age) VALUES ('foo3', 'bar', 3);`
    rs, err := db.Exec(sqlText)
    common.CheckErr(err)
    rowCount, err := rs.RowsAffected()
    common.CheckErr(err)
    fmt.Printf("inserted %d rows\n", rowCount)
}

func _select() {
    rows, err := db.Query("SELECT * FROM user limit 10")
    defer rows.Close()
    common.CheckErr(err)

    columns, _ := rows.Columns()
    fmt.Println(columns)

    scanArgs := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
    for j := range values {
        scanArgs[j] = &values[j]
    }

    record := make(map[string]string)
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
        fmt.Println(record)
    }
}

func startHttpServer() {
    http.HandleFunc("/pool", pool)
    err := http.ListenAndServe(":8090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func main() {
    startHttpServer()
}
