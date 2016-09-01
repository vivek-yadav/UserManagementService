package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Apps []App

type App struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"Id"`
	Name        string        `bson:"Name" json:"Name"`
	Description string        `bson:"Description" json:"Description"`
	Token       string        `bson:"Token" json:"Token"`
	TOC         time.Time     `bson:"TOC" json:"TOC"`
	TTL         int64         `bson:"TTL" json:"TTL"`
	Roles       []Role        `bson:"Roles" json:"Roles"`
}

type Role struct {
	Name        string `bson:"Name" json:"Name"`
	Description string `bson:"Description" json:"Description"`
	Paths       []Path `bson:"Paths" json:"Paths"`
}

type Path struct {
	Path        string `bson:"Path" json:"Path"`
	AccessLevel int8   `bson:"AccessLevel" json:"AccessLevel"`
}
