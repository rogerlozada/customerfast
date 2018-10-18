package customerdomainmodel

import "gopkg.in/mgo.v2/bson"

type (
	//Customer instances
	Customer struct {
		ID   bson.ObjectId `json:"id" bson:"_id"`
		Name string        `json:"nome" bson:"nome"`
	}
)
