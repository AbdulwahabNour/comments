package model
type Comment struct{
    ID string `json:"id" db:"id"`
    Slug string `json:"slug" db:"slug"`
    Body string `json:"body" db:"body"`
    Author string `json:"author" db:"author"`
}