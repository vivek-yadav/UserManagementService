package utils

import "gopkg.in/mgo.v2/bson"

func Selector(q ...string) (r bson.M) {
	r = make(bson.M, len(q))
	for _, s := range q {
		r[s] = 1
	}
	return
}

func Deselector(q ...string) (r bson.M) {
	r = make(bson.M, len(q))
	for _, s := range q {
		r[s] = 0
	}
	return
}
