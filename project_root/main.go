package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "project-my-test/example/rpcApp"
)

func main() {
    
    // init
    rpcClient := rpcApp.CreateTestRpc()
    
    
    // page - rpc
    http.HandleFunc("/", rpcClient.RunHttp)
    // page - api form
    http.HandleFunc("/api-form", func(w http.ResponseWriter, r *http.Request) {
        pwd, _ := os.Getwd()
        text, err := ioutil.ReadFile(pwd + "/template/api.html")
        if err != nil {
            w.Write([]byte("Не удалось считать файл"))
            return
        }
        w.Header().Set("Content-Type", "text/html")
        w.Write(text)
    })
    // page - browser favicon fix
    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(""))
    })
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
