package db

import (
	"context"

	"github.com/AbdulwahabNour/comments/internal/model"
)

 

type Service interface{
    GetComment( context.Context,  *model.Comment)(*model.Comment, error)
    PostComment( context.Context,  *model.Comment)(*model.Comment, error)
    DeleteComment( context.Context,  *model.Comment) (error)
    UpdateComment( context.Context,  *model.Comment)(*model.Comment, error)
    
    GetUser( context.Context,  int64)(*model.User, error)
    GetUserByEmail(ctx context.Context, data *model.User) (*model.User, error)
    PostUser( context.Context,  *model.User)(*model.User, error)
    DeleteUser( context.Context,  int64) (error)
    UpdateUser( context.Context,  *model.User)(*model.User, error)
}
