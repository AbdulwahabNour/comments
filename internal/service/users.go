package service

import (
	"context"
	"errors"
	"log"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/asaskevich/govalidator"
)

 

func(s *Service) GetUser(ctx context.Context, id int64)(*model.User, error){

    user, err := s.Store.GetUser(ctx, id)

    if err != nil {
        log.Println(err)
        return &model.User{}, ErrFetchingUser
    }

    return  user, err
}

func(s * Service) PostUser(ctx context.Context, user *model.User)(*model.User, error){
    
   if err := validateUserInput(user); err != nil{
        return &model.User{}, err
   }
    
   user, err := s.Store.PostUser(ctx, user)
 

    if err != nil{
        log.Println(err)
         return &model.User{}, ErrInsertUser
    }       

    return user, nil
}


func(s *Service) DeleteUser( ctx context.Context, id string) (error){
    err := s.Store.DeleteUser(ctx, id)
    if err != nil  {
        log.Println(err)
            return  ErrDeletingUser
    }
    return nil
}
func(s *Service) UpdateUser(ctx context.Context, user *model.User )(*model.User, error){
   
    if err := validateUserInput(user); err != nil{
        return &model.User{}, err
    }

    user, err := s.Store.UpdateUser(ctx,user)

    

    if err != nil{
        log.Println(err)
        return &model.User{}, ErrUpdatingUser
    }
    return  user, nil
}

 
 

func validateUserInput(user *model.User) error {
    // Check that username is not empty
    if user.Username == "" {
        return errors.New("username cannot be empty")
    }
   
    if !govalidator.MaxStringLength(user.Username, "50" )   {
        return errors.New("username must be less than 50 character")
    }

    if !govalidator.MinStringLength(user.Username, "10" )   {
        return errors.New("username must be more than 10 character")
    }
 
    // Check that password is not empty
    if user.Password == "" {
        return errors.New("password cannot be empty")
    }
    
    // // Check that email is valid
    if  !govalidator.IsEmail(user.Email) {
        return errors.New("invalid email address")
    }


    // Input data is valid
    return nil
}