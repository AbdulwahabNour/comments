package service

import (
	"errors"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/asaskevich/govalidator"
)


func runCommentValidate(comment *model.Comment, fn ...commentFunc)error{
        for _, f:= range fn{
            err :=f(comment)
            if err != nil {
                return err
             }
        }
    return nil
}
 

func maxCommentBodyLength(comment *model.Comment) error{

    if !govalidator.MaxStringLength(comment.Body, "1500"){
        return errors.New("comment body must be less than 1500 character")
    }
    return nil
}

func minCommentBodyLength(comment *model.Comment) error{
   
    if !govalidator.MinStringLength(comment.Body, "25" )   {
        return errors.New("comment body must be more than 25 character")
    }
    return nil
}

 