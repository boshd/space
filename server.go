package main

import (
	"github.com/kareemarab/space/api"
	"github.com/kareemarab/space/models"

	// "github.com/kareemarab/space/models"
	"github.com/kareemarab/space/routes"
	"github.com/urfave/negroni"
)

func main() {
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
