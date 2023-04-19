package service

import (
	"context"
	"log"

	"github.com/AbdulwahabNour/comments/internal/model"
)

 

func(s *Service) GetComment(ctx context.Context, id string)(*model.Comment, error){
    cmt, err := s.Store.GetComment(ctx, id)
    if err != nil {
        log.Println(err)
        return &model.Comment{}, ErrFetchingComment
    }

    return cmt, err
}
 

func(s *Service)  UpdateComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){
    
    if  cmt.Author == "" || cmt.Body == "" || cmt.Slug == ""{
        return &model.Comment{}, ErrEmptyFilds
    }

    comment, err := s.Store.UpdateComment(ctx, cmt)
    if err != nil{
        log.Println(err)
        return &model.Comment{}, ErrUpdatingComment
    }
    return  comment, nil
}
 
func(s *Service)  DeleteComment(ctx context.Context, uuid string) error{
    
    
     err := s.Store.DeleteComment(ctx, uuid)
     
    if err != nil{
        log.Println(err)
        return  ErrDeletingComment 
    }


    return  nil
}
func(s *Service) PostComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){
    if  cmt.Author == "" || cmt.Body == "" || cmt.Slug == ""{
        return &model.Comment{}, ErrEmptyFilds
    }
    comment, err := s.Store.PostComment(ctx, cmt)
    if err != nil {
        log.Println(err)
        return &model.Comment{}, ErrInsertComment
    }

    return comment, err
   
}