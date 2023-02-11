package config

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"user_service/ent"
)

var DB *ent.Client

func InitDB(c *Config) {

	dbUrl := c.Db.User + ":" + c.Db.Password + "@" + c.Db.Host + "/" + c.Db.Db + "?parseTime=True"

	client, err := ent.Open(c.Db.Dbms, dbUrl)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()
	DB = client

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	fmt.Println(c.Db.Dbms, " connected successfully !!")
}
