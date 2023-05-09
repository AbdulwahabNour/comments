package http

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	mySigningKey = []byte("Your-secret-key")

	ErrorNotAuthorized = errors.New("not authorized")
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		token, err := getTokenFromHeader(r)

		if err != nil {
			http.Error(w, ErrorNotAuthorized.Error(), http.StatusUnauthorized)
			return
		}

		status, err := validateToken(token)
		if !status {
			log.Println(err)
			http.Error(w, ErrorNotAuthorized.Error(), http.StatusUnauthorized)
			return
		}

		original(w, r)

	}

}

func validateToken(accesstoken string) (bool, error) {

	token, err := jwt.Parse(accesstoken, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		log.Println(err)
		return false, err
	}

	return token.Valid, nil
}

func getClaimsByToken(r *http.Request) (jwt.MapClaims, error) {

	var claims jwt.MapClaims

	accesstoken, err := getTokenFromHeader(r)
	if err != nil {
		return nil, ErrorNotAuthorized
	}

	_, err  = jwt.ParseWithClaims(accesstoken, &claims, func(token *jwt.Token) (interface{}, error) {
        return mySigningKey, nil  // Replace with your secret key
    })

	if err != nil {
		return nil, err
	}
   
	return claims, nil
}

func generateJwtToken(claims jwt.MapClaims) (string, error) {
	// var err error

	// claims["UUID"], err = rand.String(32)

	// if err != nil {
	// 	return "", err
	// }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
   
	return tokenString, nil
}

func getTokenFromHeader(r *http.Request) (string, error) {

	authHeader := r.Header["Authorization"]

	if authHeader == nil {
		return "", ErrorNotAuthorized
	}

	authHeaderParts := strings.Split(authHeader[0], " ")

	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "token" {
		return "", ErrorNotAuthorized
	}
	return authHeaderParts[1], nil
}

// func checkExpiration(claims jwt.MapClaims) (bool,error) {

// 	exp, err := claims.GetExpirationTime()

// 	if err != nil {
// 		return false, err
// 	}
 
// 	fmt.Println(exp,  exp.Before(time.Now()))
// 	return exp.Before(time.Now()), nil
// }
