package ipc

import (
    "bufio"
    "encoding/json"
    "log"
    "net"
    "os"
)

func Start(requestCh chan<- Request, socketPath string) error {
    // remove stale socket
    if _, err := os.Stat(socketPath); err == nil {
        os.Remove(socketPath)
    }
    l, err := net.Listen("unix", socketPath)
    if err != nil {
        return err
    }
    // ensure socket is removed on exit
    go func() {
        <-requestCh // sentinel not used; this is just to keep goroutine alive if needed
        l.Close()
    }()

    log.Printf("ipc: listening on %s", socketPath)
    for {
        conn, err := l.Accept()
        if err != nil {
            // listener closed
            return err
        }
        go func(c net.Conn) {
            defer c.Close()
            scanner := bufio.NewScanner(c)
            for scanner.Scan() {
                line := scanner.Bytes()
                var req Request
                if err := json.Unmarshal(line, &req); err != nil {
                    log.Printf("ipc: invalid request: %v", err)
                    continue
                }
                requestCh <- req
            }
            if err := scanner.Err(); err != nil {
                log.Printf("ipc: connection error: %v", err)
            }
        }(conn)
    }
}
