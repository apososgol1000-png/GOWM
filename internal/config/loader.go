package config

import (
    "encoding/json"
    "os"
)

func Load(path string) (*Config, error) {
    if path == "" {
        d := Default()
        return &d, nil
    }
    f, err := os.Open(path)
    if err != nil {
        // fallback to defaults if file missing
        d := Default()
        return &d, nil
    }
    defer f.Close()

    var cfg Config
    if err := json.NewDecoder(f).Decode(&cfg); err != nil {
        return nil, err
    }

    // fill missing with defaults
    d := Default()
    if cfg.BorderWidth == 0 {
        cfg.BorderWidth = d.BorderWidth
    }
    if cfg.Gaps == 0 {
        cfg.Gaps = d.Gaps
    }
    if cfg.FocusColor == "" {
        cfg.FocusColor = d.FocusColor
    }
    if cfg.Terminal == "" {
        cfg.Terminal = d.Terminal
    }

    return &cfg, nil
}
