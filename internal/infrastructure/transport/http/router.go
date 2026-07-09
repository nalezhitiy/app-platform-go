package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func NewBaseRouter(log *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	// 1. Захист від паніки (щоб один кривий хендлер не впустив весь сервер)
	r.Use(middleware.Recoverer)

	// 2. Прокидання Request ID для трейсингу логів
	r.Use(middleware.RequestID)

	// 3. Реальний IP клієнта (корисно за проксі на кшталт Nginx/Cloudflare)
	r.Use(middleware.RealIP)

	// 4. Наш кастомний мідлварь-логера для Zap
	r.Use(ZapHttpLogger(log))

	return r
}

// ZapHttpLogger адаптує наш структурований логер під запити Chi
func ZapHttpLogger(log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Створюємо обгортку над ResponseWriter, щоб витягнути HTTP статус відповіді
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			// Логуємо кінець запиту в нашому чистому форматі
			log.Info("http request processed",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status", ww.Status()),
				zap.String("ip", r.RemoteAddr),
				zap.Duration("latency", time.Since(start)),
				zap.String("req_id", middleware.GetReqID(r.Context())),
			)
		})
	}
}
