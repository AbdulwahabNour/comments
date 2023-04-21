package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)
var mySigningKey = []byte("Your-secret-key")
func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){

      return func(w http.ResponseWriter, r *http.Request){
 
            authHeader := r.Header["Authorization"]
            if authHeader == nil{
                http.Error(w, "not authoried", http.StatusUnauthorized)
                return
            }
 
            authHeaderParts := strings.Split(authHeader[0], " ")

            if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "token"{
                  http.Error(w, "not authoried", http.StatusUnauthorized)
                  return
            }

            if !validateToken(authHeaderParts[1] ){
                  http.Error(w, "not authoried", http.StatusUnauthorized)
                  return
            }

            original(w, r)

      }
 
}

func validateToken(accesstoken string)bool{
     
      token, err :=  jwt.Parse(accesstoken, func(t *jwt.Token) (interface{}, error){ 

                  if _, ok := t.Method.(*jwt.SigningMethodHMAC); ! ok{
                        return nil, errors.New("could not validate auth token")
                  }
                  return mySigningKey, nil
      })

      if err != nil{
            return false
      }
   

      return token.Valid
}

func generateJwtToken(clams jwt.MapClaims) (string,error){
            token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)

            tokenString, err := token.SignedString(mySigningKey)
            if  err != nil {
                  return "", err
            }
            return tokenString,nil
}