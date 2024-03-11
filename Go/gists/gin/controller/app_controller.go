package controller

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "luvx/gin/db"
    "luvx/gin/model"
    "net/http"
)

func HealthyCheck(c *gin.Context) {
    var user model.User
    if err := db.MySQLClient.Where("id = ?", 1).First(&user).Error; err != nil {
        panic(err)
    }

    userTable := db.MongoDatabase.Collection("user")
    filter := bson.D{{"_id", 1}}
    var result bson.M
    _ = userTable.FindOne(context.TODO(), filter).Decode(&result)

    v, _ := db.RedisClient.Get(context.Background(), "foo").Result()

    c.JSON(http.StatusOK, gin.H{
        "mysql": user,
        "mongo": result,
        "redis": v,
    })
}
