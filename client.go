package main

import (
    "bufio"
    "log"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    response, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    if response != "OK\n" {
        log.Fatal("response != OK\n" + response)
    }
}
