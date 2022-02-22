package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Writer struct {
	ID         primitive.ObjectID `bson:"_id"`
	user primitive.ObjectID             `bson:"user"`
	tracks    []primitive.ObjectID             `bson:"tracks"`
	created time.Date() `bson:created`

}
