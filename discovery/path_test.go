package discovery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsoluteURL(t *testing.T) {
	tests := []struct {
		base string
		in   string
		out  string
	}{
		{"http://mama", "./bp1.yaml", "http://mama/bp1.yaml"},
		{"http://mama", "https://papa", "https://papa"},
		{"https://mama", "bp2.yaml", "https://mama/bp2.yaml"},
		{"https://mama/hello/to/world.html", "../bp2.yaml", "https://mama/hello/bp2.yaml"},
		{"https://mama/hello/to/", "../bp2.yaml", "https://mama/hello/bp2.yaml"},
		{"https://mama/hello/", "../bp2.yaml", "https://mama/bp2.yaml"},
		{"https://mama/", "../bp2.yaml", "https://mama/bp2.yaml"},
	}

	for _, test := range tests {
		res, err := absoluteURL(test.base, test.in)
		assert.NoError(t, err)
		assert.Equal(t, test.out, res)
	}
}
