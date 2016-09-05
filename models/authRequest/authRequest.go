package models

import (
	"errors"
	"github.com/vivek-yadav/UserManagementService/db/mongo"
	"github.com/vivek-yadav/UserManagementService/models/user"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type AuthRequest struct {
	Username    string `json:"Username"`
	Email       string `json:"Email"`
	AppToken    string `json:"AppToken"`
	AccessLevel int    `json:"AccessLevel"`
	Path        string `json:"Path"`
}

func (this AuthRequest) IsAuth(req *AuthRequest, result *bool) (er error) {
	*result = false
	if this.Path != "" && req.Path == "" {
		req = &this
	}
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C("users")

	parts := strings.Split(req.Path, "/")
	patterns := make([]string, (len(parts) * 2))
	t := 0
	for i, _ := range parts {
		q := strings.Join(parts[:i+1], "/")
		patterns[t] = q + "/*"
		t++
		if q == "" {
			patterns[t] = "/"
			t++
		} else {
			patterns[t] = q
			t++
		}
	}
	var user models.User
	if req.Email != "" {
		er = uc.Find(
			bson.M{"$and": []bson.M{
				{"Email": req.Email},
				{"AppAccess.Token": req.AppToken},
				{"AppAccess.Paths.Path": bson.M{"$in": patterns}},
				{"AppAccess.Paths.AccessLevel": bson.M{"$gt": req.AccessLevel}},
			},
			},
		).One(&user)
	} else if req.Username != "" {
		er = uc.Find(
			bson.M{"$and": []bson.M{
				{"Username": req.Username},
				{"AppAccess.Token": req.AppToken},
				{"AppAccess.Paths.Path": bson.M{"$in": patterns}},
				{"AppAccess.Paths.AccessLevel": bson.M{"$gt": req.AccessLevel}},
			},
			},
		).One(&user)
	} else {
		er = errors.New("ERROR : Username or Email field not set.")
		return
	}
	if er != nil {
		er = errors.New("ERROR : No valid user found, Thus not authorized. (\n\t" + er.Error() + "\n)")
		return
	}
	for _, acc := range user.Accesses {
		if acc.Token == req.AppToken {
			for _, path := range acc.Paths {
				if path.AccessLevel == 0 {
					for _, tar := range patterns {
						if path.Path == tar {
							er = errors.New("Not authorized to access req path")
							return
						}
					}
				}
			}
		}
	}
	*result = true
	return
}
