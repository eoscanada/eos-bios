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

	net := NewNetwork("/tmp/disco", ts.URL+"/bp1.yaml")
	assert.NoError(t, net.FetchAll())
	assert.NoError(t, net.VerifyGraph())
	assert.NoError(t, net.CalculateWeights())
	assert.Equal(t, 0.5, net.discoveredPeers[ts.URL+"/bp1.yaml"].TotalWeight)
	assert.Equal(t, 1.0, net.discoveredPeers[ts.URL+"/bp2.yaml"].TotalWeight)
	assert.Equal(t, 1.0, net.discoveredPeers[ts.URL+"/bp3.yaml"].TotalWeight)
	assert.Equal(t, ts.URL+"/bp2.yaml", net.orderedPeers[0].DiscoveryURL)
	assert.Equal(t, ts.URL+"/bp3.yaml", net.orderedPeers[1].DiscoveryURL)
	assert.Equal(t, ts.URL+"/bp1.yaml", net.orderedPeers[2].DiscoveryURL)
}

func newFileServer() *httptest.Server {
	curDir, _ := os.Getwd()
	ts := httptest.NewServer(http.FileServer(http.Dir(filepath.Join(curDir, "test-data"))))
	return ts
}
