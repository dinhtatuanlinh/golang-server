package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"server/web"
)


func main() {
	// db := database.Connection()
	// query := fmt.Sprintf(`insert into users (created_at, username, email, "password") values (?, ?, ?, ?)`)

	// result := db.Exec(query, "asdfwefd", "e3rdfdsew", "refdsfsd", "asdfwefdscxv" )
	// fmt.Println(*result)
// 	rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()
// defer rows.Close()
// for rows.Next() {
//   rows.Scan(&name, &age, &email)

//   // do something
// }
	// database.CreateDatabase(db)
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
	
	err := http.ListenAndServe(":3000", r)
	fmt.Println(err)
}
