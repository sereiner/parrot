package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/higo"
	"github.com/stretchr/testify/assert"
)

type middlewareGenerator func() higo.MiddlewareFunc

func TestRedirectHTTPSRedirect(t *testing.T) {
	res := redirectTest(HTTPSRedirect, "labstack.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "https://labstack.com/", res.Header().Get(higo.HeaderLocation))
}

func TestHTTPSRedirectBehindTLSTerminationProxy(t *testing.T) {
	header := http.Header{}
	header.Set(higo.HeaderXForwardedProto, "https")
	res := redirectTest(HTTPSRedirect, "labstack.com", header)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestRedirectHTTPSWWWRedirect(t *testing.T) {
	res := redirectTest(HTTPSWWWRedirect, "labstack.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "https://www.labstack.com/", res.Header().Get(higo.HeaderLocation))
}

func TestRedirectHTTPSWWWRedirectBehindTLSTerminationProxy(t *testing.T) {
	header := http.Header{}
	header.Set(higo.HeaderXForwardedProto, "https")
	res := redirectTest(HTTPSWWWRedirect, "labstack.com", header)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestRedirectHTTPSNonWWWRedirect(t *testing.T) {
	res := redirectTest(HTTPSNonWWWRedirect, "www.labstack.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "https://labstack.com/", res.Header().Get(higo.HeaderLocation))
}

func TestRedirectHTTPSNonWWWRedirectBehindTLSTerminationProxy(t *testing.T) {
	header := http.Header{}
	header.Set(higo.HeaderXForwardedProto, "https")
	res := redirectTest(HTTPSNonWWWRedirect, "www.labstack.com", header)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestRedirectWWWRedirect(t *testing.T) {
	res := redirectTest(WWWRedirect, "labstack.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "http://www.labstack.com/", res.Header().Get(higo.HeaderLocation))
}

func TestRedirectNonWWWRedirect(t *testing.T) {
	res := redirectTest(NonWWWRedirect, "www.labstack.com", nil)

	assert.Equal(t, http.StatusMovedPermanently, res.Code)
	assert.Equal(t, "http://labstack.com/", res.Header().Get(higo.HeaderLocation))
}

func redirectTest(fn middlewareGenerator, host string, header http.Header) *httptest.ResponseRecorder {
	e := higo.New()
	next := func(c higo.Context) (err error) {
		return c.NoContent(http.StatusOK)
	}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Host = host
	if header != nil {
		req.Header = header
	}
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	fn()(next)(c)

	return res
}
