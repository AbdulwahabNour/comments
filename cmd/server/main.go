package main

import (
	"fmt"

	"github.com/AbdulwahabNour/comments/internal/db"
)

// Run is responsible for the instantiation
// and start up the application
func Run() error{
    fmt.Println("starting up the application")
    db, err := db.NewDatabase()
    if err != nil{
        fmt.Println("Failed to connect to the database %w", err)
        return err
    }
    if err := db.MigrateDB(); err != nil{
        fmt.Println("failed to migrate database", err)
        return err
    } 
    return nil
}
func main(){
    Run()
    fmt.Println("Hello")
}