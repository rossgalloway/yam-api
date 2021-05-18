package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AprYam struct {
	ID    primitive.ObjectID     `bson:"_id"`
	Value map[string]interface{} `bson:"apryam"`
}
type AprDegenerative struct {
	ID    primitive.ObjectID     `bson:"_id"`
	Value map[string]interface{} `bson:"aprdegenerative"`
}
