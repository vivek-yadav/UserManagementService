package utils

import "strings"

var AccessLevelNo int8 = 0
var AccessLevelGet int8 = 1
var AccessLevelPost int8 = 2
var AccessLevelGetPost int8 = 3
var AccessLevelGetPostPut int8 = 4
var AccessLevelGetPostDelete int8 = 5
var AccessLevelGetPostPutDelete int8 = 6

func GetAccessLevel(method string) int8 {
	switch strings.ToUpper(method) {
	case "GET":
		return AccessLevelGet
	case "POST":
		return AccessLevelGetPost
	case "PUT":
		return AccessLevelGetPostPut
	case "DELETE":
		return AccessLevelGetPostDelete
	default:
		return AccessLevelNo
	}
}
