package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Beats struct {
	ID     primitive.ObjectID   `bson:"_id"`
	producer   primitive.ObjectID   `bson:"producer"`
	name string `bson:"title"`
	album primitive.ObjectID `bson:"album"`
	image string `bson:"image_url"`
	audio string `bson:"audio_url"`
	genre primitive.ObjectID `bson:"genre"`
	created time.Date() `bson:created`

}
