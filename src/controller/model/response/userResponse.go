package response

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Age      int8   `json:"age" binding:"omitempty,min=0"`
}
