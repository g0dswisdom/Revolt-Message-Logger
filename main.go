package main

import (
    "os"
    "os/signal"
    "syscall"
    "fmt"
    "github.com/5elenay/revoltgo"
)

func main() {
    // Init a new client.
    client := revoltgo.Client{
        SelfBot: &revoltgo.SelfBot{
            Id:           "session id",
            SessionToken: "session token",
            UserId:       "user id",
        },
    }

    // Listen a on message event.
    client.OnMessage(func(m *revoltgo.Message) {
        fmt.Print(m.Content)
    })

    // Start the client.
    client.Start()

    // Wait for close.
    sc := make(chan os.Signal, 1)

    signal.Notify(
        sc,
        syscall.SIGINT,
        syscall.SIGTERM,
        os.Interrupt,
    )
    <-sc

    // Destroy client.
    client.Destroy()
}
