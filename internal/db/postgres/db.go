package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


type Database struct{
    Client *sqlx.DB
}

func NewDatabase()(*Database, error) {

    dataConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
                            os.Getenv("DB_HOST"),
                            os.Getenv("DB_PORT"),
                            os.Getenv("DB_USERNAME"),
                            os.Getenv("DB_TABLE"),
                            os.Getenv("DB_PASSWORD"),
                            os.Getenv("SSL_MODE"))
    conn, err := sqlx.Connect("postgres", dataConn)
   
    if  err != nil {
            return &Database{}, fmt.Errorf("could not  connect to the database: %w", err)
    }

    return &Database{Client:conn}, nil
}

func(d *Database)Ping(ctx context.Context) error{
    return d.Client.DB.PingContext(ctx)
}
