package gin

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "luvx/api/logs"
    "luvx/config"
    _ "luvx/gin/db"
    "luvx/gin/router"

    //"html/http"
)

func WebStart() {
    logs.Log.Infoln("(*^▽^*) 启动... (〃'▽'〃)")
    r := gin.New()
    router.Register(r)

    port := config.AppConfig.Server.Port
    if err := r.Run(":" + port); err != nil {
        fmt.Printf("Start server error,err=%v", err)
    }
}
