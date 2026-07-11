package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
)

func main() {
	port := 8080
	if s := os.Getenv("PORT"); s != "" {
		n, err := strconv.Atoi(s)
		if err != nil || n < 1 || n > 65535 {
			log.Fatal("PORT must be a number between 1 and 65535")
		}
		port = n
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, fmt.Sprintf(":%d", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

// newServer builds the Echo instance with all routes registered.
func newServer() *echo.Echo {
	e := echo.New()
	e.GET("/", handleIndex)
	e.GET("/health", handleHealth)
	return e
}

// run serves HTTP on addr until ctx is cancelled, then shuts down gracefully.
func run(ctx context.Context, addr string) error {
	srv := &http.Server{
		Addr:              addr,
		Handler:           newServer(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       2 * time.Minute,
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe()
	}()
	log.Printf("listening on http://localhost%s", addr)

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		log.Print("shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return srv.Shutdown(shutdownCtx)
	}
}

func handleIndex(c *echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>It works!</h1><p>Your Go web project is running.</p>")
}

func handleHealth(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
