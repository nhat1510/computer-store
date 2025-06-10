package types

type ForgotPasswordInput struct {
    Email string `json:"email" binding:"required,email"`
}
