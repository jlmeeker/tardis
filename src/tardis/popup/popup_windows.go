package popup

import (
	"github.com/lxn/walk"
)

func ShowPopUpWindow(msg string) {
	walk.MsgBox(nil, "Tardis Update", msg, walk.MsgBoxIconWarning)
}
