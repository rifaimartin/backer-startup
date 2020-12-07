package user

//RegisterUserInput is struct
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

//LoginInput is struct
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

//CheckEmailInput is struct
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
