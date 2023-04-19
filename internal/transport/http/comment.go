package http

import (
	"encoding/json"
	"log"

	"net/http"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/gorilla/mux"
)


 


func(h *Handler)PostComment(w http.ResponseWriter, r * http.Request){
 
    var cmt model.Comment
 
   if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil{
    log.Println(err)
     return
   }

   _, err :=  h.Service.PostComment(r.Context(), &cmt)

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
    var cmt model.Comment

     id := vars["id"]
     if id == ""{
        w.WriteHeader(http.StatusBadRequest)
        return
     }
    
     if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil{
        
        w.WriteHeader(http.StatusInternalServerError)
        return
     }

     _, err := h.Service.UpdateComment(r.Context(), &cmt)

     if err != nil{
        log.Println( err)
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
     }

     if err := json.NewEncoder(w).Encode(cmt); err != nil {

        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
        
     }

 
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
        w.WriteHeader(http.StatusBadRequest)

        return
   }

   json.NewEncoder(w).Encode(Response{Message:"Done"})
 
}


