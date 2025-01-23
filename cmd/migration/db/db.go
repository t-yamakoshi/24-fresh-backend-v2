package db

import (
	"fmt"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen"
)

func NewConfig(dsn string) (*mysql.Config, error){
	if dsn == "" {
		return nil, fmt.Errorf("dsn is empty")
	}
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func NewClient(cfg *mysql.Config) (*entgen.Client, error){
	client, err := entgen.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return client, nil
}
