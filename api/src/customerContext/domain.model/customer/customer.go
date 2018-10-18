package customerdomainmodel

import "gopkg.in/mgo.v2/bson"

type (
	//Customer instance
	Customer struct {
		ID   bson.ObjectId `json:"id" bson:"_id"`
		Name string        `json:"name" bson:"name"`
	}
)
