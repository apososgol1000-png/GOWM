package config

import "fmt"

func Validate(cfg *Config) error {
    if cfg == nil {
        return fmt.Errorf("config is nil")
    }
    if cfg.BorderWidth < 0 || cfg.BorderWidth > 100 {
        return fmt.Errorf("border width out of range")
    }
    if cfg.Gaps < 0 || cfg.Gaps > 200 {
        return fmt.Errorf("gaps out of range")
    }
    if cfg.FocusColor == "" {
        return fmt.Errorf("focus color empty")
    }
    if cfg.Terminal == "" {
        return fmt.Errorf("terminal empty")
    }
    return nil
}
