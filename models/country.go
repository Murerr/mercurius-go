package models

import "gopkg.in/mgo.v2/bson"

type Country struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	SortName    string        `bson:"sortname" json:"sortname"`
	Name        string        `bson:"name" json:"name"`
	PhoneCode   string        `bson:"phoneCode" json:"phoneCode"`
}

