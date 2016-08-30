package models

import (
	"errors"
	"github.com/vivek-yadav/UserManagementService/db/mongo"
	"github.com/vivek-yadav/UserManagementService/utils"
	"gopkg.in/mgo.v2/bson"
)

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
	q = q.Select(bson.M{"Password": 0, "Accesses": 0})

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
	q = q.Select(bson.M{"Password": 0, "Accesses": 0})

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

// This function is used to fetch detils of a user by Id
func (this User) GetById() (User, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("users")

	er := uc.Find(bson.M{"_id": this.Id}).One(&this)
	if er != nil {
		return this, errors.New("ERROR : Failed to Find User with id " + this.Id.Hex() + " (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}

// This function is used to fetch all the users
func (this Users) GetList(fields []string, page, size int64) (Users, int64, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("users")

	q := uc.Find(bson.M{})
	q = q.Select(utils.Selector(fields...))
	c, erc := q.Count()
	if erc != nil {
		return this, 0, errors.New("ERROR : Failed to Find Users (\n\t" + erc.Error() + "\n)")
	}
	total := int64(c)
	if total < (page-1)*size {
		return this, 0, errors.New("ERROR : Failed to Find Users on this page or page limit reached")
	}
	q = q.Limit(int(size))
	q = q.Skip(int((page - 1) * size))
	er := q.All(&this)
	if er != nil {
		return this, 0, errors.New("ERROR : Failed to Find Users (\n\t" + er.Error() + "\n)")
	}
	//	fmt.Println(this["Name"])
	return this, total, nil
}

func (this User) CreateUser() (User, error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C("users")

	this.Id = bson.NewObjectId()

	er := uc.Insert(this)
	if er != nil {
		return this, errors.New("ERROR : Failed to insert User (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}
