package models

import (
	"github.com/vivek-yadav/UserManagementService/models/user"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type App struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Name        string        `bson:"Name" json:"Name"`
	Description string        `bson:"Description" json:"Description"`
	Token       string        `bson:"Token" json:"Token"`
	TOC         time.Time     `bson:"TOC" json:"TOC"`
	TTL         int64         `bson:"TTL" json:"TTL"`
	Roles       []Role        `bson:"Roles" json:"Roles"`
}

type Role struct {
	Id          bson.ObjectId    `bson:"_id,omitempty" json:"_id"`
	Name        string           `bson:"Name" json:"Name"`
	Description string           `bson:"Description" json:"Description"`
	Paths       []models.AppPath `bson:"AppPath" json:"AppPath"`
}
