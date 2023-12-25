package dto

type (
	UserRegisterInput struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	UserLoginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	UserAuthResponse struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	}

	Auth struct {
		Id string
	}
)
