package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/gorilla/mux"
)

func(h *Handler)GetUser(w http.ResponseWriter, r *http.Request){
   
    id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
    
    if err != nil  {
        json.NewEncoder(w).Encode(Response{Message:"invalid id"})
        return
    }
    
 
    user, err := h.Service.GetUser(r.Context(), id)
    if err != nil{
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
    }

    if err:= json.NewEncoder(w).Encode(user); err!=nil{
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
}

func(h *Handler)PostUser(w http.ResponseWriter, r *http.Request){
    
    var user model.User
   
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
         log.Println(err)
         w.WriteHeader(http.StatusBadRequest)
        return
    }
 
    _, err := h.Service.PostUser(r.Context(), &user)

    if err !=nil{
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
    }

    if err := json.NewEncoder(w).Encode(user); err != nil{
            log.Println(err)
            w.WriteHeader(http.StatusBadRequest)
            return
    }

}

func(h *Handler)DeleteUser(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    if id == ""{
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    err := h.Service.DeleteUser(r.Context(), id)
    if err != nil{
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

   json.NewEncoder(w).Encode(Response{Message: "Done"})
}



type UserResponse struct{
    Username string `json:"username"`
    Email string `json:"email"`
    CreatedAt time.Time `json:"create_at" `
}


func(h *Handler)UpdateUser(w http.ResponseWriter, r *http.Request){
 
    var user model.User

    id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

    if err != nil {
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message:"invalid id"})
        return
    }
       

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
        log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    //check if user exist in database or not
    
    if err != nil {
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message:"invalid id"})
    }



    user.ID= int64(id);
    _, err = h.Service.UpdateUser(r.Context(), &user)
    
    if err != nil{
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
    }
   
    userRep  := UserResponse{
            Username: user.Username,
            Email: user.Email,
            CreatedAt: user.CreatedAt,
    }
    
    if err := json.NewEncoder(w).Encode(userRep); err != nil{
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}   