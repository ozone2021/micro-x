package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

func main() {
    // TODO hashing testline DELETE AFTER
    name := os.Getenv("SERVICE")

    if name == "" {
        log.Fatalln("SERVICE env var not set")
    }
    namePascal := strings.Replace(strings.ToUpper(name), "-", "_", -1)
    port := os.Getenv(fmt.Sprintf("%s_PORT", namePascal))
    if port == "" {
        log.Fatalf("%s_PORT env var not set", namePascal)
    }

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

        fmt.Printf("Logging for %s.", name)
        w.WriteHeader(http.StatusOK)
        msg := fmt.Sprintf("Runnable %s is OK", name)
        w.Write([]byte(msg))
    })

    addr := fmt.Sprintf(":%s", port)
    http.ListenAndServe(addr, nil)
}