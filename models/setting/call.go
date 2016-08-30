package settings

import (
	"errors"
	"github.com/vivek-yadav/UserManagementService/db/mongo"
	"gopkg.in/mgo.v2/bson"
)

func (this *Setting) Get() (Setting, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("settings")

	er := uc.Find(bson.M{}).One(this)
	if er != nil {
		//return this,errors.New("ERROR : Failed to Find User Management Service (UMS) Settings (\n\t"+er.Error()+"\n)")

		if this.Id == "" {
			_, err := this.CreateNew()
			if err != nil {
				return *this, errors.New("ERROR : Failed to create Settings in the ums database (\n\t" + err.Error() + "\n)")
			}
			//this = *s
		}
	}
	return *this, nil
}

func (this *Setting) CreateNew() (*Setting, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("settings")

	this.Id = bson.NewObjectId()
	this.Title = "User Management Service"
	this.IsInitDone = false

	er := uc.Insert(*this)
	if er != nil {
		return this, errors.New("ERROR : Failed to insert User (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}
