package service

import (
	"errors"

	"github.com/AbdulwahabNour/comments/internal/db"
)


var(
    ErrFetchingComment = errors.New("failed to fetch comment ")
    ErrInsertComment = errors.New("failed to insert comment ")
    ErrDeletingComment = errors.New("failed to delete comment ")
    ErrUpdatingComment = errors.New("failed to update comment ")


    ErrNotImplemented = errors.New("not implemented ")
    ErrEmptyFilds = errors.New("we don't accept empty field ")
    ErrFetchingUser = errors.New("failed to fetch user ")
    ErrGetUserByEmil = errors.New("wrong email or password")

    ErrInsertUser = errors.New("failed to insert user ")
    ErrDeletingUser = errors.New("failed to delete user ")
    ErrUpdatingUser = errors.New("failed to update user ")
)
 

type Service struct{
    Store db.Service
}

func NewService(s db.Service) *Service{
    return &Service{
        Store: s,
    }
}