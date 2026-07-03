package ipc

import (
    "encoding/json"
    "net"
    "time"
)

type Client struct {
    path string
}

func NewClient(path string) *Client {
    return &Client{path: path}
}

func (c *Client) Send(req Request) error {
    conn, err := net.DialTimeout("unix", c.path, 500*time.Millisecond)
    if err != nil {
        return err
    }
    defer conn.Close()
    enc := json.NewEncoder(conn)
    return enc.Encode(req)
}
