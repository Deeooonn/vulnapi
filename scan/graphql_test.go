package scan_test

import (
	"net/http"
	"testing"

	"github.com/cerberauth/vulnapi/internal/auth"
	"github.com/cerberauth/vulnapi/internal/request"
	"github.com/cerberauth/vulnapi/scan"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGraphQLScan(t *testing.T) {
	s, err := scan.NewGraphQLScan("http://localhost:8080", nil, nil)

	require.NoError(t, err)
	assert.Equal(t, "http://localhost:8080", s.Operations[0].Request.URL.String())
	assert.Equal(t, "POST", s.Operations[0].Request.Method)
	assert.Equal(t, []auth.SecurityScheme{auth.NewNoAuthSecurityScheme()}, s.Operations[0].SecuritySchemes)
}

func TestNewGraphQLScanWithUpperCaseAuthorizationHeader(t *testing.T) {
	header := http.Header{}
	token := "token"
	header.Add("Authorization", "Bearer "+token)
	client := request.NewClient(header, nil)

	s, err := scan.NewGraphQLScan("http://localhost:8080", client, nil)

	require.NoError(t, err)
	assert.Equal(t, "http://localhost:8080", s.Operations[0].Request.URL.String())
	assert.Equal(t, "POST", s.Operations[0].Request.Method)
	assert.Equal(t, []auth.SecurityScheme{auth.NewAuthorizationBearerSecurityScheme("default", &token)}, s.Operations[0].SecuritySchemes)
}

func TestNewGraphQLScanWithUpperCaseAuthorizationAndLowerCaseBearerHeader(t *testing.T) {
	header := http.Header{}
	token := "token"
	header.Add("Authorization", "bearer "+token)
	client := request.NewClient(header, nil)

	s, err := scan.NewGraphQLScan("http://localhost:8080", client, nil)

	require.NoError(t, err)
	assert.Equal(t, "http://localhost:8080", s.Operations[0].Request.URL.String())
	assert.Equal(t, "POST", s.Operations[0].Request.Method)
	assert.Equal(t, []auth.SecurityScheme{auth.NewAuthorizationBearerSecurityScheme("default", &token)}, s.Operations[0].SecuritySchemes)
}

func TestNewGraphQLScanWithLowerCaseAuthorizationHeader(t *testing.T) {
	header := http.Header{}
	token := "token"
	header.Add("authorization", "Bearer "+token)
	client := request.NewClient(header, nil)

	s, err := scan.NewGraphQLScan("http://localhost:8080", client, nil)

	require.NoError(t, err)
	assert.Equal(t, "http://localhost:8080", s.Operations[0].Request.URL.String())
	assert.Equal(t, "POST", s.Operations[0].Request.Method)
	assert.Equal(t, []auth.SecurityScheme{auth.NewAuthorizationBearerSecurityScheme("default", &token)}, s.Operations[0].SecuritySchemes)
}
