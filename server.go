package main

import (
	"golang-server/web"
	"log"
	"net/http"
)

func main() {
	web.WebRouter()
    log.Fatal(http.ListenAndServe(":8081", nil))

}