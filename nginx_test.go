package main

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNGINX(t *testing.T) {
	testCases := map[string]struct {
		path            string
		code            int
		proxyPassHeader string
	}{
		"Root":  {"/", 200, "https://jayson-dev.netlify.com/"},
		"Blog":  {"/blog", 200, "https://jayson-dev.netlify.com/blog"},
		"About": {"/about", 200, "https://jayson-dev.netlify.com/about"},
	}

	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			var url strings.Builder
			url.WriteString(baseURL())
			url.WriteString(tc.path)

			resp, err := http.Get(url.String())
			assert.NoError(tt, err)
			defer resp.Body.Close()

			assert.Equal(tt, tc.code, resp.StatusCode)
			foundProxyPass := resp.Header["X-Proxy-Pass"]
			require.NotEmpty(tt, foundProxyPass)
			assert.Equal(tt, tc.proxyPassHeader, foundProxyPass[0])
		})
	}
}

func baseURL() string {
	envHost := os.Getenv("HOST")
	if envHost != "" {
		return envHost
	}

	return "http://localhost:9300/"
}
