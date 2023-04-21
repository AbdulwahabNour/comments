package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/AbdulwahabNour/comments/internal/utils/hash"
)

type  userFunc func(*model.User) error

func(s *Service) GetUser(ctx context.Context, id int64)(*model.User, error){

    user, err := s.Store.GetUser(ctx, id)

    if err != nil {
        log.Println(err)
        return &model.User{}, ErrFetchingUser
    }

    return  user, err
}
func(s *Service)  GetUserByEmail(ctx context.Context, user *model.User)(*model.User, error){

        fmt.Println(user)
        err := runUservalidate(user, emailRequire, checkEmail, passwordRequire)
        if err != nil{
            
             return &model.User{}, err
        }
 
        user, err = s.Store.GetUserByEmail(ctx, user)

        

        if err != nil{
            log.Println(err)
            return &model.User{}, ErrGetUserByEmil
        }


return user, nil
}
func(s * Service) PostUser(ctx context.Context, user *model.User)(*model.User, error){
    
    err := runUservalidate(user,usernameRequire,
                                maxUsernameLength,
                                minUsernameLength,
                                passwordRequire,
                                emailRequire,
                                checkEmail)
   
   if err != nil{
        return &model.User{}, err
   }
   err = setUserHashpassword(user)
    if err != nil {
            log.Println(err)
            return &model.User{}, ErrInsertUser
    }
 
   user, err = s.Store.PostUser(ctx, user)
 

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
   
   
err := runUservalidate(user, usernameRequire,
                             maxUsernameLength,
                             minUsernameLength,
                             emailRequire,
                             checkEmail)
   
   if err != nil{
        return &model.User{}, err
   }

    if strings.TrimSpace(user.Password) != "" {
        
       err := setUserHashpassword(user)
        if err != nil {
            log.Println(err)
            return &model.User{}, ErrInsertUser
        }
    }
  

    user, err = s.Store.UpdateUser(ctx,user)

         

    if err != nil{
        log.Println(err)
        return &model.User{}, ErrUpdatingUser
    }
    return  user, nil
}

func setUserHashpassword(user *model.User) error{

    newUserPass, err := hash.HashPassword(user.Password)
    if err != nil {
     return err
    }
    user.Password =newUserPass
    return nil
}



