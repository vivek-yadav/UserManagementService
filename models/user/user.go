package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Users []User

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"Id"`
	Username string        `bson:"Username" json:"Username"`
	Password string        `bson:"Password" json:"Password"`
	Email    string        `bson:"Email" json:"Email"`
	Name     string        `bson:"Name" json:"Name"`
	TOC      time.Time     `bson:"TOC" json:"TOC"`
	TTL      int64         `bson:"TTL" json:"TTL"`
	Accesses []AppAccess   `bson:"AppAccess" json:"AppAccess"`
	IsActive bool          `bson:"IsActive" json:"IsActive"`
}

type AppAccess struct {
	Token    string    `bson:"Token" json:"Token"`
	TOC      time.Time `bson:"TOC" json:"TOC"`
	TTL      int64     `bson:"TTL" json:"TTL"`
	RoleName string    `bson:"RoleName" json:"RoleName"`
	Paths    []AppPath `bson:"Paths" json:"Paths"`
}

type AppPath struct {
	Path        string `bson:"Path" json:"Path"`
	AccessLevel int8   `bson:"AccessLevel" json:"AccessLevel"`
}
