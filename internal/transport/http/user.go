package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/AbdulwahabNour/comments/internal/utils/hash"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gorilla/mux"
)
type UserResponse struct{
    Username string `json:"username"`
    Email string `json:"email"`
}

func(h *Handler)SignIn(w http.ResponseWriter, r *http.Request){
     var creds model.UserCredentials

     if err:= json.NewDecoder(r.Body).Decode(&creds); err != nil{
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message:"something wrong happened please try again later"})
        return
     }

     user, err:= h.Service.GetUserByEmail(r.Context(), &model.User{Email:creds.Email, Password: creds.Password})

     if err != nil {
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
     }

     err = hash.CheckPassword(creds.Password, user.Password)
     if err != nil{
        fmt.Println("wrong pAss")
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
     }
 
     jwtClaims := jwt.MapClaims{
        "id": user.ID,
        "username": user.Username,
        "email": user.Email,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    }

    token, err := generateJwtToken(jwtClaims)
    if err != nil {
        fmt.Println("wrong token")
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
     }

     w.Header().Add("Authorization",fmt.Sprintf("token %s", token))

     if err != nil {
        json.NewEncoder(w).Encode(Response{Message:token})
        return
     }


}
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
    userRes :=UserResponse{
        Username: user.Username,
        Email: user.Email,
    }
    if err:= json.NewEncoder(w).Encode(userRes); err!=nil{
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
 
    fetcheduser, err := h.Service.PostUser(r.Context(), &user)

    if err !=nil{
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
    }

    if err := json.NewEncoder(w).Encode(fetcheduser); err != nil{
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
    fetcheduser, err := h.Service.UpdateUser(r.Context(), &user)
    
    if err != nil{
        log.Println(err)
        json.NewEncoder(w).Encode(Response{Message:err.Error()})
        return
    }
   
    userRep  := UserResponse{
            Username: fetcheduser.Username,
            Email: fetcheduser.Email,    
    }
    
    if err := json.NewEncoder(w).Encode(userRep); err != nil{
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
}   