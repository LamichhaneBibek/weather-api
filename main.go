package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		WeatherForcast("Bibek").Render(r.Context(), w)
	})
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
