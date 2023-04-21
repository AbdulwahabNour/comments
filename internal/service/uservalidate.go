package service

import (
	"errors"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/asaskevich/govalidator"
)

func runUservalidate(user *model.User, fn ...userFunc) error {
    for _, f:= range fn{
         err := f(user)
         if err != nil {
            return err
         }
    }

    return nil
}
func usernameRequire(user *model.User) error{
    if user.Username == "" {
        return errors.New("username cannot be empty")
    }
    return nil
}

func passwordRequire(user *model.User) error{
    if user.Password == "" {
        return errors.New("password cannot be empty")
    }
    return nil
}
func emailRequire(user *model.User) error{
    if user.Email == "" {
        return errors.New("Email cannot be empty")
    }
    return nil
}

func checkEmail(user *model.User) error{

     if  !govalidator.IsEmail(user.Email) {
        return errors.New("invalid email address")
    }
    return nil
}

func maxUsernameLength(user *model.User) error{
    if !govalidator.MaxStringLength(user.Username, "50" )   {
        return errors.New("username must be less than 50 character")
    }
    return nil
}
func minUsernameLength(user *model.User) error{
   
    if !govalidator.MinStringLength(user.Username, "10" )   {
        return errors.New("username must be more than 10 character")
    }
    return nil
}

 