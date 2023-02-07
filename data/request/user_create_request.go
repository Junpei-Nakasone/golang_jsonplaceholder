package request

type UserCreateRequest struct {
	Name     string `validate:"required min=1,max=100" json:"name"`
	UserName string `validate:"required min=1,max=100" json:"username"`
	Email    string `validate:"required min=1,max=100" json:"email"`
}
