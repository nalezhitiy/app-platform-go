package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"platform/internal/infrastructure/configs"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(router *chi.Mux, cfg *configs.Config, log *zap.Logger) *http.Server {
	server := &http.Server{
		Addr:         cfg.WebServer.Host + ":" + cfg.WebServer.Port,
		Handler:      router,
		ReadTimeout:  cfg.WebServer.Timeout.Read,
		WriteTimeout: cfg.WebServer.Timeout.Write,
		IdleTimeout:  cfg.WebServer.Timeout.Idle,
	}

	log.Info("http server created", zap.String("addr", server.Addr))
	return server
}

func ServerRegisterHooks(l fx.Lifecycle, server *http.Server, log *zap.Logger) {
	l.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}

			go func() {
				err := server.Serve(ln)
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error("web server error", zap.Error(err))
					os.Exit(1)
				}
			}()

			log.Info("web server started", zap.String("addr", server.Addr))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("stopping http server...")
			err := server.Shutdown(ctx)
			if err != nil {
				return err
			}
			log.Info("web server stopped cleanly")
			return nil
		},
	})
}
