package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/google/uuid"
)

type CommentRow struct{
    ID string `db:"id"`
    UserId int64 `db:"user_id"`
    Body sql.NullString `db:"body"`
   
}

func convertCommentRowtoComment(c *CommentRow) *model.Comment{
    return &model.Comment{

        ID: c.ID,
        UserId: c.UserId,
        Body: c.Body.String,
    }
}
func (d *Database) GetComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){
    var cmtRow *model.Comment
  
    row :=  d.Client.QueryRowContext(ctx,`SELECT * FROM comments WHERE id = $1 AND user_id = $2`, cmt.ID, cmt.UserId)

    err := row.Scan(&cmtRow.ID, &cmtRow.UserId, &cmtRow.Body)
    if err != nil{
        return &model.Comment{}, fmt.Errorf("error fetching the comment by uuid: %s, err => %w", cmt.ID, err )
    }
  
   
    return  cmtRow, nil
}
func(d *Database) PostComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){

        cmt.ID = uuid.New().String()

        rows, err := d.Client.NamedQueryContext(ctx,`INSERT INTO comments(id, user_id, body) VALUES (:id, :user_id, :body)`, cmt)
    
        if err != nil{
            return &model.Comment{},  fmt.Errorf("failed insert the comment err: %w", err)
        }
        if err := rows.Close(); err != nil{
            return &model.Comment{}, fmt.Errorf("failed to close rows: %w", err)
        }
   return cmt, nil
}

func(d *Database) DeleteComment(ctx context.Context,cmt *model.Comment) error{
 
    _, err := d.Client.ExecContext(ctx, `DELETE FROM comments WHERE id = $1 AND user_id = $2`, cmt.ID, cmt.UserId)
 
    if err != nil{
        return  fmt.Errorf( err.Error())
    }


    return nil
}

func(d *Database) UpdateComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){
 
  
    rows, err := d.Client.NamedQueryContext(ctx, ` UPDATE comments SET  body = :body WHERE id = :id AND user_id = :user_id`, cmt)
 
    if err != nil{
        return  &model.Comment{}, fmt.Errorf("error update data error is %w", err)
    }
    if err := rows.Close(); err != nil{
        return &model.Comment{}, fmt.Errorf("failed to close rows: %w", err)
    }

    return  cmt, nil
}