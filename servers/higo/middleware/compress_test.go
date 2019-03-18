package middleware

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/higo"
	"github.com/stretchr/testify/assert"
)

func TestGzip(t *testing.T) {
	e := higo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Skip if no Accept-Encoding header
	h := Gzip()(func(c higo.Context) error {
		c.Response().Write([]byte("test")) // For Content-Type sniffing
		return nil
	})
	h(c)

	assert := assert.New(t)

	assert.Equal("test", rec.Body.String())

	// Gzip
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(higo.HeaderAcceptEncoding, gzipScheme)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	h(c)
	assert.Equal(gzipScheme, rec.Header().Get(higo.HeaderContentEncoding))
	assert.Contains(rec.Header().Get(higo.HeaderContentType), higo.MIMETextPlain)
	r, err := gzip.NewReader(rec.Body)
	if assert.NoError(err) {
		buf := new(bytes.Buffer)
		defer r.Close()
		buf.ReadFrom(r)
		assert.Equal("test", buf.String())
	}
}

func TestGzipNoContent(t *testing.T) {
	e := higo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(higo.HeaderAcceptEncoding, gzipScheme)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := Gzip()(func(c higo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})
	if assert.NoError(t, h(c)) {
		assert.Empty(t, rec.Header().Get(higo.HeaderContentEncoding))
		assert.Empty(t, rec.Header().Get(higo.HeaderContentType))
		assert.Equal(t, 0, len(rec.Body.Bytes()))
	}
}

func TestGzipErrorReturned(t *testing.T) {
	e := higo.New()
	e.Use(Gzip())
	e.GET("/", func(c higo.Context) error {
		return higo.ErrNotFound
	})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(higo.HeaderAcceptEncoding, gzipScheme)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Empty(t, rec.Header().Get(higo.HeaderContentEncoding))
}

// Issue #806
func TestGzipWithStatic(t *testing.T) {
	e := higo.New()
	e.Use(Gzip())
	e.Static("/test", "../_fixture/images")
	req := httptest.NewRequest(http.MethodGet, "/test/walle.png", nil)
	req.Header.Set(higo.HeaderAcceptEncoding, gzipScheme)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	// Data is written out in chunks when Content-Length == "", so only
	// validate the content length if it's not set.
	if cl := rec.Header().Get("Content-Length"); cl != "" {
		assert.Equal(t, cl, rec.Body.Len())
	}
	r, err := gzip.NewReader(rec.Body)
	if assert.NoError(t, err) {
		defer r.Close()
		want, err := ioutil.ReadFile("../_fixture/images/walle.png")
		if assert.NoError(t, err) {
			buf := new(bytes.Buffer)
			buf.ReadFrom(r)
			assert.Equal(t, want, buf.Bytes())
		}
	}
}
