package dao

import (
    "luvx/gin/db"
    "luvx/gin/model"
)

func GetUserByUsername(username string) (*model.User, error) {
    var user model.User
    if err := db.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
