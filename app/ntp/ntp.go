package ntp

import (
    "encoding/binary"
    "fmt"
    log "github.com/sirupsen/logrus"
    "net"
    "time"
)

const ntpEpochOffset = 2208988800

type packet struct {
    Settings       uint8  // leap yr indicator, ver number, and mode
    Stratum        uint8  // stratum of local clock
    Poll           int8   // poll exponent
    Precision      int8   // precision exponent
    RootDelay      uint32 // root delay
    RootDispersion uint32 // root dispersion
    ReferenceID    uint32 // reference id
    RefTimeSec     uint32 // reference timestamp sec
    RefTimeFrac    uint32 // reference timestamp fractional
    OrigTimeSec    uint32 // origin time secs
    OrigTimeFrac   uint32 // origin time fractional
    RxTimeSec      uint32 // receive time secs
    RxTimeFrac     uint32 // receive time frac
    TxTimeSec      uint32 // transmit time secs
    TxTimeFrac     uint32 // transmit time frac
}

// GetTime
// implements a trivial NTP client over UDP.
func GetTime(params ...string) {
    start := time.Now()
    // TODO: use fastest ntp server
    host := "us.pool.ntp.org:123"
    if len(params) > 0 {
        host = params[0]
    }

    // Setup a UDP connection
    conn, err := net.Dial("udp", host)
    if err != nil {
        log.Fatalf("[NTP] failed to connect: %v", err)
    }
    defer func(conn net.Conn) {
        _ = conn.Close()
    }(conn)

    if err = conn.SetDeadline(time.Now().Add(15 * time.Second)); err != nil {
        log.Fatalf("[NTP] failed to set deadline: %v", err)
    }

    // configure request settings by specifying the first byte as
    // 00 011 011 (or 0x1B)
    // |  |   +-- client mode (3)
    // |  + ----- version (3)
    // + -------- leap year indicator, 0 no warning
    req := &packet{Settings: 0x1B}

    // send time request
    if err = binary.Write(conn, binary.BigEndian, req); err != nil {
        log.Fatalf("[NTP] failed to send request: %v", err)
    }

    // block to receive server response
    rsp := &packet{}
    if err = binary.Read(conn, binary.BigEndian, rsp); err != nil {
        log.Fatalf("[NTP] failed to read server response: %v", err)
    }

    // On POSIX-compliant OS, time is expressed
    // using the Unix time epoch (or secs since year 1970).
    // NTP seconds are counted since 1900 and therefore must
    // be corrected with an epoch offset to convert NTP seconds
    // to Unix time by removing 70 yrs of seconds (1970-1900)
    // or 2208988800 seconds.
    secs := float64(rsp.TxTimeSec) - ntpEpochOffset
    nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32 // convert fractional to nanos
    fmt.Printf("%v\n", time.Unix(int64(secs), nanos))
    fmt.Printf("used %d ms\n", time.Now().Sub(start).Milliseconds())
}
