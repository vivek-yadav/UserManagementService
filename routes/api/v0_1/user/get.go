package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/result"
	"github.com/vivek-yadav/UserManagementService/models/user"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
)

func GetList(c *gin.Context) {
	var er error
	var page, size, total int64
	var fields []string
	var exfields []string
	u := models.Users{}
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
	size, er = strconv.ParseInt(sizeS, 10, 64)

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
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}

	if len(exfields) > 0 {
		u, total, _ = u.GetList(exfields, true, andCond, orCond, page, size)
	} else {
		u, total, _ = u.GetList(fields, false, andCond, orCond, page, size)
	}

	r := result.Result{}
	r.Total = total
	if r.Total != 0 {
		r.Page = page
		r.Size = size
		r.URL = c.Request.URL.String()
	}
	r.Data = u
	c.JSON(http.StatusOK, r)
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	var u models.User
	if bson.IsObjectIdHex(id) {
		u = models.User{Id: bson.ObjectIdHex(id)}
	} else {
		u = models.User{}
	}
	u, _ = u.GetById()
	c.JSON(http.StatusOK, u)
}
