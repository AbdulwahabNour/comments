package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AbdulwahabNour/comments/internal/db"
	"github.com/gorilla/mux"
)

 
type Response struct {
    Message string 
}
type Handler struct{
    Server *http.Server
    Router *mux.Router
    Service db.Service
}

func NewHandler(service db.Service) *Handler{
     h := Handler{
        Service: service,
     }
     h.Router = mux.NewRouter()
     h.Routers()
   
     h.Router.Use(JSONMiddleWare)
     h.Router.Use(LoggingMiddleWare)
     h.Router.Use(TimeOutMiddleWare)
     
     h.Server= &http.Server{
        Addr: "0.0.0.0:8080",
        Handler:  h.Router ,
     }
    
     return &h
}

 

func (h *Handler)Routers(){
        h.Router.HandleFunc("/api/v1/comment",JWTAuth(h.PostComment)).Methods("POST")
        h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
        h.Router.HandleFunc("/api/v1/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
        h.Router.HandleFunc("/api/v1/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")
       
     
        h.Router.HandleFunc("/api/v1/user/signin", h.SignIn).Methods("GET") 
        h.Router.HandleFunc("/api/v1/user/{id}", h.GetUser).Methods("GET") 
        h.Router.HandleFunc("/api/v1/user",JWTAuth(h.PostUser)).Methods("POST")
        h.Router.HandleFunc("/api/v1/user/{id}", JWTAuth(h.UpdateUser)).Methods("PUT")
        h.Router.HandleFunc("/api/v1/user/{id}", JWTAuth(h.DeleteUser)).Methods("DELETE")
        
}
 

func (h  *Handler)Serve() error{
    go func(){
        if err := h.Server.ListenAndServe(); err != nil{
            log.Println(err.Error())
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
  ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
  defer cancle()
  h.Server.Shutdown(ctx)
  log.Println("Shutdown ")
    return nil
}