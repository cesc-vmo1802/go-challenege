package database

import (
	"github.com/globalsign/mgo/bson"
)

type MgoModel struct {
	PK        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Status    int           `json:"status" bson:"status,omitempty"`
	CreatedAt *JSONTime     `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt *JSONTime     `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt *JSONTime     `json:"deleted_at" bson:"deleted_at,omitempty"`
}
