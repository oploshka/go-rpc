package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "project-my-test/example"
)

func main() {
    
    // rpc
    http.HandleFunc("/", example.TestHttpRpc)
    
    // api page
    http.HandleFunc("/api-form", func(w http.ResponseWriter, r *http.Request) {
        dat, err := ioutil.ReadFile("./static/api.html")
        if err != nil {
            w.Write([]byte("Не удалось считать файл"))
            return
        }
        w.Header().Set("Content-Type", "text/html")
        w.Write(dat)
    })
    // browser favicon fix
    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(""))
    })
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
