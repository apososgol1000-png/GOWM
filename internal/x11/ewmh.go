Este es uno de los archivos más importantes.

EWMH significa:

Extended Window Manager Hints

Es un estándar que usan las aplicaciones para comunicarse con el Window Manager.

Gracias a EWMH funcionan correctamente cosas como:

Pantalla completa (F11 en Firefox o VLC).
Barras como Polybar, Waybar (en X11), Tint2, etc.
Docks.
Notificaciones.
Saber cuál ventana está enfocada.
Cambiar de escritorio.
Mostrar el nombre de la ventana.

Aquí manejarías propiedades como:

_NET_ACTIVE_WINDOW

_NET_CURRENT_DESKTOP

_NET_CLIENT_LIST

_NET_NUMBER_OF_DESKTOPS

_NET_WM_STATE

_NET_WM_STATE_FULLSCREEN

_NET_WM_NAME

_NET_SUPPORTED

Funciones típicas:

func InitEWMH()

func SetActiveWindow()

func SetCurrentDesktop()

func SetFullscreen()

func UpdateClientList()

Sin EWMH, muchas aplicaciones modernas no interactúan correctamente con el WM.
