package dump

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(5 * time.Second):
		_, err := fmt.Fprintf(w, "hello from Slava\n")
		if err != nil {
			return
		}
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func ContextFunc() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
