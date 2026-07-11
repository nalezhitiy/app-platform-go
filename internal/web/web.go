package web

import "github.com/danielgtaylor/huma/v2"

type API interface {
	huma.API
	GetName() string
	GetBasePath() string
}

type BaseAPI struct {
	huma.API
	Name     string
	BasePath string
}

func (b BaseAPI) GetName() string     { return b.Name }
func (b BaseAPI) GetBasePath() string { return b.BasePath }
