package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/jcobhams/mmhmm/config"
	"github.com/jcobhams/mmhmm/controllers"
	"github.com/jcobhams/mmhmm/env"
	"github.com/jcobhams/mmhmm/logging"
	"github.com/jcobhams/mmhmm/repositories"
	"github.com/jcobhams/mmhmm/services"
)

func BuildServer(
	ctx context.Context,
	sc *services.Container,
	rc *repositories.Container,
) *echo.Echo {
	ctx, logger := logging.Logger(ctx, "BuildServer")

	e := echo.New()
	e.Use(middleware.Recover(), middleware.RequestID())

	e.Use(
		requestLoggerMiddleware(logger.Named("RequestLogger")),
		// add Other Middleware here | Authentication, Authorization, RateLimiting, CORS, etc.
	)

	// Register Routes
	controllers.BindRoutes(e, sc, rc)

	return e
}

func Start(
	ctx context.Context,
	sc *services.Container,
	rc *repositories.Container,
	conf config.Config,
) {
	e := BuildServer(ctx, sc, rc)

	//Start Server
	go func() {
		if err := e.Start(conf.GetString(env.ServerAddress)); err != nil {
			e.Logger.Info("Shutting Down...")
		}
	}()

	//Wait For interrupt signal gracefully shutdown the server after timeout.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// requestLoggerMiddleware is a middleware that logs the request properties.
func requestLoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields := []zap.Field{
				zap.String("method", req.Method),
				zap.Int("status", res.Status),
				zap.String("uri", req.RequestURI),
				zap.String("remote_ip", c.RealIP()),
				zap.Time("time", time.Now().UTC()),
				zap.String("id", id),
				zap.String("host", req.Host),
				zap.String("latency", stop.Sub(start).String()),
				zap.String("referer", req.Referer()),
			}

			code := res.Status
			switch {
			case code >= 500:
				logger.With(fields...).Error("Server Error", zap.Error(err))
				return
			case code >= 400 && code < 500:
				logger.With(fields...).Warn("Client Error", zap.Error(err))
				return
			}

			logger.With(fields...).Debug("")
			return
		}
	}
}
