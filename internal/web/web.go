package web

import "github.com/danielgtaylor/huma/v2"

type APIInstance struct {
	huma.API

	Name     string
	BasePath string
}
