package httpnet

import (
	"fmt"
	"io"
	"net/http"
)

func Run() {
	http.HandleFunc("/", Create)

	http.ListenAndServe(":8080", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "error")
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("request body:", string(data))

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "aplication/json")

	w.Write(data)
}
