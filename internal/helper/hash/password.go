package hash

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var hashPepper = "Hash_Pepper"

func HashPassword(pass string) (string, error){
    // Add pepper to immune to brute force attacks and password cracking using dictionary tables and rainbow tables
    hashpass, err := bcrypt.GenerateFromPassword([]byte(pass + hashPepper), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    return string(hashpass), nil
}

func CheckPassword(pass string, hashedPassword string) error {

    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pass + hashPepper))
}