package main

import (
	"github.com/sergiohdljr/aprove-me-go/pkg/database"
	route "github.com/sergiohdljr/aprove-me-go/pkg/routes"
)

func main() {
	r := route.RouteInnit()

	database.InnitDB()

	r.Run()
}
