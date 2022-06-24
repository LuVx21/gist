package main

import (
    "fmt"
    _ "luvx/api/log"
    "path"
)

func main() {
    logfile := path.Join("logDir", "app")
    fmt.Println(logfile)

    //log.WithFields(log.Fields{
    //    "漫画id": comicId,
    //    "响应码":  resp.StatusCode,
    //}).Info("获取所有章节接口:")

    panic("hello")
}
