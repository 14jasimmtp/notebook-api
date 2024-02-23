package main

import (
	"log"

	api "github.com/14jasimmtp/notebook-api/pkg/apis"
	"github.com/14jasimmtp/notebook-api/pkg/config"
	"github.com/14jasimmtp/notebook-api/pkg/db"
	"github.com/14jasimmtp/notebook-api/pkg/repository"
	"github.com/14jasimmtp/notebook-api/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("error loading envs", err)
	}
	db := db.Connection(c.DB_URL)

	repo := repository.NewRepo(db)
	api := api.NewApi(repo)
	r := gin.Default()

	routes.Notebook(r, api)

	r.Run(":3000")

}
