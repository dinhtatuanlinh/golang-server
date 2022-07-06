package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handle interface {
	Welcome(w http.ResponseWriter, r *http.Request)
	Abc(w http.ResponseWriter, r *http.Request)
	Articles(w http.ResponseWriter, r *http.Request)
	NotFound(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}
type SendingJson struct {
	Bar string
}
type Handlers struct {
}

func (h *Handlers) Welcome(w http.ResponseWriter, r *http.Request) {
	// db := database.GetConnectionInstance()
	// status := db.CheckTableExist("labs", "users", db.DB)
	// fmt.Println(status)
	w.Write([]byte("welcome"))
}

func (h *Handlers) Abc(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	w.Write([]byte("welcome"))
}

func (h *Handlers) Articles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("this is subRoute!!!"))
}

func (h *Handlers) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("route does not exist!!!!"))
}

func (h *Handlers) Post(w http.ResponseWriter, r *http.Request) {
	str, err := json.Marshal(&SendingJson{Bar: "this route is in group!!"})
	if err != nil {
		w.Write([]byte("parse JSON err"))
	}
	w.WriteHeader(404)
	w.Write(str)
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var data user
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		w.Write([]byte("json decode err"))
	}
	

	if (data.Password != data.Repassword) {
		res, err := json.Marshal(&response{Code: 500, Message: "password and repassword do not match!"})
		if err != nil{
			w.Write([]byte("parse JSON err"))
		}
		w.WriteHeader(404)
		w.Write(res)
	}

	fmt.Println(data)
	res, err := json.Marshal(&response{Code: 200, Message: "success!"})
	if err != nil {
		w.Write([]byte("parse JSON err"))
	}
	w.WriteHeader(404)
	w.Write(res)
}

type user struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Repassword string `json:"repassword"`
}
type response struct{
	Code int
	Message string
}