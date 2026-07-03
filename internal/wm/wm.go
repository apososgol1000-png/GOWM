package wm

import (
	"astralwm/internal/config"
	"astralwm/internal/x11"
	"errors"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
)

// WM representa la instancia principal de AstralWM.
type WM struct {
	// Conexión con X11.
	X *xgbutil.XUtil

	// Configuración cargada.
	Config *config.Config

	// Estado del Window Manager.
	Running bool

	// Clientes administrados.
	Clients map[xproto.Window]*Client

	// Monitores disponibles.
	Monitors []*Monitor

	// Workspaces.
	Workspaces []*Workspace

	// Cliente enfocado.
	Focused *Client
}

// New crea una nueva instancia del Window Manager.
func New() (*WM, error) {

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	xu, err := x11.NewConnection()
	if err != nil {
		return nil, err
	}

	wm := &WM{
		X:          xu,
		Config:     cfg,
		Running:    false,
		Clients:    make(map[xproto.Window]*Client),
		Monitors:   []*Monitor{},
		Workspaces: []*Workspace{},
		Focused:    nil,
	}

	return wm, nil
}

// Run inicia AstralWM.
func (wm *WM) Run() error {

	if wm.Running {
		return errors.New("window manager ya iniciado")
	}

	wm.Running = true

	if err := wm.Init(); err != nil {
		return err
	}

	return wm.EventLoop()
}

// Stop detiene AstralWM.
func (wm *WM) Stop() {

	if !wm.Running {
		return
	}

	wm.Running = false

	if wm.X != nil {
		wm.X.Conn().Close()
	}
}

// Init inicializa todos los subsistemas.
func (wm *WM) Init() error {

	if err := wm.InitMonitors(); err != nil {
		return err
	}

	if err := wm.InitWorkspaces(); err != nil {
		return err
	}

	if err := wm.RegisterEvents(); err != nil {
		return err
	}

	if err := wm.RegisterKeybinds(); err != nil {
		return err
	}

	if err := wm.StartIPC(); err != nil {
		return err
	}

	return nil
}

// EventLoop inicia el bucle principal de eventos.
func (wm *WM) EventLoop() error {

	x11.MainLoop(wm.X)

	return nil
}

// IsRunning indica si el WM sigue activo.
func (wm *WM) IsRunning() bool {
	return wm.Running
}
