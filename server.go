package main

import (
    "fmt"
    "log"
    "net"
)

func main() {
    PORT := ":8080"
    listener, err := net.Listen("tcp", PORT)
    if err != nil {
        log.Fatal(err)
    }

    defer listener.Close()
    fmt.Println("Listening on " + PORT)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Connected with", conn.RemoteAddr().String())
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    defer conn.Close()
    conn.Write([]byte("OK\n"))
}
