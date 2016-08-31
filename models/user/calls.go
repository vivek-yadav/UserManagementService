package models

import (
	"errors"
	"github.com/vivek-yadav/UserManagementService/db/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// This function is used for login api calls.
// this(User) is set with email and password before the call.
// It returns error with cause.
func (this User) FindLoginUserWithEmail() (User, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("users")
	p := this.Password
	u := this.Email
	this.Password = ""

	q := uc.Find(bson.M{"$and": []bson.M{{"Email": u}, {"Password": p}}})
	q = q.Select(bson.M{"Password": 0, "AppAccess": 0})

	er := q.One(&this)

	if er != nil {
		return this, errors.New("ERROR : Failed to Find User by Email (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}

// This function is used for login api calls.
// this(User) is set with username and password before the call.
// It returns error with cause.
func (this User) FindLoginUserWithUsername() (User, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("users")
	p := this.Password
	u := this.Username
	this.Password = ""

	q := uc.Find(bson.M{"$and": []bson.M{{"Username": u}, {"Password": p}}})
	q = q.Select(bson.M{"Password": 0, "AppAccess": 0})

	er := q.One(&this)
	if er != nil {
		return this, errors.New("ERROR : Failed to Find User by Username (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}

// This function is used to check for authorization of the user to access specific url
// if allowed it returns success else returns the reason of error
func (this User) IsAuth(AppToken string, accessLevel int8, url string) (User, error) {
	return User{}, nil
}

func (this *User) Create() (User, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("users")

	this.Id = bson.NewObjectId()

	er := uc.Insert(this)
	if er != nil {
		return *this, errors.New("ERROR : Failed to insert User (\n\t" + er.Error() + "\n)")
	}
	return *this, nil
}

func (this *Users) Create() (Users, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("users")

	er := uc.Insert(this)
	if er != nil {
		return *this, errors.New("ERROR : Failed to insert User (\n\t" + er.Error() + "\n)")
	}
	return *this, nil
}
