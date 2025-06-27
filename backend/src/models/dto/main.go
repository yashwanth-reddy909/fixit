package dto

// TOTO: Data Transfer Object(DTO) need to be in separate file based on the entity

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type GoogleCallBackResponse struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GoogleAuthCallback struct {
	Code  string `json:"code" binding:"required"`
	State string `json:"state" binding:"required"`
}
