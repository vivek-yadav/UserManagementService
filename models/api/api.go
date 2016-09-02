package modelApi

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/db/mongo"
	"github.com/vivek-yadav/UserManagementService/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

type ApiModel interface {
	DbFetchOne(*mgo.Query) (interface{}, error)
	DbInsertOne(*mgo.Collection) (interface{}, error)
	DbUpdateOne(*mgo.Collection, bson.M, bson.M) (interface{}, error)
	DbReplaceOne(*mgo.Collection) (interface{}, error)
	DbDeleteOne(*mgo.Collection) (interface{}, error)
}
type ApiModels interface {
	DbFetchAll(*mgo.Query) (interface{}, error)
	DbInsertAll(*mgo.Collection) (interface{}, error)
	DbUpdateAll(*mgo.Collection, bson.M, bson.M) (interface{}, error)
	DbReplaceAll(*mgo.Collection) (interface{}, error)
	DbDeleteAll(*mgo.Collection) (interface{}, error)
}

func DeleteOne(collectionName string, c *gin.Context) (er error) {
	id := c.Query("id")
	if id == "" {
		id = c.Param("id")
	}
	var idHex bson.ObjectId
	if bson.IsObjectIdHex(id) {
		idHex = bson.ObjectIdHex(id)
	} else {
		er = errors.New("ERROR : Id is not proper")
		return
	}

	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)
	er = uc.RemoveId(idHex)
	if er != nil {
		er = errors.New("ERROR : Failed to Delete " + collectionName + " (\n\t" + er.Error() + "\n)")
		return
	}
	return
}

func DeleteAll(collectionName string, c *gin.Context) (er error) {
	var ids []string
	c.Bind(&ids)
	idHexs := make([]bson.ObjectId, len(ids))
	for i, v := range ids {
		if bson.IsObjectIdHex(v) {
			idHexs[i] = bson.ObjectIdHex(v)
		} else {
			er = errors.New("ERROR : Id is not proper (" + v + ")")
			return
		}
	}

	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)
	var info *mgo.ChangeInfo
	info, er = uc.RemoveAll(bson.M{"_id": bson.M{"$in": idHexs}})
	if er != nil {
		er = errors.New("ERROR : Failed to Delete " + collectionName + " (\n\t" + er.Error() + "\n)")
		return
	}
	if info.Removed == 0 {
		er = errors.New("No of " + collectionName + " Removed :" + strconv.Itoa(info.Removed) + " Out matched " + collectionName + " : " + strconv.Itoa(info.Matched))
		return
	}
	return
}

func ReplaceOne(collectionName string, dbHandler func(*mgo.Collection) (interface{}, error)) (uu interface{}, er error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)

	uu, er = dbHandler(uc)
	if er != nil {
		er = errors.New("ERROR : Failed to Replace All " + collectionName + " (\n\t" + er.Error() + "\n)")
		return
	}
	return
}

func ReplaceAll(collectionName string, dbHandler func(*mgo.Collection) (interface{}, error)) (uu interface{}, er error) {
	if er != nil {
		return
	}

	uu, er = replaceAllFromDB(collectionName, dbHandler)
	if er != nil {
		return
	}
	return
}

func replaceAllFromDB(collectionName string, queryResolver func(*mgo.Collection) (interface{}, error)) (interface{}, error) {
	var u interface{}

	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)

	u, er := queryResolver(uc)
	if er != nil {
		return u, errors.New("ERROR : Failed to Replace All " + collectionName + " (\n\t" + er.Error() + "\n)")
	}
	return u, nil
}

func UpdateOneById(collectionName string, c *gin.Context, updatedValues map[string]interface{}) (uu interface{}, er error) {
	id := c.Query("id")
	if id == "" {
		id = c.Param("id")
	}
	var idHex bson.ObjectId
	if bson.IsObjectIdHex(id) {
		idHex = bson.ObjectIdHex(id)
	} else {
		er = errors.New("ERROR : Id is not proper")
		return
	}
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)
	update := bson.M{"$set": updatedValues}
	er = uc.UpdateId(idHex, update)
	if er != nil {
		er = errors.New("ERROR : Failed to Update " + collectionName + " (\n\t" + er.Error() + "\n)")
		return
	}
	uu = updatedValues
	return
}

func Update(collectionName string, c *gin.Context, updatedValues map[string]interface{}, dbHandler func(*mgo.Collection, bson.M, bson.M) (interface{}, error)) (uu interface{}, er error) {

	and := c.Query("and")
	var andCond []map[string]string
	if and != "" {
		andBytes := []byte(and)
		er = json.Unmarshal(andBytes, &andCond)
	}

	or := c.Query("or")
	var orCond []map[string]string
	if or != "" {
		orBytes := []byte(or)
		er = json.Unmarshal(orBytes, &orCond)
	}

	if er != nil {
		return
	}

	uu, er = updateFromDB(collectionName, dbHandler, andCond, orCond, updatedValues)

	if er != nil {
		return
	}
	return
}

func updateFromDB(collectionName string, queryResolver func(*mgo.Collection, bson.M, bson.M) (interface{}, error), andCond []map[string]string, orCond []map[string]string, updatedValues map[string]interface{}) (interface{}, error) {
	var u interface{}

	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)

	find := utils.GetBsonFindArray(andCond, orCond)
	update := bson.M{"$set": updatedValues}
	u, er := queryResolver(uc, find, update)
	if er != nil {
		return u, errors.New("ERROR : Failed to Update " + collectionName + " (\n\t" + er.Error() + "\n)")
	}
	return u, nil
}

func InsertOne(collectionName string, dbHandler func(*mgo.Collection) (interface{}, error)) (uu interface{}, er error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)

	uu, er = dbHandler(uc)
	return
}

func InsertAll(collectionName string, dbHandler func(*mgo.Collection) (interface{}, error)) (uu interface{}, er error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)

	uu, er = dbHandler(uc)
	return
}

func FetchAll(collectionName string, c *gin.Context, dbHandler func(*mgo.Query) (interface{}, error)) (r Result, er error) {
	//r := Result{}
	var uu interface{}

	var page, size, total int64
	var fields []string
	var exfields []string
	var sorts []string
	sort := c.Query("sort")
	if sort != "" {
		sorts = strings.Split(sort, ",")
	}
	//u := models.Users{}
	field := c.Query("fields")
	if field != "" {
		fields = strings.Split(field, ",")
	}
	exfield := c.Query("excludeFields")
	if exfield != "" {
		exfields = strings.Split(exfield, ",")
	}
	pageS := c.Query("page")
	sizeS := c.Query("size")
	page, er = strconv.ParseInt(pageS, 10, 64)
	if er != nil {
		page = 1
		er = nil
	}
	size, er = strconv.ParseInt(sizeS, 10, 64)
	if er != nil {
		size = 10
		er = nil
	}

	and := c.Query("and")
	var andCond []map[string]string
	if and != "" {
		andBytes := []byte(and)
		er = json.Unmarshal(andBytes, &andCond)
	}

	or := c.Query("or")
	var orCond []map[string]string
	if or != "" {
		orBytes := []byte(or)
		er = json.Unmarshal(orBytes, &orCond)
	}

	if er != nil {
		return
	}

	if len(exfields) > 0 {
		uu, total, er = getListFromDB(collectionName, dbHandler, exfields, true, sorts, andCond, orCond, page, size)
	} else {
		uu, total, er = getListFromDB(collectionName, dbHandler, fields, false, sorts, andCond, orCond, page, size)
	}

	if er != nil {
		return
	}

	r.Total = total
	if r.Total != 0 {
		r.Page = page
		r.Size = size
		r.URL = c.Request.RequestURI
	}
	r.Data = uu
	return
}

func getListFromDB(collectionName string, queryResolver func(*mgo.Query) (interface{}, error), fields []string, isExclude bool, sort []string, andCond []map[string]string, orCond []map[string]string, page, size int64) (interface{}, int64, error) {
	var u interface{}

	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)

	q := uc.Find(utils.GetBsonFindArray(andCond, orCond))
	if isExclude {
		q = q.Select(utils.Deselector(fields...))
	} else {
		q = q.Select(utils.Selector(fields...))
	}

	c, erc := q.Count()
	if erc != nil {
		return u, 0, errors.New("ERROR : Failed to Find " + collectionName + " (\n\t" + erc.Error() + "\n)")
	}
	total := int64(c)
	if total/size < page-1 {
		return u, 0, errors.New("ERROR : Failed to Find " + collectionName + " on this page or page limit reached")
	}
	q = q.Sort(sort...)
	q = q.Limit(int(size))
	q = q.Skip(int((page - 1) * size))
	u, er := queryResolver(q)
	if er != nil {
		return u, 0, errors.New("ERROR : Failed to Find " + collectionName + " (\n\t" + er.Error() + "\n)")
	}
	return u, total, nil
}

func FetchById(collectionName string, c *gin.Context, dbHandler func(*mgo.Query) (interface{}, error)) (uu interface{}, er error) {

	var fields []string
	var exfields []string
	//u := models.Users{}
	field := c.Query("fields")
	if field != "" {
		fields = strings.Split(field, ",")
	}
	exfield := c.Query("excludeFields")
	if exfield != "" {
		exfields = strings.Split(exfield, ",")
	}

	id := c.Query("id")
	if id == "" {
		id = c.Param("id")
	}

	if er != nil {
		return
	}

	if len(exfields) > 0 {
		uu, er = getByIdFromDB(collectionName, dbHandler, exfields, true, id)
	} else {
		uu, er = getByIdFromDB(collectionName, dbHandler, fields, false, id)
	}
	if er != nil {
		return
	}

	return
}

func getByIdFromDB(collectionName string, queryResolver func(*mgo.Query) (interface{}, error), fields []string, isExclude bool, id string) (interface{}, error) {
	var u interface{}

	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)
	var idHex bson.ObjectId
	if bson.IsObjectIdHex(id) {
		idHex = bson.ObjectIdHex(id)
	}
	q := uc.Find(bson.M{"_id": idHex})
	if isExclude {
		q = q.Select(utils.Deselector(fields...))
	} else {
		q = q.Select(utils.Selector(fields...))
	}

	_, erc := q.Count()
	if erc != nil {
		return u, errors.New("ERROR : Failed to Find " + collectionName + " (\n\t" + erc.Error() + "\n)")
	}

	u, er := queryResolver(q)
	if er != nil {
		return u, errors.New("ERROR : Failed to Find " + collectionName + " (\n\t" + er.Error() + "\n)")
	}
	return u, nil
}

func FetchOne(collectionName string, c *gin.Context, dbHandler func(*mgo.Query) (interface{}, error)) (uu interface{}, er error) {

	var fields []string
	var exfields []string
	//u := models.Users{}
	field := c.Query("fields")
	if field != "" {
		fields = strings.Split(field, ",")
	}
	exfield := c.Query("excludeFields")
	if exfield != "" {
		exfields = strings.Split(exfield, ",")
	}

	and := c.Query("and")
	var andCond []map[string]string
	if and != "" {
		andBytes := []byte(and)
		er = json.Unmarshal(andBytes, &andCond)
	}

	or := c.Query("or")
	var orCond []map[string]string
	if or != "" {
		orBytes := []byte(or)
		er = json.Unmarshal(orBytes, &orCond)
	}

	if er != nil {
		return
	}

	if len(exfields) > 0 {
		uu, er = getFromDB(collectionName, dbHandler, exfields, true, andCond, orCond)
	} else {
		uu, er = getFromDB(collectionName, dbHandler, fields, false, andCond, orCond)
	}

	if er != nil {
		return
	}
	return
}

func getFromDB(collectionName string, queryResolver func(*mgo.Query) (interface{}, error), fields []string, isExclude bool, andCond []map[string]string, orCond []map[string]string) (interface{}, error) {
	var u interface{}

	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	defer con.Close()
	uc := con.DB("").C(collectionName)

	q := uc.Find(utils.GetBsonFindArray(andCond, orCond))
	if isExclude {
		q = q.Select(utils.Deselector(fields...))
	} else {
		q = q.Select(utils.Selector(fields...))
	}

	_, erc := q.Count()
	if erc != nil {
		return u, errors.New("ERROR : Failed to Find " + collectionName + " (\n\t" + erc.Error() + "\n)")
	}
	u, er := queryResolver(q)
	if er != nil {
		return u, errors.New("ERROR : Failed to Find " + collectionName + " (\n\t" + er.Error() + "\n)")
	}
	return u, nil
}
