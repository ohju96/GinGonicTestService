package config

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sample/ent"
)

var DB *ent.Client

func EntConfig() {
	client, err := ent.Open("mysql", "springetc:springetc12#@tcp(127.0.0.1:3306)/go_sample?parseTime=True")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	DB = client
}

type User struct {
	Email    string
	Password string
	Name     string
	Age      int
}
