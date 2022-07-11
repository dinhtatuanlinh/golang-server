package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"server/database"
	"server/models"
	"server/pkg/libs/alias"
	"server/pkg/utils"
	template "server/templates"
	"server/usecases/account/register"
	"server/validation"

	"github.com/go-chi/chi/v5"
)

type Handle interface {
	Welcome(w http.ResponseWriter, r *http.Request)
	Abc(w http.ResponseWriter, r *http.Request)
	Articles(w http.ResponseWriter, r *http.Request)
	NotFound(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
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
	var body models.User
	var resp models.Response
	
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil{
		w.Write([]byte("json decode err"))
	}

	resp = validateRegData(body)
	if len(resp.Message) == 0 {
		data := database.User{
			Username: body.Username,
			Email: body.Email,
			Password: body.Password,
			Created_at: utils.Now(alias.GMT, alias.TIME_FORMAT),
			Status: "none",
			Delete_status: "none",
		}

		err := register.Register(data)
		if err != nil{
			resp.Code = 503
			resp.Message = append(resp.Message, "err insert database")
		} else{
			resp.Code = 200
			resp.Message = append(resp.Message, "success")
		}
	}

	httpResp(w, r, resp)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var body models.User
	var resp models.Response

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil{
		w.Write([]byte("json decode err"))
	}

	data := database.User{
		Username: body.Username,
		Email: body.Email,
		Password: body.Password,
	}

	result := register.Login(data)
	if result.Username == ""{
		resp.Code = 200
		resp.Message = append(resp.Message, "user not existed")
	}else{
		resp.Code = 200
		resp.Message = append(resp.Message, "success!!")
	}

	httpResp(w, r, resp)
}

func httpResp(w http.ResponseWriter, r *http.Request, resp models.Response) {
	res, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("parse JSON err"))
	}

	switch resp.Code {
	case 200:
		template.HttpResponse(w, r, res, alias.HTTP_OK)
	case 500:
		template.HttpResponse(w, r, res, alias.HTTP_BADREQUEST)
	case 503:
		template.HttpResponse(w, r, res, alias.HTTP_SERVICEUNAVAILABLE)
	}
}

func validateRegData(data models.User) (resp models.Response){
	var validate validation.Validation
	if !validate.IsEmailValid(data.Email){
		resp.Code = 500
		resp.Message = append(resp.Message, "invalid email")
	}

	if !validate.ValidatePassword(data.Password, data.Repassword){
		resp.Code = 500
		resp.Message = append(resp.Message, "invalid password")
	}

	if !validate.IsUsernameValid(data.Username){
		resp.Code = 500
		resp.Message = append(resp.Message, "invalid username")
	}

	return
}