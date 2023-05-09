package main

import (
	"fmt"

	"github.com/AbdulwahabNour/comments/internal/db/postgres"
	"github.com/AbdulwahabNour/comments/internal/service"
	"github.com/AbdulwahabNour/comments/internal/transport/http"
)

// Run is responsible for the instantiation
// and start up the application
func Run() error{
    fmt.Println("starting up the application")
    db, err := postgres.NewDatabase()
 
    if err != nil{
        fmt.Println("Failed to connect to the database %w", err)
        return err
    }
    if err := db.MigrateDB(); err != nil{
        fmt.Println("failed to migrate database", err) 
        return err
    } 

    serv := service.NewService(db)
 

    httpHandler := http.NewHandler(serv)

    err = httpHandler.Serve()
    if err != nil{
        fmt.Println("faild to serve http ", err)
        return err
    }
 
  return nil
}
func main(){
    Run()
   
}