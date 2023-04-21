package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AbdulwahabNour/comments/internal/model"
	"github.com/google/uuid"
)

type CommentRow struct{
    ID string
    Slug sql.NullString
    Body sql.NullString
    Author sql.NullString
}

func convertCommentRowtoComment(c *CommentRow) *model.Comment{
    return &model.Comment{

        ID: c.ID,
        Slug: c.Slug.String,
        Body: c.Body.String,
        Author: c.Author.String,
    }
}
func (d *Database) GetComment(ctx context.Context, uuid string)(*model.Comment, error){
    var cmtRow CommentRow
  
    row :=  d.Client.QueryRowContext(ctx,`SELECT * FROM comments WHERE id = $1`, uuid)

    err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Author, &cmtRow.Body)
    if err != nil{
        return &model.Comment{}, fmt.Errorf("error fetching the comment by uuid: %s, err => %w", uuid, err )
    }
  
   
    return convertCommentRowtoComment(&cmtRow), nil
}
func(d *Database) PostComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){

        cmt.ID = uuid.New().String()
        postRow := CommentRow{
            ID: cmt.ID,
            Slug: sql.NullString{String:cmt.Slug, Valid: true},
            Body: sql.NullString{String:cmt.Body, Valid: true},
            Author: sql.NullString{String:cmt.Author, Valid: true},
        }
        rows, err := d.Client.NamedQueryContext(ctx,`INSERT INTO comments(id, slug, author, body) VALUES (:id, :slug, :author, :body)`, postRow)
        if err != nil{
            return &model.Comment{},  fmt.Errorf("failed insert the comment err: %w", err)
        }
        if err := rows.Close(); err != nil{
            return &model.Comment{}, fmt.Errorf("failed to close rows: %w", err)
        }
   return convertCommentRowtoComment(&postRow), nil
}

func(d *Database) DeleteComment(ctx context.Context, uuid string) error{
 
    _, err := d.Client.ExecContext(ctx, `DELETE FROM comments WHERE id = $1 `, uuid)
 
    if err != nil{
        return  fmt.Errorf( err.Error())
    }


    return nil
}

func(d *Database) UpdateComment(ctx context.Context, cmt *model.Comment)(*model.Comment, error){
 
  
    rows, err := d.Client.NamedQueryContext(ctx, ` UPDATE comments SET slug = :slug, author = :author, body = :body WHERE id = :id`, cmt)
 
    if err != nil{
        return  &model.Comment{}, fmt.Errorf("error update data error is %w", err)
    }
    if err := rows.Close(); err != nil{
        return &model.Comment{}, fmt.Errorf("failed to close rows: %w", err)
    }

    return  cmt, nil
}