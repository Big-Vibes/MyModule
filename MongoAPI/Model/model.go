package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive" // Fixed: removed /v2 for consistency
)

type NetflixEX struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
