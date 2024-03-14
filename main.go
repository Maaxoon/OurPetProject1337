package main

import (
	"fmt"
	"net/http"
)

func handleGET(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    fmt.Fprintf(w, "Hello, %s!", name)
}

func handlePOST(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }
    name := r.FormValue("name")
    fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {

	http.HandleFunc("/handle-get", handleGET)
    http.HandleFunc("/handle-post", handlePOST)

	http.ListenAndServe(":80", nil)
=======
import "fmt"

func main() {
	fmt.Println("hello world")
>>>>>>> 3ece21d0ff182f61ad670d8da6ffdfae3eb3adc1
}
