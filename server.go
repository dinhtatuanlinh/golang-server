package main

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"server/pkg/ulti"
	"server/web"
)

type Config struct {
	Config struct {
		Database struct {
			Environment struct {
				POSTGRES_DB       string `yaml:"POSTGRES_DB"`
				POSTGRES_USER     string `yaml:"POSTGRES_USER"`
				POSTGRES_PASSWORD string `yaml:"POSTGRES_PASSWORD"`
				Ports             int64  `yaml:"ports"`
			} `yaml:"environment"`
		} `yaml:"database"`
		secretkey string `yaml:"secretkey"`
		salt      string `yaml:"salt"`
	} `yaml:"config"`
}

type MyStruct struct {
    Name string
    Age  int64
}

func main() {
	myData := make(map[string]interface{})
    myData["Name"] = "Tony"
    myData["Age"] = int64(23)
	result1 := &MyStruct{}
	mapstructure.Decode(myData, result1)
	fmt.Println(result1.Name)

	config := &Config{}
	result, err := ulti.ReadFile("./configs/config_server.yaml")
	if err != nil {
		fmt.Println(err)
	} else {

		mapstructure.Decode(*result, config)
		fmt.Println(*result)
	}

	r := chi.NewRouter()

	//set cors handler for all routes
	var cors = cors.New(cors.Options{
		//AllowedOrigins: []string{"https://foo.com"}, //use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		//AllowOriginFunc: func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Requested-With", "access-token", "accept-version", "Session", "Traceparent", "Tracecontext"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, //maximum value not ignored by any of major browsers
	})
	//set cors handler for all routes
	r.Use(cors.Handler)

	r.Use(middleware.Logger)

	web.Web(r)

	http.ListenAndServe(":3000", r)
}
