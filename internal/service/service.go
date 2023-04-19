package service

import (
	"errors"

	"github.com/AbdulwahabNour/comments/internal/db"
)


var(
    ErrFetchingComment = errors.New("failed to fetch comment by id")
    ErrInsertComment = errors.New("failed to insert comment by id")
    ErrDeletingComment = errors.New("failed to delete comment by id")
    ErrUpdatingComment = errors.New("failed to update comment by id")


    ErrNotImplemented = errors.New("not implemented ")
    ErrEmptyFilds = errors.New("we don't accept empty field ")
    ErrFetchingUser = errors.New("failed to fetch user by id")

    ErrInsertUser = errors.New("failed to insert user by id")
    ErrDeletingUser = errors.New("failed to delete user by id")
    ErrUpdatingUser = errors.New("failed to update user by id")
)
 

type Service struct{
    Store db.Service
}

func NewService(s db.Service) *Service{
    return &Service{
        Store: s,
    }
}