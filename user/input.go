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

//FormCreateUserInput is struct
type FormCreateUserInput struct {
	Name 		string `form:"name" binding:"required"`
	Email 		string `form:"email" binding:"required,email"`
	Occupation 	string `form:"occupation" binding:"required"`
	Password 	string `form:"password" binding:"required"`
	Error error
}

//FormUpdateUserInput is struct
type FormUpdateUserInput struct {
	ID 			int 
	Name 		string `form:"name" binding:"required"`
	Email 		string `form:"email" binding:"required,email"`
	Occupation 	string `form:"occupation" binding:"required"`
	Error error
}

