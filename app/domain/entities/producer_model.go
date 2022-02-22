package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producer struct {
	ID     primitive.ObjectID   `bson:"_id"`
	User   primitive.ObjectID   `bson:"user"`
	albums []primitive.ObjectID `bson:"album"`
	tracks []primitive.ObjectID `bson:"tracks"`
	beats  []primitie.ObjectID  `bson:"beats"`
	image string `bson:"image_url"`
	created time.Date() `bson:created`

}
