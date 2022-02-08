package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Genre struct {
	ID     primitive.ObjectID   `bson:"_id"`
	name string `bson:"name"`
	created time.Date() `bson:created`

}
