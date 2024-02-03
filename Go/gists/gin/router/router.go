package router

import (
    "github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
    RegisterApp(r)
    RegisterUser(r)
}
