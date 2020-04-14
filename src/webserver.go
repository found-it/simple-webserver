/*
 *
 *  Simple Webserver that serves stuff
 */
package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "strconv"
    "sync"
)

// Globals
var cnt int
var mutex = &sync.Mutex{}

/*
 *  Increment function used for testing
 */
func Increment(a *int) int {
    *a++
    return *a
}


/*
 *  Callback for /increment
 */
func incrementCounter(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    Increment(&cnt)
    fmt.Fprintf(w, strconv.Itoa(cnt))
    mutex.Unlock()
}


func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

    http.HandleFunc("/increment", incrementCounter)

    log.Fatal(http.ListenAndServe(":8081", nil))
}
