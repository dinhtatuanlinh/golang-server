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
	
	str, err := json.Marshal(&SendingJson{Bar: "this route is in group!!"})
	if err != nil {
		w.Write([]byte("parse JSON err"))
	}
	w.WriteHeader(404)
	w.Write(str)
}
