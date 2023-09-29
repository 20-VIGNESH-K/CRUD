package models

type Profile struct {
	Name  string `json:"name" bson:"name" validate:"required,customValidator"`
	Age   int32  `json:"age" bson:"age"`
	Place string `json:"place" bson:"place" validate:"required"`
}
