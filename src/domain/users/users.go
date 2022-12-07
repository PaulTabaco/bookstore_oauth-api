package users

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"` // `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`  // `json:"last_name" binding:"required"`
	Email     string `json:"email"`      // `json:"email" binding:"required,email"`
	// DateCreated string `json:"date_created"`
	// Status      string `json:"status"`
	// Password    string `json:"password"` //`json:"-"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
