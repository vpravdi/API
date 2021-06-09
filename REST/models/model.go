package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Name   string
	Gender string
	Age    int
	Id     bson.ObjectId
}
