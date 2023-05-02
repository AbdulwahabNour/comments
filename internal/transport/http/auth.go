package http

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)


var (
      
      mySigningKey = []byte("Your-secret-key")

      ErrorNotAuthorized = errors.New("not authorized")
      

)
func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){

      return func(w http.ResponseWriter, r *http.Request){


            token, err := getTokenFromHeader(r)
            if err != nil{
                  http.Error(w, ErrorNotAuthorized.Error(), http.StatusUnauthorized)
                  return
              }

           
            if !validateToken(token ){
                  http.Error(w, ErrorNotAuthorized.Error(), http.StatusUnauthorized)
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

func getClaimsByToken(r *http.Request)(jwt.Claims, error) {
      
      accesstoken, err := getTokenFromHeader(r)
      if err != nil{
         return nil,  ErrorNotAuthorized
        }
          
      token, err :=  jwt.Parse(accesstoken, func(t *jwt.Token) (interface{}, error){ 

            if _, ok := t.Method.(*jwt.SigningMethodHMAC); ! ok{
                  return nil, errors.New("could not validate auth token")
            }
            return mySigningKey, nil
      })

      if err != nil{
            return nil, ErrorNotAuthorized
      }
 

      return  token.Claims, ErrorNotAuthorized
}

func generateJwtToken(claims jwt.MapClaims) (string,error){
            token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 
            tokenString, err := token.SignedString(mySigningKey)
            if  err != nil {
                  return "", err
            }
 
            return tokenString,nil
}


func getTokenFromHeader(r *http.Request)(string, error){
      
      authHeader := r.Header["Authorization"]

      if authHeader == nil{
          return "", ErrorNotAuthorized
      }

      authHeaderParts := strings.Split(authHeader[0], " ")

      if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "token"{
            return "", ErrorNotAuthorized
      }
      return authHeaderParts[1], nil
}