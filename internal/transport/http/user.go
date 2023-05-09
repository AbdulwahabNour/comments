package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/AbdulwahabNour/comments/internal/utils/hash"
	"github.com/golang-jwt/jwt/v5"
)

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var creds model.UserCredentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Message: "something wrong happened please try again later"})
		return
	}

	user, err := h.Service.GetUserByEmail(r.Context(), &model.User{Email: creds.Email})

	if err != nil {
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}

	err = hash.CheckPassword(user.Password, creds.Password)
	if err != nil {
 
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}

	jwtClaims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	token, err := generateJwtToken(jwtClaims)
 

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Message: "there is something wrong with token please try again later"})
		return
	}

	w.Header().Add("Authorization", fmt.Sprintf("token %s", token))

	if err != nil {
		json.NewEncoder(w).Encode(Response{Message: token})
		return
	}

}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	claim, err := getClaimsByToken(r)

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Message: "invalid authentication"})
		return
	}

 

	user.Email = claim["email"].(string)

	fetcheduser, err := h.Service.GetUserByEmail(r.Context(), &user)

	if err != nil {
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}

	userRes := UserResponse{
		Username: fetcheduser.Username,
		Email:    fetcheduser.Email,
	}

	if err := json.NewEncoder(w).Encode(userRes); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (h *Handler) PostUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fetcheduser, err := h.Service.PostUser(r.Context(), &user)

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}

	jwtClaims := jwt.MapClaims{
		"email":    fetcheduser.Email,
		"username": fetcheduser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token, err := generateJwtToken(jwtClaims)

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Message: "there is something wrong with token please try again later"})
		return
	}
	w.Header().Add("Authorization", fmt.Sprintf("token %s", token))
	if err := json.NewEncoder(w).Encode(fetcheduser); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
 
 
 
	fetcheduser, err  := h.GetUserFromToken(r)

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Message: "invalid authentication"})
		return
	}

    
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Message: "there is something wrong please try again later"})
		return
	}
	//////////////////////////

	err = h.Service.DeleteUser(r.Context(), fetcheduser.ID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(Response{Message: "Done"})
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User
 

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}



	// claim, err := getClaimsByToken(r)

	// if err != nil {
	// 	log.Println(err)
	// 	json.NewEncoder(w).Encode(Response{Message: "invalid authentication"})
	// 	return
	// }

	// fetcheduser, err  := h.Service.GetUserByEmail(r.Context(), &model.User{Email:claim["email"].(string) })
    
	fetcheduser,err := h.GetUserFromToken(r)

	fmt.Println("update  ",user,"\n",fetcheduser)

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Message: "invalid id"})
		return
	}

	user.ID = fetcheduser.ID

	_, err = h.Service.UpdateUser(r.Context(), &user)

	fmt.Println("update222  ",user,"\n",fetcheduser)
	if err != nil {
		log.Println(err)
		
		json.NewEncoder(w).Encode(Response{Message: err.Error()})
		return
	}
     
	jwtClaims := jwt.MapClaims{
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token, err := generateJwtToken(jwtClaims)

	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Message: "there is something wrong with token please try again later"})
		return
	}
	w.Header().Add("Authorization", fmt.Sprintf("token %s", token))

	
	userRep := UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	if err := json.NewEncoder(w).Encode(userRep); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func(h *Handler)GetUserFromToken(r *http.Request)(*model.User, error){
    
	claim, err := getClaimsByToken(r)

	if err != nil {
		return nil, err
	}
 
	user, err := h.Service.GetUserByEmail(r.Context(), &model.User{Email: claim["email"].(string)})
 
	if err != nil {
		 
		return nil, err
	}
	return user, nil
}
