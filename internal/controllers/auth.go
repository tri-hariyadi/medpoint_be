package controllers

import (
	"github.com/sev-2/raiden"
)

type AuthRequest struct { // Add your request data
}

type AuthResponse struct {
	Message string `json:"message"`
}

type AuthController struct {
	raiden.ControllerBase
	Http    string `path:"/auth" type:"custom"`
	Payload *AuthRequest
	Result  AuthResponse
}

func (c *AuthController) Post(ctx raiden.Context) error {
	c.Result.Message = "Success Login"
	return ctx.SendJson(c.Result)
}
