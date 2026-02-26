package allowed_city_entity

import "go.mongodb.org/mongo-driver/v2/bson"

type AllowedCityEntity struct {
	ID   bson.ObjectID `bson:"_id,omitempty"`
	Name string        `bson:"name"`
}