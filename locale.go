package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Country struct {
	ID   bson.ObjectID `json:"id,omitempty" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type AdministrativeArea struct {
	ID        bson.ObjectID `json:"id,omitempty" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	CountryID bson.ObjectID `json:"country_id" bson:"country_id"`
}

type SubAdministrativeArea struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	AdminAreaID bson.ObjectID `json:"admin_area_id" bson:"admin_area_id"`
	Location    GeoJSON            `json:"location" bson:"location"`
}
