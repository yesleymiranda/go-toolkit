package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeJSON_NoContent(t *testing.T) {
	// Given
	w := httptest.NewRecorder()

	// When
	err := EncodeJSON(w, nil, http.StatusNoContent)

	// Then
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, w.Code)
	require.Empty(t, w.Body)
}

func TestEncodeJSON_NilBody(t *testing.T) {
	// Given
	w := httptest.NewRecorder()

	// When
	err := EncodeJSON(w, nil, http.StatusOK)

	// Then
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "null", w.Body.String())
}

func TestEncodeJSON(t *testing.T) {
	tt := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "it should respond using a map",
			input:    make(map[string]interface{}),
			expected: "{}",
		},
		{
			name:     "it should respond using a map",
			input:    []byte(`{}`),
			expected: "{}",
		},
		{
			name:     "it should respond using a map",
			input:    strings.NewReader("{}"),
			expected: "{}",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			w := httptest.NewRecorder()

			// When
			err := EncodeJSON(w, tc.input, 200)

			// Then
			require.NoError(t, err)
			require.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-type"))
			require.Equal(t, tc.expected, w.Body.String())
			require.Equal(t, 200, w.Code)
		})
	}
}

type responseWithHeaders struct{}

func (responseWithHeaders) Headers() http.Header {
	h := make(http.Header)
	h.Set("X-Custom", "custom")
	return h
}

func TestEncodeJSON_CustomHeaders(t *testing.T) {
	// Given
	w := httptest.NewRecorder()

	// When
	err := EncodeJSON(w, responseWithHeaders{}, 200)

	// Then
	require.NoError(t, err)
	require.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-type"))
	require.Equal(t, "custom", w.Header().Get("X-Custom"))
	require.Equal(t, 200, w.Code)
}
