package cache

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"cache",
	fx.Provide(
		NewMainCache,
		NewSessionCache,
	),
	fx.Invoke(RegisterLifecycle),
)

func RegisterLifecycle(
	lc fx.Lifecycle,
	main *MainCache,
	session *SessionCache,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := main.Open(ctx); err != nil {
				return fmt.Errorf("open main cache: %w", err)
			}

			if err := session.Open(ctx); err != nil {
				_ = main.Close()
				return fmt.Errorf("open session cache: %w", err)
			}

			return nil
		},
		OnStop: func(context.Context) error {
			var mainErr, sessionErr error

			sessionErr = session.Close()
			mainErr = main.Close()

			if sessionErr != nil {
				sessionErr = fmt.Errorf("close session cache: %w", sessionErr)
			}
			if mainErr != nil {
				mainErr = fmt.Errorf("close main cache: %w", mainErr)
			}

			return errors.Join(sessionErr, mainErr)
		},
	})
}
