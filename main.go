package main

import (
	"net/http"
	"web-application-golang/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
