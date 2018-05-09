package bios

import (
	"net/url"
	"path"
	"regexp"
	"strings"
)

func absoluteURL(base, relURL string) (string, error) {
	if strings.HasPrefix(relURL, "https:") || strings.HasPrefix(relURL, "http:") {
		return relURL, nil
	}

	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	newPath := path.Clean(path.Join(path.Dir(baseURL.Path), relURL))

	baseURL.Path = newPath

	return baseURL.String(), nil
}

var weirdities = regexp.MustCompile("[^a-zA-Z0-9]")

func replaceAllWeirdities(input string) string {
	return weirdities.ReplaceAllString(input, "_")
}
