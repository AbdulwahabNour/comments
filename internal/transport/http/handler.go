package http

import (
	"context"
	"fmt"
	"net/http"

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

func (h *Handler)Routers(){
        h.Router.HandleFunc("/", h.Home).Methods("GET")
}
func (h *Handler)Home(w http.ResponseWriter, r *http.Request){
       fmt.Fprint(w, "HOME")
}

func (h  *Handler)Serve() error{
    return h.Server.ListenAndServe()
}