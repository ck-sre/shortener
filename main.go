package main

import (
	"fmt"
	"net/http"
	urlshort "shortener/shorten"
)

func main() {

	urls := map[string]string{
		"/goog": "https://google.com",
		"/yah":  "https://yahoo.com",
	}
	handlerMap := urlshort.HandlerMap(urls, failsafeMux())

	yml := `
- path: /smallurl
  url: "https://github.com/gophercises/urlshort"
- path: /smallurllast
  url: "https://github.com/gophercises/urlshort/tree/final"
`
	fmt.Println(yml)
	handlerYaml, err := urlshort.HandlerYAML([]byte(yml), handlerMap)
	if err != nil {
		fmt.Println("uh oh yaml not hanled")
		panic(err)
	}
	fmt.Println("Listening on :9080")
	http.ListenAndServe(":9080", handlerYaml)

}

func failsafeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greeting)
	return mux
}
func greeting(a http.ResponseWriter, b *http.Request) {
	fmt.Println("Dear world here we come")
}
