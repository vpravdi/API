package main

import (
	controller "API/REST/controller"

	"net/http"

	"github.com/globalsign/mgo"
	//"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controller.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
