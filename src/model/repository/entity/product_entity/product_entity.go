package product_entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductEntity struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	Name          string        `bson:"name,omitempty"`
	Description   string        `bson:"description,omitempty"`
	Value         int64         `bson:"value,omitempty"` // valor em centavos
	Images        []string      `bson:"images,omitempty"`
	Stock         int32         `bson:"stock,omitempty"`
	RatingAverage *float64      `bson:"rating_average,omitempty"`
	RatingCount   int32         `bson:"rating_count,omitempty"`
	CreatedAt     time.Time     `bson:"created_at,omitempty"`
	UpdatedAt     time.Time     `bson:"updated_at,omitempty"`
	DeletedAt     *time.Time    `bson:"deleted_at,omitempty"`
}
