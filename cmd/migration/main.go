package main

import (
	"context"
	"log"
	"os"

	"github.com/t-yamakoshi/24-fresh-backend-v2/cmd/migration/db"
)

func main() {
    dsn := "root:@tcp(localhost:23306)/"
    cfg, err := db.NewConfig(dsn)
    if err != nil {
        log.Fatalf("failed creating mysql config: %v", err)
    }
    cfg.ParseTime = true
    cfg.DBName = "24-fresh"
    client, err := db.NewClient(cfg)
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




