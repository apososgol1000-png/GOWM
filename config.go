package main

import (
	"astralwm/sdk"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("=== Ejecutando Configuración de AstralOS ===")

	// 1. Conectarse al core del Window Manager
	wm, err := sdk.Connect()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer wm.Close()

	// 2. Aplicar configuraciones iniciales
	fmt.Println("Configurando Gaps a 15px...")
	wm.SetGaps(15)

	// 3. Lanzar aplicaciones de prueba con pausas
	fmt.Println("Lanzando terminal de prueba...")
	wm.Spawn("xterm") // Cambia a "mate-terminal" si no usas xterm

	time.Sleep(2 * time.Second)

	fmt.Println("Configuración aplicada con éxito.")
}
