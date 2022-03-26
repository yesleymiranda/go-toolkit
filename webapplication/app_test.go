package webapplication_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yesleymiranda/go-toolkit/webapplication"
)

func Test_NewApp(t *testing.T) {
	app := webapplication.New("8088")

	assert.NotNil(t, app)
}

func Test_Initialize(t *testing.T) {
	assert.NotPanics(t, func() {
		app := webapplication.New("8088")
		app.Initialize()
		assert.NotNil(t, app)
	})
}
