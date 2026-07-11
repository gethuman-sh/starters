package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRoutes(t *testing.T) {
	srv := httptest.NewServer(newServer())
	defer srv.Close()

	tests := []struct {
		path        string
		wantStatus  int
		wantType    string
		wantContain string
	}{
		{"/", http.StatusOK, "text/html; charset=UTF-8", "It works!"},
		{"/health", http.StatusOK, "application/json", `"status":"ok"`},
		{"/nope", http.StatusNotFound, "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			resp, err := http.Get(srv.URL + tt.path)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Fatalf("status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}
			if tt.wantType != "" && resp.Header.Get("Content-Type") != tt.wantType {
				t.Errorf("Content-Type = %q, want %q", resp.Header.Get("Content-Type"), tt.wantType)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), tt.wantContain) {
				t.Errorf("body = %q, want it to contain %q", body, tt.wantContain)
			}
		})
	}
}

func TestRunShutsDownOnContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error, 1)
	go func() {
		errCh <- run(ctx, "127.0.0.1:0")
	}()

	// Give the server a moment to start, then trigger graceful shutdown.
	time.Sleep(100 * time.Millisecond)
	cancel()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("run() = %v, want nil after graceful shutdown", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("run() did not return after context cancellation")
	}
}
