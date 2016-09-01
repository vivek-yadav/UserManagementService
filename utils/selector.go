package utils

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
)

func Selector(q ...string) (r bson.M) {
	if len(q) < 1 {
		r = nil
		return
	}
	r = make(bson.M, len(q))
	for _, s := range q {
		r[s] = 1
	}
	return
}

func Deselector(q ...string) (r bson.M) {
	if len(q) < 1 {
		r = nil
		return
	}
	r = make(bson.M, len(q))
	for _, s := range q {
		r[s] = 0
	}
	return
}

func SelDeSel(sel []string, desel []string) (r bson.M) {
	if len(sel)+len(desel) < 1 {
		r = nil
		return
	}
	r = make(bson.M, len(sel)+len(desel))
	for _, s := range sel {
		r[s] = 1
	}
	for _, s := range desel {
		r[s] = 0
	}
	return
}

//
//func UpdateBsonFromMap(mapModel map[string]interface{}) (data bson.M){
//	data = bson.M{}
//	for key, value := range mapModel {
//		//var er error
//		//var rInt int64
//		//rInt, er = strconv.ParseInt(value, 10, 64)
//		//if er == nil {
//		//	data[key] = rInt
//		//	continue
//		//}
//		//var rBool bool
//		//rBool, er = strconv.ParseBool(value)
//		//if er == nil {
//		//	data[key]= rBool
//		//	continue
//		//}
//		//var rFloat float64
//		//rFloat, er = strconv.ParseFloat(value, 64)
//		//if er == nil {
//		//	data[key]= rFloat
//		//	continue
//		//}
//		data[key] = value
//	}
//	data = bson.M{"$set":data}
//	return
//}

func GetBsonFindArray(and []map[string]string, or []map[string]string) (query bson.M) {
	query = bson.M{}
	andArray := []bson.M{}
	for _, obj := range and {
		for key, value := range obj {
			var er error
			var rInt int64
			var opr string = ""
			if strings.HasPrefix(value, ">=") {
				values := strings.Split(value, ">=")
				opr = "$gte"
				value = values[1]
			} else if strings.HasPrefix(value, ">") {
				values := strings.Split(value, ">")
				opr = "$gt"
				value = values[1]
			} else if strings.HasPrefix(value, "<=") {
				values := strings.Split(value, "<=")
				opr = "$lte"
				value = values[1]
			} else if strings.HasPrefix(value, "<") {
				values := strings.Split(value, "<")
				opr = "$lt"
				value = values[1]
			}

			rInt, er = strconv.ParseInt(value, 10, 64)
			if er == nil {
				if opr == "" {
					andArray = append(andArray, bson.M{key: rInt})
				} else {
					andArray = append(andArray, bson.M{key: bson.M{opr: rInt}})
				}
				continue
			}
			var rBool bool
			rBool, er = strconv.ParseBool(value)
			if er == nil {
				if opr == "" {
					andArray = append(andArray, bson.M{key: rBool})
				} else {
					andArray = append(andArray, bson.M{key: bson.M{opr: rBool}})
				}
				continue
			}
			var rFloat float64
			rFloat, er = strconv.ParseFloat(value, 64)
			if er == nil {
				if opr == "" {
					andArray = append(andArray, bson.M{key: rFloat})
				} else {
					andArray = append(andArray, bson.M{key: bson.M{opr: rFloat}})
				}
				continue
			}
			andArray = append(andArray, bson.M{key: bson.M{"$regex": value}})
		}
	}

	orArray := []bson.M{}
	for _, obj := range or {
		for key, value := range obj {
			var er error
			var rInt int64
			var opr string = ""
			if strings.HasPrefix(value, ">") {
				values := strings.Split(value, ">")
				opr = "$gt"
				value = values[1]
			} else if strings.HasPrefix(value, ">=") {
				values := strings.Split(value, ">=")
				opr = "$gte"
				value = values[1]
			} else if strings.HasPrefix(value, "<") {
				values := strings.Split(value, "<")
				opr = "$lt"
				value = values[1]
			} else if strings.HasPrefix(value, "<=") {
				values := strings.Split(value, "<=")
				opr = "$lte"
				value = values[1]
			}

			rInt, er = strconv.ParseInt(value, 10, 64)
			if er == nil {
				if opr == "" {
					andArray = append(andArray, bson.M{key: rInt})
				} else {
					andArray = append(andArray, bson.M{key: bson.M{opr: rInt}})
				}
				continue
			}
			var rBool bool
			rBool, er = strconv.ParseBool(value)
			if er == nil {
				if opr == "" {
					andArray = append(andArray, bson.M{key: rBool})
				} else {
					andArray = append(andArray, bson.M{key: bson.M{opr: rBool}})
				}
				continue
			}
			var rFloat float64
			rFloat, er = strconv.ParseFloat(value, 64)
			if er == nil {
				if opr == "" {
					andArray = append(andArray, bson.M{key: rFloat})
				} else {
					andArray = append(andArray, bson.M{key: bson.M{opr: rFloat}})
				}
				continue
			}
			andArray = append(andArray, bson.M{key: bson.M{"$regex": value}})
		}
	}

	if len(andArray) > 0 && len(orArray) > 0 {
		query = bson.M{"$and": []bson.M{{"$and": andArray}, {"$or": orArray}}}
	} else if len(andArray) > 0 && len(orArray) == 0 {
		query = bson.M{"$and": andArray}
	} else if len(andArray) == 0 && len(orArray) > 0 {
		query = bson.M{"$or": orArray}
	}
	return
}
