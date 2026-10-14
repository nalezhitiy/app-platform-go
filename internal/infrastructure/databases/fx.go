package databases

import "go.uber.org/fx"

var Module = fx.Module(
	"databases",
	fx.Provide(
		NewMainPostgresSession,
		NewMainMariaDBSession,
		NewMainClickhouseSession,
	),
)
