package model

import "time"

type User struct{
    ID int64 `json:"id" db:"id"`
    Username string `json:"username" db:"username"`
    Password string `json:"password" db:"password"`
    Email string `json:"email" db:"email"`
    CreatedAt time.Time `json:"created_at" db:"create_at"`
}
type UserCredentials struct{
    Email string `json:"email"`
    Password string `json:"password"`
}

 