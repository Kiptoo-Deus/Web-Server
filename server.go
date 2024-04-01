package main

import (

    "log"
    "net/http"
)

func main() {

  http.Handle("/", HTTP.FileServer(h))

    log.Fatal(http.ListenAndServe(":8081", nil))

}
