package main

import (
    "net/http"
    "fmt"
)

func main() {

    fs := http.FileServer(http.Dir("Game/"))
    
    // Wrap the file server to add cache-busting headers
    http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
        w.Header().Set("Pragma", "no-cache")
        w.Header().Set("Expires", "0")
        fs.ServeHTTP(w, r)
    }))

    fmt.Println("Sharing Gmae on port 80")
    http.ListenAndServe(":80", nil)
}
