package hash

import (
	"testing"

	"github.com/AbdulwahabNour/comments/internal/utils/rand"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)


func TestPass(t *testing.T){
    pass, err := rand.String(15)
    require.NoError(t, err)
    require.NotEmpty(t, pass)

    hashedPass, err := HashPassword(pass)
    require.NoError(t, err)
    require.NotEmpty(t, hashedPass)

    err = CheckPassword(pass, hashedPass)
    require.NoError(t, err)

    wrongPass, err := rand.String(15)
    require.NoError(t, err)
    require.NotEmpty(t, wrongPass)
    
    err = CheckPassword(wrongPass, hashedPass)
    require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

    hashedPass2, err := HashPassword(pass)
    require.NoError(t, err)
    require.NotEmpty(t, hashedPass2)
    require.NotEqual(t, hashedPass, hashedPass2)


}


