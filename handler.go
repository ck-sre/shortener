package smallurl

import "net/http"

func HandlerMap(urlPaths map[string]string, fb http.Handler) http.HandlerFunc {
	return func(a http.ResponseWriter, b *http.Request) {
		urlPath := b.URL.Path
		if d, ok := urlPaths[urlPath]; ok {
			http.Redirect(a, b, d, http.StatusFound)
		}
		fb.ServeHTTP(a, b)
	}
}

func HandlerYAML(yml []byte, fb http.Handler) (http.HandlerFunc, error) {
	paths := map[string]string{}
	var URLPath
	return HandlerMap(paths,)
}

type URLPath struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}