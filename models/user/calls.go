package models

import (
	"errors"
	"github.com/vivek-yadav/UserManagementService/db/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func (this Users) DbFetchAll(q *mgo.Query) (interface{}, error) {
	er := q.All(&this)
	if er != nil {
		return this, errors.New("ERROR : Failed to Find Users (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}

func (this User) DbFetchOne(q *mgo.Query) (interface{}, error) {
	er := q.One(&this)
	if er != nil {
		return this, errors.New("ERROR : Failed to Find User (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}

func (this User) DbInsertOne(uc *mgo.Collection) (uu interface{}, er error) {
	this.Id = bson.NewObjectId()

	er = uc.Insert(this)
	if er != nil {
		er = errors.New("ERROR : Failed to insert User (\n\t" + er.Error() + "\n)")
		return
	}
	uu = this
	return
}

func (this Users) DbInsertAll(uc *mgo.Collection) (uu interface{}, er error) {
	list := make([]interface{}, len(this))
	for i, v := range this {
		v.Id = bson.NewObjectId()
		list[i] = v
	}

	er = uc.Insert(list...)
	if er != nil {
		er = errors.New("ERROR : Failed to insert Users (\n\t" + er.Error() + "\n)")
		return
	}
	uu = list
	return
}

func (this User) DbUpdateOne(uc *mgo.Collection, sel bson.M, updates bson.M) (uu interface{}, er error) {
	er = uc.Update(sel, updates)
	if er != nil {
		er = errors.New("ERROR : Failed to Update User (\n\t" + er.Error() + "\n)")
		return
	}
	return
}

func (this Users) DbUpdateAll(uc *mgo.Collection, sel bson.M, updates bson.M) (uu interface{}, er error) {
	var changes *mgo.ChangeInfo
	changes, er = uc.UpdateAll(sel, updates)
	if er != nil {
		er = errors.New("ERROR : Failed to Update Users (\n\t" + er.Error() + "\n ) Changes : Matched (" + strconv.Itoa(changes.Matched) + ")  Updated (" + strconv.Itoa(changes.Updated) + ") Removed (" + strconv.Itoa(changes.Removed) + ")")
		return
	}
	return
}

func (this Users) DbReplaceAll(uc *mgo.Collection) (uu interface{}, er error) {
	for _, v := range this {
		er = uc.UpdateId(v.Id, v)
		if er != nil {
			er = errors.New("ERROR : Failed to Update Users (\n\t" + er.Error() + "\n)")
			return
		}
	}
	uu = this
	return
}

func (this User) DbReplaceOne(uc *mgo.Collection) (uu interface{}, er error) {
	er = uc.UpdateId(this.Id, this)
	if er != nil {
		er = errors.New("ERROR : Failed to Update User (\n\t" + er.Error() + "\n)")
		return
	}
	uu = this
	return
}

// This function is used for login api calls.
func (this User) IsLogin(self *User, result *bool) error {
	*result = false
	if this.Password != "" && this.Username != "" && self.Password == "" && self.Username == "" {
		self = &this
	}
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C("users")
	p := self.Password
	u := self.Username
	e := self.Email
	this.Password = ""
	key := "Email"
	val := e
	if u != "" && e == "" {
		key = "Username"
		val = u
	}

	q := uc.Find(bson.M{"$and": []bson.M{{key: val}, {"Password": p}}})
	q = q.Select(bson.M{"Password": 0, "AppAccess": 0})

	er := q.One(&self)
	if er != nil {
		return errors.New("ERROR : Failed to Find User (\n\t" + er.Error() + "\n)")
	}
	return nil
}

// This function is used to check for authorization of the user to access specific url
// if allowed it returns success else returns the reason of error
func (this User) IsAuth(AppToken string, accessLevel int8, url string) (User, error) {
	return User{}, nil
}

//
//func (this *User) Create() (User, error) {
//	authDB, _ := mongo.GetAuthDB()
//	con, _ := authDB.Connect()
//	uc := con.DB("").C("users")
//
//	this.Id = bson.NewObjectId()
//
//	er := uc.Insert(this)
//	if er != nil {
//		return *this, errors.New("ERROR : Failed to insert User (\n\t" + er.Error() + "\n)")
//	}
//	return *this, nil
//}
//
//func (this *Users) Create() (Users, error) {
//	authDB, _ := mongo.GetAuthDB()
//	con, _ := authDB.Connect()
//	uc := con.DB("").C("users")
//
//	er := uc.Insert(this)
//	if er != nil {
//		return *this, errors.New("ERROR : Failed to insert User (\n\t" + er.Error() + "\n)")
//	}
//	return *this, nil
//}
