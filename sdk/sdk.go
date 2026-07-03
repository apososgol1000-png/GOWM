package sdk

type Client struct {
    conn net.Conn
}

func Connect() (*Client, error) {
    // conecta al socket IPC
}
