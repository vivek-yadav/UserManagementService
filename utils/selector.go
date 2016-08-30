package utils

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
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

func GetBsonFindArray(and []map[string]string, or []map[string]string) (query bson.M) {
	query = bson.M{}
	andArray := []bson.M{}
	for _, obj := range and {
		for key, value := range obj {
			var er error
			var rInt int64
			rInt, er = strconv.ParseInt(value, 10, 64)
			if er == nil {
				andArray = append(andArray, bson.M{key: rInt})
				continue
			}
			var rBool bool
			rBool, er = strconv.ParseBool(value)
			if er == nil {
				andArray = append(andArray, bson.M{key: rBool})
				continue
			}
			var rFloat float64
			rFloat, er = strconv.ParseFloat(value, 64)
			if er == nil {
				andArray = append(andArray, bson.M{key: rFloat})
				continue
			}
			andArray = append(andArray, bson.M{key: value})
		}
	}

	orArray := []bson.M{}
	for _, obj := range or {
		for key, value := range obj {
			var er error
			var rInt int
			rInt, er = strconv.Atoi(value)
			if er == nil {
				orArray = append(orArray, bson.M{key: rInt})
				continue
			}
			var rBool bool
			rBool, er = strconv.ParseBool(value)
			if er == nil {
				orArray = append(orArray, bson.M{key: rBool})
				continue
			}
			var rFloat float64
			rFloat, er = strconv.ParseFloat(value, 64)
			if er == nil {
				orArray = append(orArray, bson.M{key: rFloat})
				continue
			}
			orArray = append(orArray, bson.M{key: value})
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
