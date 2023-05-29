package main

import (
	"fmt"
	smallurl "https://github.com/ck89/shortener"
	"net/http"
)

func main() {

	urls := map[string]string{
		"/goog": "https://google.com",
		"/yah":  "https://yahoo.com",
	}
	handlerMap := smallurl.MapHandler(urls, failsafeMux())

	yml := `
- path:/smallurl
  url: "https://godoc.org/github.com/ck89/smallurl"
- path: /smallurllast
  url: "https://godoc.org/github.com/ck89/shortener/tree/"
`
	handlerYaml, err := smallurl.HandlerYAML([]byte(yml), handlerMap)
	if err != nil {
		fmt.Println("uh oh yaml not hanled")
		panic(err)
	}
	fmt.Println("Listening on :9080")
	http.ListenAndServe(":9080", handlerMap)

}

func failsafeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greeting)
}
func greeting(a http.ResponseWriter, b *http.Request) {
	fmt.Println("Dear world here we come")
}
