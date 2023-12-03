package httpnet

import (
	"fmt"
	"net/http"
)

func Go() {
	http.HandleFunc("/", Do)

	http.ListenAndServe(":8080", nil)
}

func Do(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
