package main

import (
	"context"
	"entgo.io/ent/entc/integration/ent"
	"fmt"
	gin "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sample/routers"
)

func main() {

	// 진 라우터 생성
	router := gin.Default()

	// 정적 페이지 템플릿 경로 지정
	//router.LoadHTMLGlob("templates/**")

	// MySQL : Ent Go 연동
	client, err := ent.Open("mysql", "springetc:springetc12#@tcp(127.0.0.1:3306)/go_sample?parseTime=True")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	defer client.Close()

	ctx := context.Background()
	// create schema
	if _, err := client.ExecContext(ctx, "CREATE SCHEMA example_schema"); err != nil {
		fmt.Println(err)
		return
	}

	// create table
	if _, err := client.ExecContext(ctx, "CREATE TABLE example_schema.example_table (id INT, name VARCHAR(255))"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Schema, table and columns created!")

	// 라우터 묶음 적용
	fmt.Println("INIT 라우터 실행")
	routers.InitRouter(router)

	// 서버 실행
	router.Run()
}
