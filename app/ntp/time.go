package ntp

import "time"

func Run() {
    ticker := time.NewTicker(60 * time.Second)
    for {
        select {
        case <-ticker.C:
            GetTime()
        }
    }
}
