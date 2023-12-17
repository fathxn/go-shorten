package dto

type (
	UserRegisterInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserLoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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
