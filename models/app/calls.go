package models

import (
	"errors"
	"github.com/vivek-yadav/UserManagementService/db/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (this *App) QueryResolverAll(q *mgo.Query) (interface{}, error) {
	er := q.All(this)
	if er != nil {
		return this, errors.New("ERROR : Failed to Find Users (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}

func (this *App) QueryResolverOne(q *mgo.Query) (interface{}, error) {
	er := q.One(this)
	if er != nil {
		return this, errors.New("ERROR : Failed to Find User (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}

func (this *App) Create() (App, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("apps")

	this.Id = bson.NewObjectId()

	er := uc.Insert(*this)
	if er != nil {
		return *this, errors.New("ERROR : Failed to insert User (\n\t" + er.Error() + "\n)")
	}
	return *this, nil
}
