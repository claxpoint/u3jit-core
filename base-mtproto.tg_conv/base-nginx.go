package main

import (
    "fmt"
    "io"
    "net/http"
    "net/http/httputil"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Host == "localhost:8080" { 
            targetURL := "http://example.com"
            targetResp, err := http.Get(targetURL + r.URL.Path)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            defer targetResp.Body.Close()

            
        -io.Copy(w, targetResp.Body)

         
            w.Header().Set("Content-Type", targetResp.Header.Get("Content-Type"))
        } else {
            http.Error(w, "Invalid request", http.StatusForbidden)
            return
        }
    })

    http.ListenAndServe(":8080", nil)
}
