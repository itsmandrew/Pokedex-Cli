package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/itsmandrew/Pokedex-Cli/models"
)

// -- Mock Cache ---
type mockCache struct {
	storage map[string][]byte
}

func (c *mockCache) Get(key string) ([]byte, bool) {
	v, ok := c.storage[key]
	return v, ok
}

func (c *mockCache) Add(key string, data []byte) {
	c.storage[key] = data
}

// --- Unit Test for fetchAndUnmarshal ---

func TestFetchAndUnmarshal_Success(t *testing.T) {
	expectedJSON := `{
		"count": 1,
		"next": null,
		"previous": null, 
		"results": [{"name": "area-1", "url": "url-1"}]
	}`

	// Mock http server

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedJSON)
	}))

	defer server.Close()

	cache := &mockCache{storage: make(map[string][]byte)}
	var config m.Config

	err := fetchAndUnmarshal(server.URL, cache, &config)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if config.Results[0].Name != "area-1" {
		t.Errorf("Unmarshaled config is incorrect: %+v", config)
	}

	// Verify it got cached
	cached, ok := cache.Get(server.URL)
	if !ok || len(cached) == 0 {
		t.Errorf("Expected response to be cached")
	}
}
