package dump

import (
	"fmt"
	"net/http"
)

func simpleHello(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "simpleHello from Slava\n")
	if err != nil {
		return
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				return
			}
		}
	}
}

// HttpServers for start should from cmd be main
func HttpServers() {
	http.HandleFunc("/hello", simpleHello)
	http.HandleFunc("/headers", headers)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
