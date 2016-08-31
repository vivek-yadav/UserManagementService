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
	//QueryResolverUpdate(*mgo.Query)(interface{},error)
}
type ApiModels interface {
	DbFetchAll(*mgo.Query) (interface{}, error)
	DbInsertAll(*mgo.Collection) (interface{}, error)
	//QueryResolverUpdate(*mgo.Query)(interface{},error)
}

func Create(collectionName string, c *gin.Context, model ApiModel) (uu interface{}, er error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C(collectionName)

	uu, er = model.DbInsertOne(uc)
	return
}

func CreateList(collectionName string, c *gin.Context, model ApiModels) (uu interface{}, er error) {
	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
	uc := con.DB("").C(collectionName)

	uu, er = model.DbInsertAll(uc)
	return
}

//func GetList(collectionName string,c *gin.Context,queryResolver func(*mgo.Query)(interface{},error)) (r Result,er error){
func GetList(collectionName string, c *gin.Context, model ApiModels) (r Result, er error) {
	//r := Result{}
	var uu interface{}

	var page, size, total int64
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
		uu, total, er = getListFromDB(collectionName, model.DbFetchAll, exfields, true, andCond, orCond, page, size)
	} else {
		uu, total, er = getListFromDB(collectionName, model.DbFetchAll, fields, false, andCond, orCond, page, size)
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

func getListFromDB(collectionName string, queryResolver func(*mgo.Query) (interface{}, error), fields []string, isExclude bool, andCond []map[string]string, orCond []map[string]string, page, size int64) (interface{}, int64, error) {
	var u interface{}

	authDB, _ := mongo.GetAuthDB()
	con, _ := authDB.Connect()
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
	q = q.Limit(int(size))
	q = q.Skip(int((page - 1) * size))
	u, er := queryResolver(q)
	if er != nil {
		return u, 0, errors.New("ERROR : Failed to Find " + collectionName + " (\n\t" + er.Error() + "\n)")
	}
	return u, total, nil
}

func GetById(collectionName string, c *gin.Context, model ApiModel) (uu interface{}, er error) {

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
		uu, er = getByIdFromDB(collectionName, model.DbFetchOne, exfields, true, id)
	} else {
		uu, er = getByIdFromDB(collectionName, model.DbFetchOne, fields, false, id)
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

func Get(collectionName string, c *gin.Context, model ApiModel) (uu interface{}, er error) {

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
		uu, er = getFromDB(collectionName, model.DbFetchOne, exfields, true, andCond, orCond)
	} else {
		uu, er = getFromDB(collectionName, model.DbFetchOne, fields, false, andCond, orCond)
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
