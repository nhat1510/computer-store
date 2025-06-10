package types

type ResetPasswordInput struct {
    Email       string `json:"email" binding:"required,email"`
    Code        string `json:"code" binding:"required"`
    NewPassword string `json:"new_password" binding:"required,min=6"`
}
