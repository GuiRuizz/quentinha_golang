package product_request

// ProductRequest represents the input data for creating a new product.
// @Summary Product Input Data
// @Description Structure containing the required fields for creating a new product.
type ProductRequest struct {

	// Product name (required, minimum 3 and maximum 120 characters).
	// Example: Quentinha Fit Frango
	Name string `json:"name" binding:"required,min=3,max=120" example:"Quentinha Fit Frango"`

	// Product description (required, minimum 10 and maximum 500 characters).
	// Example: Arroz integral, frango grelhado e legumes frescos.
	Description string `json:"description" binding:"required,min=10,max=500" example:"Arroz integral, frango grelhado e legumes frescos."`

	// Product price in cents (required, must be greater than 0).
	// Example: 2490 (R$ 24,90)
	Value int64 `json:"value" binding:"required,gt=0" example:"2490"`

	// List of product image URLs (optional).
	// Example: ["https://example.com/image1.jpg"]
	Images []string `json:"images" example:"https://example.com/image1.jpg"`

	// Product stock quantity (required, must be 0 or greater).
	// Example: 50
	Stock int32 `json:"stock" binding:"required,gte=0" example:"50"`
}

// ProductUpdateRequest represents the input data for updating a product.
// @Summary Product Update Data
// @Description Structure containing fields allowed to update a product.
type ProductUpdateRequest struct {

	// Product name (required, minimum 3 and maximum 120 characters).
	Name string `json:"name" binding:"required,min=3,max=120" example:"Quentinha Fit Atualizada"`

	// Product description (required, minimum 10 and maximum 500 characters).
	Description string `json:"description" binding:"required,min=10,max=500" example:"Nova descrição do produto."`

	// Product price in cents (must be greater than 0).
	Value int64 `json:"value" binding:"required,gt=0" example:"2790"`

	// Product images list.
	Images []string `json:"images" example:"https://example.com/image1.jpg"`

	// Product stock quantity.
	Stock int32 `json:"stock" binding:"required,gte=0" example:"100"`

	// Product avaibility
	IsActive *bool `json:"isActive" binding:"required" example:"true"`
}

// ProductReviewRequest represents the input data for adding a review to a product.
// @Summary Product Review Input
// @Description Structure containing score for product evaluation.
type ProductReviewRequest struct {

	// Review score (required, must be between 0 and 5).
	// Example: 4.5
	Score float64 `json:"score" binding:"required,gte=0,lte=5" example:"4.5"`
}
