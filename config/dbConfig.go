package config

import (
	"context"
	"entgo.io/ent/entc/integration/ent"
	"log"
)

var DB *ent.Client

func MySQLConnection() {

	client, err := ent.Open("mysql", "springetc:springetc12#@tcp(127.0.0.1:3306)/sample?parseTime=True")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	DB = client
}
