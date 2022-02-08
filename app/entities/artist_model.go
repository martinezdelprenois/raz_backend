package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Artist struct {
	ID     primitive.ObjectID   `bson:"_id"`
	User   primitive.ObjectID   `bson:"user"`
	albums []primitive.ObjectID `bson:"album"`
	tracks []primitive.ObjectID `bson:"tracks"`
	image string `bson:"image_url"`
	audio string `bson:"audio_url"`
	created time.Date() `bson:created`
}
