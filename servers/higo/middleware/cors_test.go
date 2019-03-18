package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/higo"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	e := higo.New()

	// Wildcard origin
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := CORS()(higo.NotFoundHandler)
	h(c)
	assert.Equal(t, "*", rec.Header().Get(higo.HeaderAccessControlAllowOrigin))

	// Allow origins
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	h = CORSWithConfig(CORSConfig{
		AllowOrigins: []string{"localhost"},
	})(higo.NotFoundHandler)
	req.Header.Set(higo.HeaderOrigin, "localhost")
	h(c)
	assert.Equal(t, "localhost", rec.Header().Get(higo.HeaderAccessControlAllowOrigin))

	// Preflight request
	req = httptest.NewRequest(http.MethodOptions, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	req.Header.Set(higo.HeaderOrigin, "localhost")
	req.Header.Set(higo.HeaderContentType, higo.MIMEApplicationJSON)
	cors := CORSWithConfig(CORSConfig{
		AllowOrigins:     []string{"localhost"},
		AllowCredentials: true,
		MaxAge:           3600,
	})
	h = cors(higo.NotFoundHandler)
	h(c)
	assert.Equal(t, "localhost", rec.Header().Get(higo.HeaderAccessControlAllowOrigin))
	assert.NotEmpty(t, rec.Header().Get(higo.HeaderAccessControlAllowMethods))
	assert.Equal(t, "true", rec.Header().Get(higo.HeaderAccessControlAllowCredentials))
	assert.Equal(t, "3600", rec.Header().Get(higo.HeaderAccessControlMaxAge))

	// Preflight request with `AllowOrigins` *
	req = httptest.NewRequest(http.MethodOptions, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	req.Header.Set(higo.HeaderOrigin, "localhost")
	req.Header.Set(higo.HeaderContentType, higo.MIMEApplicationJSON)
	cors = CORSWithConfig(CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           3600,
	})
	h = cors(higo.NotFoundHandler)
	h(c)
	assert.Equal(t, "localhost", rec.Header().Get(higo.HeaderAccessControlAllowOrigin))
	assert.NotEmpty(t, rec.Header().Get(higo.HeaderAccessControlAllowMethods))
	assert.Equal(t, "true", rec.Header().Get(higo.HeaderAccessControlAllowCredentials))
	assert.Equal(t, "3600", rec.Header().Get(higo.HeaderAccessControlMaxAge))
}
