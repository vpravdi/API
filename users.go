package main

import (
	model "API/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user/", createUser)
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe("localhost:8080", r)

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

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := model.User{
		Name:   "Dinesh Praveen",
		Gender: "Male",
		Age:    15,
		Id:     p.ByName("id"),
	}

	userjson, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s\n", userjson)

}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := model.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "1352"

	userjson, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, "%s\n", userjson)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "deleted\n")
}
