// +build !go1.11

package middleware

import (
	"net/http"
	"net/http/httputil"

	"github.com/labstack/higo"
)

func proxyHTTP(t *ProxyTarget, c higo.Context, config ProxyConfig) http.Handler {
	return httputil.NewSingleHostReverseProxy(t.URL)
}
