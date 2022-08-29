package networks

import (
	"github.com/gdamore/tcell/v2"
)

// Draw draws this primitive onto the screen.
func (nets *Networks) Draw(screen tcell.Screen) {
	nets.Box.DrawForSubclass(screen, nets)
	nets.Box.SetBorder(false)
	x, y, width, height := nets.GetInnerRect()
	nets.table.SetRect(x, y, width, height)
	nets.table.SetBorder(true)

	nets.table.Draw(screen)
	x, y, width, height = nets.table.GetInnerRect()
	// error dialog
	if nets.errorDialog.IsDisplay() {
		nets.errorDialog.SetRect(x, y, width, height)
		nets.errorDialog.Draw(screen)
		return
	}
	// command dialog dialog
	if nets.cmdDialog.IsDisplay() {
		nets.cmdDialog.SetRect(x, y, width, height)
		nets.cmdDialog.Draw(screen)
		return
	}
	// create dialog dialog
	if nets.createDialog.IsDisplay() {
		nets.createDialog.SetRect(x, y, width, height)
		nets.createDialog.Draw(screen)
		return
	}

	// connect dialog dialog
	if nets.connectDialog.IsDisplay() {
		nets.connectDialog.SetRect(x, y, width, height)
		nets.connectDialog.Draw(screen)

		return
	}

	// disconnect dialog dialog
	if nets.disconnectDialog.IsDisplay() {
		nets.disconnectDialog.SetRect(x, y, width, height)
		nets.disconnectDialog.Draw(screen)

		return
	}

	// message dialog
	if nets.messageDialog.IsDisplay() {
		nets.messageDialog.SetRect(x, y, width, height+1)
		nets.messageDialog.Draw(screen)
		return
	}
	// confirm dialog
	if nets.confirmDialog.IsDisplay() {
		nets.confirmDialog.SetRect(x, y, width, height)
		nets.confirmDialog.Draw(screen)
		return
	}
	// progress dialog
	if nets.progressDialog.IsDisplay() {
		nets.progressDialog.SetRect(x, y, width, height)
		nets.progressDialog.Draw(screen)
	}
}
