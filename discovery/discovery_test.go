package discovery

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscoveryDir(t *testing.T) {
	ts := newFileServer()
	defer ts.Close()

	disco := NewCache("/tmp/disco", ts.URL+"/bp1.yaml")
	assert.NoError(t, disco.EnsureExists())
	assert.NoError(t, disco.FetchAll())
	assert.NoError(t, disco.VerifyGraph())
	assert.NoError(t, disco.CalculateWeights())
	assert.Equal(t, 0.5, disco.discoveredPeers[ts.URL+"/bp1.yaml"].TotalWeight)
	assert.Equal(t, 1.0, disco.discoveredPeers[ts.URL+"/bp2.yaml"].TotalWeight)
	assert.Equal(t, 1.0, disco.discoveredPeers[ts.URL+"/bp3.yaml"].TotalWeight)
}

func newFileServer() *httptest.Server {
	curDir, _ := os.Getwd()
	ts := httptest.NewServer(http.FileServer(http.Dir(filepath.Join(curDir, "test-data"))))
	return ts
}
