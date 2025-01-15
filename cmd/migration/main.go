package main

import (
	"context"
	"log"
	"os"

	mysql "github.com/go-sql-driver/mysql"

	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen"
)

func main() {
    dsn := "root:@tcp(localhost:23306)/"
    cfg, err := mysql.ParseDSN(dsn)
    if err != nil {
        log.Fatalf("failed parsing dsn: %v", err)
    }
    cfg.ParseTime = true
    cfg.DBName = "24-fresh"
    client, err := entgen.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    ctx := context.Background()

    if err := client.Debug().Schema.Create(ctx); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    // マイグレーションによる変更を標準出力にダンプ
    if err := client.Debug().Schema.WriteTo(ctx, os.Stdout); err != nil {
        log.Fatalf("failed printing schema changes: %v", err)
    }
}




