package wm

import (
	"fmt"

	"astralwm/internal/ipc"
	"astralwm/internal/x11"

	"github.com/BurntSushi/xgbutil/event"
)

// Run inicia AstralWM.
func (wm *WM) Run() error {

	fmt.Println("===================================")
	fmt.Println("      AstralWM iniciando...")
	fmt.Println("===================================")

	// Inicializar estado.
	wm.State = NewState()

	// Inicializar monitores.
	if err := wm.InitMonitors(); err != nil {
		return err
	}

	// Inicializar workspaces.
	if err := wm.InitWorkspaces(); err != nil {
		return err
	}

	// Registrar eventos X11.
	if err := wm.RegisterEvents(); err != nil {
		return err
	}

	// Registrar atajos de teclado.
	if err := wm.RegisterKeybinds(); err != nil {
		return err
	}

	// Inicializar EWMH.
	if err := x11.InitEWMH(wm.X); err != nil {
		return err
	}

	// Iniciar servidor IPC.
	if err := ipc.Start(wm); err != nil {
		return err
	}

	// El WM ya está listo.
	wm.State.Initialized = true
	wm.Running = true

	fmt.Println("AstralWM listo.")
	fmt.Println("Esperando eventos...")

	// Bucle principal de eventos.
	event.Main(wm.X)

	return nil
}
