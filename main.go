package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/propagation"
	"mulesoft.org/go-product-api/routers"
	"mulesoft.org/go-product-api/services"
)

func main() {

	// Handle SigINT (ctrl + c) gracefully
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Setup OpenTelemetry
	sName := "go-product-api"
	sVersion := "1.0.12"
	otelShutdown, err := initTracer(ctx, sName, sVersion)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	router := gin.Default()
	router.Use(otelgin.Middleware(sName, otelgin.WithPropagators(propagation.TraceContext{})))

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"api": "Composable Commerce Product API", "version": sVersion})
	})

	// Product Routes
	pr := router.Group("/api/")
	routers.ProductRoutes(pr)

	services.ConnectDatabase()

	srvErr := make(chan error, 1)

	go func() {
		srvErr <- router.Run("0.0.0.0:8080")
	}()
	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	return
}
