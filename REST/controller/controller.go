package controller

import (
	model "API/REST/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	//"github.com/globalsign/mgo" alternate packages in case the default ones dont work
	//"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := model.User{}

	if err := uc.session.DB("go-go-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNoContent)
		fmt.Println(err)
		return
	}

	userjson, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userjson)

}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := model.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("go-go-db").C("users").Insert(u)

	uj, _ := json.Marshal(u)
	fmt.Println(uj)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
	f, err := os.Create("userdata")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(u)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("go-go-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")

}
