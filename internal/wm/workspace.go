package wm

// Workspace representa un escritorio virtual.
type Workspace struct {
	// ID del workspace.
	ID int

	// Nombre mostrado (1, 2, Web, Code, etc.).
	Name string

	// Ventanas pertenecientes al workspace.
	Clients []*Client

	// Cliente enfocado.
	Focused *Client

	// Monitor donde está actualmente.
	Monitor *Monitor

	// Layout actual.
	Layout Layout

	// Estado.
	Visible bool
}

// NewWorkspace crea un nuevo workspace.
func NewWorkspace(id int, name string) *Workspace {
	return &Workspace{
		ID:       id,
		Name:     name,
		Clients:  []*Client{},
		Visible:  false,
		Focused:  nil,
		Monitor:  nil,
	}
}

// AddClient agrega una ventana al workspace.
func (ws *Workspace) AddClient(c *Client) {
	if c == nil {
		return
	}

	ws.Clients = append(ws.Clients, c)
	c.Workspace = ws
}

// RemoveClient elimina una ventana del workspace.
func (ws *Workspace) RemoveClient(c *Client) {
	if c == nil {
		return
	}

	for i, client := range ws.Clients {
		if client == c {
			ws.Clients = append(ws.Clients[:i], ws.Clients[i+1:]...)
			break
		}
	}

	if ws.Focused == c {
		ws.Focused = nil
	}

	c.Workspace = nil
}

// ClientCount devuelve el número de ventanas.
func (ws *Workspace) ClientCount() int {
	return len(ws.Clients)
}

// Empty indica si el workspace está vacío.
func (ws *Workspace) Empty() bool {
	return len(ws.Clients) == 0
}

// SetFocused cambia el cliente enfocado.
func (ws *Workspace) SetFocused(c *Client) {
	ws.Focused = c
}

// Show marca el workspace como visible.
func (ws *Workspace) Show() {
	ws.Visible = true
}

// Hide marca el workspace como oculto.
func (ws *Workspace) Hide() {
	ws.Visible = false
}

// SetMonitor asigna el monitor.
func (ws *Workspace) SetMonitor(m *Monitor) {
	ws.Monitor = m
}

// SetLayout cambia el layout.
func (ws *Workspace) SetLayout(layout Layout) {
	ws.Layout = layout
}
