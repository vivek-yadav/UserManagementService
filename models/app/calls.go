package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func (this App) DbFetchOne(q *mgo.Query) (interface{}, error) {
	er := q.One(&this)
	if er != nil {
		return this, errors.New("ERROR : Failed to Find App (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}

func (this App) DbInsertOne(uc *mgo.Collection) (uu interface{}, er error) {
	a := this
	a.Id = bson.NewObjectId()
	//for i,_ := range a.Roles{
	//	a.Roles[i].Id = bson.NewObjectId()
	//	for j,_ := range a.Roles[i].Paths{
	//		a.Roles[i].Paths[j].Id = bson.NewObjectId()
	//	}
	//}
	er = uc.Insert(a)
	if er != nil {
		er = errors.New("ERROR : Failed to insert App (\n\t" + er.Error() + "\n)")
		return
	}
	uu = a
	return
}

func (this Apps) DbFetchAll(q *mgo.Query) (interface{}, error) {
	er := q.All(&this)
	if er != nil {
		return this, errors.New("ERROR : Failed to Find Apps (\n\t" + er.Error() + "\n)")
	}
	return this, nil
}

func (this Apps) DbInsertAll(uc *mgo.Collection) (uu interface{}, er error) {
	list := make([]interface{}, len(this))
	for i, v := range this {
		v.Id = bson.NewObjectId()
		list[i] = v
	}

	er = uc.Insert(list...)
	if er != nil {
		er = errors.New("ERROR : Failed to insert Apps (\n\t" + er.Error() + "\n)")
		return
	}
	uu = list
	return
}

func (this App) DbUpdateOne(uc *mgo.Collection, sel bson.M, updates bson.M) (uu interface{}, er error) {
	er = uc.Update(sel, updates)
	if er != nil {
		er = errors.New("ERROR : Failed to Update App (\n\t" + er.Error() + "\n)")
		return
	}
	return
}

func (this Apps) DbUpdateAll(uc *mgo.Collection, sel bson.M, updates bson.M) (uu interface{}, er error) {
	var changes *mgo.ChangeInfo
	changes, er = uc.UpdateAll(sel, updates)
	if er != nil {
		er = errors.New("ERROR : Failed to Update Apps (\n\t" + er.Error() + "\n ) Changes : Matched (" + strconv.Itoa(changes.Matched) + ")  Updated (" + strconv.Itoa(changes.Updated) + ") Removed (" + strconv.Itoa(changes.Removed) + ")")
		return
	}
	return
}

func (this Apps) DbReplaceAll(uc *mgo.Collection) (uu interface{}, er error) {
	for _, v := range this {
		er = uc.UpdateId(v.Id, v)
		if er != nil {
			er = errors.New("ERROR : Failed to Update Apps (\n\t" + er.Error() + "\n)")
			return
		}
	}
	uu = this
	return
}

func (this App) DbReplaceOne(uc *mgo.Collection) (uu interface{}, er error) {
	er = uc.UpdateId(this.Id, this)
	if er != nil {
		er = errors.New("ERROR : Failed to Update App (\n\t" + er.Error() + "\n)")
		return
	}
	uu = this
	return
}
