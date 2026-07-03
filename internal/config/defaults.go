package config

func Default() Config {
    return Config{
        BorderWidth: 2,
        Gaps:        8,
        FocusColor:  "#ff0000",
        Terminal:    "xterm",
    }
}
