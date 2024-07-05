package tree

import "go.mongodb.org/mongo-driver/bson/primitive"

type TreeLocationEntity struct {
	Latitude       float64 `bson:"latitude"`
	Longitude      float64 `bson:"longitude"`
	Address        string  `bson:"address"`
	AdditionalInfo string  `bson:"additional_info"`
}

type TreeEntity struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Species  string             `bson:"species"`
	TreeNum  int                `bson:"tree_num"`
	Age      int                `bson:"age"`
	Location TreeLocationEntity `bson:"location"`
}
