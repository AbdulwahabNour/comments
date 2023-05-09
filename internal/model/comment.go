package model
type Comment struct{
    ID string `json:"id" db:"id"`
    UserId int64 `json:"user_id" db:"user_id"`
    Body string `json:"body" db:"body"`
   
}

type CommentBody struct{
    Body string `json:"body"`
}
