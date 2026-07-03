package wm

// State representa el estado actual de AstralWM.
type State struct {
	// ¿El WM ya terminó de inicializar?
	Initialized bool

	// Workspace actualmente activo.
	CurrentWorkspace *Workspace

	// Monitor actualmente enfocado.
	CurrentMonitor *Monitor

	// Ventana enfocada.
	FocusedClient *Client

	// Layout actualmente activo.
	CurrentLayout Layout

	// Número de clientes administrados.
	ClientCount int
}

// NewState crea un nuevo estado para el WM.
func NewState() *State {
	return &State{
		Initialized: false,
		ClientCount: 0,
	}
}

// SetWorkspace cambia el workspace activo.
func (s *State) SetWorkspace(ws *Workspace) {
	s.CurrentWorkspace = ws
}

// SetMonitor cambia el monitor activo.
func (s *State) SetMonitor(m *Monitor) {
	s.CurrentMonitor = m
}

// SetFocusedClient cambia la ventana enfocada.
func (s *State) SetFocusedClient(c *Client) {
	s.FocusedClient = c
}

// SetLayout cambia el layout activo.
func (s *State) SetLayout(layout Layout) {
	s.CurrentLayout = layout
}

// IncrementClients aumenta el contador.
func (s *State) IncrementClients() {
	s.ClientCount++
}

// DecrementClients disminuye el contador.
func (s *State) DecrementClients() {
	if s.ClientCount > 0 {
		s.ClientCount--
	}
}
