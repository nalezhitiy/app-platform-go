package web

import "github.com/danielgtaylor/huma/v2"

type APIInterface interface {
	huma.API
	GetName() string
	GetBasePath() string
	GetAuthCookieName() string
}

type BaseAPI struct {
	huma.API
	Name           string
	BasePath       string
	AuthCookieName string
}

func (b BaseAPI) GetName() string     { return b.Name }
func (b BaseAPI) GetBasePath() string { return b.BasePath }
func (b BaseAPI) GetAuthCookieName() string {
	return b.AuthCookieName
}

type BaseBodyRequest[T any] struct {
	Body T
}

type BaseBodyResponse[T any] struct {
	Body T `huma:"body, application/json"`
}

func NewBaseBodyResponse[T any](body T) *BaseBodyResponse[T] {
	return &BaseBodyResponse[T]{
		Body: body,
	}
}
