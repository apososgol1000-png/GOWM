package wm

// Monitor representa una pantalla física.
type Monitor struct {
	// Identificador del monitor.
	ID int

	// Nombre (HDMI-1, eDP-1, DP-1...).
	Name string

	// Posición del monitor.
	X int
	Y int

	// Resolución.
	Width  int
	Height int

	// Área utilizable (después de reservar espacio para barras).
	WorkX      int
	WorkY      int
	WorkWidth  int
	WorkHeight int

	// Workspace actualmente visible.
	Current *Workspace

	// Todos los workspaces asignados.
	Workspaces []*Workspace

	// Estado.
	Primary bool
	Enabled bool
}

// NewMonitor crea un nuevo monitor.
func NewMonitor(id int, name string, x, y, width, height int) *Monitor {
	return &Monitor{
		ID:         id,
		Name:       name,
		X:          x,
		Y:          y,
		Width:      width,
		Height:     height,
		WorkX:      x,
		WorkY:      y,
		WorkWidth:  width,
		WorkHeight: height,
		Workspaces: []*Workspace{},
		Enabled:    true,
	}
}

// Geometry devuelve la geometría del monitor.
func (m *Monitor) Geometry() (int, int, int, int) {
	return m.X, m.Y, m.Width, m.Height
}

// WorkArea devuelve el área útil del monitor.
func (m *Monitor) WorkArea() (int, int, int, int) {
	return m.WorkX, m.WorkY, m.WorkWidth, m.WorkHeight
}

// SetWorkArea cambia el área útil.
func (m *Monitor) SetWorkArea(x, y, width, height int) {
	m.WorkX = x
	m.WorkY = y
	m.WorkWidth = width
	m.WorkHeight = height
}

// AddWorkspace agrega un workspace.
func (m *Monitor) AddWorkspace(ws *Workspace) {
	if ws == nil {
		return
	}

	m.Workspaces = append(m.Workspaces, ws)
	ws.Monitor = m

	if m.Current == nil {
		m.Current = ws
		ws.Show()
	}
}

// RemoveWorkspace elimina un workspace.
func (m *Monitor) RemoveWorkspace(ws *Workspace) {
	if ws == nil {
		return
	}

	for i, workspace := range m.Workspaces {
		if workspace == ws {
			m.Workspaces = append(
				m.Workspaces[:i],
				m.Workspaces[i+1:]...,
			)
			break
		}
	}

	if m.Current == ws {
		m.Current = nil
	}

	ws.Monitor = nil
}

// SetCurrentWorkspace cambia el workspace activo.
func (m *Monitor) SetCurrentWorkspace(ws *Workspace) {
	if ws == nil {
		return
	}

	if m.Current != nil {
		m.Current.Hide()
	}

	m.Current = ws
	ws.Show()
}

// WorkspaceCount devuelve el número de workspaces.
func (m *Monitor) WorkspaceCount() int {
	return len(m.Workspaces)
}

// IsPrimary indica si es el monitor principal.
func (m *Monitor) IsPrimary() bool {
	return m.Primary
}

// SetPrimary marca este monitor como principal.
func (m *Monitor) SetPrimary(primary bool) {
	m.Primary = primary
}

// Enable habilita el monitor.
func (m *Monitor) Enable() {
	m.Enabled = true
}

// Disable deshabilita el monitor.
func (m *Monitor) Disable() {
	m.Enabled = false
}
