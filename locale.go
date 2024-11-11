package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type AdministrativeArea struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Code      string             `json:"code" bson:"code"`
	CountryID primitive.ObjectID `json:"country_id" bson:"country_id"`
}

type SubAdministrativeArea struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	AdminAreaID primitive.ObjectID `json:"admin_area_id" bson:"admin_area_id"`
}
