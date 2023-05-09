package service

import (
	"context"
	"log"

	"github.com/AbdulwahabNour/comments/internal/model"
)

 
type commentFunc func(*model.Comment) error


func(s *Service) GetComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){

    cmt, err := s.Store.GetComment(ctx, cmt)
    
    if err != nil {
        log.Println(err)
        return &model.Comment{}, ErrFetchingComment
    }

    return cmt, err
}
 

func(s *Service)  UpdateComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){
    
    err := runCommentValidate(cmt, minCommentBodyLength, maxCommentBodyLength)
    if err != nil{
        return nil, err
      }
     
    cmt, err = s.Store.UpdateComment(ctx, cmt)
    if err != nil{
        log.Println(err)
        return &model.Comment{}, ErrUpdatingComment
    }
    return  cmt, nil
}
 
func(s *Service)  DeleteComment(ctx context.Context, cmt *model.Comment) error{
    
    
     err := s.Store.DeleteComment(ctx, cmt)
     
    if err != nil{
        log.Println(err)
        return  ErrDeletingComment 
    }


    return  nil
}
func(s *Service) PostComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){
   
    err := runCommentValidate(cmt, minCommentBodyLength, maxCommentBodyLength)
    if err != nil{
    return nil, err
    }
 
   
    comment, err := s.Store.PostComment(ctx, cmt)
    if err != nil {
        log.Println(err)
        return nil, ErrInsertComment
    }

    return comment, err
   
}