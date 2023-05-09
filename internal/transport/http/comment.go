package http

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/gorilla/mux"
)


 


func(h *Handler)PostComment(w http.ResponseWriter, r * http.Request){
 
    var cmt model.Comment
    var commentBody model.CommentBody
 

   if err := json.NewDecoder(r.Body).Decode(&commentBody); err != nil{
    log.Println(err)
     return
   }

   fetcheduser, err := h.GetUserFromToken(r)
    if err != nil {
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message: "invalid id"})
    }
    
   cmt.UserId = fetcheduser.ID
   cmt.Body= commentBody.Body
fmt.Println("userrrr ",fetcheduser)
   _, err =  h.Service.PostComment(r.Context(), &cmt)

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

 
    var cmt model.Comment
    
    id :=  mux.Vars(r)["id"]
    if  id == ""{
        json.NewEncoder(w).Encode(Response{Message:"invalid id"})
        return
    }

    fetcheduser, err := h.GetUserFromToken(r)

    if err != nil {
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message: "invalid id"})
    }
    cmt.ID = id
    cmt.UserId = fetcheduser.ID
 
    fetchCmt, err := h.Service.GetComment(r.Context(), &cmt)
    if err != nil{
        log.Println( err)
        json.NewEncoder(w).Encode(Response{Message: err.Error()})
        return
    }

    if err:= json.NewEncoder(w).Encode(fetchCmt);err != nil{
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

}

func(h *Handler)UpdateComment(w http.ResponseWriter, r * http.Request){
  
   
    var cmt model.Comment
    var commentBody model.CommentBody
 


     id :=  mux.Vars(r)["id"]
     if id == ""{
        json.NewEncoder(w).Encode(Response{Message:"invalid id"})
        return
     }
    
     if err := json.NewDecoder(r.Body).Decode(&commentBody); err != nil{
        w.WriteHeader(http.StatusInternalServerError)
        return
     }

     fetcheduser, err := h.GetUserFromToken(r)

     if err != nil {
         log.Println(err)
         json.NewEncoder(w).Encode(Response{Message: "invalid id"})
     }
     cmt.ID = id
     cmt.UserId = fetcheduser.ID
     cmt.Body = commentBody.Body


      
     _, err = h.Service.UpdateComment(r.Context(), &cmt)

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
    var cmt model.Comment
 

    id := mux.Vars(r)["id"]
    if id == ""{
       w.WriteHeader(http.StatusBadRequest)
       return
    }
    fetcheduser, err := h.GetUserFromToken(r)

    if err != nil {
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message: "invalid id"})
    }

    cmt.ID = id
    cmt.UserId = fetcheduser.ID

   if err := h.Service.DeleteComment(r.Context(), &cmt); err != nil{

        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message: err.Error()})

        return
   }

   json.NewEncoder(w).Encode(Response{Message:"Done"})
 
}


