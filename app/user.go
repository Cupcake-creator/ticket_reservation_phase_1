package app

type RegisterParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type RegisterResult struct {
	ID int64 `json:"id"`
}

func (ctx *Context) Register(params RegisterParams) (*RegisterResult, error) {
	logger := ctx.getLogger()

	if err := validateInput(params); err != nil {
		logger.Errorf("validateInput error : %s", err)
		return nil, err
	}

	// Check if the user is already registered then create new user
	isExistUser, err := ctx.DB.IsExistUser(params.Username)
	if err != nil && isExistUser {
		return nil, err
	}

	userID, err := ctx.DB.CreateUser(params.Username, params.Password, params.Email)
	if err != nil {
		return nil, err
	}

	return &RegisterResult{
		ID: userID,
	}, nil
}
