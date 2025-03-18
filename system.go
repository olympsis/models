package models

type SystemConfig struct {
	Tags   []Tag   `json:"tags"`
	Sports []Sport `json:"sports"`
}

type Tag struct {
	Name      string `json:"name" bson:"name"`
	IsEnabled bool   `json:"is_enabled" bson:"is_enabled"`
}

type Sport struct {
	Name      string   `json:"name" bson:"name"`
	Images    []string `json:"images" bson:"images"`
	IsEnabled bool     `json:"is_enabled" bson:"is_enabled"`
}
