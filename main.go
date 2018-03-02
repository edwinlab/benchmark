package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"

	"./benchmark"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	dsl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USERNAME, PASSWORD, DB_NAME)
	dbConn, err := gorm.Open("postgres", dsl)

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	defer dbConn.Close()
	app.Get("/loads", func(ctx iris.Context) {
		repo := benchmark.NewRepository(dbConn)
		service := benchmark.NewService(repo)
		service.UpdateRandom()
		ctx.Writef(`ok`)
	})

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
