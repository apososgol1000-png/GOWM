package config

type Config struct {
    BorderWidth int    `json:"border_width"`
    Gaps        int    `json:"gaps"`
    FocusColor  string `json:"focus_color"`
    Terminal    string `json:"terminal"`
}
