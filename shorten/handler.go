package shorten

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

func HandlerMap(urlPaths map[string]string, fb http.Handler) http.HandlerFunc {
	return func(a http.ResponseWriter, b *http.Request) {
		urlPath := b.URL.Path
		if d, ok := urlPaths[urlPath]; ok {
			http.Redirect(a, b, d, http.StatusFound)
			return
		}
		fb.ServeHTTP(a, b)
	}
}

func HandlerYAML(yml []byte, fb http.Handler) (http.HandlerFunc, error) {
	//var pu []URLPath
	//err := yaml.Unmarshal(yml, &pu)
	pUrls, err := YamlParser(yml)
	if err != nil {
		return nil, err
	}

	paths := createMap(pUrls)

	return HandlerMap(paths, fb), nil
}

type URLPath struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func createMap(pUrlsMap []URLPath) map[string]string {
	pToURLs := make(map[string]string)
	for _, pu := range pUrlsMap {
		pToURLs[pu.Path] = pu.URL
	}
	return pToURLs
}

func YamlParser(by []byte) ([]URLPath, error) {
	var pUrls []URLPath
	err := yaml.Unmarshal(by, &pUrls)
	if err != nil {
		return nil, err
	}
	return pUrls, nil
}
