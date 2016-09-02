package authRequest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/vivek-yadav/UserManagementService/models/authRequest"
	"net/http"
	//"net/rpc"
	//"fmt"
)

func IsAuth(c *gin.Context) {
	var er error
	r := models.AuthRequest{}
	json.NewDecoder(c.Request.Body).Decode(&r)
	var result bool

	//client, err := rpc.DialHTTP("tcp", ":7001")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code":    http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//// Synchronous call
	//err = client.Call("AuthRequest.IsAuth", &r, &result)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code":    http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//fmt.Printf("Request : %v : %v", r, result)

	er = r.IsAuth(&r, &result)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
