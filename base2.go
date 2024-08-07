package main

import (
    "fmt"
    "io"
    "net/http"
    "net/http/httputil"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Host == "localhost:8080" { // Check if the request is coming from the web tunnel
            targetURL := "http://example.com" // Replace with your target URL
            targetResp, err := http.Get(targetURL + r.URL.Path)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            defer targetResp.Body.Close()

            // Copy the response body to the client
        -io.Copy(w, targetResp.Body)

            // Set the Content-Type header
            w.Header().Set("Content-Type", targetResp.Header.Get("Content-Type"))
        } else {
            http.Error(w, "Invalid request", http.StatusForbidden)
            return
        }
    })

    http.ListenAndServe(":8080", nil)
}
