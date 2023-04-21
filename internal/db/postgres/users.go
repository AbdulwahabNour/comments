package postgres

import (
	"context"
	"fmt"

	"github.com/AbdulwahabNour/comments/internal/model"
)



 
func(db *Database) GetUser(ctx context.Context, id int64)(*model.User, error){
   var user model.User
    row := db.Client.QueryRowContext(ctx, `SELECT * FROM users WHERE id = $1`, id)
    err := row.Scan(
        &user.ID,
        &user.Username,
        &user.Password,
        &user.Email,
        &user.CreatedAt)
    if err != nil{
       
        return &model.User{}, fmt.Errorf("error fetching user by id %v err:%w", id, err)
    }
    return &user, nil
}

func(db *Database) GetUserByEmail(ctx context.Context, data *model.User)(*model.User, error){
    var user model.User
    row := db.Client.QueryRowContext(ctx, `SELECT * FROM users WHERE email = $1`, data.Email)
    err := row.Scan(
        &user.ID,
        &user.Username,
        &user.Password,
        &user.Email,
        &user.CreatedAt)

        if err != nil{
       
            return &model.User{}, fmt.Errorf("error fetching user by email and password %v err:%w", data, err)
        }
        return &user, nil
}

func(db * Database) PostUser(ctx context.Context, user *model.User)(*model.User, error){
    var u model.User
    row  := db.Client.QueryRowContext(ctx, `INSERT INTO users(username, password, email) VALUES ($1, $2, $3) RETURNING id, username, password, email, create_at`, user.Username, user.Password, user.Email)
    err :=  row.Scan(
        &u.ID,
        &u.Username,
        &u.Password,
        &u.Email,
        &u.CreatedAt)
    if err != nil{
        return &model.User{}, fmt.Errorf("failed insert user : %w", err)
    }
 
    return &u, nil
}


func(db *Database) DeleteUser( ctx context.Context, id string) (error){
    _, err := db.Client.ExecContext(ctx, `DELETE FROM users WHERE id = $1 `, id)
 
    if err != nil{
        return  fmt.Errorf(err.Error())
    }


    return nil
   
}
func(db *Database) UpdateUser(ctx context.Context, user *model.User )(*model.User, error){
     

    row := db.Client.QueryRowContext(ctx, `UPDATE users SET username=$2, password= $3, email= $4 WHERE id= $1 RETURNING id, username, password, email, create_at`, user.ID, user.Username, user.Password, user.Email)
    
 
    err :=  row.Scan(
        &user.ID,
        &user.Username,
        &user.Password,
        &user.Email,
        &user.CreatedAt)
 
     
    if err != nil{
        return &model.User{},  fmt.Errorf("error update data error is %w", err)
    }
     
     return user, nil
}