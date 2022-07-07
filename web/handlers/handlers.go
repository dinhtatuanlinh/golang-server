package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"server/models"

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
	var resp  = models.SendingJson{Bar: "this route is in group!!"}
	str, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("parse JSON err"))
	}
	w.WriteHeader(404)
	w.Write(str)
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var data models.User
	
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		w.Write([]byte("json decode err"))
	}
	
	if (data.Password != data.Repassword) {
		// resp.Code= 500
		// resp.Message = "password and repassword do not match"
		var resp  = models.Response{Code: 500, Message: "password and repassword do not match"}
		res, err := json.Marshal(&resp)
		if err != nil{
			w.Write([]byte("parse JSON err"))
		}
		w.WriteHeader(404)
		w.Write(res)
	}

	fmt.Println(data)
	var resp  = models.Response{Code: 200, Message: "success"}
	res, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("parse JSON err"))
	}
	w.WriteHeader(404)
	w.Write(res)
}