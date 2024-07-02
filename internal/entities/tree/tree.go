package tree

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	Latitude       float64 `json:"lat" bson:"lat"`
	Longitude      float64 `json:"lon" bson:"lon"`
	Address        string  `json:"address" bson:"address"`
	AdditionalInfo string  `json:"additional_info" bson:"additional_info"`
}

type Tree struct {
	ID       primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Species  string             `json:"species" bson:"species"`
	TreeNum  int                `json:"tree_num" bson:"tree_num"`
	Age      int                `json:"age" bson:"age"`
	Location Location           `json:"location" bson:"location"`
}
