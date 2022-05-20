package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/mitchellh/mapstructure"

	"server/pkg/ulti"
	"server/web"
)

type Config struct {
	Config struct {
		Database struct {
			Environment struct {
				POSTGRES_DB       string `yaml:"POSTGRES_DB"`
				POSTGRES_USER     string `yaml:"POSTGRES_USER"`
				POSTGRES_PASSWORD int64  `yaml:"POSTGRES_PASSWORD"`
				Ports             int64  `yaml:"ports"`
			} `yaml:"environment"`
		} `yaml:"database"`
		Secretkey string `yaml:"secretkey"`
		Salt      string `yaml:"salt"`
	} `yaml:"config"`
}

type MyStruct struct {
	Name string
	Age  int64
}

func main() {
	config := &Config{}
	result, err := ulti.ReadFile("./configs/config_server.yaml")
	fmt.Printf("%T", *result)
	fmt.Println()
	if err != nil {
		fmt.Println(err)
	} else {
		mapstructure.Decode(*result, config)
		fmt.Println(*result)
		fmt.Println(config.Config.Salt)
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
