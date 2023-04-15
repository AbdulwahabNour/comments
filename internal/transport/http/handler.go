package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AbdulwahabNour/comments/internal/comment"
	"github.com/gorilla/mux"
)

type CommentService interface{
    GetComment( context.Context,  string)(comment.Comment, error)
    PostComment( context.Context,  comment.Comment)(comment.Comment, error)
    DeleteComment( context.Context,  string) (error)
    UpdateComment( context.Context,  comment.Comment)(comment.Comment, error)
}

type Handler struct{
    Server *http.Server
    Router *mux.Router
    Service CommentService
}

func NewHandler(service CommentService) *Handler{
     h := Handler{
        Service: service,
     }
     h.Router = mux.NewRouter()
     h.Routers()
     
     h.Server= &http.Server{
        Addr: "0.0.0.0:8080",
        Handler: h.Router,
     }
    
     return &h
}

// curl -X POST \
//   -H "Content-Type: application/json" \
//   -d '{"slug":"Hello", "body":"body", "autor":"test"}' \
//   http://localhost:8080/api/v1/comment

func (h *Handler)Routers(){
        h.Router.HandleFunc("/api/v1/comment", h.PostComment).Methods("POST")
        h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
        h.Router.HandleFunc("/api/v1/comment/{id}", h.UpdateComment).Methods("PUT")
        h.Router.HandleFunc("/api/v1/comment/{id}", h.DeleteComment).Methods("DELETE")
}
func (h *Handler)Home(w http.ResponseWriter, r *http.Request){
       fmt.Fprint(w, "HOME")
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