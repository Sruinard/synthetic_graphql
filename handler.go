package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func retrieve_users(w http.ResponseWriter, r *http.Request) {


    // create a json object with fake user data containing name and occupation
    // and return it as a response
    message = `[
        {
            "name": "John Doe",
            "occupation": "Software Engineer"
        },
        {
            "name": "Jane Doe",
            "occupation": "Software Engineer"
        }
    ]`
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, message)
}

func main() {
    listenAddr := ":8080"
    if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
        listenAddr = ":" + val
    }
    http.HandleFunc("/api/users", retrieve_users)
    log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
    log.Fatal(http.ListenAndServe(listenAddr, nil))
}