package main

import (
	route "github.com/sergiohdljr/aprove-me-go/pkg/routes"
)

func main() {
	r := route.RouteInnit()

	r.Run()
}
