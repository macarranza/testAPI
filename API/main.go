package main

import (
    "log"
    "net/http"
)

/* Main function */
func main() {
    router := NewRouter() /* Create a router/handler for all defined routes */
    log.Fatal(http.ListenAndServe(":8080", router)) /* Starts the server listening to localhost:8080 */
}

