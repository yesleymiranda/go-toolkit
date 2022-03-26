package webapplication_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yesleymiranda/go-toolkit/webapplication"
)

func Test_NewApp(t *testing.T) {
	app := webapplication.New(&webapplication.ApplicationConfig{})
	assert.NotNil(t, app)
}

func Test_Initialize(t *testing.T) {
	assert.NotPanics(t, func() {
		app := webapplication.New(&webapplication.ApplicationConfig{})
		app.Initialize()
		assert.NotNil(t, app)
	})
}

func Test_WithPing(t *testing.T) {
	assert.NotPanics(t, func() {
		app := webapplication.New(&webapplication.ApplicationConfig{
			WithPing: true,
		})
		app.Initialize()
		assert.NotNil(t, app)
	})
}

func TestApi_ListenAndServe(t *testing.T) {
	_ = os.Setenv("PLATFORM", "local")
	ch := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	app := webapplication.New(&webapplication.ApplicationConfig{
		WithPing: true,
	})
	app.Initialize()

	go func() {
		ch <- app.ListenAndServe()
	}()

	for {
		select {
		case <-ctx.Done():
			t.Fatal("timeout waiting for ping")
		case err := <-ch:
			t.Fatalf("running application: %v", err)
		case <-time.Tick(50 * time.Millisecond):
			_, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
			if err != nil {
				t.Fatal(err)
			}
			return
		}
	}
}
