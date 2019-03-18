package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/higo"
	"github.com/stretchr/testify/assert"
)

func TestSecure(t *testing.T) {
	e := higo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := func(c higo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	// Default
	Secure()(h)(c)
	assert.Equal(t, "1; mode=block", rec.Header().Get(higo.HeaderXXSSProtection))
	assert.Equal(t, "nosniff", rec.Header().Get(higo.HeaderXContentTypeOptions))
	assert.Equal(t, "SAMEORIGIN", rec.Header().Get(higo.HeaderXFrameOptions))
	assert.Equal(t, "", rec.Header().Get(higo.HeaderStrictTransportSecurity))
	assert.Equal(t, "", rec.Header().Get(higo.HeaderContentSecurityPolicy))

	// Custom
	req.Header.Set(higo.HeaderXForwardedProto, "https")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	SecureWithConfig(SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	})(h)(c)
	assert.Equal(t, "", rec.Header().Get(higo.HeaderXXSSProtection))
	assert.Equal(t, "", rec.Header().Get(higo.HeaderXContentTypeOptions))
	assert.Equal(t, "", rec.Header().Get(higo.HeaderXFrameOptions))
	assert.Equal(t, "max-age=3600; includeSubdomains", rec.Header().Get(higo.HeaderStrictTransportSecurity))
	assert.Equal(t, "default-src 'self'", rec.Header().Get(higo.HeaderContentSecurityPolicy))
}
