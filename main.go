package main

import (
    "log"
    "github.com/beevik/ntp"
)

func main() {
    time, err := ntp.Time("pool.ntp.org")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Current time: %v", time.Format("02 January 2006 15:04:05.000"))
}
