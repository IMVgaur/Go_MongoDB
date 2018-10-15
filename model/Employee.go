package model

type Employee struct {
	Name string `json:"name" bson:"name"`
	ID   string `json:"id" bson:"id"`
}
