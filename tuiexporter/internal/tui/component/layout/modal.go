package layout

import (
	"github.com/rivo/tview"
)

func AttachModalForTreeAttributes(tree *tview.TreeView, onHide func()) {
	_ = "STUB: not implemented"
	return
}

type tableModalMapper interface {
	// GetColumnIdx returns the column index for getting the content to be shown
	// in the modal
	GetColumnIdx() int
}

func AttachModalForTableRows(table *tview.Table, mapper tableModalMapper, onHide func()) {
	_ = "STUB: not implemented"
	return
}
