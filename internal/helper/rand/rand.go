package rand

import (
	"encoding/base64"
	"math/rand"
)

/*
** Bytes will generate n random  bytes
** or return an error if there was one
 */
func Byte(n uint) ([]byte, error){
        b := make([]byte, n)
        _, err := rand.Read(b)
        if err != nil{
            return nil, err
        }
        return b, nil
}
 /*
 ** String will generate a byte slice of size numBytes
 ** and return a string that is the base64 url encoded
 */
 func String(numBytes uint)(string, error){
    b, err := Byte(numBytes)
    if err != nil {
        return "", err
    }

    return base64.URLEncoding.EncodeToString(b), nil 
}