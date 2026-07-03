package wm

import "github.com/BurntSushi/xgb/xproto"

// Client representa una ventana administrada por AstralWM.
type Client struct {
	// ID de la ventana en X11.
	Window xproto.Window

	// Geometría.
	X      int
	Y      int
	Width  int
	Height int

	// Estado.
	Mapped     bool
	Floating   bool
	Fullscreen bool
	Focused    bool
	Urgent     bool
	Minimized  bool

	// Apariencia.
	BorderWidth uint32
	BorderColor uint32

	// Workspace al que pertenece.
	Workspace *Workspace

	// Monitor donde está.
	Monitor *Monitor

	// Nombre de la ventana.
	Title string

	// Clase WM_CLASS (Firefox, Kitty, etc.).
	Class string

	// Instancia WM_CLASS.
	Instance string
}

// NewClient crea un nuevo cliente.
func NewClient(win xproto.Window) *Client {
	return &Client{
		Window:       win,
		BorderWidth:  2,
		BorderColor:  0x444444,
		Mapped:       false,
		Floating:     false,
		Fullscreen:   false,
		Focused:      false,
		Urgent:       false,
		Minimized:    false,
	}
}

// Geometry devuelve la geometría actual.
func (c *Client) Geometry() (int, int, int, int) {
	return c.X, c.Y, c.Width, c.Height
}

// SetGeometry actualiza la geometría del cliente.
func (c *Client) SetGeometry(x, y, w, h int) {
	c.X = x
	c.Y = y
	c.Width = w
	c.Height = h
}

// Focus marca el cliente como enfocado.
func (c *Client) Focus() {
	c.Focused = true
}

// Unfocus quita el foco del cliente.
func (c *Client) Unfocus() {
	c.Focused = false
}

// ToggleFloating cambia el modo flotante.
func (c *Client) ToggleFloating() {
	c.Floating = !c.Floating
}

// ToggleFullscreen cambia el modo pantalla completa.
func (c *Client) ToggleFullscreen() {
	c.Fullscreen = !c.Fullscreen
}

// Map marca la ventana como visible.
func (c *Client) Map() {
	c.Mapped = true
}

// Unmap marca la ventana como oculta.
func (c *Client) Unmap() {
	c.Mapped = false
}

// Minimize minimiza la ventana.
func (c *Client) Minimize() {
	c.Minimized = true
}

// Restore restaura la ventana.
func (c *Client) Restore() {
	c.Minimized = false
}

// SetWorkspace cambia el workspace del cliente.
func (c *Client) SetWorkspace(ws *Workspace) {
	c.Workspace = ws
}

// SetMonitor cambia el monitor del cliente.
func (c *Client) SetMonitor(m *Monitor) {
	c.Monitor = m
}
