package settings

import "gopkg.in/mgo.v2/bson"

type Setting struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"Id"`
	Title      string        `bson:"Title" json:"Title"`
	IsInitDone bool          `bson:"IsInitDone" json:"IsInitDone"`
}
