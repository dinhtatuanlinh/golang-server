package checkget

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		fmt.Println("bafsdfsdad")
		fmt.Println(id)

	})
}
