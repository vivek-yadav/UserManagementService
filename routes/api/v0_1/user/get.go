package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/user"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"strings"
)

func GetList(c *gin.Context) {
	var er error
	var page, size, total int64
	u := models.Users{}
	field := c.Query("fields")
	fields := strings.Split(field, ",")
	pageS := c.Query("page")
	sizeS := c.Query("size")
	page, er = strconv.ParseInt(pageS, 10, 32)
	size, er = strconv.ParseInt(sizeS, 10, 32)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}

	u, total, _ = u.GetList(fields, page, size)
	r := models.Result{}
	r.Total = total
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
