package client

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestUrl(t *testing.T) {
	assert := assert.New(t)
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	assert.Nil(err)
	assert.Equal(u.Scheme, "http", "Scheme if equal")
	assert.Equal(u.Host, "bing.com", "Host if equal")

}
