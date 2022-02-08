package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	first_name string             `bson:"first_name"`
	surname    string             `bson:"surname"`
	dob        time.Time          `bson:"dob"`
	password   string             `bson:"password"`
	email      string             `bson:"email"`
	area_code  string             `bson:"area_code"`
	number     string             `bson:"number"`
	created time.Date() `bson:created`

}
