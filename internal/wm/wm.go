package wm

import (
    "astralwm/internal/config"
    "astralwm/internal/ipc"
    "context"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "time"
)

type Manager struct {
    cfg   *config.Config
    reqCh chan ipc.Request
    quit  chan struct{}
}

func New() (*Manager, error) {
    cfgPath := os.Getenv("ASTRALWM_CONFIG")
    if cfgPath == "" {
        home, _ := os.UserHomeDir()
        cfgPath = filepath.Join(home, ".config", "astralwm", "config.json")
    }
    cfg, err := config.Load(cfgPath)
    if err != nil {
        cfg = config.Default()
    }
    if err := config.Validate(cfg); err != nil {
        return nil, err
    }

    m := &Manager{
        cfg:   cfg,
        reqCh: make(chan ipc.Request, 32),
        quit:  make(chan struct{}),
    }

    socket := "/tmp/astralwm.sock"
    go func() {
        if err := ipc.Start(m.reqCh, socket); err != nil {
            log.Printf("ipc server stopped: %v", err)
        }
    }()

    return m, nil
}

func (m *Manager) Run() error {
    log.Printf("AstralWM (minimal) starting with cfg: %+v", m.cfg)
    for {
        select {
        case req := <-m.reqCh:
            m.handle(req)
        case <-m.quit:
            log.Println("AstralWM: shutting down")
            return nil
        }
    }
}

func (m *Manager) handle(req ipc.Request) {
    switch req.Command {
    case "reload":
        log.Println("handling reload")
        m.reload()
    case "spawn":
        if len(req.Args) > 0 {
            m.spawn(req.Args[0], req.Args[1:]...)
        } else {
            m.spawn(m.cfg.Terminal)
        }
    default:
        log.Printf("unhandled command: %s (%v)", req.Command, req.Args)
    }
}

func (m *Manager) reload() {
    cfgPath := os.Getenv("ASTRALWM_CONFIG")
    if cfgPath == "" {
        home, _ := os.UserHomeDir()
        cfgPath = filepath.Join(home, ".config", "astralwm", "config.json")
    }
    if cfg, err := config.Load(cfgPath); err == nil {
        if err := config.Validate(cfg); err == nil {
            m.cfg = cfg
            log.Println("config reloaded")
            return
        } else {
            log.Println("config validation failed:", err)
        }
    } else {
        log.Println("config load failed:", err)
    }
}

func (m *Manager) spawn(cmd string, args ...string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    c := exec.CommandContext(ctx, cmd, args...)
    c.Stdout = os.Stdout
    c.Stderr = os.Stderr
    if err := c.Start(); err != nil {
        log.Printf("spawn %s failed: %v", cmd, err)
        return
    }
    log.Printf("spawned %s pid=%d", cmd, c.Process.Pid)
}
