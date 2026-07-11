package common

import "go.uber.org/fx"

func InjectDI(name string) fx.Option {
	return fx.Options(
		fx.Invoke(
			fx.Annotate(
				RegisterHealthcheck,
				fx.ParamTags(`name:"`+name+`"`),
			),
			fx.Annotate(
				RegisterVersion,
				fx.ParamTags(`name:"`+name+`"`),
			),
		),
	)
}

func InjectAuthDI(name string) fx.Option {
	return fx.Options(
		fx.Invoke(
			fx.Annotate(
				//RegisterAuth,
				fx.ParamTags(`name:"` + name + `"`),
			),
		),
	)
}
