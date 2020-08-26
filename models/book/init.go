package book

import "github.com/globalsign/mgo/bson"

type Book struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Title  string        `bson:"title" json:"title"`
	Genre  string        `bson:"genre" json:"genre"`
	Author string        `bson:"author" json:"author"`
}

const (
	db         = "Books"
	collection = "MovieModel"
)
