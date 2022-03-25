package main

import (
	"fmt"
	"log"
    "web"
)

func main() {

    log.Fatal(http.ListenAndServe(":8081", nil))

}