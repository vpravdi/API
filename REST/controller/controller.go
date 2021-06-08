package controller

import (
	model "API/REST/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/julienschmidt/httprouter"
	//"gopkg.in/mgo.v2"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Index</title>
			</head>
			<body>
			<a href="/user/88345">GO TO: http://localhost:8080/user/88345</a>
			</body>
			</html>`

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))

}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := model.User{
		Name:   "Dinesh Praveen",
		Gender: "Male",
		Age:    15,
		Id:     p.ByName("id"),
	}

	userjson, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userjson)

}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := model.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "007"

	uj, _ := json.Marshal(u)
	fmt.Println(uj)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "deleted\n")
}
