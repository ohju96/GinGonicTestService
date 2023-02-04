package config

import (
	"context"
	"entgo.io/ent/entc/integration/ent"
	"fmt"
	"log"
)

var DB *ent.Client

func InitDB(c *Config) {

	dbUrl := c.Db.User + ":" + c.Db.Password + "@" + c.Db.Host + "/" + c.Db.Db + "?parseTime=True"

	client, err := ent.Open(c.Db.Dbms, dbUrl)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	DB = client
	fmt.Println(c.Db.Dbms, " connected successfully !!")
}
