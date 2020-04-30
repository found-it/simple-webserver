package main

import (
    "flag"
    "fmt"
    "log"
    "net"
    "net/http"
    "os"
)

func Green(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Green")
}

func Blue(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Blue!")
}

func GetIPs(w http.ResponseWriter, r *http.Request) {
    ifaces, err := net.Interfaces()
    // handle err
    if err != nil {
        log.Fatal(err)
    }
    for _, i := range ifaces {
        addrs, err := i.Addrs()
        // handle err
        if err != nil {
            log.Fatal(err)
        }
        for _, addr := range addrs {
            var ip net.IP
            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }
            if ip == nil || ip.IsLoopback() {
                continue
            }
            ip = ip.To4()
            if ip == nil {
                continue
            }
            fmt.Fprintln(w,"Node IP: ", ip.String())
        }
    }
    name, err := os.Hostname()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintln(w, "Pod: ", name)
}


func main() {
    flag.Parse()

    http.Handle("/images", http.FileServer(http.Dir("./images/")))

    http.HandleFunc("/", GetIPs)
    http.HandleFunc("/green", Green)
    http.HandleFunc("/blue", Blue)

    log.Fatal(http.ListenAndServe(":5000", nil))
}
