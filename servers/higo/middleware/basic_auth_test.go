package middleware

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sereiner/parrot/servers/higo"
	"github.com/stretchr/testify/assert"
)

func TestBasicAuth(t *testing.T) {
	e := higo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	f := func(u, p string, c higo.Context) (bool, error) {
		if u == "joe" && p == "secret" {
			return true, nil
		}
		return false, nil
	}
	h := BasicAuth(f)(func(c higo.Context) error {
		return c.String(http.StatusOK, "test")
	})

	assert := assert.New(t)

	// Valid credentials
	auth := basic + " " + base64.StdEncoding.EncodeToString([]byte("joe:secret"))
	req.Header.Set(higo.HeaderAuthorization, auth)
	assert.NoError(h(c))

	h = BasicAuthWithConfig(BasicAuthConfig{
		Skipper:   nil,
		Validator: f,
		Realm:     "someRealm",
	})(func(c higo.Context) error {
		return c.String(http.StatusOK, "test")
	})

	// Valid credentials
	auth = basic + " " + base64.StdEncoding.EncodeToString([]byte("joe:secret"))
	req.Header.Set(higo.HeaderAuthorization, auth)
	assert.NoError(h(c))

	// Case-insensitive header scheme
	auth = strings.ToUpper(basic) + " " + base64.StdEncoding.EncodeToString([]byte("joe:secret"))
	req.Header.Set(higo.HeaderAuthorization, auth)
	assert.NoError(h(c))

	// Invalid credentials
	auth = basic + " " + base64.StdEncoding.EncodeToString([]byte("joe:invalid-password"))
	req.Header.Set(higo.HeaderAuthorization, auth)
	he := h(c).(*higo.HTTPError)
	assert.Equal(http.StatusUnauthorized, he.Code)
	assert.Equal(basic+` realm="someRealm"`, res.Header().Get(higo.HeaderWWWAuthenticate))

	// Missing Authorization header
	req.Header.Del(higo.HeaderAuthorization)
	he = h(c).(*higo.HTTPError)
	assert.Equal(http.StatusUnauthorized, he.Code)

	// Invalid Authorization header
	auth = base64.StdEncoding.EncodeToString([]byte("invalid"))
	req.Header.Set(higo.HeaderAuthorization, auth)
	he = h(c).(*higo.HTTPError)
	assert.Equal(http.StatusUnauthorized, he.Code)
}
