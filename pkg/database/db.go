// pkg/database/db.go
package database

import (
    "context"
    "fmt"
    "log"

    "github.com/jackc/pgx/v5/pgxpool"
    "leapfrog-assignment/configs"
)

var DB *pgxpool.Pool

func ConnectDB(cfg configs.Config) {
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
        cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

    pool, err := pgxpool.New(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }

    DB = pool
}
