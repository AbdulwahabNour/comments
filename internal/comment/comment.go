package comment

import (
	"context"
	"errors"
)
var(
    ErrFetchingComment = errors.New("failed to fetch commet by id")
    ErrNotImplemented = errors.New("not implemented ")
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
    GetComment(ctx context.Context, id string)(Comment, error)
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
 

func(s *Service) UpdateComment(ctx context.Context, com Comment)error{
    
    return  ErrNotImplemented
}
 
func(s *Service) DeleteComment(ctx context.Context, id  int)error{
    
    return  ErrNotImplemented
}
func(s *Service) CreateComment(ctx context.Context, com Comment)(Comment, error){
    
    return  Comment{ },ErrNotImplemented
}