package main

import (
	"dio_api_rest/src/api"
	"net/http"
)

func main() {
	r := api.ConfigServer()
	http.ListenAndServe(":8080", r)

}
