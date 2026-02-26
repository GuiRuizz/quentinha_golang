package address_entity

import "go.mongodb.org/mongo-driver/v2/bson"

type AddressEntity struct {
	ID      bson.ObjectID `bson:"_id,omitempty"`
	City    string        `bson:"city,omitempty"`
	Street  string        `bson:"street,omitempty"`
	Number  string        `bson:"numbe,omitempty"`
	CEP     string        `bson:"cep,omitempty"`
	IsValid bool          `bson:"isValid"`
}
