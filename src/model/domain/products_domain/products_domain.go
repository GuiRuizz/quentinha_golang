package products_domain

import "time"

type productDomain struct {
	id            string
	name          string
	description   string
	value         int64
	images        []string
	stock         int32
	ratingAverage *float64
	ratingCount   int32
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     *time.Time
}


func (pd *productDomain) GetID() string {
	return pd.id
}

func (pd *productDomain) GetName() string {
	return pd.name
}

func (pd *productDomain) GetDescription() string {
	return pd.description
}

func (pd *productDomain) GetValue() int64 {
	return pd.value
}

func (pd *productDomain) GetImages() []string {
	return pd.images
}

func (pd *productDomain) GetRatingAverage() *float64 {
	return pd.ratingAverage
}

func (pd *productDomain) GetRatingCount() int32 {
	return pd.ratingCount
}

func (pd *productDomain) GetStock() int32 {
	return pd.stock
}

func (pd *productDomain) GetCreatedAt() time.Time {
	return pd.createdAt
}

func (pd *productDomain) GetUpdatedAt() time.Time {
	return pd.updatedAt
}

func (pd *productDomain) GetDeletedAt() *time.Time {
	return pd.deletedAt
}
func (pd *productDomain) SetID(id string) {
	pd.id = id
}

func (pd *productDomain) SetUpdatedAt() {
	pd.updatedAt = time.Now()
}

func (pd *productDomain) SetDeleted() {
	now := time.Now()
	pd.deletedAt = &now
}

func (pd *productDomain) AddReview(score float64) {
	if score < 0 || score > 5 {
		return
	}

	if pd.ratingAverage == nil {
		pd.ratingCount = 1
		pd.ratingAverage = &score
	} else {
		total := (*pd.ratingAverage * float64(pd.ratingCount)) + score
		pd.ratingCount++
		newAverage := total / float64(pd.ratingCount)
		pd.ratingAverage = &newAverage
	}

	pd.updatedAt = time.Now()
}
