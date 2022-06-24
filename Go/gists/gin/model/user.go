package model

type UsersModel struct {
    Id       int32  `json:"id"`
    Name     string `json:"name"`
    Password string `json:"-"`
    Age      int8   `json:"age"`
}
