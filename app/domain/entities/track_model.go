package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Track struct {
	ID     primitive.ObjectID   `bson:"_id"`
	Users   primitive.ObjectID   `bson:"artist"`
	featuring []primitive.ObjectID `bson:"featuring_artists"`
	title string `bson:"title"`
	album primitie.ObjectID `bson:"album"`
	writer []primitie.ObjectID `bson:"writers"`
	producer primitie.ObjectID 	`bson:"producer"`
	image string `bson:"image_url"`
	audio string `bson:"audio_url"`
	genre primitive.ObjectID `bson:"genre"`
	created time.Date() `bson:created`


}
