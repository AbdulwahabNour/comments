package http

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/AbdulwahabNour/comments/internal/comment"
	"github.com/gorilla/mux"
)


 
type Response struct {
    Message string 
}

func(h *Handler)PostComment(w http.ResponseWriter, r * http.Request){

    var cmt comment.Comment
 
   if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil{
     return
   }

  cmt, err :=  h.Service.PostComment(r.Context(), cmt)
  if err != nil{
      log.Println(err)
      json.NewEncoder(w).Encode(Response{Message: err.Error()})
      return
  }

  if err := json.NewEncoder(w).Encode(cmt); err != nil{
    log.Println(err)
    w.WriteHeader(http.StatusBadRequest)
    return
  }
}

 
func(h *Handler)GetComment(w http.ResponseWriter, r * http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    if  id == ""{
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    cmt, err := h.Service.GetComment(r.Context(), id)
    if err != nil{
        log.Println( err)
        json.NewEncoder(w).Encode(Response{Message: err.Error()})
        return
    }

    if err:= json.NewEncoder(w).Encode(cmt);err != nil{
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

}

func(h *Handler)UpdateComment(w http.ResponseWriter, r * http.Request){
    vars := mux.Vars(r)
    var cmt comment.Comment

     id := vars["id"]
     if id == ""{
        w.WriteHeader(http.StatusBadRequest)
        return
     }
    fmt.Println("Update", id)
     if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil{
        fmt.Println("Update2", id)
        w.WriteHeader(http.StatusInternalServerError)
        return
     }

     cmt, err := h.Service.UpdateComment(r.Context(), cmt)

     if err != nil{
        log.Println( err)
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
     }

     if err :=   json.NewEncoder(w).Encode(cmt); err != nil {
        fmt.Println("Update3", id)
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
     }

     fmt.Println("Update4", id)
}

func(h *Handler)DeleteComment(w http.ResponseWriter, r * http.Request){
    vars := mux.Vars(r)

    id := vars["id"]
    if id == ""{
       w.WriteHeader(http.StatusBadRequest)
       return
    }

   if err := h.Service.DeleteComment(r.Context(), id); err != nil{
      log.Println(err)
      json.NewEncoder(w).Encode(Response{Message:err.Error()})
      return
   }
   json.NewEncoder(w).Encode(Response{Message:"Done"})
    
}