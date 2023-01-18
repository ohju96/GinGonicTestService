package config

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"user-service/ent"
)

var DB *ent.Client

func EntConfig() {

	client, err := ent.Open("mysql", "mydb:1234@tcp(192.168.199.128:3306)/go_sample?parseTime=True")

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
