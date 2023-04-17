package comment

import (
	"context"
	"errors"
	"fmt"
	"time"
)
var(
    ErrFetchingComment = errors.New("failed to fetch comment by id")
    ErrNotImplemented = errors.New("not implemented ")
    ErrEmptyFilds = errors.New("we don't accept empty field ")
)
 
// Representation of the comment structure
type Comment struct{
    ID string
    Slug string
    Body string
    Author string
}

// Defines all of the methods
// our srvice need in order to operate
type Store interface {
    GetComment( context.Context,  string)(Comment, error)
    PostComment( context.Context,  Comment)(Comment, error)
    DeleteComment( context.Context,  string) (error)
    UpdateComment( context.Context,  Comment)(Comment, error)
} 
type Service struct{
    Store Store
}

func NewService(s Store) *Service{
    return &Service{
        Store: s,
    }
}

func(s *Service) GetComment(ctx context.Context, id string)(Comment, error){
    cmt, err := s.Store.GetComment(ctx, id)
    if err != nil {
        
        return Comment{}, ErrFetchingComment
    }

    return cmt, err
}
 

func(s *Service)  UpdateComment(ctx context.Context, cmt Comment)(Comment, error){
    
    if  cmt.Author == "" || cmt.Body == "" || cmt.Slug == ""{
        return Comment{}, ErrEmptyFilds
    }

    comment, err := s.Store.UpdateComment(ctx, cmt)
    if err != nil{
        return Comment{}, err
    }
    return  comment, nil
}
 
func(s *Service)  DeleteComment(ctx context.Context, uuid string) error{
    time.Sleep(20 * time.Second)
     err := s.Store.DeleteComment(ctx, uuid)
     
    if err != nil{
        return fmt.Errorf("%w,  %w", ErrFetchingComment, err) 
    }


    return  nil
}
func(s *Service) PostComment(ctx context.Context, cmt Comment)(Comment, error){
    if  cmt.Author == "" || cmt.Body == "" || cmt.Slug == ""{
        return Comment{}, ErrEmptyFilds
    }
    comment, err := s.Store.PostComment(ctx, cmt)
    if err != nil {
        
        return Comment{}, ErrFetchingComment
    }

    return comment, err
   
}