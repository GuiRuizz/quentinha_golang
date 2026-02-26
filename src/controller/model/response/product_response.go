package response

import "time"

type ProductResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Value         int64     `json:"value"` // valor em centavos
	Images        []string  `json:"images"`
	Stock         int32     `json:"stock"`
	RatingAverage *float64  `json:"rating_average,omitempty"`
	RatingCount   int32     `json:"rating_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsActive      bool      `json:"isActive"`
}
