package auth

import (
	"context"
	"fmt"
	"net/http"
	"platform/internal/web"

	"github.com/danielgtaylor/huma/v2"
)

type LoginInput struct {
	Email    string `json:"email" example:"<EMAIL>" doc:"Email address" format:"email"`
	Password string `json:"password" example:"<PASSWORD>" doc:"Password" minLength:"8" maxLength:"128"`
}

type LoginOutput struct {
	Token string `json:"token" example:"<KEY>" doc:"Token"`
}

type LoginResponse struct {
	SetCookie http.Cookie `header:"Set-Cookie"`
	Body      LoginOutput
}

func RegisterLogin(api web.APIInterface) {
	desc := huma.Operation{
		OperationID: "login",
		Summary:     "Login to platform",
		Tags:        []string{TAG},
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/%s/login/password", TAG),
		Security:    nil,
	}

	handler := func(ctx context.Context, input *web.BaseBodyRequest[LoginInput]) (*LoginResponse, error) {
		return &LoginResponse{
			SetCookie: http.Cookie{
				Name:     api.GetAuthCookieName(),
				Value:    "refreshToken",
				HttpOnly: true,
				Path:     fmt.Sprintf("%s/%s", api.GetBasePath(), TAG),
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
				MaxAge:   3600,
			},
			Body: LoginOutput{"accessToken"},
		}, nil
	}

	huma.Register(api, desc, handler)
}
